// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package eventssubscription




// FailureEventInfo - Contains information on the event for which the subscription is not successful.
type FailureEventInfo struct {

	Event NwdafEvent `json:"event"`

	FailureCode NwdafFailureCode `json:"failureCode"`
}

// AssertFailureEventInfoRequired checks if the required fields are not zero-ed
func AssertFailureEventInfoRequired(obj FailureEventInfo) error {
	elements := map[string]interface{}{
		"event": obj.Event,
		"failureCode": obj.FailureCode,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertNwdafEventRequired(obj.Event); err != nil {
		return err
	}
	if err := AssertNwdafFailureCodeRequired(obj.FailureCode); err != nil {
		return err
	}
	return nil
}

// AssertFailureEventInfoConstraints checks if the values respects the defined constraints
func AssertFailureEventInfoConstraints(obj FailureEventInfo) error {
	if err := AssertNwdafEventConstraints(obj.Event); err != nil {
		return err
	}
	if err := AssertNwdafFailureCodeConstraints(obj.FailureCode); err != nil {
		return err
	}
	return nil
}
