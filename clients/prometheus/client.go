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
	Address        = "http://rancher-monitoring-prometheus.cattle-monitoring-system.svc.cluster.local:9090"
	ElasticAddress = "http://elasticsearch-master.default.svc.cluster.local:9200"
)

type Client struct {
	promClient v1.API
	esClient   *elasticsearch.Client
}

func NewClient() (*Client, error) {
	tp := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	promClient, err := api.NewClient(
		api.Config{
			Address:      Address,
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

func (c *Client) QueryTraces(service string, start, end time.Time) error {
	// Build the Elasticsearch query
	//query := map[string]interface{}{
	//	"query": map[string]interface{}{
	//		"bool": map[string]interface{}{
	//			"must": []map[string]interface{}{
	//				{
	//					"wildcard": map[string]interface{}{
	//						"process.serviceName": map[string]interface{}{
	//							"value": fmt.Sprintf("%s*", service),
	//						},
	//					},
	//				},
	//				{
	//					"range": map[string]interface{}{
	//						"startTimeMillis": map[string]interface{}{
	//							// Jaegar timestamps are in Unix milliseconds
	//							"gte": start.Unix() * 1000,
	//							"lte": end.Unix() * 1000,
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//	//"_source": []string{"traceID", "operationName", "process.serviceName", "startTime"},
	//	//"size":    1000,
	//}
	queryEntity := CreateESAvgQuery(service, start, end)

	// Encode query to JSON
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(queryEntity); err != nil {
		log.Fatalf("Error encoding query: %s", err)
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
		return fmt.Errorf("error getting response: %v", err)
	}
	defer res.Body.Close()

	// Decode response
	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return fmt.Errorf("error parsing response body: %v", err)
	}

	// Print results
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//	doc := hit.(map[string]interface{})["_source"]
	//	traceID := doc.(map[string]interface{})["traceID"]
	//	service := doc.(map[string]interface{})["process"]
	//	serviceName := service.(map[string]interface{})["serviceName"]
	//	op := doc.(map[string]interface{})["operationName"]
	//	start := doc.(map[string]interface{})["startTime"]
	//	duration := doc.(map[string]interface{})["duration"]
	//
	//	fmt.Printf("TraceID: %s | Service: %s | Operation: %s | StartTime: %v | Duration: %v\n",
	//		traceID, serviceName, op, start, duration)
	//}
	aggregation := r["aggregations"].(map[string]interface{})
	durationDoc := aggregation["duration"].(map[string]interface{})
	count := durationDoc["doc_count"].(int)
	if count == 0 {
		return nil
	}
	avgDurationDoc := durationDoc["avg_duration"].(map[string]interface{})
	value := avgDurationDoc["value"].(float64)
	fmt.Printf("Avg duration of service '%s' traces: %v\n", service, value)
	return nil
}

func minutesToTime(m int64) time.Time {
	return time.Unix(time.Now().Unix()/86400*86400+m*60, 0)
}

func (c *Client) QueryMetrics(service string, start, end time.Time, step time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r := v1.Range{Start: start, End: end, Step: step}
	err := c.queryCPUTotalSeconds(ctx, service, r)
	if err != nil {
		return err
	}

	err = c.queryMemory(ctx, service, r)
	if err != nil {
		return err
	}

	err = c.queryNetworkBytesReceived(ctx, service, r)
	if err != nil {
		return err
	}

	err = c.queryNetworkBytesSent(ctx, service, r)
	if err != nil {
		return err
	}

	err = c.queryReceivePacketsTotal(ctx, service, r)
	if err != nil {
		return err
	}

	err = c.queryTransmitPacketsTotal(ctx, service, r)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) queryNetworkBytesSent(ctx context.Context, service string, r v1.Range) error {
	query := fmt.Sprintf("rate(container_network_transmit_bytes_total{pod=~\"%s.*\", interface=\"eth0\"}[1m])", service)
	results, warnings, err := c.promClient.QueryRange(ctx, query, r)
	if err != nil {
		return fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}

	matrix, ok := results.(model.Matrix)
	if !ok {
		return fmt.Errorf("result vector is not a vector, but actual: %v", results.Type())
	}

	// Iterate over the vector
	for _, row := range matrix {
		log.Printf("Metric: %v\n", row.Metric)
		for _, value := range row.Values {
			log.Printf("Timestamp: %v - Value: %v\n", value.Timestamp, value.Value)
		}
	}
	log.Println("BYTES SENT TOTAL METRIC DONE")
	return nil
}

func (c *Client) queryNetworkBytesReceived(ctx context.Context, service string, r v1.Range) error {
	query := fmt.Sprintf("rate(container_network_receive_bytes_total{pod=~\"%s.*\", interface=\"eth0\"}[1m])", service)
	results, warnings, err := c.promClient.QueryRange(ctx, query, r)
	if err != nil {
		return fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}

	matrix, ok := results.(model.Matrix)
	if !ok {
		return fmt.Errorf("result vector is not a vector, but actual: %v", results.Type())
	}

	// Iterate over the vector
	for _, row := range matrix {
		log.Printf("Metric: %v\n", row.Metric)
		for _, value := range row.Values {
			log.Printf("Timestamp: %v - Value: %v\n", value.Timestamp, value.Value)
		}
	}
	log.Println("BYTES RECEIVED TOTAL METRIC DONE")
	return nil
}

func (c *Client) queryReceivePacketsTotal(ctx context.Context, service string, r v1.Range) error {
	query := fmt.Sprintf("rate(container_network_receive_packets_total{pod=~\"%s.*\", interface=\"eth0\"}[1m])", service)
	results, warnings, err := c.promClient.QueryRange(ctx, query, r)
	if err != nil {
		return fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}

	matrix, ok := results.(model.Matrix)
	if !ok {
		return fmt.Errorf("result vector is not a vector, but actual: %v", results.Type())
	}

	// Iterate over the vector
	for _, row := range matrix {
		log.Printf("Metric: %v\n", row.Metric)
		for _, value := range row.Values {
			log.Printf("Timestamp: %v - Value: %v\n", value.Timestamp, value.Value)
		}
	}
	log.Println("RECEIVE PACKETS TOTAL METRIC DONE")
	return nil
}

func (c *Client) queryTransmitPacketsTotal(ctx context.Context, service string, r v1.Range) error {
	query := fmt.Sprintf("rate(container_network_transmit_packets_total{pod=~\"%s.*\", interface=\"eth0\"}[1m])", service)
	results, warnings, err := c.promClient.QueryRange(ctx, query, r)
	if err != nil {
		return fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}

	matrix, ok := results.(model.Matrix)
	if !ok {
		return fmt.Errorf("result vector is not a vector, but actual: %v", results.Type())
	}

	// Iterate over the vector
	for _, row := range matrix {
		log.Printf("Metric: %v\n", row.Metric)
		for _, value := range row.Values {
			log.Printf("Timestamp: %v - Value: %v\n", value.Timestamp, value.Value)
		}
	}
	log.Println("TRANSMIT PACKETS TOTAL METRIC DONE")
	return nil
}

func (c *Client) queryMemory(ctx context.Context, service string, r v1.Range) error {
	query := fmt.Sprintf("avg_over_time(container_memory_usage_bytes{container=\"%s\"}[1m])", service)
	results, warnings, err := c.promClient.QueryRange(ctx, query, r)
	if err != nil {
		return fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}

	matrix, ok := results.(model.Matrix)
	if !ok {
		return fmt.Errorf("result vector is not a vector, but actual: %v", results.Type())
	}

	// Iterate over the vector
	for _, row := range matrix {
		log.Printf("Metric: %v\n", row.Metric)
		for _, value := range row.Values {
			log.Printf("Timestamp: %v - Value: %v\n", value.Timestamp, value.Value)
		}
	}
	log.Println("MEMORY METRIC DONE")
	return nil
}

func (c *Client) queryCPUTotalSeconds(ctx context.Context, service string, r v1.Range) error {
	query := fmt.Sprintf("rate(container_cpu_usage_seconds_total{container=\"%s\"}[1m])", service)
	results, warnings, err := c.promClient.QueryRange(ctx, query, r)
	if err != nil {
		return fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}

	matrix, ok := results.(model.Matrix)
	if !ok {
		return fmt.Errorf("result vector is not a vector, but actual: %v", results.Type())
	}

	// Iterate over the vector
	for _, row := range matrix {
		log.Printf("Metric: %v\n", row.Metric)
		for _, value := range row.Values {
			log.Printf("Timestamp: %v - Value: %v\n", value.Timestamp, value.Value)
		}
	}
	log.Println("CPU METRIC DONE")
	return nil
}
