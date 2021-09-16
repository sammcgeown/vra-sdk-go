// Code generated by go-swagger; DO NOT EDIT.

package provider_requests

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewPatchUsingPATCHMixin5Params creates a new PatchUsingPATCHMixin5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchUsingPATCHMixin5Params() *PatchUsingPATCHMixin5Params {
	return &PatchUsingPATCHMixin5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUsingPATCHMixin5ParamsWithTimeout creates a new PatchUsingPATCHMixin5Params object
// with the ability to set a timeout on a request.
func NewPatchUsingPATCHMixin5ParamsWithTimeout(timeout time.Duration) *PatchUsingPATCHMixin5Params {
	return &PatchUsingPATCHMixin5Params{
		timeout: timeout,
	}
}

// NewPatchUsingPATCHMixin5ParamsWithContext creates a new PatchUsingPATCHMixin5Params object
// with the ability to set a context for a request.
func NewPatchUsingPATCHMixin5ParamsWithContext(ctx context.Context) *PatchUsingPATCHMixin5Params {
	return &PatchUsingPATCHMixin5Params{
		Context: ctx,
	}
}

// NewPatchUsingPATCHMixin5ParamsWithHTTPClient creates a new PatchUsingPATCHMixin5Params object
// with the ability to set a custom HTTPClient for a request.
func NewPatchUsingPATCHMixin5ParamsWithHTTPClient(client *http.Client) *PatchUsingPATCHMixin5Params {
	return &PatchUsingPATCHMixin5Params{
		HTTPClient: client,
	}
}

/* PatchUsingPATCHMixin5Params contains all the parameters to send to the API endpoint
   for the patch using p a t c h mixin5 operation.

   Typically these are written to a http.Request.
*/
type PatchUsingPATCHMixin5Params struct {

	/* Rr.

	   rr
	*/
	Rr *models.ResumeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch using p a t c h mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchUsingPATCHMixin5Params) WithDefaults() *PatchUsingPATCHMixin5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch using p a t c h mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchUsingPATCHMixin5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch using p a t c h mixin5 params
func (o *PatchUsingPATCHMixin5Params) WithTimeout(timeout time.Duration) *PatchUsingPATCHMixin5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch using p a t c h mixin5 params
func (o *PatchUsingPATCHMixin5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch using p a t c h mixin5 params
func (o *PatchUsingPATCHMixin5Params) WithContext(ctx context.Context) *PatchUsingPATCHMixin5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch using p a t c h mixin5 params
func (o *PatchUsingPATCHMixin5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch using p a t c h mixin5 params
func (o *PatchUsingPATCHMixin5Params) WithHTTPClient(client *http.Client) *PatchUsingPATCHMixin5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch using p a t c h mixin5 params
func (o *PatchUsingPATCHMixin5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRr adds the rr to the patch using p a t c h mixin5 params
func (o *PatchUsingPATCHMixin5Params) WithRr(rr *models.ResumeRequest) *PatchUsingPATCHMixin5Params {
	o.SetRr(rr)
	return o
}

// SetRr adds the rr to the patch using p a t c h mixin5 params
func (o *PatchUsingPATCHMixin5Params) SetRr(rr *models.ResumeRequest) {
	o.Rr = rr
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUsingPATCHMixin5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Rr != nil {
		if err := r.SetBodyParam(o.Rr); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
