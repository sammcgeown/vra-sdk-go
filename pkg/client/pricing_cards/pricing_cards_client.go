// Code generated by go-swagger; DO NOT EDIT.

package pricing_cards

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new pricing cards API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for pricing cards API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePolicyUsingPOST(params *CreatePolicyUsingPOSTParams, opts ...ClientOption) (*CreatePolicyUsingPOSTOK, *CreatePolicyUsingPOSTCreated, error)

	DeletePolicyUsingDELETE(params *DeletePolicyUsingDELETEParams, opts ...ClientOption) (*DeletePolicyUsingDELETEOK, *DeletePolicyUsingDELETENoContent, error)

	GetPoliciesUsingGET(params *GetPoliciesUsingGETParams, opts ...ClientOption) (*GetPoliciesUsingGETOK, error)

	GetPolicyUsingGET(params *GetPolicyUsingGETParams, opts ...ClientOption) (*GetPolicyUsingGETOK, error)

	UpdatePolicyUsingPUT(params *UpdatePolicyUsingPUTParams, opts ...ClientOption) (*UpdatePolicyUsingPUTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePolicyUsingPOST creates a new pricing card

  Create a new pricing card based on request body and validate its field according to business rules.
*/
func (a *Client) CreatePolicyUsingPOST(params *CreatePolicyUsingPOSTParams, opts ...ClientOption) (*CreatePolicyUsingPOSTOK, *CreatePolicyUsingPOSTCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePolicyUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPolicyUsingPOST",
		Method:             "POST",
		PathPattern:        "/price/api/private/pricing-cards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePolicyUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CreatePolicyUsingPOSTOK:
		return value, nil, nil
	case *CreatePolicyUsingPOSTCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pricing_cards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePolicyUsingDELETE deletes the pricing card with specified Id

  Deletes the pricing card with the specified id
*/
func (a *Client) DeletePolicyUsingDELETE(params *DeletePolicyUsingDELETEParams, opts ...ClientOption) (*DeletePolicyUsingDELETEOK, *DeletePolicyUsingDELETENoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePolicyUsingDELETEParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePolicyUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/price/api/private/pricing-cards/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePolicyUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeletePolicyUsingDELETEOK:
		return value, nil, nil
	case *DeletePolicyUsingDELETENoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pricing_cards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPoliciesUsingGET fetches all pricing cards for private policy cloud

  Returns a paginated list of pricing cards
*/
func (a *Client) GetPoliciesUsingGET(params *GetPoliciesUsingGETParams, opts ...ClientOption) (*GetPoliciesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoliciesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPoliciesUsingGET",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-cards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoliciesUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPoliciesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPoliciesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPolicyUsingGET finds the pricing card with specified Id

  Returns the pricing card with the specified id
*/
func (a *Client) GetPolicyUsingGET(params *GetPolicyUsingGETParams, opts ...ClientOption) (*GetPolicyUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPolicyUsingGET",
		Method:             "GET",
		PathPattern:        "/price/api/private/pricing-cards/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPolicyUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPolicyUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPolicyUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePolicyUsingPUT updates the pricing card

  Updates the pricing card with the specified Id
*/
func (a *Client) UpdatePolicyUsingPUT(params *UpdatePolicyUsingPUTParams, opts ...ClientOption) (*UpdatePolicyUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePolicyUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePolicyUsingPUT",
		Method:             "PUT",
		PathPattern:        "/price/api/private/pricing-cards/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdatePolicyUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePolicyUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePolicyUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
