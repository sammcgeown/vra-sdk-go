// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BlueprintExecution BlueprintExecution
//
// swagger:model BlueprintExecution
type BlueprintExecution struct {

	// Blueprint execution failure message
	// Read Only: true
	FailureMessage string `json:"failureMessage,omitempty"`

	// Blueprint execution input properties
	// Read Only: true
	InputProperties interface{} `json:"inputProperties,omitempty"`

	// Blueprint execution output properties
	// Read Only: true
	OutputProperties interface{} `json:"outputProperties,omitempty"`

	// Blueprint execution status
	// Read Only: true
	// Enum: [CREATED STARTED FINISHED FAILED CANCELLED]
	Status string `json:"status,omitempty"`

	// Blueprint execution tasks
	// Read Only: true
	Tasks []*BlueprintTaskExecution `json:"tasks"`
}

// Validate validates this blueprint execution
func (m *BlueprintExecution) Validate(formats strfmt.Registry) error {
	var res []error

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

var blueprintExecutionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATED","STARTED","FINISHED","FAILED","CANCELLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blueprintExecutionTypeStatusPropEnum = append(blueprintExecutionTypeStatusPropEnum, v)
	}
}

const (

	// BlueprintExecutionStatusCREATED captures enum value "CREATED"
	BlueprintExecutionStatusCREATED string = "CREATED"

	// BlueprintExecutionStatusSTARTED captures enum value "STARTED"
	BlueprintExecutionStatusSTARTED string = "STARTED"

	// BlueprintExecutionStatusFINISHED captures enum value "FINISHED"
	BlueprintExecutionStatusFINISHED string = "FINISHED"

	// BlueprintExecutionStatusFAILED captures enum value "FAILED"
	BlueprintExecutionStatusFAILED string = "FAILED"

	// BlueprintExecutionStatusCANCELLED captures enum value "CANCELLED"
	BlueprintExecutionStatusCANCELLED string = "CANCELLED"
)

// prop value enum
func (m *BlueprintExecution) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, blueprintExecutionTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BlueprintExecution) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintExecution) validateTasks(formats strfmt.Registry) error {
	if swag.IsZero(m.Tasks) { // not required
		return nil
	}

	for i := 0; i < len(m.Tasks); i++ {
		if swag.IsZero(m.Tasks[i]) { // not required
			continue
		}

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this blueprint execution based on the context it is used
func (m *BlueprintExecution) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFailureMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
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

func (m *BlueprintExecution) contextValidateFailureMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "failureMessage", "body", string(m.FailureMessage)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintExecution) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *BlueprintExecution) contextValidateTasks(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tasks", "body", []*BlueprintTaskExecution(m.Tasks)); err != nil {
		return err
	}

	for i := 0; i < len(m.Tasks); i++ {

		if m.Tasks[i] != nil {
			if err := m.Tasks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tasks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlueprintExecution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlueprintExecution) UnmarshalBinary(b []byte) error {
	var res BlueprintExecution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
