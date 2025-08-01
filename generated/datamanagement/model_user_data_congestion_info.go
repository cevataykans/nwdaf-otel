// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// UserDataCongestionInfo - Represents the user data congestion information.
type UserDataCongestionInfo struct {

	NetworkArea NetworkAreaInfo `json:"networkArea"`

	CongestionInfo CongestionInfo `json:"congestionInfo"`

	Snssai Snssai `json:"snssai,omitempty"`
}

// AssertUserDataCongestionInfoRequired checks if the required fields are not zero-ed
func AssertUserDataCongestionInfoRequired(obj UserDataCongestionInfo) error {
	elements := map[string]interface{}{
		"networkArea": obj.NetworkArea,
		"congestionInfo": obj.CongestionInfo,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertNetworkAreaInfoRequired(obj.NetworkArea); err != nil {
		return err
	}
	if err := AssertCongestionInfoRequired(obj.CongestionInfo); err != nil {
		return err
	}
	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	return nil
}

// AssertUserDataCongestionInfoConstraints checks if the values respects the defined constraints
func AssertUserDataCongestionInfoConstraints(obj UserDataCongestionInfo) error {
	if err := AssertNetworkAreaInfoConstraints(obj.NetworkArea); err != nil {
		return err
	}
	if err := AssertCongestionInfoConstraints(obj.CongestionInfo); err != nil {
		return err
	}
	if err := AssertSnssaiConstraints(obj.Snssai); err != nil {
		return err
	}
	return nil
}
