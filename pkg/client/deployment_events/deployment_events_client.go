// Code generated by go-swagger; DO NOT EDIT.

package deployment_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new deployment events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployment events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDeploymentEventsUsingGET fetches deployment events

Returns the events for the deployment.
*/
func (a *Client) GetDeploymentEventsUsingGET(params *GetDeploymentEventsUsingGETParams) (*GetDeploymentEventsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploymentEventsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDeploymentEventsUsingGET",
		Method:             "GET",
		PathPattern:        "/deployment/api/deployments/{depId}/events",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeploymentEventsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDeploymentEventsUsingGETOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
