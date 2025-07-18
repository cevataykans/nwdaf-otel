// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// ContextIdList - Contains a list of context identifiers of context information of analytics subscriptions. 
type ContextIdList struct {

	ContextIds []AnalyticsContextIdentifier `json:"contextIds"`
}

// AssertContextIdListRequired checks if the required fields are not zero-ed
func AssertContextIdListRequired(obj ContextIdList) error {
	elements := map[string]interface{}{
		"contextIds": obj.ContextIds,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.ContextIds {
		if err := AssertAnalyticsContextIdentifierRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertContextIdListConstraints checks if the values respects the defined constraints
func AssertContextIdListConstraints(obj ContextIdList) error {
	for _, el := range obj.ContextIds {
		if err := AssertAnalyticsContextIdentifierConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
