// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// EthFlowDescription - Identifies an Ethernet flow.
type EthFlowDescription struct {

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DestMacAddr string `json:"destMacAddr,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`

	EthType string `json:"ethType"`

	// Defines a packet filter of an IP flow.
	FDesc string `json:"fDesc,omitempty"`

	FDir FlowDirection `json:"fDir,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	SourceMacAddr string `json:"sourceMacAddr,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`

	VlanTags []string `json:"vlanTags,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	SrcMacAddrEnd string `json:"srcMacAddrEnd,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DestMacAddrEnd string `json:"destMacAddrEnd,omitempty" validate:"regexp=^([0-9a-fA-F]{2})((-[0-9a-fA-F]{2}){5})$"`
}

// AssertEthFlowDescriptionRequired checks if the required fields are not zero-ed
func AssertEthFlowDescriptionRequired(obj EthFlowDescription) error {
	elements := map[string]interface{}{
		"ethType": obj.EthType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertFlowDirectionRequired(obj.FDir); err != nil {
		return err
	}
	return nil
}

// AssertEthFlowDescriptionConstraints checks if the values respects the defined constraints
func AssertEthFlowDescriptionConstraints(obj EthFlowDescription) error {
	if err := AssertFlowDirectionConstraints(obj.FDir); err != nil {
		return err
	}
	return nil
}
