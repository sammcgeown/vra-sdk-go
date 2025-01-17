// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EndpointValidationSpec EndpointValidationSpec
//
// Endpoint validation specification
//
// swagger:discriminator EndpointValidationSpec Endpoint validation specification
type EndpointValidationSpec interface {
	runtime.Validatable
	runtime.ContextValidatable

	// A human-friendly description.
	// Required: true
	Description() *string
	SetDescription(*string)

	// id of the endpoint, if already present
	ID() string
	SetID(string)

	// This type of Endpoint can be created, updated or deleted by admin only. If a restricted Endpoint is consumed in a pipeline, and that pipeline is executed by a non-admin user, then the execution will fail at the task which is consuming this restricted Endpoint. Only admin can then resume this pipeline to make it progress.
	// Example: false
	IsRestricted() bool
	SetIsRestricted(bool)

	// A human-friendly name used as an identifier in APIs that support this option
	// Example: My-Name
	// Required: true
	Name() *string
	SetName(*string)

	// The project this entity belongs to.
	// Example: My-Project
	Project() string
	SetProject(string)

	// Endpoint specific properties
	// Required: true
	Properties() interface{}
	SetProperties(interface{})

	// The type of this Endpoint instance.
	// Example: jenkins
	// Required: true
	Type() *string
	SetType(*string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type endpointValidationSpec struct {
	descriptionField *string

	idField string

	isRestrictedField bool

	nameField *string

	projectField string

	propertiesField interface{}

	typeField *string
}

// Description gets the description of this polymorphic type
func (m *endpointValidationSpec) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *endpointValidationSpec) SetDescription(val *string) {
	m.descriptionField = val
}

// ID gets the id of this polymorphic type
func (m *endpointValidationSpec) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *endpointValidationSpec) SetID(val string) {
	m.idField = val
}

// IsRestricted gets the is restricted of this polymorphic type
func (m *endpointValidationSpec) IsRestricted() bool {
	return m.isRestrictedField
}

// SetIsRestricted sets the is restricted of this polymorphic type
func (m *endpointValidationSpec) SetIsRestricted(val bool) {
	m.isRestrictedField = val
}

// Name gets the name of this polymorphic type
func (m *endpointValidationSpec) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *endpointValidationSpec) SetName(val *string) {
	m.nameField = val
}

// Project gets the project of this polymorphic type
func (m *endpointValidationSpec) Project() string {
	return m.projectField
}

// SetProject sets the project of this polymorphic type
func (m *endpointValidationSpec) SetProject(val string) {
	m.projectField = val
}

// Properties gets the properties of this polymorphic type
func (m *endpointValidationSpec) Properties() interface{} {
	return m.propertiesField
}

// SetProperties sets the properties of this polymorphic type
func (m *endpointValidationSpec) SetProperties(val interface{}) {
	m.propertiesField = val
}

// Type gets the type of this polymorphic type
func (m *endpointValidationSpec) Type() *string {
	return m.typeField
}

// SetType sets the type of this polymorphic type
func (m *endpointValidationSpec) SetType(val *string) {
	m.typeField = val
}

// UnmarshalEndpointValidationSpecSlice unmarshals polymorphic slices of EndpointValidationSpec
func UnmarshalEndpointValidationSpecSlice(reader io.Reader, consumer runtime.Consumer) ([]EndpointValidationSpec, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []EndpointValidationSpec
	for _, element := range elements {
		obj, err := unmarshalEndpointValidationSpec(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalEndpointValidationSpec unmarshals polymorphic EndpointValidationSpec
func UnmarshalEndpointValidationSpec(reader io.Reader, consumer runtime.Consumer) (EndpointValidationSpec, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalEndpointValidationSpec(data, consumer)
}

func unmarshalEndpointValidationSpec(data []byte, consumer runtime.Consumer) (EndpointValidationSpec, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Endpoint validation specification property.
	var getType struct {
		EndpointValidationSpecification string `json:"Endpoint validation specification"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Endpoint validation specification", "body", getType.EndpointValidationSpecification); err != nil {
		return nil, err
	}

	// The value of Endpoint validation specification is used to determine which type to create and unmarshal the data into
	switch getType.EndpointValidationSpecification {
	case "EndpointValidationSpec":
		var result endpointValidationSpec
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Endpoint validation specification value: %q", getType.EndpointValidationSpecification)
}

// Validate validates this endpoint validation spec
func (m *endpointValidationSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
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

func (m *endpointValidationSpec) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *endpointValidationSpec) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *endpointValidationSpec) validateProperties(formats strfmt.Registry) error {

	if m.Properties == nil {
		return errors.Required("properties", "body", nil)
	}

	return nil
}

func (m *endpointValidationSpec) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this endpoint validation spec based on context it is used
func (m *endpointValidationSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
