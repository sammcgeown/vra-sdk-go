// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PKSCluster PKSCluster
//
// swagger:model PKSCluster
type PKSCluster struct {

	// hostname address
	HostnameAddress string `json:"hostnameAddress,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// kubernetes master ips
	KubernetesMasterIps []string `json:"kubernetes_master_ips"`

	// last action
	LastAction string `json:"last_action,omitempty"`

	// last action description
	LastActionDescription string `json:"last_action_description,omitempty"`

	// last action state
	LastActionState string `json:"last_action_state,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// parameters
	Parameters interface{} `json:"parameters,omitempty"`

	// plan name
	PlanName string `json:"plan_name,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this p k s cluster
func (m *PKSCluster) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p k s cluster based on context it is used
func (m *PKSCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PKSCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PKSCluster) UnmarshalBinary(b []byte) error {
	var res PKSCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
