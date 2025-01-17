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

// CustomIntegration CustomIntegration
//
// Represents a Custom Integration.
//
// swagger:discriminator CustomIntegration Represents a Custom Integration.
type CustomIntegration interface {
	runtime.Validatable
	runtime.ContextValidatable

	// This field is provided for backward compatibility. Contains the same value as the 'createdAt' field as a UNIX timestamp in microseconds
	// Example: 1568625938000000
	CreateTimeInMicros() int64
	SetCreateTimeInMicros(int64)

	// Partial URL that provides details of the resource.
	// Example: /codestream/api/\u003cprefix\u003e/8365ef3b-8bf3-48aa-bd5d-7113fcff827c
	Link() string
	SetLink(string)

	// Contains project id of the entity
	// Example: abcd-abcd-abcd
	ProjectID() string
	SetProjectID(string)

	// This field is provided for backward compatibility. Contains the same value as the 'updatedAt' field as a UNIX timestamp in microseconds
	// Example: 1568625938000000
	UpdateTimeInMicros() int64
	SetUpdateTimeInMicros(int64)

	// Changes from the previous version.
	// Example: Modified input property.
	ChangeLog() string
	SetChangeLog(string)

	// Date when the entity was created. The date is in ISO 8601 with time zone
	// Example: 2019-09-16 09:25:38.065065+00
	CreatedAt() string
	SetCreatedAt(string)

	// The user that created this entity
	// Example: exampleuser
	CreatedBy() string
	SetCreatedBy(string)

	// A human-friendly description.
	// Required: true
	Description() *string
	SetDescription(*string)

	// The id of this resource.
	// Example: 8365ef3b-8bf3-48aa-bd5d-7113fcff827c
	ID() string
	SetID(string)

	// A human-friendly name used as an identifier in APIs that support this option
	// Example: My-Name
	// Required: true
	Name() *string
	SetName(*string)

	// The id of the parent of this resource.
	// Example: 1abd1fd6-ae2c-459c-ab75-8c595631a11f
	ParentID() string
	SetParentID(string)

	// The project this entity belongs to.
	// Example: My-Project
	Project() string
	SetProject(string)

	// Release status of the Custom Integration
	// Example: NONE/RELEASED/DEPRECATED
	Status() string
	SetStatus(string)

	// Date when the entity was last updated. The date is in ISO 8601 with time zone.
	// Example: 2019-09-16 09:25:38.065065+00
	UpdatedAt() string
	SetUpdatedAt(string)

	// The user that last updated this entity
	// Example: exampleuser
	UpdatedBy() string
	SetUpdatedBy(string)

	// Version of the resource.
	// Example: v1
	Version() string
	SetVersion(string)

	// YAML describing Custom integration details.
	// Example: ---\nruntime: \"nodejs\"\ncode: |\n    var context = require(\"./context.js\")\n    var start = Date.now();\n    var message = context.getInput(\"message\");\n    console.log(\"starting timer is good \" + message);\n    \n    setTimeout(function() {\n      var millis = Date.now() - start;\n      console.log(\"seconds elapsed = \" + Math.floor(millis/1000));\n      context.setOutput(\"time\", millis);\n    }, 2000);\ninputProperties:\n  - name: message\n    type: text\n    title: Message\n    placeHolder: Time\n    defaultValue: \n    bindable: true\n    labelInfo: true\n    labelMessage: What is time\n    \noutputProperties:\n  - name: time\n    type: label\n    title: Time
	Yaml() string
	SetYaml(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type customIntegration struct {
	createTimeInMicrosField int64

	linkField string

	projectIdField string

	updateTimeInMicrosField int64

	changeLogField string

	createdAtField string

	createdByField string

	descriptionField *string

	idField string

	nameField *string

	parentIdField string

	projectField string

	statusField string

	updatedAtField string

	updatedByField string

	versionField string

	yamlField string
}

// CreateTimeInMicros gets the create time in micros of this polymorphic type
func (m *customIntegration) CreateTimeInMicros() int64 {
	return m.createTimeInMicrosField
}

// SetCreateTimeInMicros sets the create time in micros of this polymorphic type
func (m *customIntegration) SetCreateTimeInMicros(val int64) {
	m.createTimeInMicrosField = val
}

// Link gets the link of this polymorphic type
func (m *customIntegration) Link() string {
	return m.linkField
}

// SetLink sets the link of this polymorphic type
func (m *customIntegration) SetLink(val string) {
	m.linkField = val
}

// ProjectID gets the project Id of this polymorphic type
func (m *customIntegration) ProjectID() string {
	return m.projectIdField
}

// SetProjectID sets the project Id of this polymorphic type
func (m *customIntegration) SetProjectID(val string) {
	m.projectIdField = val
}

// UpdateTimeInMicros gets the update time in micros of this polymorphic type
func (m *customIntegration) UpdateTimeInMicros() int64 {
	return m.updateTimeInMicrosField
}

// SetUpdateTimeInMicros sets the update time in micros of this polymorphic type
func (m *customIntegration) SetUpdateTimeInMicros(val int64) {
	m.updateTimeInMicrosField = val
}

// ChangeLog gets the change log of this polymorphic type
func (m *customIntegration) ChangeLog() string {
	return m.changeLogField
}

// SetChangeLog sets the change log of this polymorphic type
func (m *customIntegration) SetChangeLog(val string) {
	m.changeLogField = val
}

// CreatedAt gets the created at of this polymorphic type
func (m *customIntegration) CreatedAt() string {
	return m.createdAtField
}

// SetCreatedAt sets the created at of this polymorphic type
func (m *customIntegration) SetCreatedAt(val string) {
	m.createdAtField = val
}

// CreatedBy gets the created by of this polymorphic type
func (m *customIntegration) CreatedBy() string {
	return m.createdByField
}

// SetCreatedBy sets the created by of this polymorphic type
func (m *customIntegration) SetCreatedBy(val string) {
	m.createdByField = val
}

// Description gets the description of this polymorphic type
func (m *customIntegration) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *customIntegration) SetDescription(val *string) {
	m.descriptionField = val
}

// ID gets the id of this polymorphic type
func (m *customIntegration) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *customIntegration) SetID(val string) {
	m.idField = val
}

