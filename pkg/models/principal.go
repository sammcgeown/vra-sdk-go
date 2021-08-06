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

// Principal Principal
//
// A representation of a user or group.
//
// swagger:discriminator Principal A representation of a user or group.
type Principal interface {
	runtime.Validatable
	runtime.ContextValidatable

	// The email of the user or name of the group.
	// Example: administrator@vmware.com
	// Required: true
	Email() *string
	SetEmail(*string)

	// Type of the principal. Currently supported 'user' (default) and 'group'.
	// Example: user
	Type() string
	SetType(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type principal struct {
	emailField *string

	typeField string
}

// Email gets the email of this polymorphic type
func (m *principal) Email() *string {
	return m.emailField
}

// SetEmail sets the email of this polymorphic type
func (m *principal) SetEmail(val *string) {
	m.emailField = val
}

// Type gets the type of this polymorphic type
func (m *principal) Type() string {
	return m.typeField
}

// SetType sets the type of this polymorphic type
func (m *principal) SetType(val string) {
	m.typeField = val
}

// UnmarshalPrincipalSlice unmarshals polymorphic slices of Principal
func UnmarshalPrincipalSlice(reader io.Reader, consumer runtime.Consumer) ([]Principal, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Principal
	for _, element := range elements {
		obj, err := unmarshalPrincipal(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalPrincipal unmarshals polymorphic Principal
func UnmarshalPrincipal(reader io.Reader, consumer runtime.Consumer) (Principal, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalPrincipal(data, consumer)
}

func unmarshalPrincipal(data []byte, consumer runtime.Consumer) (Principal, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the A representation of a user or group. property.
	var getType struct {
		ARepresentationOfaUserOrGroup string `json:"A representation of a user or group."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("A representation of a user or group.", "body", getType.ARepresentationOfaUserOrGroup); err != nil {
		return nil, err
	}

	// The value of A representation of a user or group. is used to determine which type to create and unmarshal the data into
	switch getType.ARepresentationOfaUserOrGroup {
	case "Principal":
		var result principal
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid A representation of a user or group. value: %q", getType.ARepresentationOfaUserOrGroup)
}

// Validate validates this principal
func (m *principal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *principal) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this principal based on context it is used
func (m *principal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}