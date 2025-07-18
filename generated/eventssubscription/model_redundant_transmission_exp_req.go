// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package eventssubscription




// RedundantTransmissionExpReq - Represents other redundant transmission experience analytics requirements.
type RedundantTransmissionExpReq struct {

	RedTOrderCriter RedTransExpOrderingCriterion `json:"redTOrderCriter,omitempty"`

	Order MatchingDirection `json:"order,omitempty"`
}

// AssertRedundantTransmissionExpReqRequired checks if the required fields are not zero-ed
func AssertRedundantTransmissionExpReqRequired(obj RedundantTransmissionExpReq) error {
	if err := AssertRedTransExpOrderingCriterionRequired(obj.RedTOrderCriter); err != nil {
		return err
	}
	if err := AssertMatchingDirectionRequired(obj.Order); err != nil {
		return err
	}
	return nil
}

// AssertRedundantTransmissionExpReqConstraints checks if the values respects the defined constraints
func AssertRedundantTransmissionExpReqConstraints(obj RedundantTransmissionExpReq) error {
	if err := AssertRedTransExpOrderingCriterionConstraints(obj.RedTOrderCriter); err != nil {
		return err
	}
	if err := AssertMatchingDirectionConstraints(obj.Order); err != nil {
		return err
	}
	return nil
}
