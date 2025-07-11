// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// AfEventExposureSubsc - Represents an Individual Application Event Subscription resource.
type AfEventExposureSubsc struct {

	DataAccProfId string `json:"dataAccProfId,omitempty"`

	EventsSubs []EventsSubs `json:"eventsSubs"`

	EventsRepInfo ReportingInformation `json:"eventsRepInfo"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`

	NotifId string `json:"notifId"`

	EventNotifs []AfEventNotification `json:"eventNotifs,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`
}

// AssertAfEventExposureSubscRequired checks if the required fields are not zero-ed
func AssertAfEventExposureSubscRequired(obj AfEventExposureSubsc) error {
	elements := map[string]interface{}{
		"eventsSubs": obj.EventsSubs,
		"eventsRepInfo": obj.EventsRepInfo,
		"notifUri": obj.NotifUri,
		"notifId": obj.NotifId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.EventsSubs {
		if err := AssertEventsSubsRequired(el); err != nil {
			return err
		}
	}
	if err := AssertReportingInformationRequired(obj.EventsRepInfo); err != nil {
		return err
	}
	for _, el := range obj.EventNotifs {
		if err := AssertAfEventNotificationRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAfEventExposureSubscConstraints checks if the values respects the defined constraints
func AssertAfEventExposureSubscConstraints(obj AfEventExposureSubsc) error {
	for _, el := range obj.EventsSubs {
		if err := AssertEventsSubsConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertReportingInformationConstraints(obj.EventsRepInfo); err != nil {
		return err
	}
	for _, el := range obj.EventNotifs {
		if err := AssertAfEventNotificationConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
