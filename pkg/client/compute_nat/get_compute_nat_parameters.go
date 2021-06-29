// Code generated by go-swagger; DO NOT EDIT.

package compute_nat

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

// NewGetComputeNatParams creates a new GetComputeNatParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetComputeNatParams() *GetComputeNatParams {
	return &GetComputeNatParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetComputeNatParamsWithTimeout creates a new GetComputeNatParams object
// with the ability to set a timeout on a request.
func NewGetComputeNatParamsWithTimeout(timeout time.Duration) *GetComputeNatParams {
	return &GetComputeNatParams{
		timeout: timeout,
	}
}

// NewGetComputeNatParamsWithContext creates a new GetComputeNatParams object
// with the ability to set a context for a request.
func NewGetComputeNatParamsWithContext(ctx context.Context) *GetComputeNatParams {
	return &GetComputeNatParams{
		Context: ctx,
	}
}

// NewGetComputeNatParamsWithHTTPClient creates a new GetComputeNatParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetComputeNatParamsWithHTTPClient(client *http.Client) *GetComputeNatParams {
	return &GetComputeNatParams{
		HTTPClient: client,
	}
}

/* GetComputeNatParams contains all the parameters to send to the API endpoint
   for the get compute nat operation.

   Typically these are written to a http.Request.
*/
type GetComputeNatParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the Compute Nat resource.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get compute nat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComputeNatParams) WithDefaults() *GetComputeNatParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get compute nat params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetComputeNatParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get compute nat params
func (o *GetComputeNatParams) WithTimeout(timeout time.Duration) *GetComputeNatParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get compute nat params
func (o *GetComputeNatParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get compute nat params
func (o *GetComputeNatParams) WithContext(ctx context.Context) *GetComputeNatParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get compute nat params
func (o *GetComputeNatParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get compute nat params
func (o *GetComputeNatParams) WithHTTPClient(client *http.Client) *GetComputeNatParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get compute nat params
func (o *GetComputeNatParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get compute nat params
func (o *GetComputeNatParams) WithAPIVersion(aPIVersion *string) *GetComputeNatParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get compute nat params
func (o *GetComputeNatParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get compute nat params
func (o *GetComputeNatParams) WithID(id string) *GetComputeNatParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get compute nat params
func (o *GetComputeNatParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetComputeNatParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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