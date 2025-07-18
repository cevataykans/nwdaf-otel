// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// SacEventState - Represents the state of a subscribed event
type SacEventState struct {

	Active bool `json:"active"`

	RemainReports int32 `json:"remainReports,omitempty"`

	// indicating a time in seconds.
	RemainDuration int32 `json:"remainDuration,omitempty"`
}

// AssertSacEventStateRequired checks if the required fields are not zero-ed
func AssertSacEventStateRequired(obj SacEventState) error {
	elements := map[string]interface{}{
		"active": obj.Active,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSacEventStateConstraints checks if the values respects the defined constraints
func AssertSacEventStateConstraints(obj SacEventState) error {
	return nil
}
