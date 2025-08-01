// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// NfInstanceIdCond - Subscription to a given NF Instance Id
type NfInstanceIdCond struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId"`
}

// AssertNfInstanceIdCondRequired checks if the required fields are not zero-ed
func AssertNfInstanceIdCondRequired(obj NfInstanceIdCond) error {
	elements := map[string]interface{}{
		"nfInstanceId": obj.NfInstanceId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertNfInstanceIdCondConstraints checks if the values respects the defined constraints
func AssertNfInstanceIdCondConstraints(obj NfInstanceIdCond) error {
	return nil
}
