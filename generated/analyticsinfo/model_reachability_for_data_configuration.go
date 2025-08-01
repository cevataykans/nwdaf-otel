// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




type ReachabilityForDataConfiguration struct {

	ReportCfg ReachabilityForDataReportConfig `json:"reportCfg"`

	// indicating a time in seconds.
	MinInterval int32 `json:"minInterval,omitempty"`
}

// AssertReachabilityForDataConfigurationRequired checks if the required fields are not zero-ed
func AssertReachabilityForDataConfigurationRequired(obj ReachabilityForDataConfiguration) error {
	elements := map[string]interface{}{
		"reportCfg": obj.ReportCfg,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertReachabilityForDataReportConfigRequired(obj.ReportCfg); err != nil {
		return err
	}
	return nil
}

// AssertReachabilityForDataConfigurationConstraints checks if the values respects the defined constraints
func AssertReachabilityForDataConfigurationConstraints(obj ReachabilityForDataConfiguration) error {
	if err := AssertReachabilityForDataReportConfigConstraints(obj.ReportCfg); err != nil {
		return err
	}
	return nil
}
