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

// PageOfDeploymentResourceType PageOfResourceType
//
// swagger:model PageOfDeploymentResourceType
type PageOfDeploymentResourceType struct {

	// content
	Content []*DeploymentResourceType `json:"content"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// first
	First bool `json:"first,omitempty"`

	// last
	Last bool `json:"last,omitempty"`

	// number
	Number int32 `json:"number,omitempty"`

	// number of elements
	NumberOfElements int32 `json:"numberOfElements,omitempty"`

	// size
	Size int32 `json:"size,omitempty"`

	// sort
	Sort *Sort `json:"sort,omitempty"`

	// total elements
	TotalElements int64 `json:"totalElements,omitempty"`

	// total pages
	TotalPages int32 `json:"totalPages,omitempty"`
}

// Validate validates this page of deployment resource type
func (m *PageOfDeploymentResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageOfDeploymentResourceType) validateContent(formats strfmt.Registry) error {
	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {
		if swag.IsZero(m.Content[i]) { // not required
			continue
		}

		if m.Content[i] != nil {
			if err := m.Content[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PageOfDeploymentResourceType) validateSort(formats strfmt.Registry) error {
	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	if m.Sort != nil {
		if err := m.Sort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sort")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this page of deployment resource type based on the context it is used
func (m *PageOfDeploymentResourceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PageOfDeploymentResourceType) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Content); i++ {

		if m.Content[i] != nil {
			if err := m.Content[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PageOfDeploymentResourceType) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

	if m.Sort != nil {
		if err := m.Sort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sort")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PageOfDeploymentResourceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PageOfDeploymentResourceType) UnmarshalBinary(b []byte) error {
	var res PageOfDeploymentResourceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}