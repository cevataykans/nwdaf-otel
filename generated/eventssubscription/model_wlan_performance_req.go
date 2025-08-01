// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package eventssubscription




// WlanPerformanceReq - Represents other WLAN performance analytics requirements.
type WlanPerformanceReq struct {

	SsIds []string `json:"ssIds,omitempty"`

	BssIds []string `json:"bssIds,omitempty"`

	WlanOrderCriter WlanOrderingCriterion `json:"wlanOrderCriter,omitempty"`

	Order MatchingDirection `json:"order,omitempty"`
}

// AssertWlanPerformanceReqRequired checks if the required fields are not zero-ed
func AssertWlanPerformanceReqRequired(obj WlanPerformanceReq) error {
	if err := AssertWlanOrderingCriterionRequired(obj.WlanOrderCriter); err != nil {
		return err
	}
	if err := AssertMatchingDirectionRequired(obj.Order); err != nil {
		return err
	}
	return nil
}

// AssertWlanPerformanceReqConstraints checks if the values respects the defined constraints
func AssertWlanPerformanceReqConstraints(obj WlanPerformanceReq) error {
	if err := AssertWlanOrderingCriterionConstraints(obj.WlanOrderCriter); err != nil {
		return err
	}
	if err := AssertMatchingDirectionConstraints(obj.Order); err != nil {
		return err
	}
	return nil
}
