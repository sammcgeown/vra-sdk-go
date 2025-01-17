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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StageExecution StageExecution
//
// Contains Stage Execution details.
//
// swagger:discriminator StageExecution Contains Stage Execution details.
type StageExecution interface {
	runtime.Validatable
	runtime.ContextValidatable

	// Execution duration of the Stage Execution (in micros).
	// Example: 1568625938000000
	DurationInMicros() int64
	SetDurationInMicros(int64)

	// End time of the Stage Execution (in micros)..
	// Example: 1568625938000000
	EndTime() int64
	SetEndTime(int64)

	// Start time of the Stage Execution (in micros)..
	// Example: 1568625938000000
	StartTime() int64
	SetStartTime(int64)

	// The id of this Stage.
	// Example: 8365ef3b-8bf3-48aa-bd5d-7113fcff827c
	ID() string
	SetID(string)

	// A human-friendly name used as an identifier For the Stage.
	// Example: My-Name
	// Required: true
	Name() *string
	SetName(*string)

	Notifications() []Notification
	SetNotifications([]Notification)

	RollbackConfiguration() RollbackConfiguration
	SetRollbackConfiguration(RollbackConfiguration)

	RollbackResponse() RollbackResponse
	SetRollbackResponse(RollbackResponse)

	// Execution status of the Stage Execution.
	// Example: COMPLETED
	// Enum: [NOT_STARTED STARTED RUNNING CANCELING WAITING RESUMING PAUSING PAUSED CANCELED COMPLETED FAILED SKIPPED QUEUED FAILED_CONTINUE ROLLING_BACK ROLLBACK_FAILED PREPARING_WORKSPACE ROLLBACK_COMPLETED]
	Status() string
	SetStatus(string)

	// Execution status message of the Stage Execution.
	// Example: Executing Stage0
	StatusMessage() string
	SetStatusMessage(string)

	// Ordering of the various Tasks within the Stage.
	// Example: \"taskOrder\": [\n        \"Test1,Test2\",\n        \"Test3\"\n      ],
	TaskOrder() []string
	SetTaskOrder([]string)

	// Represents the various Tasks in the Stage.
	Tasks() map[string]TaskExecution
	SetTasks(map[string]TaskExecution)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type stageExecution struct {
	durationInMicrosField int64

	endTimeField int64

	startTimeField int64

	idField string

	nameField *string

	notificationsField []Notification

	rollbackConfigurationField RollbackConfiguration

	rollbackResponseField RollbackResponse

	statusField string

	statusMessageField string

	taskOrderField []string

	tasksField map[string]TaskExecution
}

// DurationInMicros gets the duration in micros of this polymorphic type
func (m *stageExecution) DurationInMicros() int64 {
	return m.durationInMicrosField
}

// SetDurationInMicros sets the duration in micros of this polymorphic type
func (m *stageExecution) SetDurationInMicros(val int64) {
	m.durationInMicrosField = val
}

// EndTime gets the end time of this polymorphic type
func (m *stageExecution) EndTime() int64 {
	return m.endTimeField
}

// SetEndTime sets the end time of this polymorphic type
func (m *stageExecution) SetEndTime(val int64) {
	m.endTimeField = val
}

// StartTime gets the start time of this polymorphic type
func (m *stageExecution) StartTime() int64 {
	return m.startTimeField
}

// SetStartTime sets the start time of this polymorphic type
func (m *stageExecution) SetStartTime(val int64) {
	m.startTimeField = val
}

// ID gets the id of this polymorphic type
func (m *stageExecution) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *stageExecution) SetID(val string) {
	m.idField = val
}

// Name gets the name of this polymorphic type
func (m *stageExecution) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *stageExecution) SetName(val *string) {
	m.nameField = val
}

// Notifications gets the notifications of this polymorphic type
func (m *stageExecution) Notifications() []Notification {
	return m.notificationsField
}

// SetNotifications sets the notifications of this polymorphic type
func (m *stageExecution) SetNotifications(val []Notification) {
	m.notificationsField = val
}

