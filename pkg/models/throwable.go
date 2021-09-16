// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Throwable Throwable
//
// swagger:model Throwable
type Throwable struct {

	// cause
	Cause *Throwable `json:"cause,omitempty"`

	// stack trace
	StackTrace []*StackTraceElement `json:"stackTrace"`
}

// Validate validates this throwable
func (m *Throwable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackTrace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Throwable) validateCause(formats strfmt.Registry) error {
	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	if m.Cause != nil {
		if err := m.Cause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *Throwable) validateStackTrace(formats strfmt.Registry) error {
	if swag.IsZero(m.StackTrace) { // not required
		return nil
	}

	for i := 0; i < len(m.StackTrace); i++ {
		if swag.IsZero(m.StackTrace[i]) { // not required
			continue
		}

		if m.StackTrace[i] != nil {
			if err := m.StackTrace[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stackTrace" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this throwable based on the context it is used
func (m *Throwable) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCause(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStackTrace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Throwable) contextValidateCause(ctx context.Context, formats strfmt.Registry) error {

	if m.Cause != nil {
		if err := m.Cause.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *Throwable) contextValidateStackTrace(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StackTrace); i++ {

		if m.StackTrace[i] != nil {
			if err := m.StackTrace[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("stackTrace" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Throwable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Throwable) UnmarshalBinary(b []byte) error {
	var res Throwable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}