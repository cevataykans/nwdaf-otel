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



type M5QoSSpecification struct {

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MarBwDlBitRate string `json:"marBwDlBitRate" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MarBwUlBitRate string `json:"marBwUlBitRate" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MinDesBwDlBitRate string `json:"minDesBwDlBitRate,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MinDesBwUlBitRate string `json:"minDesBwUlBitRate,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MirBwDlBitRate string `json:"mirBwDlBitRate" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MirBwUlBitRate string `json:"mirBwUlBitRate" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	DesLatency int32 `json:"desLatency,omitempty"`

	DesLoss int32 `json:"desLoss,omitempty"`
}

// AssertM5QoSSpecificationRequired checks if the required fields are not zero-ed
func AssertM5QoSSpecificationRequired(obj M5QoSSpecification) error {
	elements := map[string]interface{}{
		"marBwDlBitRate": obj.MarBwDlBitRate,
		"marBwUlBitRate": obj.MarBwUlBitRate,
		"mirBwDlBitRate": obj.MirBwDlBitRate,
		"mirBwUlBitRate": obj.MirBwUlBitRate,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertM5QoSSpecificationConstraints checks if the values respects the defined constraints
func AssertM5QoSSpecificationConstraints(obj M5QoSSpecification) error {
	if obj.DesLatency < 0 {
		return &ParsingError{Param: "DesLatency", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.DesLoss < 0 {
		return &ParsingError{Param: "DesLoss", Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}
