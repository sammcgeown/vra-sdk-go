// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewEnumeratePrivateImagesVCFParams creates a new EnumeratePrivateImagesVCFParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEnumeratePrivateImagesVCFParams() *EnumeratePrivateImagesVCFParams {
	return &EnumeratePrivateImagesVCFParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEnumeratePrivateImagesVCFParamsWithTimeout creates a new EnumeratePrivateImagesVCFParams object
// with the ability to set a timeout on a request.
func NewEnumeratePrivateImagesVCFParamsWithTimeout(timeout time.Duration) *EnumeratePrivateImagesVCFParams {
	return &EnumeratePrivateImagesVCFParams{
		timeout: timeout,
	}
}

// NewEnumeratePrivateImagesVCFParamsWithContext creates a new EnumeratePrivateImagesVCFParams object
// with the ability to set a context for a request.
func NewEnumeratePrivateImagesVCFParamsWithContext(ctx context.Context) *EnumeratePrivateImagesVCFParams {
	return &EnumeratePrivateImagesVCFParams{
		Context: ctx,
	}
}

// NewEnumeratePrivateImagesVCFParamsWithHTTPClient creates a new EnumeratePrivateImagesVCFParams object
// with the ability to set a custom HTTPClient for a request.
func NewEnumeratePrivateImagesVCFParamsWithHTTPClient(client *http.Client) *EnumeratePrivateImagesVCFParams {
	return &EnumeratePrivateImagesVCFParams{
		HTTPClient: client,
	}
}

/* EnumeratePrivateImagesVCFParams contains all the parameters to send to the API endpoint
   for the enumerate private images v c f operation.

   Typically these are written to a http.Request.
*/
type EnumeratePrivateImagesVCFParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   Id of VCF cloud account to enumerate
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the enumerate private images v c f params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumeratePrivateImagesVCFParams) WithDefaults() *EnumeratePrivateImagesVCFParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the enumerate private images v c f params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EnumeratePrivateImagesVCFParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) WithTimeout(timeout time.Duration) *EnumeratePrivateImagesVCFParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) WithContext(ctx context.Context) *EnumeratePrivateImagesVCFParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) WithHTTPClient(client *http.Client) *EnumeratePrivateImagesVCFParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) WithAPIVersion(aPIVersion *string) *EnumeratePrivateImagesVCFParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) WithID(id string) *EnumeratePrivateImagesVCFParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enumerate private images v c f params
func (o *EnumeratePrivateImagesVCFParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EnumeratePrivateImagesVCFParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
