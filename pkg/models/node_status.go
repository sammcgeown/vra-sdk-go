// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeStatus NodeStatus
//
// swagger:model NodeStatus
type NodeStatus struct {

	// allocatable
	Allocatable interface{} `json:"allocatable,omitempty"`

	// capacity
	Capacity interface{} `json:"capacity,omitempty"`

	// phase
	Phase string `json:"phase,omitempty"`
}

// Validate validates this node status
func (m *NodeStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this node status based on context it is used
func (m *NodeStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeStatus) UnmarshalBinary(b []byte) error {
	var res NodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}