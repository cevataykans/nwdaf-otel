// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// PresenceInfo - If the additionalPraId IE is present, this IE shall state the presence information of the UE for the individual PRA identified by the additionalPraId IE;  If the additionalPraId IE is not present, this IE shall state the presence information of the UE for the PRA identified by the praId IE. 
type PresenceInfo struct {

	// Represents an identifier of the Presence Reporting Area (see clause 28.10 of 3GPP  TS 23.003.  This IE shall be present  if the Area of Interest subscribed or reported is a Presence Reporting Area or a Set of Core Network predefined Presence Reporting Areas. When present, it shall be encoded as a string representing an integer in the following ranges: 0 to 8 388 607 for UE-dedicated PRA 8 388 608 to 16 777 215 for Core Network predefined PRA Examples: PRA ID 123 is encoded as \"123\" PRA ID 11 238 660 is encoded as \"11238660\" 
	PraId string `json:"praId,omitempty"`

	// This IE may be present if the praId IE is present and if it contains a PRA identifier referring to a set of Core Network predefined Presence Reporting Areas. When present, this IE shall contain a PRA Identifier of an individual PRA within the Set of Core Network predefined Presence Reporting Areas indicated by the praId IE.  
	AdditionalPraId string `json:"additionalPraId,omitempty"`

	PresenceState PresenceState `json:"presenceState,omitempty"`

	// Represents the list of tracking areas that constitutes the area. This IE shall be present if the subscription or  the event report is for tracking UE presence in the tracking areas. For non 3GPP access the TAI shall be the N3GPP TAI.  
	TrackingAreaList []Tai `json:"trackingAreaList,omitempty"`

	// Represents the list of EUTRAN cell Ids that constitutes the area. This IE shall be present if the Area of Interest subscribed is a list of EUTRAN cell Ids.  
	EcgiList []Ecgi `json:"ecgiList,omitempty"`

	// Represents the list of NR cell Ids that constitutes the area. This IE shall be present if the Area of Interest subscribed is a list of NR cell Ids.  
	NcgiList []Ncgi `json:"ncgiList,omitempty"`

	// Represents the list of NG RAN node identifiers that constitutes the area. This IE shall be present if the Area of Interest subscribed is a list of NG RAN node identifiers.  
	GlobalRanNodeIdList []GlobalRanNodeId `json:"globalRanNodeIdList,omitempty"`

	// Represents the list of eNodeB identifiers that constitutes the area. This IE shall be  present if the Area of Interest subscribed is a list of eNodeB identifiers. 
	GlobaleNbIdList []GlobalRanNodeId `json:"globaleNbIdList,omitempty"`
}

// AssertPresenceInfoRequired checks if the required fields are not zero-ed
func AssertPresenceInfoRequired(obj PresenceInfo) error {
	if err := AssertPresenceStateRequired(obj.PresenceState); err != nil {
		return err
	}
	for _, el := range obj.TrackingAreaList {
		if err := AssertTaiRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EcgiList {
		if err := AssertEcgiRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NcgiList {
		if err := AssertNcgiRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.GlobalRanNodeIdList {
		if err := AssertGlobalRanNodeIdRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.GlobaleNbIdList {
		if err := AssertGlobalRanNodeIdRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPresenceInfoConstraints checks if the values respects the defined constraints
func AssertPresenceInfoConstraints(obj PresenceInfo) error {
	if err := AssertPresenceStateConstraints(obj.PresenceState); err != nil {
		return err
	}
	for _, el := range obj.TrackingAreaList {
		if err := AssertTaiConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.EcgiList {
		if err := AssertEcgiConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NcgiList {
		if err := AssertNcgiConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.GlobalRanNodeIdList {
		if err := AssertGlobalRanNodeIdConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.GlobaleNbIdList {
		if err := AssertGlobalRanNodeIdConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
