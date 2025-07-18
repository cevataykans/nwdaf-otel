// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// SnssaiEasdfInfoItem - Set of parameters supported by EASDF for a given S-NSSAI
type SnssaiEasdfInfoItem struct {

	SNssai ExtSnssai `json:"sNssai"`

	DnnEasdfInfoList []DnnEasdfInfoItem `json:"dnnEasdfInfoList"`
}

// AssertSnssaiEasdfInfoItemRequired checks if the required fields are not zero-ed
func AssertSnssaiEasdfInfoItemRequired(obj SnssaiEasdfInfoItem) error {
	elements := map[string]interface{}{
		"sNssai": obj.SNssai,
		"dnnEasdfInfoList": obj.DnnEasdfInfoList,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertExtSnssaiRequired(obj.SNssai); err != nil {
		return err
	}
	for _, el := range obj.DnnEasdfInfoList {
		if err := AssertDnnEasdfInfoItemRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertSnssaiEasdfInfoItemConstraints checks if the values respects the defined constraints
func AssertSnssaiEasdfInfoItemConstraints(obj SnssaiEasdfInfoItem) error {
	if err := AssertExtSnssaiConstraints(obj.SNssai); err != nil {
		return err
	}
	for _, el := range obj.DnnEasdfInfoList {
		if err := AssertDnnEasdfInfoItemConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
