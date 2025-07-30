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

func (c *Client) QueryCPUTotalSeconds(start, end time.Time, step time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	results, warnings, err := c.promClient.QueryRange(ctx, "rate(container_cpu_usage_seconds_total{container=\"bessd\"}[1m])", v1.Range{
		Start: start,
		End:   end,
		Step:  step,
	})
	if err != nil {
		return fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}

	matrix, ok := results.(model.Matrix)
	if !ok {
		return fmt.Errorf("result vector is not a vector, but actual: ", results.Type())
	}

	// Iterate over the vector
	for _, row := range matrix {
		log.Printf("Metric: %v\n", row.Metric)
		for _, value := range row.Values {
			log.Printf("Timestamp: %v - Value: %v\n", value.Timestamp, value.Value)
		}
	}
	return nil
}
