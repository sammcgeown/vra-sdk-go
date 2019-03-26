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

// DeploymentResource DeploymentResource
// swagger:model DeploymentResource
type DeploymentResource struct {

	// Resource dependency on other resources
	DependsOn []string `json:"dependsOn"`

	// Resource ID
	ID string `json:"id,omitempty"`

	// Resource metadata
	Metadata interface{} `json:"metadata,omitempty"`

	// Resource name
	Name string `json:"name,omitempty"`

	// Resource properties
	Properties interface{} `json:"properties,omitempty"`

	// Resource state
	// Enum: [PARTIAL TAINTED OK]
	State string `json:"state,omitempty"`

	// Resource type
	Type string `json:"type,omitempty"`
}

// Validate validates this deployment resource
func (m *DeploymentResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var deploymentResourceTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PARTIAL","TAINTED","OK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		deploymentResourceTypeStatePropEnum = append(deploymentResourceTypeStatePropEnum, v)
	}
}

const (

	// DeploymentResourceStatePARTIAL captures enum value "PARTIAL"
	DeploymentResourceStatePARTIAL string = "PARTIAL"

	// DeploymentResourceStateTAINTED captures enum value "TAINTED"
	DeploymentResourceStateTAINTED string = "TAINTED"

	// DeploymentResourceStateOK captures enum value "OK"
	DeploymentResourceStateOK string = "OK"
)

// prop value enum
func (m *DeploymentResource) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, deploymentResourceTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DeploymentResource) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentResource) UnmarshalBinary(b []byte) error {
	var res DeploymentResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
