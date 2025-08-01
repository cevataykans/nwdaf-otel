// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo


import (
	"errors"
)



// GnbId - Provides the G-NB identifier.
type GnbId struct {

	// Unsigned integer representing the bit length of the gNB ID as defined in clause 9.3.1.6 of 3GPP TS 38.413 [11], within the range 22 to 32. 
	BitLength int32 `json:"bitLength"`

	// This represents the identifier of the gNB. The value of the gNB ID shall be encoded in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\", \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The padding 0 shall be added to make multiple nibbles,  the most significant character representing the padding 0 if required together with the 4 most significant bits of the gNB ID shall appear first in the string, and the character representing the 4 least significant bit of the gNB ID shall appear last in the string. 
	GNBValue string `json:"gNBValue" validate:"regexp=^[A-Fa-f0-9]{6,8}$"`
}

// AssertGnbIdRequired checks if the required fields are not zero-ed
func AssertGnbIdRequired(obj GnbId) error {
	elements := map[string]interface{}{
		"bitLength": obj.BitLength,
		"gNBValue": obj.GNBValue,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertGnbIdConstraints checks if the values respects the defined constraints
func AssertGnbIdConstraints(obj GnbId) error {
	if obj.BitLength < 22 {
		return &ParsingError{Param: "BitLength", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.BitLength > 32 {
		return &ParsingError{Param: "BitLength", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
