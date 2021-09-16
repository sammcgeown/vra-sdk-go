// Code generated by go-swagger; DO NOT EDIT.

package supervisor_namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new supervisor namespaces API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for supervisor namespaces API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateUsingPOST2Mixin5(params *CreateUsingPOST2Mixin5Params, opts ...ClientOption) (*CreateUsingPOST2Mixin5OK, error)

	DeleteUsingDELETE1Mixin5(params *DeleteUsingDELETE1Mixin5Params, opts ...ClientOption) (*DeleteUsingDELETE1Mixin5OK, error)

	GetNamespaceQuotasUsingGET(params *GetNamespaceQuotasUsingGETParams, opts ...ClientOption) (*GetNamespaceQuotasUsingGETOK, error)

	GetNamespaceUsingGET(params *GetNamespaceUsingGETParams, opts ...ClientOption) (*GetNamespaceUsingGETOK, error)

	ListUsingGET5(params *ListUsingGET5Params, opts ...ClientOption) (*ListUsingGET5OK, error)

	RegisterUsingPUT2(params *RegisterUsingPUT2Params, opts ...ClientOption) (*RegisterUsingPUT2OK, error)

	SetNamespaceQuotasUsingPUT(params *SetNamespaceQuotasUsingPUTParams, opts ...ClientOption) (*SetNamespaceQuotasUsingPUTOK, error)

	SyncStatusUsingGET(params *SyncStatusUsingGETParams, opts ...ClientOption) (*SyncStatusUsingGETOK, error)

	SyncUsingPOST(params *SyncUsingPOSTParams, opts ...ClientOption) (*SyncUsingPOSTOK, error)

	UpdateUsingPATCHMixin5(params *UpdateUsingPATCHMixin5Params, opts ...ClientOption) (*UpdateUsingPATCHMixin5OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateUsingPOST2Mixin5 creates supervisor namespace

  Create Supervisor Namespace
*/
func (a *Client) CreateUsingPOST2Mixin5(params *CreateUsingPOST2Mixin5Params, opts ...ClientOption) (*CreateUsingPOST2Mixin5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUsingPOST2Mixin5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "createUsingPOST_2Mixin5",
		Method:             "POST",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUsingPOST2Mixin5Reader{formats: a.formats},
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
	success, ok := result.(*CreateUsingPOST2Mixin5OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUsingPOST_2Mixin5: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsingDELETE1Mixin5 makes not managed a supervisor namespace and optionally delete it

  Make not managed a Supervisor Namespace and delete it from cluster if destroy parameter is true
*/
func (a *Client) DeleteUsingDELETE1Mixin5(params *DeleteUsingDELETE1Mixin5Params, opts ...ClientOption) (*DeleteUsingDELETE1Mixin5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsingDELETE1Mixin5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteUsingDELETE_1Mixin5",
		Method:             "DELETE",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces/{selfLinkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsingDELETE1Mixin5Reader{formats: a.formats},
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
	success, ok := result.(*DeleteUsingDELETE1Mixin5OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteUsingDELETE_1Mixin5: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNamespaceQuotasUsingGET gets supervisor namespace quotas by the id from document self link

  Retrieve a Supervisor Namespace Quota by id from documentSelfLink
*/
func (a *Client) GetNamespaceQuotasUsingGET(params *GetNamespaceQuotasUsingGETParams, opts ...ClientOption) (*GetNamespaceQuotasUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespaceQuotasUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNamespaceQuotasUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces/{selfLinkId}/resource-quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNamespaceQuotasUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetNamespaceQuotasUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNamespaceQuotasUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNamespaceUsingGET finds a supervisor namespace by the id from document self link

  Retrieve a Supervisor Namespace by id from documentSelfLink
*/
func (a *Client) GetNamespaceUsingGET(params *GetNamespaceUsingGETParams, opts ...ClientOption) (*GetNamespaceUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespaceUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNamespaceUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces/{selfLinkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNamespaceUsingGETReader{formats: a.formats},
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
	success, ok := result.(*GetNamespaceUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNamespaceUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListUsingGET5 gets all supervisor namespaces

  Get all Supervisor Namespaces
*/
func (a *Client) ListUsingGET5(params *ListUsingGET5Params, opts ...ClientOption) (*ListUsingGET5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListUsingGET5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "listUsingGET_5",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListUsingGET5Reader{formats: a.formats},
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
	success, ok := result.(*ListUsingGET5OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listUsingGET_5: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RegisterUsingPUT2 makes a supervisor namespace a managed entity

  The body shall contain valid projectId, DocumentSelfLink and list of viewer and editor user and groups.
*/
func (a *Client) RegisterUsingPUT2(params *RegisterUsingPUT2Params, opts ...ClientOption) (*RegisterUsingPUT2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterUsingPUT2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "registerUsingPUT_2",
		Method:             "PUT",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces/{namespaceSelfLinkId}/register",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegisterUsingPUT2Reader{formats: a.formats},
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
	success, ok := result.(*RegisterUsingPUT2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for registerUsingPUT_2: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetNamespaceQuotasUsingPUT sets supervisor namespace quotas by the id from document self link

  Set a Supervisor Namespace Quota by id from documentSelfLink
*/
func (a *Client) SetNamespaceQuotasUsingPUT(params *SetNamespaceQuotasUsingPUTParams, opts ...ClientOption) (*SetNamespaceQuotasUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetNamespaceQuotasUsingPUTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setNamespaceQuotasUsingPUT",
		Method:             "PUT",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces/{selfLinkId}/resource-quotas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetNamespaceQuotasUsingPUTReader{formats: a.formats},
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
	success, ok := result.(*SetNamespaceQuotasUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setNamespaceQuotasUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SyncStatusUsingGET retrieves sync status

  Retrieve sync status.
*/
func (a *Client) SyncStatusUsingGET(params *SyncStatusUsingGETParams, opts ...ClientOption) (*SyncStatusUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncStatusUsingGETParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "syncStatusUsingGET",
		Method:             "GET",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces/{namespaceSelfLinkId}/principals/{requestId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SyncStatusUsingGETReader{formats: a.formats},
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
	success, ok := result.(*SyncStatusUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for syncStatusUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SyncUsingPOST syncs supervisor namespace principals access list

  The body shall contain list of viewer and editor user and groups.
*/
func (a *Client) SyncUsingPOST(params *SyncUsingPOSTParams, opts ...ClientOption) (*SyncUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "syncUsingPOST",
		Method:             "POST",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces/{namespaceSelfLinkId}/principals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SyncUsingPOSTReader{formats: a.formats},
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
	success, ok := result.(*SyncUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for syncUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateUsingPATCHMixin5 updates supervisor namespace

  Update Supervisor Namespace
*/
func (a *Client) UpdateUsingPATCHMixin5(params *UpdateUsingPATCHMixin5Params, opts ...ClientOption) (*UpdateUsingPATCHMixin5OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUsingPATCHMixin5Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateUsingPATCHMixin5",
		Method:             "PATCH",
		PathPattern:        "/cmx/api/resources/supervisor-namespaces/{selfLinkId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateUsingPATCHMixin5Reader{formats: a.formats},
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
	success, ok := result.(*UpdateUsingPATCHMixin5OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateUsingPATCHMixin5: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
