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

// NewDeleteTerraformVersionUsingDELETE1Params creates a new DeleteTerraformVersionUsingDELETE1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTerraformVersionUsingDELETE1Params() *DeleteTerraformVersionUsingDELETE1Params {
	return &DeleteTerraformVersionUsingDELETE1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTerraformVersionUsingDELETE1ParamsWithTimeout creates a new DeleteTerraformVersionUsingDELETE1Params object
// with the ability to set a timeout on a request.
func NewDeleteTerraformVersionUsingDELETE1ParamsWithTimeout(timeout time.Duration) *DeleteTerraformVersionUsingDELETE1Params {
	return &DeleteTerraformVersionUsingDELETE1Params{
		timeout: timeout,
	}
}

// NewDeleteTerraformVersionUsingDELETE1ParamsWithContext creates a new DeleteTerraformVersionUsingDELETE1Params object
// with the ability to set a context for a request.
func NewDeleteTerraformVersionUsingDELETE1ParamsWithContext(ctx context.Context) *DeleteTerraformVersionUsingDELETE1Params {
	return &DeleteTerraformVersionUsingDELETE1Params{
		Context: ctx,
	}
}

// NewDeleteTerraformVersionUsingDELETE1ParamsWithHTTPClient creates a new DeleteTerraformVersionUsingDELETE1Params object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTerraformVersionUsingDELETE1ParamsWithHTTPClient(client *http.Client) *DeleteTerraformVersionUsingDELETE1Params {
	return &DeleteTerraformVersionUsingDELETE1Params{
		HTTPClient: client,
	}
}

/* DeleteTerraformVersionUsingDELETE1Params contains all the parameters to send to the API endpoint
   for the delete terraform version using d e l e t e 1 operation.

   Typically these are written to a http.Request.
*/
type DeleteTerraformVersionUsingDELETE1Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* VersionID.

	   versionId

	   Format: uuid
	*/
	VersionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete terraform version using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTerraformVersionUsingDELETE1Params) WithDefaults() *DeleteTerraformVersionUsingDELETE1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete terraform version using d e l e t e 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTerraformVersionUsingDELETE1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) WithTimeout(timeout time.Duration) *DeleteTerraformVersionUsingDELETE1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) WithContext(ctx context.Context) *DeleteTerraformVersionUsingDELETE1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) WithHTTPClient(client *http.Client) *DeleteTerraformVersionUsingDELETE1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) WithAPIVersion(aPIVersion *string) *DeleteTerraformVersionUsingDELETE1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithVersionID adds the versionID to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) WithVersionID(versionID strfmt.UUID) *DeleteTerraformVersionUsingDELETE1Params {
	o.SetVersionID(versionID)
	return o
}

// SetVersionID adds the versionId to the delete terraform version using d e l e t e 1 params
func (o *DeleteTerraformVersionUsingDELETE1Params) SetVersionID(versionID strfmt.UUID) {
	o.VersionID = versionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTerraformVersionUsingDELETE1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param versionId
	if err := r.SetPathParam("versionId", o.VersionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}