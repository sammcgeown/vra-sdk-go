// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Sort Sort
//
// swagger:model Sort
type Sort struct {

	// empty
	Empty bool `json:"empty,omitempty"`

	// sorted
	Sorted bool `json:"sorted,omitempty"`

	// unsorted
	Unsorted bool `json:"unsorted,omitempty"`
}

// Validate validates this sort
func (m *Sort) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sort based on context it is used
func (m *Sort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Sort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Sort) UnmarshalBinary(b []byte) error {
	var res Sort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
