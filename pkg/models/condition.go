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

// Condition Condition
//
// Definition of a condition a constraint may have.
//
// swagger:discriminator Condition Definition of a condition a constraint may have.
type Condition interface {
	runtime.Validatable
	runtime.ContextValidatable

	// enforcement
	// Enum: [HARD SOFT]
	Enforcement() string
	SetEnforcement(string)

	// expression
	Expression() *Tag
	SetExpression(*Tag)

	// occurrence
	// Enum: [MUST_OCCUR MUST_NOT_OCCUR]
	Occurrence() string
	SetOccurrence(string)

	// type
	// Enum: [TAG]
	Type() string
	SetType(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type condition struct {
	enforcementField string

	expressionField *Tag

	occurrenceField string

	typeField string
}

// Enforcement gets the enforcement of this polymorphic type
func (m *condition) Enforcement() string {
	return m.enforcementField
}

// SetEnforcement sets the enforcement of this polymorphic type
func (m *condition) SetEnforcement(val string) {
	m.enforcementField = val
}

// Expression gets the expression of this polymorphic type
func (m *condition) Expression() *Tag {
	return m.expressionField
}

// SetExpression sets the expression of this polymorphic type
func (m *condition) SetExpression(val *Tag) {
	m.expressionField = val
}

// Occurrence gets the occurrence of this polymorphic type
func (m *condition) Occurrence() string {
	return m.occurrenceField
}

// SetOccurrence sets the occurrence of this polymorphic type
func (m *condition) SetOccurrence(val string) {
	m.occurrenceField = val
}

// Type gets the type of this polymorphic type
func (m *condition) Type() string {
	return m.typeField
}

// SetType sets the type of this polymorphic type
func (m *condition) SetType(val string) {
	m.typeField = val
}

// UnmarshalConditionSlice unmarshals polymorphic slices of Condition
func UnmarshalConditionSlice(reader io.Reader, consumer runtime.Consumer) ([]Condition, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Condition
	for _, element := range elements {
		obj, err := unmarshalCondition(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalCondition unmarshals polymorphic Condition
func UnmarshalCondition(reader io.Reader, consumer runtime.Consumer) (Condition, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalCondition(data, consumer)
}

func unmarshalCondition(data []byte, consumer runtime.Consumer) (Condition, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the Definition of a condition a constraint may have. property.
	var getType struct {
		DefinitionOfaConditionaConstraintMayHave string `json:"Definition of a condition a constraint may have."`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("Definition of a condition a constraint may have.", "body", getType.DefinitionOfaConditionaConstraintMayHave); err != nil {
		return nil, err
	}

	// The value of Definition of a condition a constraint may have. is used to determine which type to create and unmarshal the data into
	switch getType.DefinitionOfaConditionaConstraintMayHave {
	case "Condition":
		var result condition
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid Definition of a condition a constraint may have. value: %q", getType.DefinitionOfaConditionaConstraintMayHave)
}

// Validate validates this condition
func (m *condition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnforcement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpression(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOccurrence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conditionTypeEnforcementPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HARD","SOFT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeEnforcementPropEnum = append(conditionTypeEnforcementPropEnum, v)
	}
}

const (

	// ConditionEnforcementHARD captures enum value "HARD"
	ConditionEnforcementHARD string = "HARD"

	// ConditionEnforcementSOFT captures enum value "SOFT"
	ConditionEnforcementSOFT string = "SOFT"
)

// prop value enum
func (m *condition) validateEnforcementEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conditionTypeEnforcementPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *condition) validateEnforcement(formats strfmt.Registry) error {
	if swag.IsZero(m.Enforcement()) { // not required
		return nil
	}

	// value enum
	if err := m.validateEnforcementEnum("enforcement", "body", m.Enforcement()); err != nil {
		return err
	}

	return nil
}

func (m *condition) validateExpression(formats strfmt.Registry) error {
	if swag.IsZero(m.Expression()) { // not required
		return nil
	}

	if m.Expression() != nil {
		if err := m.Expression().Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expression")
			}
			return err
		}
	}

	return nil
}

var conditionTypeOccurrencePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MUST_OCCUR","MUST_NOT_OCCUR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeOccurrencePropEnum = append(conditionTypeOccurrencePropEnum, v)
	}
}

const (

	// ConditionOccurrenceMUSTOCCUR captures enum value "MUST_OCCUR"
	ConditionOccurrenceMUSTOCCUR string = "MUST_OCCUR"

	// ConditionOccurrenceMUSTNOTOCCUR captures enum value "MUST_NOT_OCCUR"
	ConditionOccurrenceMUSTNOTOCCUR string = "MUST_NOT_OCCUR"
)

// prop value enum
func (m *condition) validateOccurrenceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conditionTypeOccurrencePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *condition) validateOccurrence(formats strfmt.Registry) error {
	if swag.IsZero(m.Occurrence()) { // not required
		return nil
	}

	// value enum
	if err := m.validateOccurrenceEnum("occurrence", "body", m.Occurrence()); err != nil {
		return err
	}

	return nil
}

var conditionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TAG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeTypePropEnum = append(conditionTypeTypePropEnum, v)
	}
}

const (

	// ConditionTypeTAG captures enum value "TAG"
	ConditionTypeTAG string = "TAG"
)

// prop value enum
func (m *condition) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conditionTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *condition) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type()) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type()); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this condition based on the context it is used
func (m *condition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExpression(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *condition) contextValidateExpression(ctx context.Context, formats strfmt.Registry) error {

	if m.Expression() != nil {
		if err := m.Expression().ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expression")
			}
			return err
		}
	}

	return nil
}
