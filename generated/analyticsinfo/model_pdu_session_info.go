// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// PduSessionInfo - Represents session information.
type PduSessionInfo struct {

	// The identifier of the N4 session for the reported PDU Session.
	N4SessId string `json:"n4SessId,omitempty"`

	// indicating a time in seconds.
	SessInactiveTimer int32 `json:"sessInactiveTimer,omitempty"`

	PduSessStatus PduSessionStatus `json:"pduSessStatus,omitempty"`
}

// AssertPduSessionInfoRequired checks if the required fields are not zero-ed
func AssertPduSessionInfoRequired(obj PduSessionInfo) error {
	if err := AssertPduSessionStatusRequired(obj.PduSessStatus); err != nil {
		return err
	}
	return nil
}

// AssertPduSessionInfoConstraints checks if the values respects the defined constraints
func AssertPduSessionInfoConstraints(obj PduSessionInfo) error {
	if err := AssertPduSessionStatusConstraints(obj.PduSessStatus); err != nil {
		return err
	}
	return nil
}
