// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePolicyUsingPOST1(params *CreatePolicyUsingPOST1Params, opts ...ClientOption) (*CreatePolicyUsingPOST1OK, error)

	DeletePolicyUsingDELETE1(params *DeletePolicyUsingDELETE1Params, opts ...ClientOption) (*DeletePolicyUsingDELETE1NoContent, error)

	GetPoliciesUsingGET1(params *GetPoliciesUsingGET1Params, opts ...ClientOption) (*GetPoliciesUsingGET1OK, error)

	GetPolicyUsingGET1(params *GetPolicyUsingGET1Params, opts ...ClientOption) (*GetPolicyUsingGET1OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePolicyUsingPOST1 creates a new policy or updates an existing policy

  Create a new policy or update an existing policy based on request body and validate its fields according to business rules.
*/
func (a *Client) CreatePolicyUsingPOST1(params *CreatePolicyUsingPOST1Params, opts ...ClientOption) (*CreatePolicyUsingPOST1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePolicyUsingPOST1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createPolicyUsingPOST_1",
		Method:             "POST",
		PathPattern:        "/policy/api/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreatePolicyUsingPOST1Reader{formats: a.formats},
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
	success, ok := result.(*CreatePolicyUsingPOST1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPolicyUsingPOST_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePolicyUsingDELETE1 deletes a policy

  Delete a specified policy corresponding to its unique id.
*/
func (a *Client) DeletePolicyUsingDELETE1(params *DeletePolicyUsingDELETE1Params, opts ...ClientOption) (*DeletePolicyUsingDELETE1NoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePolicyUsingDELETE1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deletePolicyUsingDELETE_1",
		Method:             "DELETE",
		PathPattern:        "/policy/api/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeletePolicyUsingDELETE1Reader{formats: a.formats},
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
	success, ok := result.(*DeletePolicyUsingDELETE1NoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePolicyUsingDELETE_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPoliciesUsingGET1 returns a paginated list of policies

  Find all the policies associated with current org.
*/
func (a *Client) GetPoliciesUsingGET1(params *GetPoliciesUsingGET1Params, opts ...ClientOption) (*GetPoliciesUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPoliciesUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPoliciesUsingGET_1",
		Method:             "GET",
		PathPattern:        "/policy/api/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPoliciesUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetPoliciesUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPoliciesUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPolicyUsingGET1 returns a specified policy

  Find a specific policy based on the input policy id.
*/
func (a *Client) GetPolicyUsingGET1(params *GetPolicyUsingGET1Params, opts ...ClientOption) (*GetPolicyUsingGET1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPolicyUsingGET1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getPolicyUsingGET_1",
		Method:             "GET",
		PathPattern:        "/policy/api/policies/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPolicyUsingGET1Reader{formats: a.formats},
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
	success, ok := result.(*GetPolicyUsingGET1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPolicyUsingGET_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
