package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	analyticsinfoAPI "nwdaf-otel/generated/analyticsinfo"
	"nwdaf-otel/server/analyticsinfo"
)

type Server interface {
	Setup()
	Start() chan error
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
	log.Println("AnalyticsInfo Server setup complete")
}

func (s *analyticsInfoServer) Start() chan error {
	errChan := make(chan error)
	go func() {
		srv := http.Server{
			Addr:    ":8080",
			Handler: s.mux,
		}
		log.Println("Listening and serving HTTP on " + srv.Addr)
		err := srv.ListenAndServe()
		log.Println("Server stopped")
		errChan <- err
	}()
	return errChan
}
