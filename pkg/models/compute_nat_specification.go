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

// ComputeNatSpecification compute nat specification
//
// swagger:model ComputeNatSpecification
type ComputeNatSpecification struct {

	// Additional custom properties that may be used to extend this resource.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// The id of the deployment that is associated with this resource
	// Example: 123e4567-e89b-12d3-a456-426655440000
	DeploymentID string `json:"deploymentId,omitempty"`

	// Id of the Compute Gateway to which the Compute Nat resource will be attached
	// Required: true
	Gateway *string `json:"gateway"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// List of NAT Rules
	// Required: true
	NatRules []*NatRule `json:"natRules"`

	// The id of the project the current user belongs to.
	// Example: e058
	// Required: true
	ProjectID *string `json:"projectId"`
}

// Validate validates this compute nat specification
func (m *ComputeNatSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNatRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeNatSpecification) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *ComputeNatSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ComputeNatSpecification) validateNatRules(formats strfmt.Registry) error {

	if err := validate.Required("natRules", "body", m.NatRules); err != nil {
		return err
	}

	for i := 0; i < len(m.NatRules); i++ {
		if swag.IsZero(m.NatRules[i]) { // not required
			continue
		}

		if m.NatRules[i] != nil {
			if err := m.NatRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("natRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeNatSpecification) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this compute nat specification based on the context it is used
func (m *ComputeNatSpecification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNatRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeNatSpecification) contextValidateNatRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NatRules); i++ {

		if m.NatRules[i] != nil {
			if err := m.NatRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("natRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputeNatSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputeNatSpecification) UnmarshalBinary(b []byte) error {
	var res ComputeNatSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}