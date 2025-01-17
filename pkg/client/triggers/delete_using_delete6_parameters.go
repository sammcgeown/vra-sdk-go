// Code generated by go-swagger; DO NOT EDIT.

package triggers

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

// NewDeleteUsingDELETE6Params creates a new DeleteUsingDELETE6Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUsingDELETE6Params() *DeleteUsingDELETE6Params {
	return &DeleteUsingDELETE6Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsingDELETE6ParamsWithTimeout creates a new DeleteUsingDELETE6Params object
// with the ability to set a timeout on a request.
func NewDeleteUsingDELETE6ParamsWithTimeout(timeout time.Duration) *DeleteUsingDELETE6Params {
	return &DeleteUsingDELETE6Params{
		timeout: timeout,
	}
}

// NewDeleteUsingDELETE6ParamsWithContext creates a new DeleteUsingDELETE6Params object
// with the ability to set a context for a request.
func NewDeleteUsingDELETE6ParamsWithContext(ctx context.Context) *DeleteUsingDELETE6Params {
	return &DeleteUsingDELETE6Params{
		Context: ctx,
	}
}

// NewDeleteUsingDELETE6ParamsWithHTTPClient creates a new DeleteUsingDELETE6Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUsingDELETE6ParamsWithHTTPClient(client *http.Client) *DeleteUsingDELETE6Params {
	return &DeleteUsingDELETE6Params{
		HTTPClient: client,
	}
}

/* DeleteUsingDELETE6Params contains all the parameters to send to the API endpoint
   for the delete using d e l e t e 6 operation.

   Typically these are written to a http.Request.
*/
type DeleteUsingDELETE6Params struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete using d e l e t e 6 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE6Params) WithDefaults() *DeleteUsingDELETE6Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete using d e l e t e 6 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE6Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) WithTimeout(timeout time.Duration) *DeleteUsingDELETE6Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) WithContext(ctx context.Context) *DeleteUsingDELETE6Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) WithHTTPClient(client *http.Client) *DeleteUsingDELETE6Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) WithAuthorization(authorization string) *DeleteUsingDELETE6Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) WithAPIVersion(aPIVersion string) *DeleteUsingDELETE6Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) WithID(id string) *DeleteUsingDELETE6Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete using d e l e t e 6 params
func (o *DeleteUsingDELETE6Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsingDELETE6Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
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
