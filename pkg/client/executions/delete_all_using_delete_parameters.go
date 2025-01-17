// Code generated by go-swagger; DO NOT EDIT.

package executions

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

// NewDeleteAllUsingDELETEParams creates a new DeleteAllUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteAllUsingDELETEParams() *DeleteAllUsingDELETEParams {
	return &DeleteAllUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAllUsingDELETEParamsWithTimeout creates a new DeleteAllUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteAllUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteAllUsingDELETEParams {
	return &DeleteAllUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteAllUsingDELETEParamsWithContext creates a new DeleteAllUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteAllUsingDELETEParamsWithContext(ctx context.Context) *DeleteAllUsingDELETEParams {
	return &DeleteAllUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteAllUsingDELETEParamsWithHTTPClient creates a new DeleteAllUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteAllUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteAllUsingDELETEParams {
	return &DeleteAllUsingDELETEParams{
		HTTPClient: client,
	}
}

/* DeleteAllUsingDELETEParams contains all the parameters to send to the API endpoint
   for the delete all using d e l e t e operation.

   Typically these are written to a http.Request.
*/
type DeleteAllUsingDELETEParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete all using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllUsingDELETEParams) WithDefaults() *DeleteAllUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete all using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteAllUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteAllUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) WithContext(ctx context.Context) *DeleteAllUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteAllUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) WithAuthorization(authorization string) *DeleteAllUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) WithAPIVersion(aPIVersion string) *DeleteAllUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete all using d e l e t e params
func (o *DeleteAllUsingDELETEParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAllUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