// RollbackConfiguration gets the rollback configuration of this polymorphic type
func (m *stageExecution) RollbackConfiguration() RollbackConfiguration {
	return m.rollbackConfigurationField
}

// SetRollbackConfiguration sets the rollback configuration of this polymorphic type
func (m *stageExecution) SetRollbackConfiguration(val RollbackConfiguration) {
	m.rollbackConfigurationField = val
}

// RollbackResponse gets the rollback response of this polymorphic type
func (m *stageExecution) RollbackResponse() RollbackResponse {
	return m.rollbackResponseField
}

// SetRollbackResponse sets the rollback response of this polymorphic type
func (m *stageExecution) SetRollbackResponse(val RollbackResponse) {
	m.rollbackResponseField = val
}

// Status gets the status of this polymorphic type
func (m *stageExecution) Status() string {
	return m.statusField
}

// SetStatus sets the status of this polymorphic type
func (m *stageExecution) SetStatus(val string) {
	m.statusField = val
}

// StatusMessage gets the status message of this polymorphic type
func (m *stageExecution) StatusMessage() string {
	return m.statusMessageField
}

// SetStatusMessage sets the status message of this polymorphic type
func (m *stageExecution) SetStatusMessage(val string) {
	m.statusMessageField = val
}

// TaskOrder gets the task order of this polymorphic type
func (m *stageExecution) TaskOrder() []string {
	return m.taskOrderField
}

// SetTaskOrder sets the task order of this polymorphic type
func (m *stageExecution) SetTaskOrder(val []string) {
	m.taskOrderField = val
}

// Tasks gets the tasks of this polymorphic type
func (m *stageExecution) Tasks() map[string]TaskExecution {
	return m.tasksField
}

// SetTasks sets the tasks of this polymorphic type
func (m *stageExecution) SetTasks(val map[string]TaskExecution) {
	m.tasksField = val
}

