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

// MeteringPolicyAssignmentInfo MeteringPolicyAssignmentInfo
//
// swagger:model MeteringPolicyAssignmentInfo
type MeteringPolicyAssignmentInfo struct {

	// count
	Count int32 `json:"count,omitempty"`

	// entity type
	// Enum: [ALL PROJECT CLOUDZONE]
	EntityType string `json:"entityType,omitempty"`
}

// Validate validates this metering policy assignment info
func (m *MeteringPolicyAssignmentInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var meteringPolicyAssignmentInfoTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALL","PROJECT","CLOUDZONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		meteringPolicyAssignmentInfoTypeEntityTypePropEnum = append(meteringPolicyAssignmentInfoTypeEntityTypePropEnum, v)
	}
}

const (

	// MeteringPolicyAssignmentInfoEntityTypeALL captures enum value "ALL"
	MeteringPolicyAssignmentInfoEntityTypeALL string = "ALL"

	// MeteringPolicyAssignmentInfoEntityTypePROJECT captures enum value "PROJECT"
	MeteringPolicyAssignmentInfoEntityTypePROJECT string = "PROJECT"

	// MeteringPolicyAssignmentInfoEntityTypeCLOUDZONE captures enum value "CLOUDZONE"
	MeteringPolicyAssignmentInfoEntityTypeCLOUDZONE string = "CLOUDZONE"
)

// prop value enum
func (m *MeteringPolicyAssignmentInfo) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, meteringPolicyAssignmentInfoTypeEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MeteringPolicyAssignmentInfo) validateEntityType(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this metering policy assignment info based on context it is used
func (m *MeteringPolicyAssignmentInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MeteringPolicyAssignmentInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MeteringPolicyAssignmentInfo) UnmarshalBinary(b []byte) error {
	var res MeteringPolicyAssignmentInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
