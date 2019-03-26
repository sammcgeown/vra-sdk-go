// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Resource Resource
//
// A resource part of a deployment.
// swagger:model Resource
type Resource struct {

	// Creation time
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// A list of other resources this resource depends on
	DependsOn []string `json:"dependsOn"`

	// A description of the resource
	Description string `json:"description,omitempty"`

	// Expense associated with the deployment.
	// Read Only: true
	Expense *Expense `json:"expense,omitempty"`

	// Unique identifier of the resource
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// Name of the resource
	// Required: true
	Name *string `json:"name"`

	// properties
	Properties interface{} `json:"properties,omitempty"`

	// The current state of the resource
	// Enum: [PARTIAL TAINTED OK]
	State string `json:"state,omitempty"`

	// The current sync status
	// Enum: [SUCCESS MISSING STALE]
	SyncStatus string `json:"syncStatus,omitempty"`

	// Type of the resource
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this resource
func (m *Resource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpense(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSyncStatus(formats); err != nil {
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

func (m *Resource) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateExpense(formats strfmt.Registry) error {

	if swag.IsZero(m.Expense) { // not required
		return nil
	}

	if m.Expense != nil {
		if err := m.Expense.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expense")
			}
			return err
		}
	}

	return nil
}

func (m *Resource) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var resourceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PARTIAL","TAINTED","OK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeStatePropEnum = append(resourceTypeStatePropEnum, v)
	}
}

const (

	// ResourceStatePARTIAL captures enum value "PARTIAL"
	ResourceStatePARTIAL string = "PARTIAL"

	// ResourceStateTAINTED captures enum value "TAINTED"
	ResourceStateTAINTED string = "TAINTED"

	// ResourceStateOK captures enum value "OK"
	ResourceStateOK string = "OK"
)

// prop value enum
func (m *Resource) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Resource) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var resourceTypeSyncStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","MISSING","STALE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		resourceTypeSyncStatusPropEnum = append(resourceTypeSyncStatusPropEnum, v)
	}
}

const (

	// ResourceSyncStatusSUCCESS captures enum value "SUCCESS"
	ResourceSyncStatusSUCCESS string = "SUCCESS"

	// ResourceSyncStatusMISSING captures enum value "MISSING"
	ResourceSyncStatusMISSING string = "MISSING"

	// ResourceSyncStatusSTALE captures enum value "STALE"
	ResourceSyncStatusSTALE string = "STALE"
)

// prop value enum
func (m *Resource) validateSyncStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, resourceTypeSyncStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Resource) validateSyncStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.SyncStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateSyncStatusEnum("syncStatus", "body", m.SyncStatus); err != nil {
		return err
	}

	return nil
}

func (m *Resource) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Resource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Resource) UnmarshalBinary(b []byte) error {
	var res Resource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
