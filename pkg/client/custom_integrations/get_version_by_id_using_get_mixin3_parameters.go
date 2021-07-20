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

// NewGetVersionByIDUsingGETMixin3Params creates a new GetVersionByIDUsingGETMixin3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVersionByIDUsingGETMixin3Params() *GetVersionByIDUsingGETMixin3Params {
	return &GetVersionByIDUsingGETMixin3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionByIDUsingGETMixin3ParamsWithTimeout creates a new GetVersionByIDUsingGETMixin3Params object
// with the ability to set a timeout on a request.
func NewGetVersionByIDUsingGETMixin3ParamsWithTimeout(timeout time.Duration) *GetVersionByIDUsingGETMixin3Params {
	return &GetVersionByIDUsingGETMixin3Params{
		timeout: timeout,
	}
}

// NewGetVersionByIDUsingGETMixin3ParamsWithContext creates a new GetVersionByIDUsingGETMixin3Params object
// with the ability to set a context for a request.
func NewGetVersionByIDUsingGETMixin3ParamsWithContext(ctx context.Context) *GetVersionByIDUsingGETMixin3Params {
	return &GetVersionByIDUsingGETMixin3Params{
		Context: ctx,
	}
}

// NewGetVersionByIDUsingGETMixin3ParamsWithHTTPClient creates a new GetVersionByIDUsingGETMixin3Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetVersionByIDUsingGETMixin3ParamsWithHTTPClient(client *http.Client) *GetVersionByIDUsingGETMixin3Params {
	return &GetVersionByIDUsingGETMixin3Params{
		HTTPClient: client,
	}
}

/* GetVersionByIDUsingGETMixin3Params contains all the parameters to send to the API endpoint
   for the get version by Id using g e t mixin3 operation.

   Typically these are written to a http.Request.
*/
type GetVersionByIDUsingGETMixin3Params struct {

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

// WithDefaults hydrates default values in the get version by Id using g e t mixin3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionByIDUsingGETMixin3Params) WithDefaults() *GetVersionByIDUsingGETMixin3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get version by Id using g e t mixin3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVersionByIDUsingGETMixin3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) WithTimeout(timeout time.Duration) *GetVersionByIDUsingGETMixin3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) WithContext(ctx context.Context) *GetVersionByIDUsingGETMixin3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) WithHTTPClient(client *http.Client) *GetVersionByIDUsingGETMixin3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) WithAuthorization(authorization string) *GetVersionByIDUsingGETMixin3Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) WithAPIVersion(aPIVersion string) *GetVersionByIDUsingGETMixin3Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) WithID(id string) *GetVersionByIDUsingGETMixin3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) WithVersion(version string) *GetVersionByIDUsingGETMixin3Params {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get version by Id using g e t mixin3 params
func (o *GetVersionByIDUsingGETMixin3Params) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionByIDUsingGETMixin3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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