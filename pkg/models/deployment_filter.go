// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentFilter Filter
//
// swagger:model DeploymentFilter
type DeploymentFilter struct {

	// content
	Content []*FilterEntry `json:"content"`

	// empty
	Empty bool `json:"empty,omitempty"`

	// first
	First bool `json:"first,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// last
	Last bool `json:"last,omitempty"`

	// name
	Name string `json:"name,omitempty"`

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

	// type
	// Enum: [MULTISELECT DATE_RANGE BOOLEAN]
	Type string `json:"type,omitempty"`
}

// Validate validates this deployment filter
func (m *DeploymentFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
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

func (m *DeploymentFilter) validateContent(formats strfmt.Registry) error {
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

func (m *DeploymentFilter) validateSort(formats strfmt.Registry) error {
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

var deploymentFilterTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MULTISELECT","DATE_RANGE","BOOLEAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentFilterTypeTypePropEnum = append(deploymentFilterTypeTypePropEnum, v)
	}
}

const (

	// DeploymentFilterTypeMULTISELECT captures enum value "MULTISELECT"
	DeploymentFilterTypeMULTISELECT string = "MULTISELECT"

	// DeploymentFilterTypeDATERANGE captures enum value "DATE_RANGE"
	DeploymentFilterTypeDATERANGE string = "DATE_RANGE"

	// DeploymentFilterTypeBOOLEAN captures enum value "BOOLEAN"
	DeploymentFilterTypeBOOLEAN string = "BOOLEAN"
)

// prop value enum
func (m *DeploymentFilter) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, deploymentFilterTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentFilter) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this deployment filter based on the context it is used
func (m *DeploymentFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *DeploymentFilter) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

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

func (m *DeploymentFilter) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

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
func (m *DeploymentFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentFilter) UnmarshalBinary(b []byte) error {
	var res DeploymentFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
