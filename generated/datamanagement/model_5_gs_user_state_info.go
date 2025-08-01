// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// Model5GsUserStateInfo - Represents the 5GS User state of the UE for an access type
type Model5GsUserStateInfo struct {

	Var5gsUserState Model5GsUserState `json:"5gsUserState"`

	AccessType AccessType `json:"accessType"`
}

// AssertModel5GsUserStateInfoRequired checks if the required fields are not zero-ed
func AssertModel5GsUserStateInfoRequired(obj Model5GsUserStateInfo) error {
	elements := map[string]interface{}{
		"5gsUserState": obj.Var5gsUserState,
		"accessType": obj.AccessType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertModel5GsUserStateRequired(obj.Var5gsUserState); err != nil {
		return err
	}
	return nil
}

// AssertModel5GsUserStateInfoConstraints checks if the values respects the defined constraints
func AssertModel5GsUserStateInfoConstraints(obj Model5GsUserStateInfo) error {
	if err := AssertModel5GsUserStateConstraints(obj.Var5gsUserState); err != nil {
		return err
	}
	return nil
}
