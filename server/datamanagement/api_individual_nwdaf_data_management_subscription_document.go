// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement

import (
	"encoding/json"
	"net/http"
	datamanagementAPI "nwdaf-otel/generated/datamanagement"
	"strings"

	"github.com/gorilla/mux"
)

// IndividualNWDAFDataManagementSubscriptionDocumentAPIController binds http requests to an api service and writes the service results to the http response
type IndividualNWDAFDataManagementSubscriptionDocumentAPIController struct {
	service      datamanagementAPI.IndividualNWDAFDataManagementSubscriptionDocumentAPIServicer
	errorHandler datamanagementAPI.ErrorHandler
}

// IndividualNWDAFDataManagementSubscriptionDocumentAPIOption for how the controller is set up.
type IndividualNWDAFDataManagementSubscriptionDocumentAPIOption func(*IndividualNWDAFDataManagementSubscriptionDocumentAPIController)

// WithIndividualNWDAFDataManagementSubscriptionDocumentAPIErrorHandler inject ErrorHandler into controller
func WithIndividualNWDAFDataManagementSubscriptionDocumentAPIErrorHandler(h datamanagementAPI.ErrorHandler) IndividualNWDAFDataManagementSubscriptionDocumentAPIOption {
	return func(c *IndividualNWDAFDataManagementSubscriptionDocumentAPIController) {
		c.errorHandler = h
	}
}

// NewIndividualNWDAFDataManagementSubscriptionDocumentAPIController creates a default api controller
func NewIndividualNWDAFDataManagementSubscriptionDocumentAPIController(s datamanagementAPI.IndividualNWDAFDataManagementSubscriptionDocumentAPIServicer, opts ...IndividualNWDAFDataManagementSubscriptionDocumentAPIOption) *IndividualNWDAFDataManagementSubscriptionDocumentAPIController {
	controller := &IndividualNWDAFDataManagementSubscriptionDocumentAPIController{
		service:      s,
		errorHandler: datamanagementAPI.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the IndividualNWDAFDataManagementSubscriptionDocumentAPIController
func (c *IndividualNWDAFDataManagementSubscriptionDocumentAPIController) Routes() datamanagementAPI.Routes {
	return datamanagementAPI.Routes{
		"UpdateNWDAFDataSubscription": datamanagementAPI.Route{
			"UpdateNWDAFDataSubscription",
			strings.ToUpper("Put"),
			"/nnwdaf-datamanagement/v1/subscriptions/{subscriptionId}",
			c.UpdateNWDAFDataSubscription,
		},
		"DeleteNWDAFDataSubscription": datamanagementAPI.Route{
			"DeleteNWDAFDataSubscription",
			strings.ToUpper("Delete"),
			"/nnwdaf-datamanagement/v1/subscriptions/{subscriptionId}",
			c.DeleteNWDAFDataSubscription,
		},
	}
}

// OrderedRoutes returns all the api routes in a deterministic order for the IndividualNWDAFDataManagementSubscriptionDocumentAPIController
func (c *IndividualNWDAFDataManagementSubscriptionDocumentAPIController) OrderedRoutes() []datamanagementAPI.Route {
	return []datamanagementAPI.Route{
		datamanagementAPI.Route{
			"UpdateNWDAFDataSubscription",
			strings.ToUpper("Put"),
			"/nnwdaf-datamanagement/v1/subscriptions/{subscriptionId}",
			c.UpdateNWDAFDataSubscription,
		},
		datamanagementAPI.Route{
			"DeleteNWDAFDataSubscription",
			strings.ToUpper("Delete"),
			"/nnwdaf-datamanagement/v1/subscriptions/{subscriptionId}",
			c.DeleteNWDAFDataSubscription,
		},
	}
}



// UpdateNWDAFDataSubscription - Update an existing Individual NWDAF Data Subscription.
func (c *IndividualNWDAFDataManagementSubscriptionDocumentAPIController) UpdateNWDAFDataSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &datamanagementAPI.RequiredError{"subscriptionId"}, nil)
		return
	}
	var nnwdafDataManagementSubscParam datamanagementAPI.NnwdafDataManagementSubsc
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&nnwdafDataManagementSubscParam); err != nil {
		c.errorHandler(w, r, &datamanagementAPI.ParsingError{Err: err}, nil)
		return
	}
	if err := datamanagementAPI.AssertNnwdafDataManagementSubscRequired(nnwdafDataManagementSubscParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := datamanagementAPI.AssertNnwdafDataManagementSubscConstraints(nnwdafDataManagementSubscParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.UpdateNWDAFDataSubscription(r.Context(), subscriptionIdParam, &nnwdafDataManagementSubscParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = datamanagementAPI.EncodeJSONResponse(result.Body, &result.Code, w)
}

// DeleteNWDAFDataSubscription - unsubscribe from notifications
func (c *IndividualNWDAFDataManagementSubscriptionDocumentAPIController) DeleteNWDAFDataSubscription(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	subscriptionIdParam := params["subscriptionId"]
	if subscriptionIdParam == "" {
		c.errorHandler(w, r, &datamanagementAPI.RequiredError{"subscriptionId"}, nil)
		return
	}
	result, err := c.service.DeleteNWDAFDataSubscription(r.Context(), subscriptionIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = datamanagementAPI.EncodeJSONResponse(result.Body, &result.Code, w)
}
