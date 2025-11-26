package externalscaler

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	exscaler "nwdaf-otel/generated/externalscaler"
	"sync"
)

type ScalingServer struct {
	kedaServer  *grpc.Server
	nwdafServer *http.Server
	wg          sync.WaitGroup
}

func NewExternalScalerServer() *ScalingServer {
	scalingServer := &ScalingServer{
		kedaServer: grpc.NewServer(),
		nwdafServer: &http.Server{
			Addr: ":8081",
		},
	}
	exscaler.RegisterExternalScalerServer(scalingServer.kedaServer, &scaler{})
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

	listener, err := net.Listen("tcp", ":8080")
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
	_ = s.nwdafServer.ListenAndServe()
}
