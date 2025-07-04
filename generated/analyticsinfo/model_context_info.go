// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




type ContextInfo struct {

	OrigHeaders []string `json:"origHeaders,omitempty"`

	RequestHeaders []string `json:"requestHeaders,omitempty"`
}

// AssertContextInfoRequired checks if the required fields are not zero-ed
func AssertContextInfoRequired(obj ContextInfo) error {
	return nil
}

// AssertContextInfoConstraints checks if the values respects the defined constraints
func AssertContextInfoConstraints(obj ContextInfo) error {
	return nil
}
