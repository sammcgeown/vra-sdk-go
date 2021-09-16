// Code generated by go-swagger; DO NOT EDIT.

package p_k_s_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDestroyClusterUsingDELETEParams creates a new DestroyClusterUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDestroyClusterUsingDELETEParams() *DestroyClusterUsingDELETEParams {
	return &DestroyClusterUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDestroyClusterUsingDELETEParamsWithTimeout creates a new DestroyClusterUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDestroyClusterUsingDELETEParamsWithTimeout(timeout time.Duration) *DestroyClusterUsingDELETEParams {
	return &DestroyClusterUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDestroyClusterUsingDELETEParamsWithContext creates a new DestroyClusterUsingDELETEParams object
// with the ability to set a context for a request.
func NewDestroyClusterUsingDELETEParamsWithContext(ctx context.Context) *DestroyClusterUsingDELETEParams {
	return &DestroyClusterUsingDELETEParams{
		Context: ctx,
	}
}

// NewDestroyClusterUsingDELETEParamsWithHTTPClient creates a new DestroyClusterUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDestroyClusterUsingDELETEParamsWithHTTPClient(client *http.Client) *DestroyClusterUsingDELETEParams {
	return &DestroyClusterUsingDELETEParams{
		HTTPClient: client,
	}
}

/* DestroyClusterUsingDELETEParams contains all the parameters to send to the API endpoint
   for the destroy cluster using d e l e t e operation.

   Typically these are written to a http.Request.
*/
type DestroyClusterUsingDELETEParams struct {

	/* ClusterID.

	   clusterId
	*/
	ClusterID string

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the destroy cluster using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DestroyClusterUsingDELETEParams) WithDefaults() *DestroyClusterUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the destroy cluster using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DestroyClusterUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) WithTimeout(timeout time.Duration) *DestroyClusterUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) WithContext(ctx context.Context) *DestroyClusterUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) WithHTTPClient(client *http.Client) *DestroyClusterUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) WithClusterID(clusterID string) *DestroyClusterUsingDELETEParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithID adds the id to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) WithID(id string) *DestroyClusterUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the destroy cluster using d e l e t e params
func (o *DestroyClusterUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DestroyClusterUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", o.ClusterID); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
