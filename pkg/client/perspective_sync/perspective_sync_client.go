// Code generated by go-swagger; DO NOT EDIT.

package perspective_sync

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new perspective sync API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for perspective sync API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	SyncPerspectiveGroupUsingPOST(params *SyncPerspectiveGroupUsingPOSTParams, opts ...ClientOption) (*SyncPerspectiveGroupUsingPOSTOK, *SyncPerspectiveGroupUsingPOSTAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SyncPerspectiveGroupUsingPOST ons demand perspective sync

  To do on demand perspective sync for within the given org
*/
func (a *Client) SyncPerspectiveGroupUsingPOST(params *SyncPerspectiveGroupUsingPOSTParams, opts ...ClientOption) (*SyncPerspectiveGroupUsingPOSTOK, *SyncPerspectiveGroupUsingPOSTAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncPerspectiveGroupUsingPOSTParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "syncPerspectiveGroupUsingPOST",
		Method:             "POST",
		PathPattern:        "/price/api/cloudhealth/perspective-sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SyncPerspectiveGroupUsingPOSTReader{formats: a.formats},
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
	case *SyncPerspectiveGroupUsingPOSTOK:
		return value, nil, nil
	case *SyncPerspectiveGroupUsingPOSTAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for perspective_sync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
