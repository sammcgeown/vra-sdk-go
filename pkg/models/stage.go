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

// Stage Stage
//
// Data type for the Stage.
//
// swagger:discriminator Stage Data type for the Stage.
type Stage interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Stage description.
	// Example: Image Publish Stage
	Description() string
	SetDescription(string)

	// Tags are useful for ease in searching, grouping or filtering.
	// Example: ["Tag1","Tag2"]
	Tags() []string
	SetTags([]string)

	// Order in which tasks will be executed.
	// Example: ["Task1","Task0"]
	TaskOrder() []string
	SetTaskOrder([]string)

	// Map representing details of various tasks present in the stage.
	// Example: {"Task1":{"ignoreFailure":false,"input":{"action":"get","headers":{"Accept":"application/json","Content-Type":"application/json"},"url":"https://www.vmware.com"},"preCondition":"","type":"REST"}}
	Tasks() map[string]Task
	SetTasks(map[string]Task)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type stage struct {
	descriptionField string

	tagsField []string

	taskOrderField []string

	tasksField map[string]Task
}

// Description gets the description of this polymorphic type
func (m *stage) Description() string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *stage) SetDescription(val string) {
	m.descriptionField = val
}

// Tags gets the tags of this polymorphic type
func (m *stage) Tags() []string {
	return m.tagsField
}

// SetTags sets the tags of this polymorphic type
func (m *stage) SetTags(val []string) {
	m.tagsField = val
}

// TaskOrder gets the task order of this polymorphic type
func (m *stage) TaskOrder() []string {
	return m.taskOrderField
}

// SetTaskOrder sets the task order of this polymorphic type
func (m *stage) SetTaskOrder(val []string) {
	m.taskOrderField = val
}

// Tasks gets the tasks of this polymorphic type
func (m *stage) Tasks() map[string]Task {
	return m.tasksField
}

// SetTasks sets the tasks of this polymorphic type
func (m *stage) SetTasks(val map[string]Task) {
	m.tasksField = val
}

// UnmarshalStageSlice unmarshals polymorphic slices of Stage
func UnmarshalStageSlice(reader io.Reader, consumer runtime.Consumer) ([]Stage, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Stage
	for _, element := range elements {
		obj, err := unmarshalStage(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalStage unmarshals polymorphic Stage
func UnmarshalStage(reader io.Reader, consumer runtime.Consumer) (Stage, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalStage(data, consumer)
}

func unmarshalStage(data []byte, consumer runtime.Consumer) (Stage, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Data type for the Stage. property.
	var getType struct {
		DataTypeForTheStage string `json:"Data type for the Stage."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Data type for the Stage.", "body", getType.DataTypeForTheStage); err != nil {
		return nil, err
	}

	// The value of Data type for the Stage. is used to determine which type to create and unmarshal the data into
	switch getType.DataTypeForTheStage {
	case "Stage":
		var result stage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Data type for the Stage. value: %q", getType.DataTypeForTheStage)
}

// Validate validates this stage
func (m *stage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *stage) validateTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.Tasks()) { // not required
		return nil
	}

	for k := range m.Tasks() {

		if err := validate.Required("tasks"+"."+k, "body", m.Tasks()[k]); err != nil {
			return err
		}
		if val, ok := m.Tasks()[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this stage based on the context it is used
func (m *stage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *stage) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Tasks() {

		if val, ok := m.Tasks()[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}