// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// NefInfo - Information of an NEF NF Instance
type NefInfo struct {

	// Identity of the NEF
	NefId string `json:"nefId,omitempty"`

	PfdData PfdData `json:"pfdData,omitempty"`

	AfEeData AfEventExposureData `json:"afEeData,omitempty"`

	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`

	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`

	ServedFqdnList []string `json:"servedFqdnList,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	DnaiList []string `json:"dnaiList,omitempty"`

	UnTrustAfInfoList []UnTrustAfInfo `json:"unTrustAfInfoList,omitempty"`

	UasNfFunctionalityInd bool `json:"uasNfFunctionalityInd,omitempty"`
}

// AssertNefInfoRequired checks if the required fields are not zero-ed
func AssertNefInfoRequired(obj NefInfo) error {
	if err := AssertPfdDataRequired(obj.PfdData); err != nil {
		return err
	}
	if err := AssertAfEventExposureDataRequired(obj.AfEeData); err != nil {
		return err
	}
	for _, el := range obj.GpsiRanges {
		if err := AssertIdentityRangeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ExternalGroupIdentifiersRanges {
		if err := AssertIdentityRangeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TaiList {
		if err := AssertTaiRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TaiRangeList {
		if err := AssertTaiRangeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UnTrustAfInfoList {
		if err := AssertUnTrustAfInfoRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertNefInfoConstraints checks if the values respects the defined constraints
func AssertNefInfoConstraints(obj NefInfo) error {
	if err := AssertPfdDataConstraints(obj.PfdData); err != nil {
		return err
	}
	if err := AssertAfEventExposureDataConstraints(obj.AfEeData); err != nil {
		return err
	}
	for _, el := range obj.GpsiRanges {
		if err := AssertIdentityRangeConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ExternalGroupIdentifiersRanges {
		if err := AssertIdentityRangeConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TaiList {
		if err := AssertTaiConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.TaiRangeList {
		if err := AssertTaiRangeConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UnTrustAfInfoList {
		if err := AssertUnTrustAfInfoConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
