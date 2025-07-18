// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




type ServiceDataFlowDescription struct {

	FlowDescription IpPacketFilterSet `json:"flowDescription,omitempty"`

	DomainName string `json:"domainName,omitempty"`
}

// AssertServiceDataFlowDescriptionRequired checks if the required fields are not zero-ed
func AssertServiceDataFlowDescriptionRequired(obj ServiceDataFlowDescription) error {
	if err := AssertIpPacketFilterSetRequired(obj.FlowDescription); err != nil {
		return err
	}
	return nil
}

// AssertServiceDataFlowDescriptionConstraints checks if the values respects the defined constraints
func AssertServiceDataFlowDescriptionConstraints(obj ServiceDataFlowDescription) error {
	if err := AssertIpPacketFilterSetConstraints(obj.FlowDescription); err != nil {
		return err
	}
	return nil
}
