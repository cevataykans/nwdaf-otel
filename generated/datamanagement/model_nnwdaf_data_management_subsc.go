// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// NnwdafDataManagementSubsc - Represents an Individual NWDAF Data Management Subscription resource.
type NnwdafDataManagementSubsc struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AdrfId string `json:"adrfId,omitempty"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	AdrfSetId string `json:"adrfSetId,omitempty"`

	AnaSub NnwdafEventsSubscription `json:"anaSub,omitempty"`

	// The purposes of data collection. This attribute may only be provided if user consent is reqiured depending on local policy and regulations and the consumer has not checked user consent. 
	DataCollectPurposes []DataCollectionPurpose `json:"dataCollectPurposes,omitempty"`

	DataSub *DataSubscription `json:"dataSub,omitempty"`

	FormatInstruct FormattingInstruction `json:"formatInstruct,omitempty"`

	NotifCorrId string `json:"notifCorrId"`

	// String providing an URI formatted according to RFC 3986.
	NotificURI string `json:"notificURI"`

	ProcInstruct ProcessingInstruction `json:"procInstruct,omitempty"`

	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	TargetNfId string `json:"targetNfId,omitempty"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	TargetNfSetId string `json:"targetNfSetId,omitempty"`

	TimePeriod TimeWindow `json:"timePeriod,omitempty"`
}

// AssertNnwdafDataManagementSubscRequired checks if the required fields are not zero-ed
func AssertNnwdafDataManagementSubscRequired(obj NnwdafDataManagementSubsc) error {
	elements := map[string]interface{}{
		"notifCorrId": obj.NotifCorrId,
		"notificURI": obj.NotificURI,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertNnwdafEventsSubscriptionRequired(obj.AnaSub); err != nil {
		return err
	}
	for _, el := range obj.DataCollectPurposes {
		if err := AssertDataCollectionPurposeRequired(el); err != nil {
			return err
		}
	}
	if obj.DataSub != nil {
		if err := AssertDataSubscriptionRequired(*obj.DataSub); err != nil {
			return err
		}
	}
	if err := AssertFormattingInstructionRequired(obj.FormatInstruct); err != nil {
		return err
	}
	if err := AssertProcessingInstructionRequired(obj.ProcInstruct); err != nil {
		return err
	}
	if err := AssertTimeWindowRequired(obj.TimePeriod); err != nil {
		return err
	}
	return nil
}

// AssertNnwdafDataManagementSubscConstraints checks if the values respects the defined constraints
func AssertNnwdafDataManagementSubscConstraints(obj NnwdafDataManagementSubsc) error {
	if err := AssertNnwdafEventsSubscriptionConstraints(obj.AnaSub); err != nil {
		return err
	}
	for _, el := range obj.DataCollectPurposes {
		if err := AssertDataCollectionPurposeConstraints(el); err != nil {
			return err
		}
	}
    if obj.DataSub != nil {
     	if err := AssertDataSubscriptionConstraints(*obj.DataSub); err != nil {
     		return err
     	}
    }
	if err := AssertFormattingInstructionConstraints(obj.FormatInstruct); err != nil {
		return err
	}
	if err := AssertProcessingInstructionConstraints(obj.ProcInstruct); err != nil {
		return err
	}
	if err := AssertTimeWindowConstraints(obj.TimePeriod); err != nil {
		return err
	}
	return nil
}
