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

// NewGetVmcCloudAccountParams creates a new GetVmcCloudAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVmcCloudAccountParams() *GetVmcCloudAccountParams {
	return &GetVmcCloudAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVmcCloudAccountParamsWithTimeout creates a new GetVmcCloudAccountParams object
// with the ability to set a timeout on a request.
func NewGetVmcCloudAccountParamsWithTimeout(timeout time.Duration) *GetVmcCloudAccountParams {
	return &GetVmcCloudAccountParams{
		timeout: timeout,
	}
}

// NewGetVmcCloudAccountParamsWithContext creates a new GetVmcCloudAccountParams object
// with the ability to set a context for a request.
func NewGetVmcCloudAccountParamsWithContext(ctx context.Context) *GetVmcCloudAccountParams {
	return &GetVmcCloudAccountParams{
		Context: ctx,
	}
}

// NewGetVmcCloudAccountParamsWithHTTPClient creates a new GetVmcCloudAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVmcCloudAccountParamsWithHTTPClient(client *http.Client) *GetVmcCloudAccountParams {
	return &GetVmcCloudAccountParams{
		HTTPClient: client,
	}
}

/* GetVmcCloudAccountParams contains all the parameters to send to the API endpoint
   for the get vmc cloud account operation.

   Typically these are written to a http.Request.
*/
type GetVmcCloudAccountParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the Cloud Account
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vmc cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVmcCloudAccountParams) WithDefaults() *GetVmcCloudAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vmc cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVmcCloudAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) WithTimeout(timeout time.Duration) *GetVmcCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) WithContext(ctx context.Context) *GetVmcCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) WithHTTPClient(client *http.Client) *GetVmcCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) WithAPIVersion(aPIVersion *string) *GetVmcCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) WithID(id string) *GetVmcCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get vmc cloud account params
func (o *GetVmcCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetVmcCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
