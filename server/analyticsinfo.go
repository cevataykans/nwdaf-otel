package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	analyticsinfoAPI "nwdaf-otel/generated/analyticsinfo"
	"nwdaf-otel/server/analyticsinfo"
)

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
	log.Println("Server setup complete")
}

func (s *analyticsInfoServer) Start() chan error {
	errChan := make(chan error)
	go func() {
		// TODO: get PORT from config
		// TODO: configure TLS, use mkcert certificates
		srv := http.Server{
			Addr:    ":8080",
			Handler: s.mux,
		}
		errChan <- srv.ListenAndServe()
		log.Println("Server stopped")
	}()
	return errChan
}
