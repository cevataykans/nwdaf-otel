// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// TrafficDescriptor - Represents the Traffic Descriptor
type TrafficDescriptor struct {

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	SNssai Snssai `json:"sNssai,omitempty"`

	DddTrafficDescriptorList []DddTrafficDescriptor `json:"dddTrafficDescriptorList,omitempty"`
}

// AssertTrafficDescriptorRequired checks if the required fields are not zero-ed
func AssertTrafficDescriptorRequired(obj TrafficDescriptor) error {
	if err := AssertSnssaiRequired(obj.SNssai); err != nil {
		return err
	}
	for _, el := range obj.DddTrafficDescriptorList {
		if err := AssertDddTrafficDescriptorRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertTrafficDescriptorConstraints checks if the values respects the defined constraints
func AssertTrafficDescriptorConstraints(obj TrafficDescriptor) error {
	if err := AssertSnssaiConstraints(obj.SNssai); err != nil {
		return err
	}
	for _, el := range obj.DddTrafficDescriptorList {
		if err := AssertDddTrafficDescriptorConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
