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



// PerfData - Represents DN performance data.
type PerfData struct {

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	AvgTrafficRate string `json:"avgTrafficRate,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxTrafficRate string `json:"maxTrafficRate,omitempty" validate:"regexp=^\\\\d+(\\\\.\\\\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	AvePacketDelay int32 `json:"avePacketDelay,omitempty"`

	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	MaxPacketDelay int32 `json:"maxPacketDelay,omitempty"`

	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	AvgPacketLossRate int32 `json:"avgPacketLossRate,omitempty"`
}

// AssertPerfDataRequired checks if the required fields are not zero-ed
func AssertPerfDataRequired(obj PerfData) error {
	return nil
}

// AssertPerfDataConstraints checks if the values respects the defined constraints
func AssertPerfDataConstraints(obj PerfData) error {
	if obj.AvePacketDelay < 1 {
		return &ParsingError{Param: "AvePacketDelay", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MaxPacketDelay < 1 {
		return &ParsingError{Param: "MaxPacketDelay", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AvgPacketLossRate < 0 {
		return &ParsingError{Param: "AvgPacketLossRate", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AvgPacketLossRate > 1000 {
		return &ParsingError{Param: "AvgPacketLossRate", Err: errors.New(errMsgMaxValueConstraint)}
	}
	return nil
}
