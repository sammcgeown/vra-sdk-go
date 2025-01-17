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

// Execution Execution
//
// Execution
//
// swagger:discriminator Execution Execution
type Execution interface {
	runtime.Validatable
	runtime.ContextValidatable

	// This field is provided for backward compatibility. Contains the same value as the 'createdAt' field as a UNIX timestamp in microseconds
	// Example: 1568625938000000
	CreateTimeInMicros() int64
	SetCreateTimeInMicros(int64)

	// duration in micros
	DurationInMicros() int64
	SetDurationInMicros(int64)

	// executed by
	ExecutedBy() string
	SetExecutedBy(string)

	// input meta
	InputMeta() map[string]PropertyMetaData
	SetInputMeta(map[string]PropertyMetaData)

	// Partial URL that provides details of the resource.
	// Example: /codestream/api/\u003cprefix\u003e/8365ef3b-8bf3-48aa-bd5d-7113fcff827c
	Link() string
	SetLink(string)

	// nested
	Nested() bool
	SetNested(bool)

	// output meta
	OutputMeta() map[string]PropertyMetaData
	SetOutputMeta(map[string]PropertyMetaData)

	// pipeline link
	PipelineLink() string
	SetPipelineLink(string)

	// Contains project id of the entity
	// Example: abcd-abcd-abcd
	ProjectID() string
	SetProjectID(string)

	// request time in micros
	RequestTimeInMicros() int64
	SetRequestTimeInMicros(int64)

	// rollback
	Rollback() bool
	SetRollback(bool)

	// source
	Source() string
	SetSource(string)

	// total duration in micros
	TotalDurationInMicros() int64
	SetTotalDurationInMicros(int64)

	// triggered by
	TriggeredBy() string
	SetTriggeredBy(string)

	// This field is provided for backward compatibility. Contains the same value as the 'updatedAt' field as a UNIX timestamp in microseconds
	// Example: 1568625938000000
	UpdateTimeInMicros() int64
	SetUpdateTimeInMicros(int64)

	// comments
	Comments() string
	SetComments(string)

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

	// icon
	Icon() string
	SetIcon(string)

	// The id of this resource.
	// Example: 8365ef3b-8bf3-48aa-bd5d-7113fcff827c
	ID() string
	SetID(string)

	// index
	Index() int64
	SetIndex(int64)

	// input
	Input() interface{}
	SetInput(interface{})

	// A human-friendly name used as an identifier in APIs that support this option
	// Example: My-Name
	// Required: true
	Name() *string
	SetName(*string)

	Notifications() []Notification
	SetNotifications([]Notification)

	// output
	Output() interface{}
	SetOutput(interface{})

	// The project this entity belongs to.
	// Example: My-Project
	Project() string
	SetProject(string)

	// reason
	Reason() string
	SetReason(string)

	// resumed at
	ResumedAt() string
	SetResumedAt(string)

	// stage order
	StageOrder() []string
	SetStageOrder([]string)

	// stages
	Stages() map[string]StageExecution
	SetStages(map[string]StageExecution)

	Starred() PipelineStarredProperty
	SetStarred(PipelineStarredProperty)

	// status
	// Enum: [NOT_STARTED STARTED RUNNING CANCELING WAITING RESUMING PAUSING PAUSED CANCELED COMPLETED FAILED SKIPPED QUEUED FAILED_CONTINUE ROLLING_BACK ROLLBACK_FAILED PREPARING_WORKSPACE ROLLBACK_COMPLETED]
	Status() string
	SetStatus(string)

	// status message
	StatusMessage() string
	SetStatusMessage(string)

	// A set of tag keys and optional values that were set on on the resource.
	// Example: [{"key":"env","value":"dev"}]
	Tags() []string
	SetTags([]string)

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

	// workspace results
	WorkspaceResults() []*WorkspaceResult
	SetWorkspaceResults([]*WorkspaceResult)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type execution struct {
	createTimeInMicrosField int64

	durationInMicrosField int64

	executedByField string

	inputMetaField map[string]PropertyMetaData

	linkField string

	nestedField bool

	outputMetaField map[string]PropertyMetaData

	pipelineLinkField string

	projectIdField string

	requestTimeInMicrosField int64

	rollbackField bool

	sourceField string

	totalDurationInMicrosField int64

	triggeredByField string

	updateTimeInMicrosField int64

	commentsField string

	createdAtField string

	createdByField string

	descriptionField *string

	iconField string

	idField string

	indexField int64

	inputField interface{}

	nameField *string

	notificationsField []Notification

	outputField interface{}

	projectField string

	reasonField string

	resumedAtField string

	stageOrderField []string

	stagesField map[string]StageExecution

	starredField PipelineStarredProperty

	statusField string

	statusMessageField string

	tagsField []string

	updatedAtField string

	updatedByField string

	versionField string

	workspaceResultsField []*WorkspaceResult
}

// CreateTimeInMicros gets the create time in micros of this polymorphic type
func (m *execution) CreateTimeInMicros() int64 {
	return m.createTimeInMicrosField
}

// SetCreateTimeInMicros sets the create time in micros of this polymorphic type
func (m *execution) SetCreateTimeInMicros(val int64) {
	m.createTimeInMicrosField = val
}

// DurationInMicros gets the duration in micros of this polymorphic type
func (m *execution) DurationInMicros() int64 {
	return m.durationInMicrosField
}

// SetDurationInMicros sets the duration in micros of this polymorphic type
func (m *execution) SetDurationInMicros(val int64) {
	m.durationInMicrosField = val
}

// ExecutedBy gets the executed by of this polymorphic type
func (m *execution) ExecutedBy() string {
	return m.executedByField
}

// SetExecutedBy sets the executed by of this polymorphic type
func (m *execution) SetExecutedBy(val string) {
	m.executedByField = val
}

// InputMeta gets the input meta of this polymorphic type
func (m *execution) InputMeta() map[string]PropertyMetaData {
	return m.inputMetaField
}

// SetInputMeta sets the input meta of this polymorphic type
func (m *execution) SetInputMeta(val map[string]PropertyMetaData) {
	m.inputMetaField = val
}

// Link gets the link of this polymorphic type
func (m *execution) Link() string {
	return m.linkField
}

// SetLink sets the link of this polymorphic type
func (m *execution) SetLink(val string) {
	m.linkField = val
}

// Nested gets the nested of this polymorphic type
func (m *execution) Nested() bool {
	return m.nestedField
}

// SetNested sets the nested of this polymorphic type
func (m *execution) SetNested(val bool) {
	m.nestedField = val
}

// OutputMeta gets the output meta of this polymorphic type
func (m *execution) OutputMeta() map[string]PropertyMetaData {
	return m.outputMetaField
}

// SetOutputMeta sets the output meta of this polymorphic type
func (m *execution) SetOutputMeta(val map[string]PropertyMetaData) {
	m.outputMetaField = val
}

// PipelineLink gets the pipeline link of this polymorphic type
func (m *execution) PipelineLink() string {
	return m.pipelineLinkField
}

// SetPipelineLink sets the pipeline link of this polymorphic type
func (m *execution) SetPipelineLink(val string) {
	m.pipelineLinkField = val
}

// ProjectID gets the project Id of this polymorphic type
func (m *execution) ProjectID() string {
	return m.projectIdField
}

// SetProjectID sets the project Id of this polymorphic type
func (m *execution) SetProjectID(val string) {
	m.projectIdField = val
}

// RequestTimeInMicros gets the request time in micros of this polymorphic type
func (m *execution) RequestTimeInMicros() int64 {
	return m.requestTimeInMicrosField
}

// SetRequestTimeInMicros sets the request time in micros of this polymorphic type
func (m *execution) SetRequestTimeInMicros(val int64) {
	m.requestTimeInMicrosField = val
}

// Rollback gets the rollback of this polymorphic type
func (m *execution) Rollback() bool {
	return m.rollbackField
}

// SetRollback sets the rollback of this polymorphic type
func (m *execution) SetRollback(val bool) {
	m.rollbackField = val
}

// Source gets the source of this polymorphic type
func (m *execution) Source() string {
	return m.sourceField
}

// SetSource sets the source of this polymorphic type
func (m *execution) SetSource(val string) {
	m.sourceField = val
}

// TotalDurationInMicros gets the total duration in micros of this polymorphic type
func (m *execution) TotalDurationInMicros() int64 {
	return m.totalDurationInMicrosField
}

// SetTotalDurationInMicros sets the total duration in micros of this polymorphic type
func (m *execution) SetTotalDurationInMicros(val int64) {
	m.totalDurationInMicrosField = val
}

// TriggeredBy gets the triggered by of this polymorphic type
func (m *execution) TriggeredBy() string {
	return m.triggeredByField
}

// SetTriggeredBy sets the triggered by of this polymorphic type
func (m *execution) SetTriggeredBy(val string) {
	m.triggeredByField = val
}

// UpdateTimeInMicros gets the update time in micros of this polymorphic type
func (m *execution) UpdateTimeInMicros() int64 {
	return m.updateTimeInMicrosField
}

// SetUpdateTimeInMicros sets the update time in micros of this polymorphic type
func (m *execution) SetUpdateTimeInMicros(val int64) {
	m.updateTimeInMicrosField = val
}

// Comments gets the comments of this polymorphic type
func (m *execution) Comments() string {
	return m.commentsField
}

// SetComments sets the comments of this polymorphic type
func (m *execution) SetComments(val string) {
	m.commentsField = val
}

// CreatedAt gets the created at of this polymorphic type
func (m *execution) CreatedAt() string {
	return m.createdAtField
}

// SetCreatedAt sets the created at of this polymorphic type
func (m *execution) SetCreatedAt(val string) {
	m.createdAtField = val
}

// CreatedBy gets the created by of this polymorphic type
func (m *execution) CreatedBy() string {
	return m.createdByField
}

// SetCreatedBy sets the created by of this polymorphic type
func (m *execution) SetCreatedBy(val string) {
	m.createdByField = val
}

// Description gets the description of this polymorphic type
func (m *execution) Description() *string {
	return m.descriptionField
}

// SetDescription sets the description of this polymorphic type
func (m *execution) SetDescription(val *string) {
	m.descriptionField = val
}

// Icon gets the icon of this polymorphic type
func (m *execution) Icon() string {
	return m.iconField
}

// SetIcon sets the icon of this polymorphic type
func (m *execution) SetIcon(val string) {
	m.iconField = val
}

// ID gets the id of this polymorphic type
func (m *execution) ID() string {
	return m.idField
}

// SetID sets the id of this polymorphic type
func (m *execution) SetID(val string) {
	m.idField = val
}

// Index gets the index of this polymorphic type
func (m *execution) Index() int64 {
	return m.indexField
}

// SetIndex sets the index of this polymorphic type
func (m *execution) SetIndex(val int64) {
	m.indexField = val
}

// Input gets the input of this polymorphic type
func (m *execution) Input() interface{} {
	return m.inputField
}

// SetInput sets the input of this polymorphic type
func (m *execution) SetInput(val interface{}) {
	m.inputField = val
}

// Name gets the name of this polymorphic type
func (m *execution) Name() *string {
	return m.nameField
}

// SetName sets the name of this polymorphic type
func (m *execution) SetName(val *string) {
	m.nameField = val
}

// Notifications gets the notifications of this polymorphic type
func (m *execution) Notifications() []Notification {
	return m.notificationsField
}

// SetNotifications sets the notifications of this polymorphic type
func (m *execution) SetNotifications(val []Notification) {
	m.notificationsField = val
}

// Output gets the output of this polymorphic type
func (m *execution) Output() interface{} {
	return m.outputField
}

// SetOutput sets the output of this polymorphic type
func (m *execution) SetOutput(val interface{}) {
	m.outputField = val
}

// Project gets the project of this polymorphic type
func (m *execution) Project() string {
	return m.projectField
}

// SetProject sets the project of this polymorphic type
func (m *execution) SetProject(val string) {
	m.projectField = val
}

// Reason gets the reason of this polymorphic type
func (m *execution) Reason() string {
	return m.reasonField
}

// SetReason sets the reason of this polymorphic type
func (m *execution) SetReason(val string) {
	m.reasonField = val
}

// ResumedAt gets the resumed at of this polymorphic type
func (m *execution) ResumedAt() string {
	return m.resumedAtField
}

// SetResumedAt sets the resumed at of this polymorphic type
func (m *execution) SetResumedAt(val string) {
	m.resumedAtField = val
}

// StageOrder gets the stage order of this polymorphic type
func (m *execution) StageOrder() []string {
	return m.stageOrderField
}

// SetStageOrder sets the stage order of this polymorphic type
func (m *execution) SetStageOrder(val []string) {
	m.stageOrderField = val
}

// Stages gets the stages of this polymorphic type
func (m *execution) Stages() map[string]StageExecution {
	return m.stagesField
}

// SetStages sets the stages of this polymorphic type
func (m *execution) SetStages(val map[string]StageExecution) {
	m.stagesField = val
}

// Starred gets the starred of this polymorphic type
func (m *execution) Starred() PipelineStarredProperty {
	return m.starredField
}

// SetStarred sets the starred of this polymorphic type
func (m *execution) SetStarred(val PipelineStarredProperty) {
	m.starredField = val
}

// Status gets the status of this polymorphic type
func (m *execution) Status() string {
	return m.statusField
}

// SetStatus sets the status of this polymorphic type
func (m *execution) SetStatus(val string) {
	m.statusField = val
}

// StatusMessage gets the status message of this polymorphic type
func (m *execution) StatusMessage() string {
	return m.statusMessageField
}

// SetStatusMessage sets the status message of this polymorphic type
func (m *execution) SetStatusMessage(val string) {
	m.statusMessageField = val
}

// Tags gets the tags of this polymorphic type
func (m *execution) Tags() []string {
	return m.tagsField
}

// SetTags sets the tags of this polymorphic type
func (m *execution) SetTags(val []string) {
	m.tagsField = val
}

// UpdatedAt gets the updated at of this polymorphic type
func (m *execution) UpdatedAt() string {
	return m.updatedAtField
}

// SetUpdatedAt sets the updated at of this polymorphic type
func (m *execution) SetUpdatedAt(val string) {
	m.updatedAtField = val
}

// UpdatedBy gets the updated by of this polymorphic type
func (m *execution) UpdatedBy() string {
	return m.updatedByField
}

// SetUpdatedBy sets the updated by of this polymorphic type
func (m *execution) SetUpdatedBy(val string) {
	m.updatedByField = val
}

// Version gets the version of this polymorphic type
func (m *execution) Version() string {
	return m.versionField
}

// SetVersion sets the version of this polymorphic type
func (m *execution) SetVersion(val string) {
	m.versionField = val
}

// WorkspaceResults gets the workspace results of this polymorphic type
func (m *execution) WorkspaceResults() []*WorkspaceResult {
	return m.workspaceResultsField
}

// SetWorkspaceResults sets the workspace results of this polymorphic type
func (m *execution) SetWorkspaceResults(val []*WorkspaceResult) {
	m.workspaceResultsField = val
}

// UnmarshalExecutionSlice unmarshals polymorphic slices of Execution
func UnmarshalExecutionSlice(reader io.Reader, consumer runtime.Consumer) ([]Execution, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Execution
	for _, element := range elements {
		obj, err := unmarshalExecution(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalExecution unmarshals polymorphic Execution
func UnmarshalExecution(reader io.Reader, consumer runtime.Consumer) (Execution, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalExecution(data, consumer)
}

func unmarshalExecution(data []byte, consumer runtime.Consumer) (Execution, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Execution property.
	var getType struct {
		Execution string `json:"Execution"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Execution", "body", getType.Execution); err != nil {
		return nil, err
	}

	// The value of Execution is used to determine which type to create and unmarshal the data into
	switch getType.Execution {
	case "Execution":
		var result execution
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Execution value: %q", getType.Execution)
}

// Validate validates this execution
func (m *execution) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInputMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarred(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspaceResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *execution) validateInputMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.InputMeta()) { // not required
		return nil
	}

	for k := range m.InputMeta() {

		if err := validate.Required("_inputMeta"+"."+k, "body", m.InputMeta()[k]); err != nil {
			return err
		}
		if val, ok := m.InputMeta()[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *execution) validateOutputMeta(formats strfmt.Registry) error {
	if swag.IsZero(m.OutputMeta()) { // not required
		return nil
	}

	for k := range m.OutputMeta() {

		if err := validate.Required("_outputMeta"+"."+k, "body", m.OutputMeta()[k]); err != nil {
			return err
		}
		if val, ok := m.OutputMeta()[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *execution) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description()); err != nil {
		return err
	}

	return nil
}

func (m *execution) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name()); err != nil {
		return err
	}

	return nil
}

func (m *execution) validateNotifications(formats strfmt.Registry) error {
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

func (m *execution) validateStages(formats strfmt.Registry) error {
	if swag.IsZero(m.Stages()) { // not required
		return nil
	}

	for k := range m.Stages() {

		if err := validate.Required("stages"+"."+k, "body", m.Stages()[k]); err != nil {
			return err
		}
		if val, ok := m.Stages()[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *execution) validateStarred(formats strfmt.Registry) error {
	if swag.IsZero(m.Starred()) { // not required
		return nil
	}

	if err := m.Starred().Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("starred")
		}
		return err
	}

	return nil
}

var executionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_STARTED","STARTED","RUNNING","CANCELING","WAITING","RESUMING","PAUSING","PAUSED","CANCELED","COMPLETED","FAILED","SKIPPED","QUEUED","FAILED_CONTINUE","ROLLING_BACK","ROLLBACK_FAILED","PREPARING_WORKSPACE","ROLLBACK_COMPLETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		executionTypeStatusPropEnum = append(executionTypeStatusPropEnum, v)
	}
}

const (

	// ExecutionStatusNOTSTARTED captures enum value "NOT_STARTED"
	ExecutionStatusNOTSTARTED string = "NOT_STARTED"

	// ExecutionStatusSTARTED captures enum value "STARTED"
	ExecutionStatusSTARTED string = "STARTED"

	// ExecutionStatusRUNNING captures enum value "RUNNING"
	ExecutionStatusRUNNING string = "RUNNING"

	// ExecutionStatusCANCELING captures enum value "CANCELING"
	ExecutionStatusCANCELING string = "CANCELING"

	// ExecutionStatusWAITING captures enum value "WAITING"
	ExecutionStatusWAITING string = "WAITING"

	// ExecutionStatusRESUMING captures enum value "RESUMING"
	ExecutionStatusRESUMING string = "RESUMING"

	// ExecutionStatusPAUSING captures enum value "PAUSING"
	ExecutionStatusPAUSING string = "PAUSING"

	// ExecutionStatusPAUSED captures enum value "PAUSED"
	ExecutionStatusPAUSED string = "PAUSED"

	// ExecutionStatusCANCELED captures enum value "CANCELED"
	ExecutionStatusCANCELED string = "CANCELED"

	// ExecutionStatusCOMPLETED captures enum value "COMPLETED"
	ExecutionStatusCOMPLETED string = "COMPLETED"

	// ExecutionStatusFAILED captures enum value "FAILED"
	ExecutionStatusFAILED string = "FAILED"

	// ExecutionStatusSKIPPED captures enum value "SKIPPED"
	ExecutionStatusSKIPPED string = "SKIPPED"

	// ExecutionStatusQUEUED captures enum value "QUEUED"
	ExecutionStatusQUEUED string = "QUEUED"

	// ExecutionStatusFAILEDCONTINUE captures enum value "FAILED_CONTINUE"
	ExecutionStatusFAILEDCONTINUE string = "FAILED_CONTINUE"

	// ExecutionStatusROLLINGBACK captures enum value "ROLLING_BACK"
	ExecutionStatusROLLINGBACK string = "ROLLING_BACK"

	// ExecutionStatusROLLBACKFAILED captures enum value "ROLLBACK_FAILED"
	ExecutionStatusROLLBACKFAILED string = "ROLLBACK_FAILED"

	// ExecutionStatusPREPARINGWORKSPACE captures enum value "PREPARING_WORKSPACE"
	ExecutionStatusPREPARINGWORKSPACE string = "PREPARING_WORKSPACE"

	// ExecutionStatusROLLBACKCOMPLETED captures enum value "ROLLBACK_COMPLETED"
	ExecutionStatusROLLBACKCOMPLETED string = "ROLLBACK_COMPLETED"
)

// prop value enum
func (m *execution) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, executionTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *execution) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status()) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status()); err != nil {
		return err
	}

	return nil
}

func (m *execution) validateWorkspaceResults(formats strfmt.Registry) error {
	if swag.IsZero(m.WorkspaceResults()) { // not required
		return nil
	}

	for i := 0; i < len(m.WorkspaceResults()); i++ {
		if swag.IsZero(m.workspaceResultsField[i]) { // not required
			continue
		}

		if m.workspaceResultsField[i] != nil {
			if err := m.workspaceResultsField[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaceResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this execution based on the context it is used
func (m *execution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInputMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOutputMeta(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNotifications(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStarred(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspaceResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *execution) contextValidateInputMeta(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.InputMeta() {

		if val, ok := m.InputMeta()[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *execution) contextValidateOutputMeta(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.OutputMeta() {

		if val, ok := m.OutputMeta()[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *execution) contextValidateNotifications(ctx context.Context, formats strfmt.Registry) error {

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

func (m *execution) contextValidateStages(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Stages() {

		if val, ok := m.Stages()[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *execution) contextValidateStarred(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Starred().ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("starred")
		}
		return err
	}

	return nil
}

func (m *execution) contextValidateWorkspaceResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WorkspaceResults()); i++ {

		if m.workspaceResultsField[i] != nil {
			if err := m.workspaceResultsField[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("workspaceResults" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}
