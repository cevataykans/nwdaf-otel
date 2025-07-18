// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_MLModelProvision
 *
 * Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 */

package mlmodelprovision




// IpAddr - Contains an IP adresse.
type IpAddr struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`

	Ipv6Prefix Ipv6Prefix `json:"ipv6Prefix,omitempty"`
}

// AssertIpAddrRequired checks if the required fields are not zero-ed
func AssertIpAddrRequired(obj IpAddr) error {
	if err := AssertIpv6AddrRequired(obj.Ipv6Addr); err != nil {
		return err
	}
	if err := AssertIpv6PrefixRequired(obj.Ipv6Prefix); err != nil {
		return err
	}
	return nil
}

// AssertIpAddrConstraints checks if the values respects the defined constraints
func AssertIpAddrConstraints(obj IpAddr) error {
	if err := AssertIpv6AddrConstraints(obj.Ipv6Addr); err != nil {
		return err
	}
	if err := AssertIpv6PrefixConstraints(obj.Ipv6Prefix); err != nil {
		return err
	}
	return nil
}
