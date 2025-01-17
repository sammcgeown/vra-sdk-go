// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new endpoints API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoints API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUsingPOST2(params *CreateUsingPOST2Params, opts ...ClientOption) (*CreateUsingPOST2OK, error)

	DeleteUsingDELETE1(params *DeleteUsingDELETE1Params, opts ...ClientOption) (*DeleteUsingDELETE1OK, error)

	DeleteUsingDELETE2(params *DeleteUsingDELETE2Params, opts ...ClientOption) (*DeleteUsingDELETE2OK, error)

	GetAllUsingGET2(params *GetAllUsingGET2Params, opts ...ClientOption) (*GetAllUsingGET2OK, error)

	GetByIDUsingGET(params *GetByIDUsingGETParams, opts ...ClientOption) (*GetByIDUsingGETOK, error)

	GetByNameUsingGET(params *GetByNameUsingGETParams, opts ...ClientOption) (*GetByNameUsingGETOK, error)

	GetEndpointPropertiesUsingGET(params *GetEndpointPropertiesUsingGETParams, opts ...ClientOption) (*GetEndpointPropertiesUsingGETOK, error)

	GetEndpointTilesUsingGET(params *GetEndpointTilesUsingGETParams, opts ...ClientOption) (*GetEndpointTilesUsingGETOK, error)

	GetUsingGETMixin3(params *GetUsingGETMixin3Params, opts ...ClientOption) (*GetUsingGETMixin3OK, error)

	UpdateByIDUsingPUT1(params *UpdateByIDUsingPUT1Params, opts ...ClientOption) (*UpdateByIDUsingPUT1OK, error)

	UpdateByNameUsingPUT1(params *UpdateByNameUsingPUT1Params, opts ...ClientOption) (*UpdateByNameUsingPUT1OK, error)

	ValidateEndpointUsingPOST(params *ValidateEndpointUsingPOSTParams, opts ...ClientOption) (*ValidateEndpointUsingPOSTOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateUsingPOST2 creates an endpoint

  Create an Endpoint based on the given project
*/
func (a *Client) CreateUsingPOST2(params *CreateUsingPOST2Params, opts ...ClientOption) (*CreateUsingPOST2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUsingPOST2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUsingPOST_2",
		Method:             "POST",
		PathPattern:        "/codestream/api/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUsingPOST2Reader{formats: a.formats},
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
	success, ok := result.(*CreateUsingPOST2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUsingPOST_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsingDELETE1 deletes an endpoint by id

  Delete an Endpoint with the given id
*/
func (a *Client) DeleteUsingDELETE1(params *DeleteUsingDELETE1Params, opts ...ClientOption) (*DeleteUsingDELETE1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsingDELETE1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUsingDELETE_1",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/endpoints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsingDELETE1Reader{formats: a.formats},
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
	success, ok := result.(*DeleteUsingDELETE1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUsingDELETE_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsingDELETE2 deletes an endpoint by project and name

  Delete an Endpoint with the given name
*/
func (a *Client) DeleteUsingDELETE2(params *DeleteUsingDELETE2Params, opts ...ClientOption) (*DeleteUsingDELETE2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsingDELETE2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUsingDELETE_2",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/endpoints/{project}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsingDELETE2Reader{formats: a.formats},
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
	success, ok := result.(*DeleteUsingDELETE2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUsingDELETE_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAllUsingGET2 gets all endpoints

  Get all Endpoints with specified paging and filter parameters
*/
func (a *Client) GetAllUsingGET2(params *GetAllUsingGET2Params, opts ...ClientOption) (*GetAllUsingGET2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllUsingGET2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getAllUsingGET_2",
		Method:             "GET",
		PathPattern:        "/codestream/api/endpoints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAllUsingGET2Reader{formats: a.formats},
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
	success, ok := result.(*GetAllUsingGET2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getAllUsingGET_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetByIDUsingGET gets an endpoint

  Get an Endpoint with the given id
*/
func (a *Client) GetByIDUsingGET(params *GetByIDUsingGETParams, opts ...ClientOption) (*GetByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetByIDUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/endpoints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetByIDUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetByNameUsingGET gets an endpoint by project and name

  Get an Endpoint with the given project and name
*/
func (a *Client) GetByNameUsingGET(params *GetByNameUsingGETParams, opts ...ClientOption) (*GetByNameUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetByNameUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getByNameUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/endpoints/{project}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetByNameUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetByNameUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getByNameUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEndpointPropertiesUsingGET gets endpoint properties

  Get endpoint properties with the given endpoint type
*/
func (a *Client) GetEndpointPropertiesUsingGET(params *GetEndpointPropertiesUsingGETParams, opts ...ClientOption) (*GetEndpointPropertiesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointPropertiesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEndpointPropertiesUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/endpoint-tiles/{type}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEndpointPropertiesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetEndpointPropertiesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEndpointPropertiesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetEndpointTilesUsingGET gets all endpoint tiles

  Get all supported endpoint tiles
*/
func (a *Client) GetEndpointTilesUsingGET(params *GetEndpointTilesUsingGETParams, opts ...ClientOption) (*GetEndpointTilesUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointTilesUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getEndpointTilesUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/endpoint-tiles",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetEndpointTilesUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetEndpointTilesUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getEndpointTilesUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsingGETMixin3 gets endpoint certificate

  Get endpoint certificate chain for validation
*/
func (a *Client) GetUsingGETMixin3(params *GetUsingGETMixin3Params, opts ...ClientOption) (*GetUsingGETMixin3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsingGETMixin3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "getUsingGETMixin3",
		Method:             "GET",
		PathPattern:        "/codestream/api/endpoint-certificate",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsingGETMixin3Reader{formats: a.formats},
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
	success, ok := result.(*GetUsingGETMixin3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getUsingGETMixin3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateByIDUsingPUT1 updates an endpoint by id

  Update an Endpoint with the given id
*/
func (a *Client) UpdateByIDUsingPUT1(params *UpdateByIDUsingPUT1Params, opts ...ClientOption) (*UpdateByIDUsingPUT1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateByIDUsingPUT1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateByIdUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/codestream/api/endpoints/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateByIDUsingPUT1Reader{formats: a.formats},
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
	success, ok := result.(*UpdateByIDUsingPUT1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateByIdUsingPUT_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateByNameUsingPUT1 updates an endpoint by project and name

  Update an Endpoint with the given project and name
*/
func (a *Client) UpdateByNameUsingPUT1(params *UpdateByNameUsingPUT1Params, opts ...ClientOption) (*UpdateByNameUsingPUT1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateByNameUsingPUT1Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateByNameUsingPUT_1",
		Method:             "PUT",
		PathPattern:        "/codestream/api/endpoints/{project}/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateByNameUsingPUT1Reader{formats: a.formats},
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
	success, ok := result.(*UpdateByNameUsingPUT1OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateByNameUsingPUT_1: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ValidateEndpointUsingPOST validates endpoint

  Validates the given endpoint
*/
func (a *Client) ValidateEndpointUsingPOST(params *ValidateEndpointUsingPOSTParams, opts ...ClientOption) (*ValidateEndpointUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateEndpointUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "validateEndpointUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/endpoint-validation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ValidateEndpointUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*ValidateEndpointUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for validateEndpointUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
