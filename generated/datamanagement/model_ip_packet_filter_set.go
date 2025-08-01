// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




type IpPacketFilterSet struct {

	SrcIp string `json:"srcIp,omitempty"`

	DstIp string `json:"dstIp,omitempty"`

	Protocol int32 `json:"protocol,omitempty"`

	SrcPort int32 `json:"srcPort,omitempty"`

	DstPort int32 `json:"dstPort,omitempty"`

	ToSTc string `json:"toSTc,omitempty"`

	FlowLabel int32 `json:"flowLabel,omitempty"`

	Spi int32 `json:"spi,omitempty"`

	Direction string `json:"direction"`
}

// AssertIpPacketFilterSetRequired checks if the required fields are not zero-ed
func AssertIpPacketFilterSetRequired(obj IpPacketFilterSet) error {
	elements := map[string]interface{}{
		"direction": obj.Direction,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertIpPacketFilterSetConstraints checks if the values respects the defined constraints
func AssertIpPacketFilterSetConstraints(obj IpPacketFilterSet) error {
	return nil
}
