package externalscaler

import (
	"context"
	"math/rand"
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

const MetricName = "udm_latency"

// Documentation: https://keda.sh/docs/2.18/concepts/external-scalers/
type scaler struct {
	pb.UnimplementedExternalScalerServer
}

func (s *scaler) IsActive(ctx context.Context, req *pb.ScaledObjectRef) (*pb.IsActiveResponse, error) {
	// Call your microservice
	//resp, err := http.Get("http://my-microservice.default.svc.cluster.local/my-latency")
	//if err != nil {
	//	return &pb.IsActiveResponse{Result: false}, nil
	//}
	//defer resp.Body.Close()
	//
	//body, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	//payload := USGSResponse{}
	//err = json.Unmarshal(body, &payload)
	//if err != nil {
	//	return nil, status.Error(codes.Internal, err.Error())
	//}

	// Example: Active when latency > 200ms
	// NWDAF can directly control when scaling should be active!
	return &pb.IsActiveResponse{Result: rand.Float64()*100 > 0.9}, nil
}

func (s *scaler) StreamIsActive(req *pb.ScaledObjectRef, kedaServer pb.ExternalScaler_StreamIsActiveServer) error {
	//longitude := req.ScalerMetadata["longitude"]
	//latitude := req.ScalerMetadata["latitude"]
	//
	//if len(longitude) == 0 || len(latitude) == 0 {
	//	return status.Error(codes.InvalidArgument, "longitude and latitude must be specified")
	//}

	for {
		select {
		case <-kedaServer.Context().Done():
			// call cancelled
			return nil
		case <-time.Tick(time.Second):
			//earthquakeCount, err := getEarthQuakeCount(longitude, latitude)
			//if err != nil {
			//	// log error
			//	continue
			//}

			if rand.Float64()*100 > 0.9 {
				_ = kedaServer.Send(&pb.IsActiveResponse{
					Result: true,
				})
			}
		}
	}
}

func (s *scaler) GetMetricSpec(ctx context.Context, req *pb.ScaledObjectRef) (*pb.GetMetricSpecResponse, error) {
	// Provide target value
	return &pb.GetMetricSpecResponse{
		MetricSpecs: []*pb.MetricSpec{
			{
				MetricName:      MetricName,
				TargetSizeFloat: 2000, //200, // target: 200ms
			},
		},
	}, nil
}

func (s *scaler) GetMetrics(ctx context.Context, metricReq *pb.GetMetricsRequest) (*pb.GetMetricsResponse, error) {
	return &pb.GetMetricsResponse{
		MetricValues: []*pb.MetricValue{
			{
				MetricName:       MetricName,
				MetricValueFloat: rand.Float64() * 100,
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