// UnmarshalStageExecutionSlice unmarshals polymorphic slices of StageExecution
func UnmarshalStageExecutionSlice(reader io.Reader, consumer runtime.Consumer) ([]StageExecution, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []StageExecution
	for _, element := range elements {
		obj, err := unmarshalStageExecution(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalStageExecution unmarshals polymorphic StageExecution
func UnmarshalStageExecution(reader io.Reader, consumer runtime.Consumer) (StageExecution, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalStageExecution(data, consumer)
}

func unmarshalStageExecution(data []byte, consumer runtime.Consumer) (StageExecution, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Contains Stage Execution details. property.
	var getType struct {
		ContainsStageExecutionDetails string `json:"Contains Stage Execution details."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Contains Stage Execution details.", "body", getType.ContainsStageExecutionDetails); err != nil {
		return nil, err
	}

	// The value of Contains Stage Execution details. is used to determine which type to create and unmarshal the data into
	switch getType.ContainsStageExecutionDetails {
	case "StageExecution":
		var result stageExecution
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Contains Stage Execution details. value: %q", getType.ContainsStageExecutionDetails)
}

// Validate validates this stage execution
func (m *stageExecution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRollbackConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRollbackResponse(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTasks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *stageExecution) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *stageExecution) validateNotifications(formats strfmt.Registry) error {
	if swag.IsZero(m.Notifications()) { // not required
		return nil
	}

	for i := 0; i < len(m.Notifications()); i++ {

		if err := m.notificationsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *stageExecution) validateRollbackConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.RollbackConfiguration()) { // not required
		return nil
	}

	if err := m.RollbackConfiguration().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rollbackConfiguration")
		}
		return err
	}

	return nil
}

func (m *stageExecution) validateRollbackResponse(formats strfmt.Registry) error {
	if swag.IsZero(m.RollbackResponse()) { // not required
		return nil
	}

	if err := m.RollbackResponse().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rollbackResponse")
		}
		return err
	}

	return nil
}

var stageExecutionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_STARTED","STARTED","RUNNING","CANCELING","WAITING","RESUMING","PAUSING","PAUSED","CANCELED","COMPLETED","FAILED","SKIPPED","QUEUED","FAILED_CONTINUE","ROLLING_BACK","ROLLBACK_FAILED","PREPARING_WORKSPACE","ROLLBACK_COMPLETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stageExecutionTypeStatusPropEnum = append(stageExecutionTypeStatusPropEnum, v)
	}
}

const (

	// StageExecutionStatusNOTSTARTED captures enum value "NOT_STARTED"
	StageExecutionStatusNOTSTARTED string = "NOT_STARTED"

	// StageExecutionStatusSTARTED captures enum value "STARTED"
	StageExecutionStatusSTARTED string = "STARTED"

	// StageExecutionStatusRUNNING captures enum value "RUNNING"
	StageExecutionStatusRUNNING string = "RUNNING"

	// StageExecutionStatusCANCELING captures enum value "CANCELING"
	StageExecutionStatusCANCELING string = "CANCELING"

	// StageExecutionStatusWAITING captures enum value "WAITING"
	StageExecutionStatusWAITING string = "WAITING"

	// StageExecutionStatusRESUMING captures enum value "RESUMING"
	StageExecutionStatusRESUMING string = "RESUMING"

	// StageExecutionStatusPAUSING captures enum value "PAUSING"
	StageExecutionStatusPAUSING string = "PAUSING"

	// StageExecutionStatusPAUSED captures enum value "PAUSED"
	StageExecutionStatusPAUSED string = "PAUSED"

	// StageExecutionStatusCANCELED captures enum value "CANCELED"
	StageExecutionStatusCANCELED string = "CANCELED"

	// StageExecutionStatusCOMPLETED captures enum value "COMPLETED"
	StageExecutionStatusCOMPLETED string = "COMPLETED"

	// StageExecutionStatusFAILED captures enum value "FAILED"
	StageExecutionStatusFAILED string = "FAILED"

	// StageExecutionStatusSKIPPED captures enum value "SKIPPED"
	StageExecutionStatusSKIPPED string = "SKIPPED"

	// StageExecutionStatusQUEUED captures enum value "QUEUED"
	StageExecutionStatusQUEUED string = "QUEUED"

	// StageExecutionStatusFAILEDCONTINUE captures enum value "FAILED_CONTINUE"
	StageExecutionStatusFAILEDCONTINUE string = "FAILED_CONTINUE"

	// StageExecutionStatusROLLINGBACK captures enum value "ROLLING_BACK"
	StageExecutionStatusROLLINGBACK string = "ROLLING_BACK"

	// StageExecutionStatusROLLBACKFAILED captures enum value "ROLLBACK_FAILED"
	StageExecutionStatusROLLBACKFAILED string = "ROLLBACK_FAILED"

	// StageExecutionStatusPREPARINGWORKSPACE captures enum value "PREPARING_WORKSPACE"
	StageExecutionStatusPREPARINGWORKSPACE string = "PREPARING_WORKSPACE"

	// StageExecutionStatusROLLBACKCOMPLETED captures enum value "ROLLBACK_COMPLETED"
	StageExecutionStatusROLLBACKCOMPLETED string = "ROLLBACK_COMPLETED"
)

// prop value enum
func (m *stageExecution) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stageExecutionTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *stageExecution) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status()) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status()); err != nil {
		return err
	}

	return nil
}

func (m *stageExecution) validateTasks(formats strfmt.Registry) error {
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

// ContextValidate validate this stage execution based on the context it is used
func (m *stageExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRollbackConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRollbackResponse(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTasks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *stageExecution) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Notifications()); i++ {

		if err := m.notificationsField[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notifications" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *stageExecution) contextValidateRollbackConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RollbackConfiguration().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rollbackConfiguration")
		}
		return err
	}

	return nil
}

func (m *stageExecution) contextValidateRollbackResponse(ctx context.Context, formats strfmt.Registry) error {

	if err := m.RollbackResponse().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rollbackResponse")
		}
		return err
	}

	return nil
}

func (m *stageExecution) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Tasks() {

		if val, ok := m.Tasks()[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}
