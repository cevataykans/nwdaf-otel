package nwdaf

import (
	"fmt"
	"net/http"
)

const SubscriptionId string = "subscriptionId"

type eventSubscriptionService struct {
}

// TODO: decide if analytics subscription transfer not supported ?
func (s *eventSubscriptionService) Setup(mux *http.ServeMux) {
	mux.HandleFunc(fmt.Sprintf("/nnwdaf-eventssubscription/v1/subscriptions"), s.handleCreate)
	mux.HandleFunc(fmt.Sprintf("/nnwdaf-eventssubscription/v1/subscriptions/{%s}", SubscriptionId), s.handleUpdate)
}

func (s *eventSubscriptionService) handleCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// TODO: create subcription, write header location
	w.WriteHeader(http.StatusCreated)
}

func (s *eventSubscriptionService) handleUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		s.handlePut(w, r)
		return
	}
	if r.Method == "DELETE" {
		s.handleDelete(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func (s *eventSubscriptionService) handlePut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Subscription ID: ", r.PathValue(SubscriptionId))
	// TODO: check subscription, update it
	w.WriteHeader(http.StatusNoContent)
}

func (s *eventSubscriptionService) handleDelete(w http.ResponseWriter, r *http.Request) {
	// TODO: check & delete subscription
	fmt.Println("Subscription ID: ", r.PathValue(SubscriptionId))
	// TODO: check subscription, update it
	w.WriteHeader(http.StatusNoContent)
}
