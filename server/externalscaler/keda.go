package externalscaler

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	pb "nwdaf-otel/generated/externalscaler"
	"time"
)

/*
Code for KEDA NWDAF External Scaler Integration
Flow:
KEDA -> External Scaler (this package) [KEDA asks for service specific resources for scaling check]
External Scaler -> NWDAF [query relevant metric value]
NWDAF -> External Scaler [returns NWDAF controlled metric values, dynamically changing]
External Scaler -> KEDA [Based on the returned values, NWDAF implicitly controls scaling with KEDA acting as the operator]
*/

const (
	MetricName       = "udm_max_latency"
	LatencyEndpoint  = "http://nwdaf-analytics-info.aether-5gc.svc.cluster.local:8080/latency/udm"
	LatencyThreshold = float64(4.0)
)

// Documentation: https://keda.sh/docs/2.18/concepts/external-scalers/
type Scaler struct {
	pb.UnimplementedExternalScalerServer
}

func (s *Scaler) getLatency(ctx context.Context) (float64, error) {
	r, err := http.NewRequestWithContext(ctx, "GET", LatencyEndpoint, nil)
	if err != nil {
		return 0, fmt.Errorf("error creating get latency request: %w", err)
	}

	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return 0, fmt.Errorf("error doing get latency request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error getting latency status: %d", res.StatusCode)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading response body: %w", err)
	}

	var payload float64
	err = json.Unmarshal(body, &payload)
	if err != nil {
		return 0, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return payload, nil
}

func (s *Scaler) IsActive(ctx context.Context, req *pb.ScaledObjectRef) (*pb.IsActiveResponse, error) {
	log.Println("IsActive")
	// The value can be a latency model that can contain values, thresholds ...
	value, err := s.getLatency(ctx)
	if err != nil {
		log.Printf("error getting latency value when serving IsActive: %v", err)
		return &pb.IsActiveResponse{Result: false}, nil
	}

	/* 	Example: Active when latency > 4.0 seconds
	NWDAF can directly control when scaling should be active!
	e.g. NWDAF can also return a threshold value, or just boolean true or false -> can do the evaluation inside!
	*/
	return &pb.IsActiveResponse{Result: value > LatencyThreshold}, nil
}

func (s *Scaler) StreamIsActive(req *pb.ScaledObjectRef, kedaServer pb.ExternalScaler_StreamIsActiveServer) error {
	//longitude := req.ScalerMetadata["longitude"]
	//latitude := req.ScalerMetadata["latitude"]
	//
	//if len(longitude) == 0 || len(latitude) == 0 {
	//	return status.Error(codes.InvalidArgument, "longitude and latitude must be specified")
	//}
	log.Println("StreamIsActive")
	for {
		select {
		case <-kedaServer.Context().Done():
			// call cancelled
			return nil
		case <-time.Tick(time.Second):
			value, err := s.getLatency(kedaServer.Context())
			if err != nil {
				log.Printf("error getting latency value when serving StreamIsActive: %v", err)
				continue
			}
			_ = kedaServer.Send(&pb.IsActiveResponse{
				Result: value > LatencyThreshold,
			})
		}
	}
}

func (s *Scaler) GetMetricSpec(ctx context.Context, req *pb.ScaledObjectRef) (*pb.GetMetricSpecResponse, error) {
	// Provide target value
	log.Println("GetMetricSpec")
	return &pb.GetMetricSpecResponse{
		MetricSpecs: []*pb.MetricSpec{
			{
				MetricName: MetricName,
				// 4 seconds
				TargetSize: 4000, //200, // target: 200ms
			},
		},
	}, nil
}

func (s *Scaler) GetMetrics(ctx context.Context, metricReq *pb.GetMetricsRequest) (*pb.GetMetricsResponse, error) {
	log.Println("GetMetrics")
	value, err := s.getLatency(ctx)
	if err != nil {
		log.Printf("error getting latency value when serving GetMetrics: %v", err)
		value = 0.0
	}
	return &pb.GetMetricsResponse{
		MetricValues: []*pb.MetricValue{
			{
				MetricName:  MetricName,
				MetricValue: int64(value * 1000),
			},
		},
	}, nil

	//resp, _ := http.Get("http://my-microservice.default.svc.cluster.local/my-latency")
	//defer resp.Body.Close()
	//
	//return &pb.GetMetricsResponse{
	//	MetricValues: []*pb.MetricValue{
	//		{
	//			MetricName:       metricReq.MetricName,
	//			MetricValueFloat: value * 1000,
	//		},
	//	},
	//}, nil
}
