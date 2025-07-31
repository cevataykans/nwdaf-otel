package prometheus

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"log"
	"net/http"
	"time"
)

const (
	Address = "http://rancher-monitoring-prometheus.cattle-monitoring-system.svc.cluster.local:9090"
)

type Client struct {
	promClient v1.API
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
	return &Client{
		apiClient,
	}, nil
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
