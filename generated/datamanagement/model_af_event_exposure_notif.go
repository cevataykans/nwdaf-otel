// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// AfEventExposureNotif - Represents notifications on application event(s) that occurred for an Individual Application Event Subscription resource. 
type AfEventExposureNotif struct {

	NotifId string `json:"notifId"`

	EventNotifs []AfEventNotification `json:"eventNotifs"`
}

// AssertAfEventExposureNotifRequired checks if the required fields are not zero-ed
func AssertAfEventExposureNotifRequired(obj AfEventExposureNotif) error {
	elements := map[string]interface{}{
		"notifId": obj.NotifId,
		"eventNotifs": obj.EventNotifs,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.EventNotifs {
		if err := AssertAfEventNotificationRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAfEventExposureNotifConstraints checks if the values respects the defined constraints
func AssertAfEventExposureNotifConstraints(obj AfEventExposureNotif) error {
	for _, el := range obj.EventNotifs {
		if err := AssertAfEventNotificationConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
