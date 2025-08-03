package prometheus

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"log"
	"net/http"
	"time"
)

const (
	PrometheusAddress = "http://rancher-monitoring-prometheus.cattle-monitoring-system.svc.cluster.local:9090"
	ElasticAddress    = "http://elasticsearch-master.default.svc.cluster.local:9200"
)

type MetricResults struct {
	CpuTotalSeconds             float64
	MemoryTotalBytes            float64
	NetworkReceiveBytesTotal    float64
	NetworkTransmitBytesTotal   float64
	NetworkReceivePacketsTotal  float64
	NetworkTransmitPacketsTotal float64
}

type Client struct {
	promClient v1.API
	esClient   *elasticsearch.Client
}

func NewClient() (*Client, error) {
	// Create Prometheus client
	tp := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	promClient, err := api.NewClient(
		api.Config{
			Address:      PrometheusAddress,
			RoundTripper: tp,
		})
	if err != nil {
		return nil, fmt.Errorf("error creating prometheus client: %v", err)
	}
	apiClient := v1.NewAPI(promClient)

	// Create Elasticsearch client
	esClient, err := elasticsearch.NewClient(elasticsearch.Config{Addresses: []string{ElasticAddress}})
	if err != nil {
		return nil, fmt.Errorf("error creating elastic search client: %v", err)
	}
	return &Client{
		apiClient,
		esClient,
	}, nil
}

func (c *Client) QueryTraces(service string, start, end time.Time) (float64, error) {
	queryEntity := CreateESAvgQuery(service, start, end)
	// Encode query to JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(queryEntity); err != nil {
		return 0, fmt.Errorf("Error encoding query: %s", err)
	}

	// Perform the search request
	es := c.esClient
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("jaeger-span-*"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return 0, fmt.Errorf("error getting response: %v", err)
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			log.Printf("error closing response body: %v\n", err)
		}
	}()
	var avgRes ElasticsearchResponse
	if err := json.NewDecoder(res.Body).Decode(&avgRes); err != nil {
		return 0, fmt.Errorf("failed to decode response: %v", err)
	}

	if avgRes.TimedOut {
		return 0, fmt.Errorf("avg query timed out for service: %s", service)
	}
	log.Printf("Query took %v, scanned documents: %v\n", avgRes.Took, avgRes.Hits.Total.Value)

	avgDuration := avgRes.Aggregations.DurationAgg.AvgDuration.Value
	if avgRes.Aggregations.DurationAgg.DocCount == 0 || avgDuration == nil {
		return 0, nil
	}
	log.Printf("Avg duration of service '%s' traces: %v\n", service, *avgDuration)
	return *avgDuration, nil
}

func (c *Client) QueryMetrics(service string, start, end time.Time, step time.Duration) (MetricResults, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r := v1.Range{Start: start, End: end, Step: step}
	results := MetricResults{}
	res, err := c.queryCPUTotalSeconds(ctx, service, r)
	if err != nil {
		return results, err
	}
	results.CpuTotalSeconds = res

	res, err = c.queryMemory(ctx, service, r)
	if err != nil {
		return results, err
	}
	results.MemoryTotalBytes = res

	res, err = c.queryNetworkBytesReceived(ctx, service, r)
	if err != nil {
		return results, err
	}
	results.NetworkReceiveBytesTotal = res

	res, err = c.queryNetworkBytesSent(ctx, service, r)
	if err != nil {
		return results, err
	}
	results.NetworkTransmitBytesTotal = res

	res, err = c.queryReceivePacketsTotal(ctx, service, r)
	if err != nil {
		return results, err
	}
	results.NetworkReceivePacketsTotal = res

	res, err = c.queryTransmitPacketsTotal(ctx, service, r)
	if err != nil {
		return results, err
	}
	results.NetworkTransmitPacketsTotal = res
	return results, nil
}

func (c *Client) queryNetworkBytesSent(ctx context.Context, service string, r v1.Range) (float64, error) {
	query := fmt.Sprintf("rate(container_network_transmit_bytes_total{pod=~\"%s.*\", interface=\"eth0\"}[1m])", service)
	return c.queryPrometheus(ctx, query, r)
}

func (c *Client) queryNetworkBytesReceived(ctx context.Context, service string, r v1.Range) (float64, error) {
	query := fmt.Sprintf("rate(container_network_receive_bytes_total{pod=~\"%s.*\", interface=\"eth0\"}[1m])", service)
	return c.queryPrometheus(ctx, query, r)
}

func (c *Client) queryReceivePacketsTotal(ctx context.Context, service string, r v1.Range) (float64, error) {
	query := fmt.Sprintf("rate(container_network_receive_packets_total{pod=~\"%s.*\", interface=\"eth0\"}[1m])", service)
	return c.queryPrometheus(ctx, query, r)
}

func (c *Client) queryTransmitPacketsTotal(ctx context.Context, service string, r v1.Range) (float64, error) {
	query := fmt.Sprintf("rate(container_network_transmit_packets_total{pod=~\"%s.*\", interface=\"eth0\"}[1m])", service)
	return c.queryPrometheus(ctx, query, r)
}

func (c *Client) queryMemory(ctx context.Context, service string, r v1.Range) (float64, error) {
	query := fmt.Sprintf("avg_over_time(container_memory_usage_bytes{container=\"%s\"}[1m])", service)
	return c.queryPrometheus(ctx, query, r)
}

func (c *Client) queryCPUTotalSeconds(ctx context.Context, service string, r v1.Range) (float64, error) {
	query := fmt.Sprintf("rate(container_cpu_usage_seconds_total{container=\"%s\"}[1m])", service)
	return c.queryPrometheus(ctx, query, r)
}

func (c *Client) queryPrometheus(ctx context.Context, query string, r v1.Range) (float64, error) {
	results, warnings, err := c.promClient.QueryRange(ctx, query, r)
	if err != nil {
		return 0, fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}

	matrix, ok := results.(model.Matrix)
	if !ok {
		return 0, fmt.Errorf("result vector is not of type Matrix, but actual: %v", results.Type())
	}
	if len(matrix) > 1 {
		log.Printf("warning, matrix has more results then service with count: %v", len(matrix))
	}

	value := 0.0
	// Iterate over the vector
	//for _, row := range matrix {
	//	log.Printf("Metric: %v\n", row.Metric)
	//	for _, value := range row.Values {
	//		log.Printf("Timestamp: %v - Value: %v\n", value.Timestamp, value.Value)
	//	}
	//}
	if len(matrix) > 0 && len(matrix[0].Values) > 0 {
		end := len(matrix[0].Values) - 1
		value = float64(matrix[0].Values[end].Value)
		ts := int64(matrix[0].Values[end].Timestamp)
		log.Printf("Timestamp: %v - Value: %v\n", ts, value)
	}
	return value, nil
}
