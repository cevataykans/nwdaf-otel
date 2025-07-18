// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// NfTypeCond - Subscription to a set of NFs based on their NF Type
type NfTypeCond struct {

	NfType NfType `json:"nfType"`
}

// AssertNfTypeCondRequired checks if the required fields are not zero-ed
func AssertNfTypeCondRequired(obj NfTypeCond) error {
	elements := map[string]interface{}{
		"nfType": obj.NfType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertNfTypeRequired(obj.NfType); err != nil {
		return err
	}
	return nil
}

// AssertNfTypeCondConstraints checks if the values respects the defined constraints
func AssertNfTypeCondConstraints(obj NfTypeCond) error {
	if err := AssertNfTypeConstraints(obj.NfType); err != nil {
		return err
	}
	return nil
}
