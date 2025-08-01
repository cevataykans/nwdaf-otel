// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo




// ScpInfo - Information of an SCP Instance
type ScpInfo struct {

	// A map (list of key-value pairs) where the key of the map shall be the string identifying an SCP domain 
	ScpDomainInfoList map[string]ScpDomainInfo `json:"scpDomainInfoList,omitempty"`

	ScpPrefix string `json:"scpPrefix,omitempty"`

	// Port numbers for HTTP and HTTPS. The key of the map shall be \"http\" or \"https\". 
	ScpPorts map[string]int32 `json:"scpPorts,omitempty"`

	AddressDomains []string `json:"addressDomains,omitempty"`

	Ipv4Addresses []string `json:"ipv4Addresses,omitempty"`

	Ipv6Prefixes []Ipv6Prefix `json:"ipv6Prefixes,omitempty"`

	Ipv4AddrRanges []Ipv4AddressRange `json:"ipv4AddrRanges,omitempty"`

	Ipv6PrefixRanges []Ipv6PrefixRange `json:"ipv6PrefixRanges,omitempty"`

	ServedNfSetIdList []string `json:"servedNfSetIdList,omitempty"`

	RemotePlmnList []PlmnId `json:"remotePlmnList,omitempty"`

	RemoteSnpnList []PlmnIdNid `json:"remoteSnpnList,omitempty"`

	IpReachability IpReachability `json:"ipReachability,omitempty"`

	ScpCapabilities []ScpCapability `json:"scpCapabilities,omitempty"`
}

// AssertScpInfoRequired checks if the required fields are not zero-ed
func AssertScpInfoRequired(obj ScpInfo) error {
	for _, el := range obj.Ipv6Prefixes {
		if err := AssertIpv6PrefixRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ipv4AddrRanges {
		if err := AssertIpv4AddressRangeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ipv6PrefixRanges {
		if err := AssertIpv6PrefixRangeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RemotePlmnList {
		if err := AssertPlmnIdRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RemoteSnpnList {
		if err := AssertPlmnIdNidRequired(el); err != nil {
			return err
		}
	}
	if err := AssertIpReachabilityRequired(obj.IpReachability); err != nil {
		return err
	}
	for _, el := range obj.ScpCapabilities {
		if err := AssertScpCapabilityRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertScpInfoConstraints checks if the values respects the defined constraints
func AssertScpInfoConstraints(obj ScpInfo) error {
	for _, el := range obj.Ipv6Prefixes {
		if err := AssertIpv6PrefixConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ipv4AddrRanges {
		if err := AssertIpv4AddressRangeConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Ipv6PrefixRanges {
		if err := AssertIpv6PrefixRangeConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RemotePlmnList {
		if err := AssertPlmnIdConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.RemoteSnpnList {
		if err := AssertPlmnIdNidConstraints(el); err != nil {
			return err
		}
	}
	if err := AssertIpReachabilityConstraints(obj.IpReachability); err != nil {
		return err
	}
	for _, el := range obj.ScpCapabilities {
		if err := AssertScpCapabilityConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
