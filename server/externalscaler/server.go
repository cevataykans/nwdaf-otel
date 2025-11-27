package externalscaler

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	externalScaler "nwdaf-otel/generated/externalscaler"
	"sync"
)

const (
	HttpPort = 8080
	GrpcPort = 8081
)

type ScalingServer struct {
	kedaServer  *grpc.Server
	nwdafServer *http.Server
	wg          sync.WaitGroup
}

func NewExternalScalerServer() *ScalingServer {
	grpcServer := grpc.NewServer()
	externalScaler.RegisterExternalScalerServer(grpcServer, &Scaler{})
	scalingServer := &ScalingServer{
		kedaServer: grpcServer,
		nwdafServer: &http.Server{
			Addr: fmt.Sprintf(":%v", HttpPort),
		},
	}
	return scalingServer
}

func (s *ScalingServer) Start() {
	go s.serveGrpc()
	go s.serveHTTP()
}

func (s *ScalingServer) Stop(ctx context.Context) {
	go s.kedaServer.GracefulStop()
	go s.nwdafServer.Shutdown(ctx)
	s.wg.Wait()
}

func (s *ScalingServer) serveGrpc() {
	s.wg.Add(1)
	defer s.wg.Done()

	log.Printf("Starting gRPC server on port %v", GrpcPort)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", GrpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = s.kedaServer.Serve(listener)
	if err != nil {
		log.Printf("failed to serve: %v", err)
	}
}

func (s *ScalingServer) serveHTTP() {
	s.wg.Add(1)
	defer s.wg.Done()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		err := r.Body.Close()
		if err != nil {
			log.Println("err on closing health check request body: " + err.Error())
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("ok"))
		if err != nil {
			log.Println("err on writing health check response 'ok': " + err.Error())
		}
	})
	s.nwdafServer.Handler = mux
	log.Printf("Starting HTTP server on port %v", HttpPort)
	_ = s.nwdafServer.ListenAndServe()
}
