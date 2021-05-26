// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

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

// NewGetAwsStorageProfileParams creates a new GetAwsStorageProfileParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAwsStorageProfileParams() *GetAwsStorageProfileParams {
	return &GetAwsStorageProfileParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAwsStorageProfileParamsWithTimeout creates a new GetAwsStorageProfileParams object
// with the ability to set a timeout on a request.
func NewGetAwsStorageProfileParamsWithTimeout(timeout time.Duration) *GetAwsStorageProfileParams {
	return &GetAwsStorageProfileParams{
		timeout: timeout,
	}
}

// NewGetAwsStorageProfileParamsWithContext creates a new GetAwsStorageProfileParams object
// with the ability to set a context for a request.
func NewGetAwsStorageProfileParamsWithContext(ctx context.Context) *GetAwsStorageProfileParams {
	return &GetAwsStorageProfileParams{
		Context: ctx,
	}
}

// NewGetAwsStorageProfileParamsWithHTTPClient creates a new GetAwsStorageProfileParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAwsStorageProfileParamsWithHTTPClient(client *http.Client) *GetAwsStorageProfileParams {
	return &GetAwsStorageProfileParams{
		HTTPClient: client,
	}
}

/* GetAwsStorageProfileParams contains all the parameters to send to the API endpoint
   for the get aws storage profile operation.

   Typically these are written to a http.Request.
*/
type GetAwsStorageProfileParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of storage profile.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get aws storage profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAwsStorageProfileParams) WithDefaults() *GetAwsStorageProfileParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get aws storage profile params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAwsStorageProfileParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get aws storage profile params
func (o *GetAwsStorageProfileParams) WithTimeout(timeout time.Duration) *GetAwsStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get aws storage profile params
func (o *GetAwsStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get aws storage profile params
func (o *GetAwsStorageProfileParams) WithContext(ctx context.Context) *GetAwsStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get aws storage profile params
func (o *GetAwsStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get aws storage profile params
func (o *GetAwsStorageProfileParams) WithHTTPClient(client *http.Client) *GetAwsStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get aws storage profile params
func (o *GetAwsStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get aws storage profile params
func (o *GetAwsStorageProfileParams) WithAPIVersion(aPIVersion *string) *GetAwsStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get aws storage profile params
func (o *GetAwsStorageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get aws storage profile params
func (o *GetAwsStorageProfileParams) WithID(id string) *GetAwsStorageProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get aws storage profile params
func (o *GetAwsStorageProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAwsStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
