// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// IdentityRange - A range of GPSIs (subscriber identities), either based on a numeric range, or based on regular-expression matching 
type IdentityRange struct {

	Start string `json:"start,omitempty" validate:"regexp=^[0-9]+$"`

	End string `json:"end,omitempty" validate:"regexp=^[0-9]+$"`

	Pattern string `json:"pattern,omitempty"`
}

// AssertIdentityRangeRequired checks if the required fields are not zero-ed
func AssertIdentityRangeRequired(obj IdentityRange) error {
	return nil
}

// AssertIdentityRangeConstraints checks if the values respects the defined constraints
func AssertIdentityRangeConstraints(obj IdentityRange) error {
	return nil
}
