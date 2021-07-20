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

// ServiceRequest ServiceRequest
//
// Request object for actions such as cloning.
//
// swagger:discriminator ServiceRequest Request object for actions such as cloning.
type ServiceRequest interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Description for the cloned entity.
	// Example: cloned entity
	Description() string
	SetDescription(string)

	// Name field for the cloned entity.
	// Example: Pipeline-1
	Name() string
	SetName(string)

	// Action to be performed on the service.
	// Example: CLONE
	RequestType() string
	SetRequestType(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type serviceRequest struct {
	descriptionField string

	nameField string

	requestTypeField string
}

// Description gets the description of this polymorphic type
func (m *serviceRequest) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *serviceRequest) SetDescription(val string) {
	m.descriptionField = val
}

// Name gets the name of this polymorphic type
func (m *serviceRequest) Name() string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *serviceRequest) SetName(val string) {
	m.nameField = val
}

// RequestType gets the request type of this polymorphic type
func (m *serviceRequest) RequestType() string {
	return m.requestTypeField
}

// SetRequestType sets the request type of this polymorphic type
func (m *serviceRequest) SetRequestType(val string) {
	m.requestTypeField = val
}

// UnmarshalServiceRequestSlice unmarshals polymorphic slices of ServiceRequest
func UnmarshalServiceRequestSlice(reader io.Reader, consumer runtime.Consumer) ([]ServiceRequest, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ServiceRequest
	for _, element := range elements {
		obj, err := unmarshalServiceRequest(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalServiceRequest unmarshals polymorphic ServiceRequest
func UnmarshalServiceRequest(reader io.Reader, consumer runtime.Consumer) (ServiceRequest, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalServiceRequest(data, consumer)
}

func unmarshalServiceRequest(data []byte, consumer runtime.Consumer) (ServiceRequest, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Request object for actions such as cloning. property.
	var getType struct {
		RequestObjectForActionsSuchAsCloning string `json:"Request object for actions such as cloning."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Request object for actions such as cloning.", "body", getType.RequestObjectForActionsSuchAsCloning); err != nil {
		return nil, err
	}

	// The value of Request object for actions such as cloning. is used to determine which type to create and unmarshal the data into
	switch getType.RequestObjectForActionsSuchAsCloning {
	case "ServiceRequest":
		var result serviceRequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Request object for actions such as cloning. value: %q", getType.RequestObjectForActionsSuchAsCloning)
}

// Validate validates this service request
func (m *serviceRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service request based on context it is used
func (m *serviceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}