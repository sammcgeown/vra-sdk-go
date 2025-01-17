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
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DockerRegistryEvents DockerRegistryEvents
//
// List of Docker Registry Events.
//
// swagger:discriminator DockerRegistryEvents List of Docker Registry Events.
type DockerRegistryEvents interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Number of resources within the current page.
	Count() int64
	SetCount(int64)

	// Details of the queried resources.
	// Example: \"documents\": {\n        \"/codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337\": {\n            \"project\": \"test-project\",\n            \"id\": \"b80254a7-fcff-4918-ad88-501d08096337\",\n            \"name\": \"test\",\n            \"updatedAt\": \"2019-09-23 13:48:54.483\",\n            \"tags\": [],\n            \"_link\": \"/codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337\",\n            \"_updateTimeInMicros\": 1569226734483000,\n            \"_createTimeInMicros\": 1569226720988000,\n            \"index\": 1,\n            \"notifications\": [],\n            \"comments\": \"\",\n            \"starred\": {},\n            \"input\": {},\n            \"output\": {},\n            \"stageOrder\": [],\n            \"stages\": {},\n            \"status\": \"COMPLETED\",\n            \"statusMessage\": \"Execution Completed.\",\n            \"_durationInMicros\": 13495000,\n            \"_requestTimeInMicros\": 1569226720988000,\n            \"_executedBy\": \"exampleuser\",\n            \"_pipelineLink\": \"/codestream/api/pipelines/b49898f9-d42d-4f19-8bda-e77a373c41b9\",\n            \"_nested\": false,\n            \"_rollback\": false,\n            \"_inputMeta\": {},\n            \"_outputMeta\": {},\n            \"workspaceResults\": []\n        }\n    }
	Documents() map[string]DockerRegistryEvent
	SetDocuments(map[string]DockerRegistryEvent)

	// Partial URLs representing the links to the queried resources.
	// Example: /codestream/api/executions/b80254a7-fcff-4918-ad88-501d08096337
	Links() []string
	SetLinks([]string)

	// Number of resources across all pages.
	TotalCount() int64
	SetTotalCount(int64)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type dockerRegistryEvents struct {
	countField int64

	documentsField map[string]DockerRegistryEvent

	linksField []string

	totalCountField int64
}

// Count gets the count of this polymorphic type
func (m *dockerRegistryEvents) Count() int64 {
	return m.countField
}

// SetCount sets the count of this polymorphic type
func (m *dockerRegistryEvents) SetCount(val int64) {
	m.countField = val
}

// Documents gets the documents of this polymorphic type
func (m *dockerRegistryEvents) Documents() map[string]DockerRegistryEvent {
	return m.documentsField
}

// SetDocuments sets the documents of this polymorphic type
func (m *dockerRegistryEvents) SetDocuments(val map[string]DockerRegistryEvent) {
	m.documentsField = val
}

// Links gets the links of this polymorphic type
func (m *dockerRegistryEvents) Links() []string {
	return m.linksField
}

// SetLinks sets the links of this polymorphic type
func (m *dockerRegistryEvents) SetLinks(val []string) {
	m.linksField = val
}

// TotalCount gets the total count of this polymorphic type
func (m *dockerRegistryEvents) TotalCount() int64 {
	return m.totalCountField
}

// SetTotalCount sets the total count of this polymorphic type
func (m *dockerRegistryEvents) SetTotalCount(val int64) {
	m.totalCountField = val
}

// UnmarshalDockerRegistryEventsSlice unmarshals polymorphic slices of DockerRegistryEvents
func UnmarshalDockerRegistryEventsSlice(reader io.Reader, consumer runtime.Consumer) ([]DockerRegistryEvents, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []DockerRegistryEvents
	for _, element := range elements {
		obj, err := unmarshalDockerRegistryEvents(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalDockerRegistryEvents unmarshals polymorphic DockerRegistryEvents
func UnmarshalDockerRegistryEvents(reader io.Reader, consumer runtime.Consumer) (DockerRegistryEvents, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalDockerRegistryEvents(data, consumer)
}

func unmarshalDockerRegistryEvents(data []byte, consumer runtime.Consumer) (DockerRegistryEvents, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the List of Docker Registry Events. property.
	var getType struct {
		ListOfDockerRegistryEvents string `json:"List of Docker Registry Events."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("List of Docker Registry Events.", "body", getType.ListOfDockerRegistryEvents); err != nil {
		return nil, err
	}

	// The value of List of Docker Registry Events. is used to determine which type to create and unmarshal the data into
	switch getType.ListOfDockerRegistryEvents {
	case "DockerRegistryEvents":
		var result dockerRegistryEvents
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid List of Docker Registry Events. value: %q", getType.ListOfDockerRegistryEvents)
}

// Validate validates this docker registry events
func (m *dockerRegistryEvents) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDocuments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *dockerRegistryEvents) validateDocuments(formats strfmt.Registry) error {
	if swag.IsZero(m.Documents()) { // not required
		return nil
	}

	for k := range m.Documents() {

		if err := validate.Required("documents"+"."+k, "body", m.Documents()[k]); err != nil {
			return err
		}
		if val, ok := m.Documents()[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this docker registry events based on the context it is used
func (m *dockerRegistryEvents) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDocuments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *dockerRegistryEvents) contextValidateDocuments(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Documents() {

		if val, ok := m.Documents()[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}
