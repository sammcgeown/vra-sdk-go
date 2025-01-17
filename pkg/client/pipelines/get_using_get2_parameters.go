// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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

// NewGetUsingGET2Params creates a new GetUsingGET2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsingGET2Params() *GetUsingGET2Params {
	return &GetUsingGET2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsingGET2ParamsWithTimeout creates a new GetUsingGET2Params object
// with the ability to set a timeout on a request.
func NewGetUsingGET2ParamsWithTimeout(timeout time.Duration) *GetUsingGET2Params {
	return &GetUsingGET2Params{
		timeout: timeout,
	}
}

// NewGetUsingGET2ParamsWithContext creates a new GetUsingGET2Params object
// with the ability to set a context for a request.
func NewGetUsingGET2ParamsWithContext(ctx context.Context) *GetUsingGET2Params {
	return &GetUsingGET2Params{
		Context: ctx,
	}
}

// NewGetUsingGET2ParamsWithHTTPClient creates a new GetUsingGET2Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsingGET2ParamsWithHTTPClient(client *http.Client) *GetUsingGET2Params {
	return &GetUsingGET2Params{
		HTTPClient: client,
	}
}

/* GetUsingGET2Params contains all the parameters to send to the API endpoint
   for the get using g e t 2 operation.

   Typically these are written to a http.Request.
*/
type GetUsingGET2Params struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   The ID of the Pipeline
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGET2Params) WithDefaults() *GetUsingGET2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGET2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get using g e t 2 params
func (o *GetUsingGET2Params) WithTimeout(timeout time.Duration) *GetUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get using g e t 2 params
func (o *GetUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get using g e t 2 params
func (o *GetUsingGET2Params) WithContext(ctx context.Context) *GetUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get using g e t 2 params
func (o *GetUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get using g e t 2 params
func (o *GetUsingGET2Params) WithHTTPClient(client *http.Client) *GetUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get using g e t 2 params
func (o *GetUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get using g e t 2 params
func (o *GetUsingGET2Params) WithAuthorization(authorization string) *GetUsingGET2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get using g e t 2 params
func (o *GetUsingGET2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get using g e t 2 params
func (o *GetUsingGET2Params) WithAPIVersion(aPIVersion string) *GetUsingGET2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get using g e t 2 params
func (o *GetUsingGET2Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get using g e t 2 params
func (o *GetUsingGET2Params) WithID(id string) *GetUsingGET2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get using g e t 2 params
func (o *GetUsingGET2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
