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
	"github.com/go-openapi/validate"
)

// CloudAccountAwsSpecification Specification for a Amazon cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region or data center where the associated cloud account resources are hosted.
//
// swagger:model CloudAccountAwsSpecification
type CloudAccountAwsSpecification struct {

	// Aws Access key ID
	// Example: ACDC55DB4MFH6ADG75KK
	// Required: true
	AccessKeyID *string `json:"accessKeyId"`

	// Create default cloud zones for the enabled regions.
	// Example: true
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// A set of Region names to enable provisioning on. Refer to /iaas/cloud-accounts-aws/region-enumeration..
	// Example: [ \"us-east-1\", \"ap-northeast-1\" ]
	// Required: true
	RegionIds []string `json:"regionIds"`

	// Aws Secret Access Key
	// Example: gfsScK345sGGaVdds222dasdfDDSSasdfdsa34fS
	// Required: true
	SecretAccessKey *string `json:"secretAccessKey"`

	// A set of tag keys and optional values to set on the Cloud Account
	// Example: [ { \"key\" : \"env\", \"value\": \"dev\" } ]
	Tags []*Tag `json:"tags"`
}

// Validate validates this cloud account aws specification
func (m *CloudAccountAwsSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKeyID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountAwsSpecification) validateAccessKeyID(formats strfmt.Registry) error {

	if err := validate.Required("accessKeyId", "body", m.AccessKeyID); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountAwsSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountAwsSpecification) validateRegionIds(formats strfmt.Registry) error {

	if err := validate.Required("regionIds", "body", m.RegionIds); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountAwsSpecification) validateSecretAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("secretAccessKey", "body", m.SecretAccessKey); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountAwsSpecification) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this cloud account aws specification based on the context it is used
func (m *CloudAccountAwsSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountAwsSpecification) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CloudAccountAwsSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountAwsSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountAwsSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
