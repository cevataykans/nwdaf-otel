// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package eventssubscription

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// IndividualNWDAFEventSubscriptionTransferDocumentAPIController binds http requests to an api service and writes the service results to the http response
type IndividualNWDAFEventSubscriptionTransferDocumentAPIController struct {
	service IndividualNWDAFEventSubscriptionTransferDocumentAPIServicer
	errorHandler ErrorHandler
}

// IndividualNWDAFEventSubscriptionTransferDocumentAPIOption for how the controller is set up.
type IndividualNWDAFEventSubscriptionTransferDocumentAPIOption func(*IndividualNWDAFEventSubscriptionTransferDocumentAPIController)

// WithIndividualNWDAFEventSubscriptionTransferDocumentAPIErrorHandler inject ErrorHandler into controller
func WithIndividualNWDAFEventSubscriptionTransferDocumentAPIErrorHandler(h ErrorHandler) IndividualNWDAFEventSubscriptionTransferDocumentAPIOption {
	return func(c *IndividualNWDAFEventSubscriptionTransferDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualNWDAFEventSubscriptionTransferDocumentAPIController creates a default api controller
func NewIndividualNWDAFEventSubscriptionTransferDocumentAPIController(s IndividualNWDAFEventSubscriptionTransferDocumentAPIServicer, opts ...IndividualNWDAFEventSubscriptionTransferDocumentAPIOption) *IndividualNWDAFEventSubscriptionTransferDocumentAPIController {
	controller := &IndividualNWDAFEventSubscriptionTransferDocumentAPIController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualNWDAFEventSubscriptionTransferDocumentAPIController
func (c *IndividualNWDAFEventSubscriptionTransferDocumentAPIController) Routes() Routes {
	return Routes{
		"UpdateNWDAFEventSubscriptionTransfer": Route{
			"UpdateNWDAFEventSubscriptionTransfer",
			strings.ToUpper("Put"),
			"/nnwdaf-eventssubscription/v1/transfers/{transferId}",
			c.UpdateNWDAFEventSubscriptionTransfer,
		},
		"DeleteNWDAFEventSubscriptionTransfer": Route{
			"DeleteNWDAFEventSubscriptionTransfer",
			strings.ToUpper("Delete"),
			"/nnwdaf-eventssubscription/v1/transfers/{transferId}",
			c.DeleteNWDAFEventSubscriptionTransfer,
		},
	}
}

// OrderedRoutes returns all the api routes in a deterministic order for the IndividualNWDAFEventSubscriptionTransferDocumentAPIController
func (c *IndividualNWDAFEventSubscriptionTransferDocumentAPIController) OrderedRoutes() []Route {
	return []Route{
		Route{
			"UpdateNWDAFEventSubscriptionTransfer",
			strings.ToUpper("Put"),
			"/nnwdaf-eventssubscription/v1/transfers/{transferId}",
			c.UpdateNWDAFEventSubscriptionTransfer,
		},
		Route{
			"DeleteNWDAFEventSubscriptionTransfer",
			strings.ToUpper("Delete"),
			"/nnwdaf-eventssubscription/v1/transfers/{transferId}",
			c.DeleteNWDAFEventSubscriptionTransfer,
		},
	}
}



// UpdateNWDAFEventSubscriptionTransfer - Update an existing Individual NWDAF Event Subscription Transfer
func (c *IndividualNWDAFEventSubscriptionTransferDocumentAPIController) UpdateNWDAFEventSubscriptionTransfer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transferIdParam := params["transferId"]
	if transferIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"transferId"}, nil)
		return
	}
	var analyticsSubscriptionsTransferParam AnalyticsSubscriptionsTransfer
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&analyticsSubscriptionsTransferParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertAnalyticsSubscriptionsTransferRequired(analyticsSubscriptionsTransferParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := AssertAnalyticsSubscriptionsTransferConstraints(analyticsSubscriptionsTransferParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateNWDAFEventSubscriptionTransfer(r.Context(), transferIdParam, analyticsSubscriptionsTransferParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteNWDAFEventSubscriptionTransfer - Delete an existing Individual NWDAF Event Subscription Transfer
func (c *IndividualNWDAFEventSubscriptionTransferDocumentAPIController) DeleteNWDAFEventSubscriptionTransfer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	transferIdParam := params["transferId"]
	if transferIdParam == "" {
		c.errorHandler(w, r, &RequiredError{"transferId"}, nil)
		return
	}
	result, err := c.service.DeleteNWDAFEventSubscriptionTransfer(r.Context(), transferIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = EncodeJSONResponse(result.Body, &result.Code, w)
}
