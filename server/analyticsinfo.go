package nwdaf

import (
	"net/http"
)

type analyticsInfoService struct {
}

// TODO: decide if analytics subscription transfer not supported ?
func (s *analyticsInfoService) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/nnwdaf-analyticsinfo/v1/analytics", s.handleGetAnalytics)
	mux.HandleFunc("/nnwdaf-analyticsinfo/v1/context", s.handleGetContext)
}

func (s *analyticsInfoService) handleGetAnalytics(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// TODO: parse query params, execute logic
	w.WriteHeader(http.StatusNoContent) // if requested analytics is not found
	w.WriteHeader(http.StatusOK)        // if content is found
}

func (s *analyticsInfoService) handleGetContext(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// TODO: parse query params, execute logic
	w.WriteHeader(http.StatusNoContent) // if requested context is not found
	w.WriteHeader(http.StatusOK)        // if context is found
}
