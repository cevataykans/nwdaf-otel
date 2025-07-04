// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// PlmnRange - Range of PLMN IDs
type PlmnRange struct {

	Start string `json:"start,omitempty" validate:"regexp=^[0-9]{3}[0-9]{2,3}$"`

	End string `json:"end,omitempty" validate:"regexp=^[0-9]{3}[0-9]{2,3}$"`

	Pattern string `json:"pattern,omitempty"`
}

// AssertPlmnRangeRequired checks if the required fields are not zero-ed
func AssertPlmnRangeRequired(obj PlmnRange) error {
	return nil
}

// AssertPlmnRangeConstraints checks if the values respects the defined constraints
func AssertPlmnRangeConstraints(obj PlmnRange) error {
	return nil
}
