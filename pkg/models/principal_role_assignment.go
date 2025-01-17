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

// PrincipalRoleAssignment PrincipalRoleAssignment
//
// List of user ids to add and remove from a project
//
// swagger:discriminator PrincipalRoleAssignment List of user ids to add and remove from a project
type PrincipalRoleAssignment interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Ids of users to add
	Add() []string
	SetAdd([]string)

	// Ids of users to remove
	Remove() []string
	SetRemove([]string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type principalRoleAssignment struct {
	addField []string

	removeField []string
}

// Add gets the add of this polymorphic type
func (m *principalRoleAssignment) Add() []string {
	return m.addField
}

// SetAdd sets the add of this polymorphic type
func (m *principalRoleAssignment) SetAdd(val []string) {
	m.addField = val
}

// Remove gets the remove of this polymorphic type
func (m *principalRoleAssignment) Remove() []string {
	return m.removeField
}

// SetRemove sets the remove of this polymorphic type
func (m *principalRoleAssignment) SetRemove(val []string) {
	m.removeField = val
}

// UnmarshalPrincipalRoleAssignmentSlice unmarshals polymorphic slices of PrincipalRoleAssignment
func UnmarshalPrincipalRoleAssignmentSlice(reader io.Reader, consumer runtime.Consumer) ([]PrincipalRoleAssignment, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []PrincipalRoleAssignment
	for _, element := range elements {
		obj, err := unmarshalPrincipalRoleAssignment(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalPrincipalRoleAssignment unmarshals polymorphic PrincipalRoleAssignment
func UnmarshalPrincipalRoleAssignment(reader io.Reader, consumer runtime.Consumer) (PrincipalRoleAssignment, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalPrincipalRoleAssignment(data, consumer)
}

func unmarshalPrincipalRoleAssignment(data []byte, consumer runtime.Consumer) (PrincipalRoleAssignment, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the List of user ids to add and remove from a project property.
	var getType struct {
		ListOfUserIdsToAddAndRemoveFromaProject string `json:"List of user ids to add and remove from a project"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("List of user ids to add and remove from a project", "body", getType.ListOfUserIdsToAddAndRemoveFromaProject); err != nil {
		return nil, err
	}

	// The value of List of user ids to add and remove from a project is used to determine which type to create and unmarshal the data into
	switch getType.ListOfUserIdsToAddAndRemoveFromaProject {
	case "PrincipalRoleAssignment":
		var result principalRoleAssignment
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid List of user ids to add and remove from a project value: %q", getType.ListOfUserIdsToAddAndRemoveFromaProject)
}

// Validate validates this principal role assignment
func (m *principalRoleAssignment) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this principal role assignment based on context it is used
func (m *principalRoleAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
