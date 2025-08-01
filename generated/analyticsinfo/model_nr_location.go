// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo


import (
	"time"
	"errors"
)



// NrLocation - Contains the NR user location.
type NrLocation struct {

	Tai Tai `json:"tai"`

	Ncgi Ncgi `json:"ncgi"`

	IgnoreNcgi bool `json:"ignoreNcgi,omitempty"`

	// The value represents the elapsed time in minutes since the last network contact of the mobile station. Value \"0\" indicates that the location information was obtained after a successful paging procedure for Active Location Retrieval when the UE is in idle mode or after a successful  NG-RAN location reporting procedure with the eNB when the UE is in connected mode. Any other value than \"0\" indicates that the location information is the last known one. See 3GPP TS 29.002 clause 17.7.8. 
	AgeOfLocationInformation int32 `json:"ageOfLocationInformation,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	UeLocationTimestamp time.Time `json:"ueLocationTimestamp,omitempty"`

	// Refer to geographical Information. See 3GPP TS 23.032 clause 7.3.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used. 
	GeographicalInformation string `json:"geographicalInformation,omitempty" validate:"regexp=^[0-9A-F]{16}$"`

	// Refers to Calling Geodetic Location. See ITU-T Recommendation Q.763 (1999) [24] clause 3.88.2. Only the description of an ellipsoid point with uncertainty circle is allowed to be used. 
	GeodeticInformation string `json:"geodeticInformation,omitempty" validate:"regexp=^[0-9A-F]{20}$"`

	GlobalGnbId *GlobalRanNodeId `json:"globalGnbId,omitempty"`
}

// AssertNrLocationRequired checks if the required fields are not zero-ed
func AssertNrLocationRequired(obj NrLocation) error {
	elements := map[string]interface{}{
		"tai": obj.Tai,
		"ncgi": obj.Ncgi,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertTaiRequired(obj.Tai); err != nil {
		return err
	}
	if err := AssertNcgiRequired(obj.Ncgi); err != nil {
		return err
	}
	if obj.GlobalGnbId != nil {
		if err := AssertGlobalRanNodeIdRequired(*obj.GlobalGnbId); err != nil {
			return err
		}
	}
	return nil
}

// AssertNrLocationConstraints checks if the values respects the defined constraints
func AssertNrLocationConstraints(obj NrLocation) error {
	if err := AssertTaiConstraints(obj.Tai); err != nil {
		return err
	}
	if err := AssertNcgiConstraints(obj.Ncgi); err != nil {
		return err
	}
	if obj.AgeOfLocationInformation < 0 {
		return &ParsingError{Param: "AgeOfLocationInformation", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AgeOfLocationInformation > 32767 {
		return &ParsingError{Param: "AgeOfLocationInformation", Err: errors.New(errMsgMaxValueConstraint)}
	}
    if obj.GlobalGnbId != nil {
     	if err := AssertGlobalRanNodeIdConstraints(*obj.GlobalGnbId); err != nil {
     		return err
     	}
    }
	return nil
}
