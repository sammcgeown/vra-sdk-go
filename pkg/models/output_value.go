// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OutputValue OutputValue
//
// swagger:model OutputValue
type OutputValue struct {

	// The description of the output value in the Terraform configuration.
	Description string `json:"description,omitempty"`

	// The name of the output value in the Terraform configuration.
	Name string `json:"name,omitempty"`

	// A flag indicating that the field should be obscured because of security concerns.
	Sensitive bool `json:"sensitive,omitempty"`
}

// Validate validates this output value
func (m *OutputValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this output value based on context it is used
func (m *OutputValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OutputValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OutputValue) UnmarshalBinary(b []byte) error {
	var res OutputValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}