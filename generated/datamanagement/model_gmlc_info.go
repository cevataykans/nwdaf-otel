// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// GmlcInfo - Information of a GMLC NF Instance
type GmlcInfo struct {

	ServingClientTypes []ExternalClientType `json:"servingClientTypes,omitempty"`

	GmlcNumbers []string `json:"gmlcNumbers,omitempty"`
}

// AssertGmlcInfoRequired checks if the required fields are not zero-ed
func AssertGmlcInfoRequired(obj GmlcInfo) error {
	for _, el := range obj.ServingClientTypes {
		if err := AssertExternalClientTypeRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGmlcInfoConstraints checks if the values respects the defined constraints
func AssertGmlcInfoConstraints(obj GmlcInfo) error {
	for _, el := range obj.ServingClientTypes {
		if err := AssertExternalClientTypeConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
