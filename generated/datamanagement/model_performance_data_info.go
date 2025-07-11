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
	"time"
)



// PerformanceDataInfo - Contains Performance Data Analytics related information collection.
type PerformanceDataInfo struct {

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	UeIpAddr *IpAddr `json:"ueIpAddr,omitempty"`

	IpTrafficFilter FlowInfo `json:"ipTrafficFilter,omitempty"`

	UserLoc UserLocation `json:"userLoc,omitempty"`

	AppLocs []string `json:"appLocs,omitempty"`

	AsAddr AddrFqdn `json:"asAddr,omitempty"`

	PerfData PerformanceData `json:"perfData"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
}

// AssertPerformanceDataInfoRequired checks if the required fields are not zero-ed
func AssertPerformanceDataInfoRequired(obj PerformanceDataInfo) error {
	elements := map[string]interface{}{
		"perfData": obj.PerfData,
		"timeStamp": obj.TimeStamp,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.UeIpAddr != nil {
		if err := AssertIpAddrRequired(*obj.UeIpAddr); err != nil {
			return err
		}
	}
	if err := AssertFlowInfoRequired(obj.IpTrafficFilter); err != nil {
		return err
	}
	if err := AssertUserLocationRequired(obj.UserLoc); err != nil {
		return err
	}
	if err := AssertAddrFqdnRequired(obj.AsAddr); err != nil {
		return err
	}
	if err := AssertPerformanceDataRequired(obj.PerfData); err != nil {
		return err
	}
	return nil
}

// AssertPerformanceDataInfoConstraints checks if the values respects the defined constraints
func AssertPerformanceDataInfoConstraints(obj PerformanceDataInfo) error {
    if obj.UeIpAddr != nil {
     	if err := AssertIpAddrConstraints(*obj.UeIpAddr); err != nil {
     		return err
     	}
    }
	if err := AssertFlowInfoConstraints(obj.IpTrafficFilter); err != nil {
		return err
	}
	if err := AssertUserLocationConstraints(obj.UserLoc); err != nil {
		return err
	}
	if err := AssertAddrFqdnConstraints(obj.AsAddr); err != nil {
		return err
	}
	if err := AssertPerformanceDataConstraints(obj.PerfData); err != nil {
		return err
	}
	return nil
}
