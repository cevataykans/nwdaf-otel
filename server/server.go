package nwdaf

import "net/http"

type server struct {
	mux *http.ServeMux

	analyticsService    *analyticsInfoService
	subscriptionService *eventSubscriptionService
}

func (s *server) SetupServices() {
	s.analyticsService.Setup(s.mux)
	s.subscriptionService.Setup(s.mux)
}

func (s *server) Start() chan error {
	errChan := make(chan error)
	go func() {
		// TODO: get PORT from config
		// TODO: configure TLS, use mkcert certificates
		errChan <- http.ListenAndServe(":8080", s.mux)
		close(errChan)
	}()
	return errChan
}
