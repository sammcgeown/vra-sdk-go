// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

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

// NewGetTerraformVersionUsingGET1Params creates a new GetTerraformVersionUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTerraformVersionUsingGET1Params() *GetTerraformVersionUsingGET1Params {
	return &GetTerraformVersionUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTerraformVersionUsingGET1ParamsWithTimeout creates a new GetTerraformVersionUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetTerraformVersionUsingGET1ParamsWithTimeout(timeout time.Duration) *GetTerraformVersionUsingGET1Params {
	return &GetTerraformVersionUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetTerraformVersionUsingGET1ParamsWithContext creates a new GetTerraformVersionUsingGET1Params object
// with the ability to set a context for a request.
func NewGetTerraformVersionUsingGET1ParamsWithContext(ctx context.Context) *GetTerraformVersionUsingGET1Params {
	return &GetTerraformVersionUsingGET1Params{
		Context: ctx,
	}
}

// NewGetTerraformVersionUsingGET1ParamsWithHTTPClient creates a new GetTerraformVersionUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetTerraformVersionUsingGET1ParamsWithHTTPClient(client *http.Client) *GetTerraformVersionUsingGET1Params {
	return &GetTerraformVersionUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetTerraformVersionUsingGET1Params contains all the parameters to send to the API endpoint
   for the get terraform version using get1 operation.

   Typically these are written to a http.Request.
*/
type GetTerraformVersionUsingGET1Params struct {

	/* VersionID.

	   versionId

	   Format: uuid
	*/
	VersionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get terraform version using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformVersionUsingGET1Params) WithDefaults() *GetTerraformVersionUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get terraform version using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformVersionUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get terraform version using get1 params
func (o *GetTerraformVersionUsingGET1Params) WithTimeout(timeout time.Duration) *GetTerraformVersionUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get terraform version using get1 params
func (o *GetTerraformVersionUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get terraform version using get1 params
func (o *GetTerraformVersionUsingGET1Params) WithContext(ctx context.Context) *GetTerraformVersionUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get terraform version using get1 params
func (o *GetTerraformVersionUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get terraform version using get1 params
func (o *GetTerraformVersionUsingGET1Params) WithHTTPClient(client *http.Client) *GetTerraformVersionUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get terraform version using get1 params
func (o *GetTerraformVersionUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVersionID adds the versionID to the get terraform version using get1 params
func (o *GetTerraformVersionUsingGET1Params) WithVersionID(versionID strfmt.UUID) *GetTerraformVersionUsingGET1Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the get terraform version using get1 params
func (o *GetTerraformVersionUsingGET1Params) SetVersionID(versionID strfmt.UUID) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTerraformVersionUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
