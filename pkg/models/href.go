// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Href href
//
// swagger:model Href
type Href struct {

	// href
	Href string `json:"href,omitempty"`

	// hrefs
	Hrefs []string `json:"hrefs"`
}

// Validate validates this href
func (m *Href) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this href based on context it is used
func (m *Href) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Href) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Href) UnmarshalBinary(b []byte) error {
	var res Href
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
