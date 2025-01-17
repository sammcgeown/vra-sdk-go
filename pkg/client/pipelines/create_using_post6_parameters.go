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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateUsingPOST6Params creates a new CreateUsingPOST6Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUsingPOST6Params() *CreateUsingPOST6Params {
	return &CreateUsingPOST6Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUsingPOST6ParamsWithTimeout creates a new CreateUsingPOST6Params object
// with the ability to set a timeout on a request.
func NewCreateUsingPOST6ParamsWithTimeout(timeout time.Duration) *CreateUsingPOST6Params {
	return &CreateUsingPOST6Params{
		timeout: timeout,
	}
}

// NewCreateUsingPOST6ParamsWithContext creates a new CreateUsingPOST6Params object
// with the ability to set a context for a request.
func NewCreateUsingPOST6ParamsWithContext(ctx context.Context) *CreateUsingPOST6Params {
	return &CreateUsingPOST6Params{
		Context: ctx,
	}
}

// NewCreateUsingPOST6ParamsWithHTTPClient creates a new CreateUsingPOST6Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUsingPOST6ParamsWithHTTPClient(client *http.Client) *CreateUsingPOST6Params {
	return &CreateUsingPOST6Params{
		HTTPClient: client,
	}
}

/* CreateUsingPOST6Params contains all the parameters to send to the API endpoint
   for the create using p o s t 6 operation.

   Typically these are written to a http.Request.
*/
type CreateUsingPOST6Params struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* Body.

	   Pipeline specification
	*/
	Body models.PipelineSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create using p o s t 6 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOST6Params) WithDefaults() *CreateUsingPOST6Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create using p o s t 6 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOST6Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) WithTimeout(timeout time.Duration) *CreateUsingPOST6Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) WithContext(ctx context.Context) *CreateUsingPOST6Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) WithHTTPClient(client *http.Client) *CreateUsingPOST6Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) WithAuthorization(authorization string) *CreateUsingPOST6Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) WithAPIVersion(aPIVersion string) *CreateUsingPOST6Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) WithBody(body models.PipelineSpec) *CreateUsingPOST6Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create using p o s t 6 params
func (o *CreateUsingPOST6Params) SetBody(body models.PipelineSpec) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUsingPOST6Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
