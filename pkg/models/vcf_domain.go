// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VcfDomain VcfDomain
//
// swagger:model VcfDomain
type VcfDomain struct {

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// nsx resource
	NsxResource *NsxAccount `json:"nsxResource,omitempty"`

	// status
	// Enum: [ACTIVE CONFIGURED NOT_CONFIGURED NOT_ACTIVE]
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// vsphere
	Vsphere *VsphereAccount `json:"vsphere,omitempty"`
}

// Validate validates this vcf domain
func (m *VcfDomain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNsxResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcfDomain) validateNsxResource(formats strfmt.Registry) error {
	if swag.IsZero(m.NsxResource) { // not required
		return nil
	}

	if m.NsxResource != nil {
		if err := m.NsxResource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxResource")
			}
			return err
		}
	}

	return nil
}

var vcfDomainTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","CONFIGURED","NOT_CONFIGURED","NOT_ACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vcfDomainTypeStatusPropEnum = append(vcfDomainTypeStatusPropEnum, v)
	}
}

const (

	// VcfDomainStatusACTIVE captures enum value "ACTIVE"
	VcfDomainStatusACTIVE string = "ACTIVE"

	// VcfDomainStatusCONFIGURED captures enum value "CONFIGURED"
	VcfDomainStatusCONFIGURED string = "CONFIGURED"

	// VcfDomainStatusNOTCONFIGURED captures enum value "NOT_CONFIGURED"
	VcfDomainStatusNOTCONFIGURED string = "NOT_CONFIGURED"

	// VcfDomainStatusNOTACTIVE captures enum value "NOT_ACTIVE"
	VcfDomainStatusNOTACTIVE string = "NOT_ACTIVE"
)

// prop value enum
func (m *VcfDomain) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vcfDomainTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VcfDomain) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *VcfDomain) validateVsphere(formats strfmt.Registry) error {
	if swag.IsZero(m.Vsphere) { // not required
		return nil
	}

	if m.Vsphere != nil {
		if err := m.Vsphere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vcf domain based on the context it is used
func (m *VcfDomain) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNsxResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVsphere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcfDomain) contextValidateNsxResource(ctx context.Context, formats strfmt.Registry) error {

	if m.NsxResource != nil {
		if err := m.NsxResource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nsxResource")
			}
			return err
		}
	}

	return nil
}

func (m *VcfDomain) contextValidateVsphere(ctx context.Context, formats strfmt.Registry) error {

	if m.Vsphere != nil {
		if err := m.Vsphere.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcfDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcfDomain) UnmarshalBinary(b []byte) error {
	var res VcfDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}