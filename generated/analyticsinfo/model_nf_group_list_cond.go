// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// NfGroupListCond - Subscription to a set of NFs based on their Group Ids
type NfGroupListCond struct {

	ConditionType string `json:"conditionType"`

	NfType string `json:"nfType"`

	NfGroupIdList []string `json:"nfGroupIdList"`
}

// AssertNfGroupListCondRequired checks if the required fields are not zero-ed
func AssertNfGroupListCondRequired(obj NfGroupListCond) error {
	elements := map[string]interface{}{
		"conditionType": obj.ConditionType,
		"nfType": obj.NfType,
		"nfGroupIdList": obj.NfGroupIdList,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertNfGroupListCondConstraints checks if the values respects the defined constraints
func AssertNfGroupListCondConstraints(obj NfGroupListCond) error {
	return nil
}
