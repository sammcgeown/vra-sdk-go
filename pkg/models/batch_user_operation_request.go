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

// BatchUserOperationRequest BatchUserOperationRequest
//
// Batch user operation request.
//
// swagger:discriminator BatchUserOperationRequest Batch user operation request.
type BatchUserOperationRequest interface {
	runtime.Validatable
	runtime.ContextValidatable

	// The list of user-op ids to be batch approved/rejected.
	// Required: true
	Ids() []string
	SetIds([]string)

	// The response message which the responder would like to give.
	// Example: Approved on Monday.
	ResponseMessage() string
	SetResponseMessage(string)

	// The status of approval requests.
	// Example: Approved
	// Required: true
	Status() *string
	SetStatus(*string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type batchUserOperationRequest struct {
	idsField []string

	responseMessageField string

	statusField *string
}

// Ids gets the ids of this polymorphic type
func (m *batchUserOperationRequest) Ids() []string {
	return m.idsField
}

// SetIds sets the ids of this polymorphic type
func (m *batchUserOperationRequest) SetIds(val []string) {
	m.idsField = val
}

// ResponseMessage gets the response message of this polymorphic type
func (m *batchUserOperationRequest) ResponseMessage() string {
	return m.responseMessageField
}

// SetResponseMessage sets the response message of this polymorphic type
func (m *batchUserOperationRequest) SetResponseMessage(val string) {
	m.responseMessageField = val
}

// Status gets the status of this polymorphic type
func (m *batchUserOperationRequest) Status() *string {
	return m.statusField
}

// SetStatus sets the status of this polymorphic type
func (m *batchUserOperationRequest) SetStatus(val *string) {
	m.statusField = val
}

// UnmarshalBatchUserOperationRequestSlice unmarshals polymorphic slices of BatchUserOperationRequest
func UnmarshalBatchUserOperationRequestSlice(reader io.Reader, consumer runtime.Consumer) ([]BatchUserOperationRequest, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []BatchUserOperationRequest
	for _, element := range elements {
		obj, err := unmarshalBatchUserOperationRequest(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalBatchUserOperationRequest unmarshals polymorphic BatchUserOperationRequest
func UnmarshalBatchUserOperationRequest(reader io.Reader, consumer runtime.Consumer) (BatchUserOperationRequest, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalBatchUserOperationRequest(data, consumer)
}

func unmarshalBatchUserOperationRequest(data []byte, consumer runtime.Consumer) (BatchUserOperationRequest, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Batch user operation request. property.
	var getType struct {
		BatchUserOperationRequest string `json:"Batch user operation request."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Batch user operation request.", "body", getType.BatchUserOperationRequest); err != nil {
		return nil, err
	}

	// The value of Batch user operation request. is used to determine which type to create and unmarshal the data into
	switch getType.BatchUserOperationRequest {
	case "BatchUserOperationRequest":
		var result batchUserOperationRequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Batch user operation request. value: %q", getType.BatchUserOperationRequest)
}

// Validate validates this batch user operation request
func (m *batchUserOperationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *batchUserOperationRequest) validateIds(formats strfmt.Registry) error {

	if err := validate.Required("ids", "body", m.Ids()); err != nil {
		return err
	}

	return nil
}

func (m *batchUserOperationRequest) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this batch user operation request based on context it is used
func (m *batchUserOperationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
