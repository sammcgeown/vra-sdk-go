// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContentDefinition ContentDefinition
//
// Represents a catalog item or content source that is linked to a project via an entitlement.
// swagger:model ContentDefinition
type ContentDefinition struct {

	// Description of either the catalog item or the catalog source
	Description string `json:"description,omitempty"`

	// Icon id of associated catalog item (if association is with catalog item)
	// Format: uuid
	IconID strfmt.UUID `json:"iconId,omitempty"`

	// Id of either the catalog source or catalog item.
	// Required: true
	// Format: uuid
	ID *strfmt.UUID `json:"id"`

	// Name of either the catalog item or the catalog source
	Name string `json:"name,omitempty"`

	// Number of items in the associated catalog source
	NumItems int32 `json:"numItems,omitempty"`

	// Catalog source name
	SourceName string `json:"sourceName,omitempty"`

	// Catalog source type
	SourceType string `json:"sourceType,omitempty"`

	// Content definition type
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this content definition
func (m *ContentDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIconID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentDefinition) validateIconID(formats strfmt.Registry) error {

	if swag.IsZero(m.IconID) { // not required
		return nil
	}

	if err := validate.FormatOf("iconId", "body", "uuid", m.IconID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContentDefinition) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContentDefinition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentDefinition) UnmarshalBinary(b []byte) error {
	var res ContentDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
