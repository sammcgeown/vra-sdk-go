// Code generated by go-swagger; DO NOT EDIT.

package supervisor_namespaces

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

// NewCreateUsingPOST2Mixin5Params creates a new CreateUsingPOST2Mixin5Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUsingPOST2Mixin5Params() *CreateUsingPOST2Mixin5Params {
	return &CreateUsingPOST2Mixin5Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUsingPOST2Mixin5ParamsWithTimeout creates a new CreateUsingPOST2Mixin5Params object
// with the ability to set a timeout on a request.
func NewCreateUsingPOST2Mixin5ParamsWithTimeout(timeout time.Duration) *CreateUsingPOST2Mixin5Params {
	return &CreateUsingPOST2Mixin5Params{
		timeout: timeout,
	}
}

// NewCreateUsingPOST2Mixin5ParamsWithContext creates a new CreateUsingPOST2Mixin5Params object
// with the ability to set a context for a request.
func NewCreateUsingPOST2Mixin5ParamsWithContext(ctx context.Context) *CreateUsingPOST2Mixin5Params {
	return &CreateUsingPOST2Mixin5Params{
		Context: ctx,
	}
}

// NewCreateUsingPOST2Mixin5ParamsWithHTTPClient creates a new CreateUsingPOST2Mixin5Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUsingPOST2Mixin5ParamsWithHTTPClient(client *http.Client) *CreateUsingPOST2Mixin5Params {
	return &CreateUsingPOST2Mixin5Params{
		HTTPClient: client,
	}
}

/* CreateUsingPOST2Mixin5Params contains all the parameters to send to the API endpoint
   for the create using p o s t 2 mixin5 operation.

   Typically these are written to a http.Request.
*/
type CreateUsingPOST2Mixin5Params struct {

	/* Dto.

	   dto
	*/
	Dto *models.SupervisorNamespaceCreateDTO

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create using p o s t 2 mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOST2Mixin5Params) WithDefaults() *CreateUsingPOST2Mixin5Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create using p o s t 2 mixin5 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOST2Mixin5Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create using p o s t 2 mixin5 params
func (o *CreateUsingPOST2Mixin5Params) WithTimeout(timeout time.Duration) *CreateUsingPOST2Mixin5Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create using p o s t 2 mixin5 params
func (o *CreateUsingPOST2Mixin5Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create using p o s t 2 mixin5 params
func (o *CreateUsingPOST2Mixin5Params) WithContext(ctx context.Context) *CreateUsingPOST2Mixin5Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create using p o s t 2 mixin5 params
func (o *CreateUsingPOST2Mixin5Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create using p o s t 2 mixin5 params
func (o *CreateUsingPOST2Mixin5Params) WithHTTPClient(client *http.Client) *CreateUsingPOST2Mixin5Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create using p o s t 2 mixin5 params
func (o *CreateUsingPOST2Mixin5Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDto adds the dto to the create using p o s t 2 mixin5 params
func (o *CreateUsingPOST2Mixin5Params) WithDto(dto *models.SupervisorNamespaceCreateDTO) *CreateUsingPOST2Mixin5Params {
	o.SetDto(dto)
	return o
}

// SetDto adds the dto to the create using p o s t 2 mixin5 params
func (o *CreateUsingPOST2Mixin5Params) SetDto(dto *models.SupervisorNamespaceCreateDTO) {
	o.Dto = dto
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUsingPOST2Mixin5Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Dto != nil {
		if err := r.SetBodyParam(o.Dto); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}