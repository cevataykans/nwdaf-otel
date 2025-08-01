// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement


import (
	"time"
	"errors"
)



// NsmfEventExposure - Represents an Individual SMF Notification Subscription resource. The serviveName property corresponds to the serviceName in the main body of the specification. 
type NsmfEventExposure struct {

	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi,omitempty" validate:"regexp=^(imsi-[0-9]{5,15}|nai-.+|gci-.+|gli-.+|.+)$"`

	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi string `json:"gpsi,omitempty" validate:"regexp=^(msisdn-[0-9]{5,15}|extid-[^@]+@[^@]+|.+)$"`

	// Any UE indication. This IE shall be present if the event subscription is applicable to  any UE. Default value \"false\" is used, if not present. 
	AnyUeInd bool `json:"anyUeInd,omitempty"`

	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	GroupId string `json:"groupId,omitempty" validate:"regexp=^[A-Fa-f0-9]{8}-[0-9]{3}-[0-9]{2,3}-([A-Fa-f0-9][A-Fa-f0-9]){1,10}$"`

	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSeId int32 `json:"pduSeId,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`

	// Identifies an Individual SMF Notification Subscription. To enable that the value is used as part of a URI, the string shall only contain characters allowed according to the \"lower-with-hyphen\" naming convention defined in 3GPP TS 29.501. In an OpenAPI schema, the format shall be designated as \"SubId\". 
	SubId string `json:"subId,omitempty"`

	// Notification Correlation ID assigned by the NF service consumer.
	NotifId string `json:"notifId"`

	// String providing an URI formatted according to RFC 3986.
	NotifUri string `json:"notifUri"`

	// Alternate or backup IPv4 address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty"`

	// Alternate or backup IPv6 address(es) where to send Notifications.
	AltNotifIpv6Addrs []Ipv6Addr `json:"altNotifIpv6Addrs,omitempty"`

	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty"`

	// Subscribed events
	EventSubs []EventSubscription1 `json:"eventSubs"`

	EventNotifs []EventNotification1 `json:"eventNotifs,omitempty"`

	ImmeRep bool `json:"ImmeRep,omitempty"`

	NotifMethod NotificationMethod1 `json:"notifMethod,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReportNbr int32 `json:"maxReportNbr,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	// indicating a time in seconds.
	RepPeriod int32 `json:"repPeriod,omitempty"`

	Guami Guami `json:"guami,omitempty"`

	ServiveName ServiceName `json:"serviveName,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures string `json:"supportedFeatures,omitempty" validate:"regexp=^[A-Fa-f0-9]*$"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SampRatio int32 `json:"sampRatio,omitempty"`

	// Criteria for partitioning the UEs before applying the sampling ratio.
	PartitionCriteria []PartitioningCriteria `json:"partitionCriteria,omitempty"`

	// indicating a time in seconds.
	GrpRepTime int32 `json:"grpRepTime,omitempty"`

	NotifFlag NotificationFlag `json:"notifFlag,omitempty"`
}

// AssertNsmfEventExposureRequired checks if the required fields are not zero-ed
func AssertNsmfEventExposureRequired(obj NsmfEventExposure) error {
	elements := map[string]interface{}{
		"notifId": obj.NotifId,
		"notifUri": obj.NotifUri,
		"eventSubs": obj.EventSubs,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSnssaiRequired(obj.Snssai); err != nil {
		return err
	}
	for _, el := range obj.AltNotifIpv6Addrs {
		if err := AssertIpv6AddrRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EventSubs {
		if err := AssertEventSubscription1Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EventNotifs {
		if err := AssertEventNotification1Required(el); err != nil {
			return err
		}
	}
	if err := AssertNotificationMethod1Required(obj.NotifMethod); err != nil {
		return err
	}
	if err := AssertGuamiRequired(obj.Guami); err != nil {
		return err
	}
	if err := AssertServiceNameRequired(obj.ServiveName); err != nil {
		return err
	}
	for _, el := range obj.PartitionCriteria {
		if err := AssertPartitioningCriteriaRequired(el); err != nil {
			return err
		}
	}
	if err := AssertNotificationFlagRequired(obj.NotifFlag); err != nil {
		return err
	}
	return nil
}

// AssertNsmfEventExposureConstraints checks if the values respects the defined constraints
func AssertNsmfEventExposureConstraints(obj NsmfEventExposure) error {
	if obj.PduSeId < 0 {
		return &ParsingError{Param: "PduSeId", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.PduSeId > 255 {
		return &ParsingError{Param: "PduSeId", Err: errors.New(errMsgMaxValueConstraint)}
	}
	if err := AssertSnssaiConstraints(obj.Snssai); err != nil {
		return err
	}
	for _, el := range obj.AltNotifIpv6Addrs {
		if err := AssertIpv6AddrConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EventSubs {
		if err := AssertEventSubscription1Constraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EventNotifs {
		if err := AssertEventNotification1Constraints(el); err != nil {
			return err
		}
	}
	if err := AssertNotificationMethod1Constraints(obj.NotifMethod); err != nil {
		return err
	}
	if obj.MaxReportNbr < 0 {
		return &ParsingError{Param: "MaxReportNbr", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertGuamiConstraints(obj.Guami); err != nil {
		return err
	}
	if err := AssertServiceNameConstraints(obj.ServiveName); err != nil {
		return err
	}
	if obj.SampRatio < 1 {
		return &ParsingError{Param: "SampRatio", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.SampRatio > 100 {
		return &ParsingError{Param: "SampRatio", Err: errors.New(errMsgMaxValueConstraint)}
	}
	for _, el := range obj.PartitionCriteria {
		if err := AssertPartitioningCriteriaConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertNotificationFlagConstraints(obj.NotifFlag); err != nil {
		return err
	}
	return nil
}
