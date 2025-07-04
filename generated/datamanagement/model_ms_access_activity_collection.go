// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// MsAccessActivityCollection - Contains Media Streaming access activity collected for an UE Application via AF.
type MsAccessActivityCollection struct {

	MsAccActs []MediaStreamingAccessRecord `json:"msAccActs"`
}

// AssertMsAccessActivityCollectionRequired checks if the required fields are not zero-ed
func AssertMsAccessActivityCollectionRequired(obj MsAccessActivityCollection) error {
	elements := map[string]interface{}{
		"msAccActs": obj.MsAccActs,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.MsAccActs {
		if err := AssertMediaStreamingAccessRecordRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertMsAccessActivityCollectionConstraints checks if the values respects the defined constraints
func AssertMsAccessActivityCollectionConstraints(obj MsAccessActivityCollection) error {
	for _, el := range obj.MsAccActs {
		if err := AssertMediaStreamingAccessRecordConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
