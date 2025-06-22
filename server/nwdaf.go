package nwdaf

import "net/http"

type NWDAF interface {
	SetupServices()
	Start() chan error
}

func New(cfg Config) NWDAF {
	return &server{
		mux:                 http.NewServeMux(),
		analyticsService:    &analyticsInfoService{},
		subscriptionService: &eventSubscriptionService{},
	}
}
