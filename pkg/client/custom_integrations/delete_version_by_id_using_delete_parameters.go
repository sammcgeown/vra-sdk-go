// Code generated by go-swagger; DO NOT EDIT.

package custom_integrations

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

// NewDeleteVersionByIDUsingDELETEParams creates a new DeleteVersionByIDUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVersionByIDUsingDELETEParams() *DeleteVersionByIDUsingDELETEParams {
	return &DeleteVersionByIDUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVersionByIDUsingDELETEParamsWithTimeout creates a new DeleteVersionByIDUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteVersionByIDUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteVersionByIDUsingDELETEParams {
	return &DeleteVersionByIDUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteVersionByIDUsingDELETEParamsWithContext creates a new DeleteVersionByIDUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteVersionByIDUsingDELETEParamsWithContext(ctx context.Context) *DeleteVersionByIDUsingDELETEParams {
	return &DeleteVersionByIDUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteVersionByIDUsingDELETEParamsWithHTTPClient creates a new DeleteVersionByIDUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVersionByIDUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteVersionByIDUsingDELETEParams {
	return &DeleteVersionByIDUsingDELETEParams{
		HTTPClient: client,
	}
}

/* DeleteVersionByIDUsingDELETEParams contains all the parameters to send to the API endpoint
   for the delete version by Id using d e l e t e operation.

   Typically these are written to a http.Request.
*/
type DeleteVersionByIDUsingDELETEParams struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   The ID of the Custom Integration
	*/
	ID string

	/* Version.

	   The version of the Custom Integration
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete version by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVersionByIDUsingDELETEParams) WithDefaults() *DeleteVersionByIDUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete version by Id using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVersionByIDUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteVersionByIDUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) WithContext(ctx context.Context) *DeleteVersionByIDUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteVersionByIDUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) WithAuthorization(authorization string) *DeleteVersionByIDUsingDELETEParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) WithAPIVersion(aPIVersion string) *DeleteVersionByIDUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) WithID(id string) *DeleteVersionByIDUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) WithVersion(version string) *DeleteVersionByIDUsingDELETEParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete version by Id using d e l e t e params
func (o *DeleteVersionByIDUsingDELETEParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVersionByIDUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}