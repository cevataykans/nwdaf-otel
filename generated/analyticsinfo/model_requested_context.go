// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// RequestedContext - Contains types of analytics context information.
type RequestedContext struct {

	// List of analytics context types.
	Contexts []ContextType `json:"contexts"`
}

// AssertRequestedContextRequired checks if the required fields are not zero-ed
func AssertRequestedContextRequired(obj RequestedContext) error {
	elements := map[string]interface{}{
		"contexts": obj.Contexts,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Contexts {
		if err := AssertContextTypeRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRequestedContextConstraints checks if the values respects the defined constraints
func AssertRequestedContextConstraints(obj RequestedContext) error {
	for _, el := range obj.Contexts {
		if err := AssertContextTypeConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
