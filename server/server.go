package server

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	analyticsinfoAPI "nwdaf-otel/generated/analyticsinfo"
	"nwdaf-otel/server/analyticsinfo"
	"time"
)

type Server interface {
	Setup()
	Start(shutdownChn chan struct{}) chan error
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	err := r.Body.Close()
	if err != nil {
		log.Println("err on closing health check request body: " + err.Error())
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("ok"))
	if err != nil {
		log.Println("err on writing health check response 'ok': " + err.Error())
	}
}

// TODO: accept a config, that will point to port, certificate e.g. options
type analyticsInfoServer struct {
	mux *mux.Router
	srv *http.Server
}

func NewAnalyticsInfoServer() Server {
	return &analyticsInfoServer{mux: nil}
}

func (s *analyticsInfoServer) Setup() {
	analyticsDocumentService := analyticsinfo.NewNWDAFAnalyticsDocumentAPIService()
	analyticsDocumentController := analyticsinfo.NewNWDAFAnalyticsDocumentAPIController(analyticsDocumentService)
	contextDocumentService := analyticsinfo.NewNWDAFContextDocumentAPIService()
	contextDocumentController := analyticsinfo.NewNWDAFContextDocumentAPIController(contextDocumentService)
	s.mux = analyticsinfoAPI.NewRouter(analyticsDocumentController, contextDocumentController)

	// Handle health check probes
	s.mux.HandleFunc("/health", handleHealthCheck)
}

func (s *analyticsInfoServer) Start(shutdownChn chan struct{}) chan error {
	errChan := make(chan error)
	go s.serve(errChan)
	go s.stopGracefully(shutdownChn)
	return errChan
}

func (s *analyticsInfoServer) serve(errChan chan error) {
	s.srv = &http.Server{
		Addr:    ":8080",
		Handler: s.mux,
	}
	log.Println("Listening and serving HTTP on " + s.srv.Addr)
	if err := s.srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		errChan <- err
	}
	close(errChan)
}

func (s *analyticsInfoServer) stopGracefully(shutdownChn chan struct{}) {
	<-shutdownChn
	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := s.srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server http shutdown error: %v", err)
	}
}
