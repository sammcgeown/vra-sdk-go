// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new variables API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for variables API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUsingPOST7(params *CreateUsingPOST7Params, opts ...ClientOption) (*CreateUsingPOST7OK, error)

	DeleteUsingDELETE11(params *DeleteUsingDELETE11Params, opts ...ClientOption) (*DeleteUsingDELETE11OK, error)

	DeleteUsingDELETE12(params *DeleteUsingDELETE12Params, opts ...ClientOption) (*DeleteUsingDELETE12OK, error)

	GetAllUsingGET11(params *GetAllUsingGET11Params, opts ...ClientOption) (*GetAllUsingGET11OK, error)

	GetUsingGET5(params *GetUsingGET5Params, opts ...ClientOption) (*GetUsingGET5OK, error)

	GetUsingGET6(params *GetUsingGET6Params, opts ...ClientOption) (*GetUsingGET6OK, error)

	UpdateUsingPUT4(params *UpdateUsingPUT4Params, opts ...ClientOption) (*UpdateUsingPUT4OK, error)

	UpdateUsingPUT5(params *UpdateUsingPUT5Params, opts ...ClientOption) (*UpdateUsingPUT5OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateUsingPOST7 creates a variable

  Creates a Variable based on project name
*/
func (a *Client) CreateUsingPOST7(params *CreateUsingPOST7Params, opts ...ClientOption) (*CreateUsingPOST7OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUsingPOST7Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUsingPOST_7",
		Method:             "POST",
		PathPattern:        "/codestream/api/variables",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUsingPOST7Reader{formats: a.formats},
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
	success, ok := result.(*CreateUsingPOST7OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUsingPOST_7: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsingDELETE11 deletes a variable by Id

  Deletes a Variable with the given Id
*/
func (a *Client) DeleteUsingDELETE11(params *DeleteUsingDELETE11Params, opts ...ClientOption) (*DeleteUsingDELETE11OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsingDELETE11Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUsingDELETE_11",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/variables/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsingDELETE11Reader{formats: a.formats},
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
	success, ok := result.(*DeleteUsingDELETE11OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUsingDELETE_11: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsingDELETE12 deletes a variable by project and name

  Deletes a Variable with the given name
*/
func (a *Client) DeleteUsingDELETE12(params *DeleteUsingDELETE12Params, opts ...ClientOption) (*DeleteUsingDELETE12OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsingDELETE12Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUsingDELETE_12",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/variables/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsingDELETE12Reader{formats: a.formats},
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
	success, ok := result.(*DeleteUsingDELETE12OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUsingDELETE_12: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllUsingGET11 gets all variables

  Get all Variables with specified paging and filter parameters.
*/
func (a *Client) GetAllUsingGET11(params *GetAllUsingGET11Params, opts ...ClientOption) (*GetAllUsingGET11OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllUsingGET11Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllUsingGET_11",
		Method:             "GET",
		PathPattern:        "/codestream/api/variables",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllUsingGET11Reader{formats: a.formats},
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
	success, ok := result.(*GetAllUsingGET11OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllUsingGET_11: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsingGET5 gets a variable

  Gets a Variable with the given id
*/
func (a *Client) GetUsingGET5(params *GetUsingGET5Params, opts ...ClientOption) (*GetUsingGET5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsingGET5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsingGET_5",
		Method:             "GET",
		PathPattern:        "/codestream/api/variables/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsingGET5Reader{formats: a.formats},
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
	success, ok := result.(*GetUsingGET5OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsingGET_5: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsingGET6 gets a variable by project and name

  Get an Variable with the given project and name
*/
func (a *Client) GetUsingGET6(params *GetUsingGET6Params, opts ...ClientOption) (*GetUsingGET6OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsingGET6Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsingGET_6",
		Method:             "GET",
		PathPattern:        "/codestream/api/variables/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsingGET6Reader{formats: a.formats},
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
	success, ok := result.(*GetUsingGET6OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsingGET_6: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateUsingPUT4 updates a variable by id

  Updates a Variable with the given id
*/
func (a *Client) UpdateUsingPUT4(params *UpdateUsingPUT4Params, opts ...ClientOption) (*UpdateUsingPUT4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUsingPUT4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUsingPUT_4",
		Method:             "PUT",
		PathPattern:        "/codestream/api/variables/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUsingPUT4Reader{formats: a.formats},
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
	success, ok := result.(*UpdateUsingPUT4OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUsingPUT_4: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateUsingPUT5 updates a variable by project and name

  Update an Variable with the given project and name
*/
func (a *Client) UpdateUsingPUT5(params *UpdateUsingPUT5Params, opts ...ClientOption) (*UpdateUsingPUT5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUsingPUT5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUsingPUT_5",
		Method:             "PUT",
		PathPattern:        "/codestream/api/variables/{project}/{name}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUsingPUT5Reader{formats: a.formats},
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
	success, ok := result.(*UpdateUsingPUT5OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUsingPUT_5: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
