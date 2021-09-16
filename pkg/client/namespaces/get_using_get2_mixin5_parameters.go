// Code generated by go-swagger; DO NOT EDIT.

package namespaces

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

// NewGetUsingGET2Mixin5Params creates a new GetUsingGET2Mixin5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUsingGET2Mixin5Params() *GetUsingGET2Mixin5Params {
	return &GetUsingGET2Mixin5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsingGET2Mixin5ParamsWithTimeout creates a new GetUsingGET2Mixin5Params object
// with the ability to set a timeout on a request.
func NewGetUsingGET2Mixin5ParamsWithTimeout(timeout time.Duration) *GetUsingGET2Mixin5Params {
	return &GetUsingGET2Mixin5Params{
		timeout: timeout,
	}
}

// NewGetUsingGET2Mixin5ParamsWithContext creates a new GetUsingGET2Mixin5Params object
// with the ability to set a context for a request.
func NewGetUsingGET2Mixin5ParamsWithContext(ctx context.Context) *GetUsingGET2Mixin5Params {
	return &GetUsingGET2Mixin5Params{
		Context: ctx,
	}
}

// NewGetUsingGET2Mixin5ParamsWithHTTPClient creates a new GetUsingGET2Mixin5Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetUsingGET2Mixin5ParamsWithHTTPClient(client *http.Client) *GetUsingGET2Mixin5Params {
	return &GetUsingGET2Mixin5Params{
		HTTPClient: client,
	}
}

/* GetUsingGET2Mixin5Params contains all the parameters to send to the API endpoint
   for the get using g e t 2 mixin5 operation.

   Typically these are written to a http.Request.
*/
type GetUsingGET2Mixin5Params struct {

	/* ID.

	   id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get using g e t 2 mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGET2Mixin5Params) WithDefaults() *GetUsingGET2Mixin5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get using g e t 2 mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUsingGET2Mixin5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get using g e t 2 mixin5 params
func (o *GetUsingGET2Mixin5Params) WithTimeout(timeout time.Duration) *GetUsingGET2Mixin5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get using g e t 2 mixin5 params
func (o *GetUsingGET2Mixin5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get using g e t 2 mixin5 params
func (o *GetUsingGET2Mixin5Params) WithContext(ctx context.Context) *GetUsingGET2Mixin5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get using g e t 2 mixin5 params
func (o *GetUsingGET2Mixin5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get using g e t 2 mixin5 params
func (o *GetUsingGET2Mixin5Params) WithHTTPClient(client *http.Client) *GetUsingGET2Mixin5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get using g e t 2 mixin5 params
func (o *GetUsingGET2Mixin5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get using g e t 2 mixin5 params
func (o *GetUsingGET2Mixin5Params) WithID(id strfmt.UUID) *GetUsingGET2Mixin5Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get using g e t 2 mixin5 params
func (o *GetUsingGET2Mixin5Params) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsingGET2Mixin5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