// Name gets the name of this polymorphic type
func (m *customIntegration) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *customIntegration) SetName(val *string) {
	m.nameField = val
}

// ParentID gets the parent Id of this polymorphic type
func (m *customIntegration) ParentID() string {
	return m.parentIdField
}

// SetParentID sets the parent Id of this polymorphic type
func (m *customIntegration) SetParentID(val string) {
	m.parentIdField = val
}

// Project gets the project of this polymorphic type
func (m *customIntegration) Project() string {
	return m.projectField
}

// SetProject sets the project of this polymorphic type
func (m *customIntegration) SetProject(val string) {
	m.projectField = val
}

// Status gets the status of this polymorphic type
func (m *customIntegration) Status() string {
	return m.statusField
}

// SetStatus sets the status of this polymorphic type
func (m *customIntegration) SetStatus(val string) {
	m.statusField = val
}

// UpdatedAt gets the updated at of this polymorphic type
func (m *customIntegration) UpdatedAt() string {
	return m.updatedAtField
}

// SetUpdatedAt sets the updated at of this polymorphic type
func (m *customIntegration) SetUpdatedAt(val string) {
	m.updatedAtField = val
}

// UpdatedBy gets the updated by of this polymorphic type
func (m *customIntegration) UpdatedBy() string {
	return m.updatedByField
}

// SetUpdatedBy sets the updated by of this polymorphic type
func (m *customIntegration) SetUpdatedBy(val string) {
	m.updatedByField = val
}

// Version gets the version of this polymorphic type
func (m *customIntegration) Version() string {
	return m.versionField
}

// SetVersion sets the version of this polymorphic type
func (m *customIntegration) SetVersion(val string) {
	m.versionField = val
}

// Yaml gets the yaml of this polymorphic type
func (m *customIntegration) Yaml() string {
	return m.yamlField
}

// SetYaml sets the yaml of this polymorphic type
func (m *customIntegration) SetYaml(val string) {
	m.yamlField = val
}

// UnmarshalCustomIntegrationSlice unmarshals polymorphic slices of CustomIntegration
func UnmarshalCustomIntegrationSlice(reader io.Reader, consumer runtime.Consumer) ([]CustomIntegration, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []CustomIntegration
	for _, element := range elements {
		obj, err := unmarshalCustomIntegration(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalCustomIntegration unmarshals polymorphic CustomIntegration
func UnmarshalCustomIntegration(reader io.Reader, consumer runtime.Consumer) (CustomIntegration, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalCustomIntegration(data, consumer)
}

func unmarshalCustomIntegration(data []byte, consumer runtime.Consumer) (CustomIntegration, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Represents a Custom Integration. property.
	var getType struct {
		RepresentsaCustomIntegration string `json:"Represents a Custom Integration."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Represents a Custom Integration.", "body", getType.RepresentsaCustomIntegration); err != nil {
		return nil, err
	}

	// The value of Represents a Custom Integration. is used to determine which type to create and unmarshal the data into
	switch getType.RepresentsaCustomIntegration {
	case "CustomIntegration":
		var result customIntegration
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Represents a Custom Integration. value: %q", getType.RepresentsaCustomIntegration)
}

// Validate validates this custom integration
func (m *customIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *customIntegration) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *customIntegration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom integration based on context it is used
func (m *customIntegration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
