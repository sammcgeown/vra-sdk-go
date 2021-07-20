// Code generated by go-swagger; DO NOT EDIT.

package user_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteUsingDELETE10(params *DeleteUsingDELETE10Params, opts ...ClientOption) (*DeleteUsingDELETE10OK, error)

	GetAllUsingGET10(params *GetAllUsingGET10Params, opts ...ClientOption) (*GetAllUsingGET10OK, error)

	GetUsingGET4(params *GetUsingGET4Params, opts ...ClientOption) (*GetUsingGET4OK, error)

	ModifyPatchUserOperationUsingPATCH(params *ModifyPatchUserOperationUsingPATCHParams, opts ...ClientOption) (*ModifyPatchUserOperationUsingPATCHOK, error)

	ModifyPostUserOperationUsingPOST(params *ModifyPostUserOperationUsingPOSTParams, opts ...ClientOption) (*ModifyPostUserOperationUsingPOSTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteUsingDELETE10 deletes a user operation by id

  Delete a User Operation with the given id
*/
func (a *Client) DeleteUsingDELETE10(params *DeleteUsingDELETE10Params, opts ...ClientOption) (*DeleteUsingDELETE10OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsingDELETE10Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUsingDELETE_10",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/user-operations/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsingDELETE10Reader{formats: a.formats},
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
	success, ok := result.(*DeleteUsingDELETE10OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUsingDELETE_10: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllUsingGET10 gets all user operations

  Get all User operations with specified paging and filter parameters.
*/
func (a *Client) GetAllUsingGET10(params *GetAllUsingGET10Params, opts ...ClientOption) (*GetAllUsingGET10OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllUsingGET10Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllUsingGET_10",
		Method:             "GET",
		PathPattern:        "/codestream/api/user-operations",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllUsingGET10Reader{formats: a.formats},
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
	success, ok := result.(*GetAllUsingGET10OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllUsingGET_10: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsingGET4 gets a user operation

  Get a User Operation with the given id
*/
func (a *Client) GetUsingGET4(params *GetUsingGET4Params, opts ...ClientOption) (*GetUsingGET4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsingGET4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsingGET_4",
		Method:             "GET",
		PathPattern:        "/codestream/api/user-operations/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsingGET4Reader{formats: a.formats},
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
	success, ok := result.(*GetUsingGET4OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsingGET_4: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ModifyPatchUserOperationUsingPATCH modifies a user operation

  Modify a User Operation with the given id
*/
func (a *Client) ModifyPatchUserOperationUsingPATCH(params *ModifyPatchUserOperationUsingPATCHParams, opts ...ClientOption) (*ModifyPatchUserOperationUsingPATCHOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyPatchUserOperationUsingPATCHParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifyPatchUserOperationUsingPATCH",
		Method:             "PATCH",
		PathPattern:        "/codestream/api/user-operations/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyPatchUserOperationUsingPATCHReader{formats: a.formats},
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
	success, ok := result.(*ModifyPatchUserOperationUsingPATCHOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifyPatchUserOperationUsingPATCH: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ModifyPostUserOperationUsingPOST modifies a user operation

  Modify a User Operation with the given id
*/
func (a *Client) ModifyPostUserOperationUsingPOST(params *ModifyPostUserOperationUsingPOSTParams, opts ...ClientOption) (*ModifyPostUserOperationUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyPostUserOperationUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "modifyPostUserOperationUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/user-operations/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyPostUserOperationUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*ModifyPostUserOperationUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for modifyPostUserOperationUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}