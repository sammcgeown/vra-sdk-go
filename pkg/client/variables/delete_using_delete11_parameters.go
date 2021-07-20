// Code generated by go-swagger; DO NOT EDIT.

package variables

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

// NewDeleteUsingDELETE11Params creates a new DeleteUsingDELETE11Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteUsingDELETE11Params() *DeleteUsingDELETE11Params {
	return &DeleteUsingDELETE11Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUsingDELETE11ParamsWithTimeout creates a new DeleteUsingDELETE11Params object
// with the ability to set a timeout on a request.
func NewDeleteUsingDELETE11ParamsWithTimeout(timeout time.Duration) *DeleteUsingDELETE11Params {
	return &DeleteUsingDELETE11Params{
		timeout: timeout,
	}
}

// NewDeleteUsingDELETE11ParamsWithContext creates a new DeleteUsingDELETE11Params object
// with the ability to set a context for a request.
func NewDeleteUsingDELETE11ParamsWithContext(ctx context.Context) *DeleteUsingDELETE11Params {
	return &DeleteUsingDELETE11Params{
		Context: ctx,
	}
}

// NewDeleteUsingDELETE11ParamsWithHTTPClient creates a new DeleteUsingDELETE11Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteUsingDELETE11ParamsWithHTTPClient(client *http.Client) *DeleteUsingDELETE11Params {
	return &DeleteUsingDELETE11Params{
		HTTPClient: client,
	}
}

/* DeleteUsingDELETE11Params contains all the parameters to send to the API endpoint
   for the delete using d e l e t e 11 operation.

   Typically these are written to a http.Request.
*/
type DeleteUsingDELETE11Params struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   The ID of the Variable
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete using d e l e t e 11 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE11Params) WithDefaults() *DeleteUsingDELETE11Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete using d e l e t e 11 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteUsingDELETE11Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) WithTimeout(timeout time.Duration) *DeleteUsingDELETE11Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) WithContext(ctx context.Context) *DeleteUsingDELETE11Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) WithHTTPClient(client *http.Client) *DeleteUsingDELETE11Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) WithAuthorization(authorization string) *DeleteUsingDELETE11Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) WithAPIVersion(aPIVersion string) *DeleteUsingDELETE11Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) WithID(id string) *DeleteUsingDELETE11Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete using d e l e t e 11 params
func (o *DeleteUsingDELETE11Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUsingDELETE11Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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