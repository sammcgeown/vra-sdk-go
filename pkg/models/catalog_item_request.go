// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CatalogItemRequest CatalogItemRequest
//
// A request to create a deployment based on a catalog item
// swagger:model CatalogItemRequest
type CatalogItemRequest struct {

	// Name of the requested deployment
	DeploymentName string `json:"deploymentName,omitempty"`

	// Input parameters for the request. These must be compliant with the schema of the corresponding catalog item
	Inputs interface{} `json:"inputs,omitempty"`

	// Project to be used for the request
	ProjectID string `json:"projectId,omitempty"`

	// Reason for request
	Reason string `json:"reason,omitempty"`

	// Version of the catalog item. e.g. v2.0
	Version string `json:"version,omitempty"`
}

// Validate validates this catalog item request
func (m *CatalogItemRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CatalogItemRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogItemRequest) UnmarshalBinary(b []byte) error {
	var res CatalogItemRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
