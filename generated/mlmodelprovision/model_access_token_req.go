// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 */

package mlmodelprovision




// AccessTokenReq - Contains information related to the access token request
type AccessTokenReq struct {

	GrantType string `json:"grant_type"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfInstanceId string `json:"nfInstanceId"`

	NfType NfType `json:"nfType,omitempty"`

	TargetNfType NfType `json:"targetNfType,omitempty"`

	Scope string `json:"scope" validate:"regexp=^([a-zA-Z0-9_:-]+)( [a-zA-Z0-9_:-]+)*$"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	TargetNfInstanceId string `json:"targetNfInstanceId,omitempty"`

	RequesterPlmn PlmnId `json:"requesterPlmn,omitempty"`

	RequesterPlmnList []PlmnId `json:"requesterPlmnList,omitempty"`

	RequesterSnssaiList []Snssai `json:"requesterSnssaiList,omitempty"`

	// Fully Qualified Domain Name
	RequesterFqdn string `json:"requesterFqdn,omitempty" validate:"regexp=^([0-9A-Za-z]([-0-9A-Za-z]{0,61}[0-9A-Za-z])?\\\\.)+[A-Za-z]{2,63}\\\\.?$"`

	RequesterSnpnList []PlmnIdNid `json:"requesterSnpnList,omitempty"`

	TargetPlmn PlmnId `json:"targetPlmn,omitempty"`

	TargetSnpn PlmnIdNid `json:"targetSnpn,omitempty"`

	TargetSnssaiList []Snssai `json:"targetSnssaiList,omitempty"`

	TargetNsiList []string `json:"targetNsiList,omitempty"`

	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	TargetNfSetId string `json:"targetNfSetId,omitempty"`

	// NF Service Set Identifier (see clause 28.12 of 3GPP TS 23.003) formatted as the following  string \"set<Set ID>.sn<Service Name>.nfi<NF Instance ID>.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.sn<ServiceName>.nfi<NFInstanceID>.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)   <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NID> encoded as defined in clause 5.4.2 (\"Nid\" data type definition)  <NFInstanceId> encoded as defined in clause 5.3.2  <ServiceName> encoded as defined in 3GPP TS 29.510  <Set ID> encoded as a string of characters consisting of alphabetic    characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that shall end    with either an alphabetic character or a digit. 
	TargetNfServiceSetId string `json:"targetNfServiceSetId,omitempty"`

	// String providing an URI formatted according to RFC 3986.
	HnrfAccessTokenUri string `json:"hnrfAccessTokenUri,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SourceNfInstanceId string `json:"sourceNfInstanceId,omitempty"`
}

// AssertAccessTokenReqRequired checks if the required fields are not zero-ed
func AssertAccessTokenReqRequired(obj AccessTokenReq) error {
	elements := map[string]interface{}{
		"grant_type": obj.GrantType,
		"nfInstanceId": obj.NfInstanceId,
		"scope": obj.Scope,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertNfTypeRequired(obj.NfType); err != nil {
		return err
	}
	if err := AssertNfTypeRequired(obj.TargetNfType); err != nil {
		return err
	}
	if err := AssertPlmnIdRequired(obj.RequesterPlmn); err != nil {
		return err
	}
	for _, el := range obj.RequesterPlmnList {
		if err := AssertPlmnIdRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RequesterSnssaiList {
		if err := AssertSnssaiRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RequesterSnpnList {
		if err := AssertPlmnIdNidRequired(el); err != nil {
			return err
		}
	}
	if err := AssertPlmnIdRequired(obj.TargetPlmn); err != nil {
		return err
	}
	if err := AssertPlmnIdNidRequired(obj.TargetSnpn); err != nil {
		return err
	}
	for _, el := range obj.TargetSnssaiList {
		if err := AssertSnssaiRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAccessTokenReqConstraints checks if the values respects the defined constraints
func AssertAccessTokenReqConstraints(obj AccessTokenReq) error {
	if err := AssertNfTypeConstraints(obj.NfType); err != nil {
		return err
	}
	if err := AssertNfTypeConstraints(obj.TargetNfType); err != nil {
		return err
	}
	if err := AssertPlmnIdConstraints(obj.RequesterPlmn); err != nil {
		return err
	}
	for _, el := range obj.RequesterPlmnList {
		if err := AssertPlmnIdConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RequesterSnssaiList {
		if err := AssertSnssaiConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RequesterSnpnList {
		if err := AssertPlmnIdNidConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertPlmnIdConstraints(obj.TargetPlmn); err != nil {
		return err
	}
	if err := AssertPlmnIdNidConstraints(obj.TargetSnpn); err != nil {
		return err
	}
	for _, el := range obj.TargetSnssaiList {
		if err := AssertSnssaiConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
