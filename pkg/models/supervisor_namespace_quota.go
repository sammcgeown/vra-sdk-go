// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SupervisorNamespaceQuota SupervisorNamespaceQuota
//
// swagger:model SupervisorNamespaceQuota
type SupervisorNamespaceQuota struct {

	// definition
	Definition int64 `json:"definition,omitempty"`

	// resource
	Resource string `json:"resource,omitempty"`
}

// Validate validates this supervisor namespace quota
func (m *SupervisorNamespaceQuota) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this supervisor namespace quota based on context it is used
func (m *SupervisorNamespaceQuota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SupervisorNamespaceQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupervisorNamespaceQuota) UnmarshalBinary(b []byte) error {
	var res SupervisorNamespaceQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
