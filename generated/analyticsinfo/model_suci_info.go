// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// SuciInfo - SUCI information containing Routing Indicator and Home Network Public Key ID
type SuciInfo struct {

	RoutingInds []string `json:"routingInds,omitempty"`

	HNwPubKeyIds []int32 `json:"hNwPubKeyIds,omitempty"`
}

// AssertSuciInfoRequired checks if the required fields are not zero-ed
func AssertSuciInfoRequired(obj SuciInfo) error {
	return nil
}

// AssertSuciInfoConstraints checks if the values respects the defined constraints
func AssertSuciInfoConstraints(obj SuciInfo) error {
	return nil
}
