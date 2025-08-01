// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// NcgiTai - List of NR cell ids, with their pertaining TAIs
type NcgiTai struct {

	Tai Tai `json:"tai"`

	// List of List of NR cell ids
	CellList []Ncgi `json:"cellList"`
}

// AssertNcgiTaiRequired checks if the required fields are not zero-ed
func AssertNcgiTaiRequired(obj NcgiTai) error {
	elements := map[string]interface{}{
		"tai": obj.Tai,
		"cellList": obj.CellList,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertTaiRequired(obj.Tai); err != nil {
		return err
	}
	for _, el := range obj.CellList {
		if err := AssertNcgiRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertNcgiTaiConstraints checks if the values respects the defined constraints
func AssertNcgiTaiConstraints(obj NcgiTai) error {
	if err := AssertTaiConstraints(obj.Tai); err != nil {
		return err
	}
	for _, el := range obj.CellList {
		if err := AssertNcgiConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
