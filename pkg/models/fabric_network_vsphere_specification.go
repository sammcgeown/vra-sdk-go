// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FabricNetworkVsphereSpecification Specification for updating a Vsphere FabricNetwork
//
// swagger:model FabricNetworkVsphereSpecification
type FabricNetworkVsphereSpecification struct {

	// Network CIDR to be used.
	// Example: 10.1.2.0/24
	Cidr string `json:"cidr,omitempty"`

	// IPv4 default gateway to be used.
	// Example: 10.1.2.1
	DefaultGateway string `json:"defaultGateway,omitempty"`

	// IPv6 default gateway to be used.
	// Example: 2001:eeee:6bd:2a::1
	DefaultIPV6Gateway string `json:"defaultIpv6Gateway,omitempty"`

	// A list of DNS search domains that were set on this resource instance.
	// Example: [vmware.com]
	DNSSearchDomains []string `json:"dnsSearchDomains"`

	// A list of DNS server addresses that were set on this resource instance.
	// Example: [1.1.1.1]
	DNSServerAddresses []string `json:"dnsServerAddresses"`

	// Domain value.
	// Example: sqa.local
	Domain string `json:"domain,omitempty"`

	// Network IPv6 CIDR to be used.
	// Example: 2001:eeee:6bd:2a::1/64
	IPV6Cidr string `json:"ipv6Cidr,omitempty"`

	// Indicates whether this is the default subnet for the zone.
	IsDefault *bool `json:"isDefault,omitempty"`

	// Indicates whether the sub-network supports public IP assignment.
	IsPublic *bool `json:"isPublic,omitempty"`

	// A set of tag keys and optional values that were set on this resource instance.
	// Example: [ { \"key\" : \"fast-network\", \"value\": \"true\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this fabric network vsphere specification
func (m *FabricNetworkVsphereSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricNetworkVsphereSpecification) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this fabric network vsphere specification based on the context it is used
func (m *FabricNetworkVsphereSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricNetworkVsphereSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricNetworkVsphereSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricNetworkVsphereSpecification) UnmarshalBinary(b []byte) error {
	var res FabricNetworkVsphereSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
