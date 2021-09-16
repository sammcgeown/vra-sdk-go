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

// K8SZoneResourceAssignment K8SZoneResourceAssignment
//
// swagger:model K8SZoneResourceAssignment
type K8SZoneResourceAssignment struct {

	// created millis
	CreatedMillis int64 `json:"createdMillis,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// org Id
	OrgID string `json:"orgId,omitempty"`

	// priority
	Priority int32 `json:"priority,omitempty"`

	// resource Id
	ResourceID string `json:"resourceId,omitempty"`

	// resource type
	// Enum: [PKS_PLAN]
	ResourceType string `json:"resourceType,omitempty"`

	// tag links
	TagLinks []string `json:"tagLinks"`

	// tags
	Tags []*TagState `json:"tags"`

	// updated millis
	UpdatedMillis int64 `json:"updatedMillis,omitempty"`
}

// Validate validates this k8 s zone resource assignment
func (m *K8SZoneResourceAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
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

func (m *K8SZoneResourceAssignment) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

var k8SZoneResourceAssignmentTypeResourceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PKS_PLAN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		k8SZoneResourceAssignmentTypeResourceTypePropEnum = append(k8SZoneResourceAssignmentTypeResourceTypePropEnum, v)
	}
}

const (

	// K8SZoneResourceAssignmentResourceTypePKSPLAN captures enum value "PKS_PLAN"
	K8SZoneResourceAssignmentResourceTypePKSPLAN string = "PKS_PLAN"
)

// prop value enum
func (m *K8SZoneResourceAssignment) validateResourceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, k8SZoneResourceAssignmentTypeResourceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *K8SZoneResourceAssignment) validateResourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceTypeEnum("resourceType", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *K8SZoneResourceAssignment) validateTags(formats strfmt.Registry) error {
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

// ContextValidate validate this k8 s zone resource assignment based on the context it is used
func (m *K8SZoneResourceAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *K8SZoneResourceAssignment) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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
func (m *K8SZoneResourceAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *K8SZoneResourceAssignment) UnmarshalBinary(b []byte) error {
	var res K8SZoneResourceAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
