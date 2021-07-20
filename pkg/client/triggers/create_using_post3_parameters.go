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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateUsingPOST3Params creates a new CreateUsingPOST3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateUsingPOST3Params() *CreateUsingPOST3Params {
	return &CreateUsingPOST3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateUsingPOST3ParamsWithTimeout creates a new CreateUsingPOST3Params object
// with the ability to set a timeout on a request.
func NewCreateUsingPOST3ParamsWithTimeout(timeout time.Duration) *CreateUsingPOST3Params {
	return &CreateUsingPOST3Params{
		timeout: timeout,
	}
}

// NewCreateUsingPOST3ParamsWithContext creates a new CreateUsingPOST3Params object
// with the ability to set a context for a request.
func NewCreateUsingPOST3ParamsWithContext(ctx context.Context) *CreateUsingPOST3Params {
	return &CreateUsingPOST3Params{
		Context: ctx,
	}
}

// NewCreateUsingPOST3ParamsWithHTTPClient creates a new CreateUsingPOST3Params object
// with the ability to set a custom HTTPClient for a request.
func NewCreateUsingPOST3ParamsWithHTTPClient(client *http.Client) *CreateUsingPOST3Params {
	return &CreateUsingPOST3Params{
		HTTPClient: client,
	}
}

/* CreateUsingPOST3Params contains all the parameters to send to the API endpoint
   for the create using p o s t 3 operation.

   Typically these are written to a http.Request.
*/
type CreateUsingPOST3Params struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* GerritListenerspec.

	   gerritListenerspec
	*/
	GerritListenerspec models.GerritListenerSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create using p o s t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOST3Params) WithDefaults() *CreateUsingPOST3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create using p o s t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateUsingPOST3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) WithTimeout(timeout time.Duration) *CreateUsingPOST3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) WithContext(ctx context.Context) *CreateUsingPOST3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) WithHTTPClient(client *http.Client) *CreateUsingPOST3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) WithAuthorization(authorization string) *CreateUsingPOST3Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) WithAPIVersion(aPIVersion string) *CreateUsingPOST3Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGerritListenerspec adds the gerritListenerspec to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) WithGerritListenerspec(gerritListenerspec models.GerritListenerSpec) *CreateUsingPOST3Params {
	o.SetGerritListenerspec(gerritListenerspec)
	return o
}

// SetGerritListenerspec adds the gerritListenerspec to the create using p o s t 3 params
func (o *CreateUsingPOST3Params) SetGerritListenerspec(gerritListenerspec models.GerritListenerSpec) {
	o.GerritListenerspec = gerritListenerspec
}

// WriteToRequest writes these params to a swagger request
func (o *CreateUsingPOST3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetBodyParam(o.GerritListenerspec); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}