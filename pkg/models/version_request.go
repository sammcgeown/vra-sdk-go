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

// VersionRequest VersionRequest
//
// Request object for version of a Custom Integration.
//
// swagger:discriminator VersionRequest Request object for version of a Custom Integration.
type VersionRequest interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Changelog describing the changes between this and previous versions.
	// Example: Modified API signatures.
	ChangeLog() string
	SetChangeLog(string)

	// Description of the version.
	// Example: This is the latest version.
	Description() string
	SetDescription(string)

	// Version of the Custom Integration.
	// Example: v1
	Version() string
	SetVersion(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type versionRequest struct {
	changeLogField string

	descriptionField string

	versionField string
}

// ChangeLog gets the change log of this polymorphic type
func (m *versionRequest) ChangeLog() string {
	return m.changeLogField
}

// SetChangeLog sets the change log of this polymorphic type
func (m *versionRequest) SetChangeLog(val string) {
	m.changeLogField = val
}

// Description gets the description of this polymorphic type
func (m *versionRequest) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *versionRequest) SetDescription(val string) {
	m.descriptionField = val
}

// Version gets the version of this polymorphic type
func (m *versionRequest) Version() string {
	return m.versionField
}

// SetVersion sets the version of this polymorphic type
func (m *versionRequest) SetVersion(val string) {
	m.versionField = val
}

// UnmarshalVersionRequestSlice unmarshals polymorphic slices of VersionRequest
func UnmarshalVersionRequestSlice(reader io.Reader, consumer runtime.Consumer) ([]VersionRequest, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []VersionRequest
	for _, element := range elements {
		obj, err := unmarshalVersionRequest(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalVersionRequest unmarshals polymorphic VersionRequest
func UnmarshalVersionRequest(reader io.Reader, consumer runtime.Consumer) (VersionRequest, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalVersionRequest(data, consumer)
}

func unmarshalVersionRequest(data []byte, consumer runtime.Consumer) (VersionRequest, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Request object for version of a Custom Integration. property.
	var getType struct {
		RequestObjectForVersionOfaCustomIntegration string `json:"Request object for version of a Custom Integration."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Request object for version of a Custom Integration.", "body", getType.RequestObjectForVersionOfaCustomIntegration); err != nil {
		return nil, err
	}

	// The value of Request object for version of a Custom Integration. is used to determine which type to create and unmarshal the data into
	switch getType.RequestObjectForVersionOfaCustomIntegration {
	case "VersionRequest":
		var result versionRequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Request object for version of a Custom Integration. value: %q", getType.RequestObjectForVersionOfaCustomIntegration)
}

// Validate validates this version request
func (m *versionRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this version request based on context it is used
func (m *versionRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
