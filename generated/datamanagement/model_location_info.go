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
	"errors"
)



// LocationInfo - Represents UE location information.
type LocationInfo struct {

	Loc UserLocation `json:"loc"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	Ratio int32 `json:"ratio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}

// AssertLocationInfoRequired checks if the required fields are not zero-ed
func AssertLocationInfoRequired(obj LocationInfo) error {
	elements := map[string]interface{}{
		"loc": obj.Loc,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertUserLocationRequired(obj.Loc); err != nil {
		return err
	}
	return nil
}

// AssertLocationInfoConstraints checks if the values respects the defined constraints
func AssertLocationInfoConstraints(obj LocationInfo) error {
	if err := AssertUserLocationConstraints(obj.Loc); err != nil {
		return err
	}
	if obj.Ratio < 1 {
		return &ParsingError{Param: "Ratio", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.Ratio > 100 {
		return &ParsingError{Param: "Ratio", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if obj.Confidence < 0 {
		return &ParsingError{Param: "Confidence", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
