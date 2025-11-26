package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"nwdaf-otel/clients/prometheus"
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

func handleUDMMetricRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/metric" {

	}
}

// TODO: accept a config, that will point to port, certificate e.g. options
type analyticsInfoServer struct {
	mux        *mux.Router
	srv        *http.Server
	latencySrv latencyHandler
}

type latencyHandler struct {
	client *prometheus.Client
	mux    *mux.Router
}

func (lh latencyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lh.mux.ServeHTTP(w, r)
}

func createLatencyHandler(pClient *prometheus.Client) latencyHandler {
	latencyMux := mux.NewRouter()
	latencyMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// return actual latency value
		val, err := pClient.QueryUDMLatency()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		bytes, err := json.Marshal(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(bytes)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("err on writing latency response 'ok': " + err.Error())
		}
	})
	latencyMux.HandleFunc("/metricspecs", func(w http.ResponseWriter, r *http.Request) {
		// return metric specs -> make it static for thesis, dynamic requires a statistical, ML approach
	})
	latencyMux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		// return metrics -> make it static for thesis, dynamic requires a statistical, ML approach
	})
	return latencyHandler{
		client: pClient,
		mux:    latencyMux,
	}
}

func NewAnalyticsInfoServer(pClient *prometheus.Client) Server {
	return &analyticsInfoServer{
		mux:        nil,
		latencySrv: createLatencyHandler(pClient),
	}
}

func (s *analyticsInfoServer) Setup() {
	analyticsDocumentService := analyticsinfo.NewNWDAFAnalyticsDocumentAPIService()
	analyticsDocumentController := analyticsinfo.NewNWDAFAnalyticsDocumentAPIController(analyticsDocumentService)
	contextDocumentService := analyticsinfo.NewNWDAFContextDocumentAPIService()
	contextDocumentController := analyticsinfo.NewNWDAFContextDocumentAPIController(contextDocumentService)
	s.mux = analyticsinfoAPI.NewRouter(analyticsDocumentController, contextDocumentController)

	// Handle health check probes
	s.mux.HandleFunc("/health", handleHealthCheck)
	s.mux.Handle("/latency/udm", s.latencySrv)
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
