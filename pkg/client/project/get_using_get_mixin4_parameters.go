// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewGetUsingGETMixin4Params creates a new GetUsingGETMixin4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsingGETMixin4Params() *GetUsingGETMixin4Params {
	return &GetUsingGETMixin4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsingGETMixin4ParamsWithTimeout creates a new GetUsingGETMixin4Params object
// with the ability to set a timeout on a request.
func NewGetUsingGETMixin4ParamsWithTimeout(timeout time.Duration) *GetUsingGETMixin4Params {
	return &GetUsingGETMixin4Params{
		timeout: timeout,
	}
}

// NewGetUsingGETMixin4ParamsWithContext creates a new GetUsingGETMixin4Params object
// with the ability to set a context for a request.
func NewGetUsingGETMixin4ParamsWithContext(ctx context.Context) *GetUsingGETMixin4Params {
	return &GetUsingGETMixin4Params{
		Context: ctx,
	}
}

// NewGetUsingGETMixin4ParamsWithHTTPClient creates a new GetUsingGETMixin4Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsingGETMixin4ParamsWithHTTPClient(client *http.Client) *GetUsingGETMixin4Params {
	return &GetUsingGETMixin4Params{
		HTTPClient: client,
	}
}

/* GetUsingGETMixin4Params contains all the parameters to send to the API endpoint
   for the get using g e t mixin4 operation.

   Typically these are written to a http.Request.
*/
type GetUsingGETMixin4Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format. For versioning information refer to /project-service/api/about.
	*/
	APIVersion *string

	/* ID.

	   Тhe id of the project.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get using g e t mixin4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGETMixin4Params) WithDefaults() *GetUsingGETMixin4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get using g e t mixin4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGETMixin4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) WithTimeout(timeout time.Duration) *GetUsingGETMixin4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) WithContext(ctx context.Context) *GetUsingGETMixin4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) WithHTTPClient(client *http.Client) *GetUsingGETMixin4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) WithAPIVersion(aPIVersion *string) *GetUsingGETMixin4Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) WithID(id string) *GetUsingGETMixin4Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get using g e t mixin4 params
func (o *GetUsingGETMixin4Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsingGETMixin4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string

		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {

			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}
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
