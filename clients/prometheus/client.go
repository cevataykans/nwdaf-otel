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
	c := &http.Client{
		Transport: tp,
	}
	promClient, err := api.NewClient(
		api.Config{
			Address:      Address,
			Client:       c,
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

func (c *Client) QueryCPUTotalSeconds() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	currentSecond := time.Now().UTC().Second()
	results, warnings, err := c.promClient.QueryRange(ctx, "rate(container_cpu_usage_seconds_total{container=\"bessd\"}[1m])", v1.Range{
		Start: time.Unix(int64(currentSecond)-60, 0).UTC(),
		End:   time.Unix(int64(currentSecond), 0).UTC(),
	})
	if err != nil {
		return fmt.Errorf("error querying Prometheus: %v", err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}

	values, ok := results.(model.Vector)
	if !ok {
		return fmt.Errorf("result vector is not a vector, but actual: ", results.Type())
	}

	// Iterate over the vector
	for _, sample := range values {
		fmt.Printf("Metric: %v\n", sample.Metric)
		fmt.Printf("Value: %v\n", sample.Value)
		fmt.Printf("Timestamp: %v\n", sample.Timestamp.Time())
	}
	return nil
}
