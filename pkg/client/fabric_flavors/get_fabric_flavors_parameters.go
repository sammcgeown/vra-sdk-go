// Code generated by go-swagger; DO NOT EDIT.

package fabric_flavors

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

// NewGetFabricFlavorsParams creates a new GetFabricFlavorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFabricFlavorsParams() *GetFabricFlavorsParams {
	return &GetFabricFlavorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricFlavorsParamsWithTimeout creates a new GetFabricFlavorsParams object
// with the ability to set a timeout on a request.
func NewGetFabricFlavorsParamsWithTimeout(timeout time.Duration) *GetFabricFlavorsParams {
	return &GetFabricFlavorsParams{
		timeout: timeout,
	}
}

// NewGetFabricFlavorsParamsWithContext creates a new GetFabricFlavorsParams object
// with the ability to set a context for a request.
func NewGetFabricFlavorsParamsWithContext(ctx context.Context) *GetFabricFlavorsParams {
	return &GetFabricFlavorsParams{
		Context: ctx,
	}
}

// NewGetFabricFlavorsParamsWithHTTPClient creates a new GetFabricFlavorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFabricFlavorsParamsWithHTTPClient(client *http.Client) *GetFabricFlavorsParams {
	return &GetFabricFlavorsParams{
		HTTPClient: client,
	}
}

/* GetFabricFlavorsParams contains all the parameters to send to the API endpoint
   for the get fabric flavors operation.

   Typically these are written to a http.Request.
*/
type GetFabricFlavorsParams struct {

	/* DollarFilter.

	   Add a filter to return limited results
	*/
	DollarFilter *string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get fabric flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFabricFlavorsParams) WithDefaults() *GetFabricFlavorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get fabric flavors params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFabricFlavorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get fabric flavors params
func (o *GetFabricFlavorsParams) WithTimeout(timeout time.Duration) *GetFabricFlavorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fabric flavors params
func (o *GetFabricFlavorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fabric flavors params
func (o *GetFabricFlavorsParams) WithContext(ctx context.Context) *GetFabricFlavorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fabric flavors params
func (o *GetFabricFlavorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fabric flavors params
func (o *GetFabricFlavorsParams) WithHTTPClient(client *http.Client) *GetFabricFlavorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fabric flavors params
func (o *GetFabricFlavorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarFilter adds the dollarFilter to the get fabric flavors params
func (o *GetFabricFlavorsParams) WithDollarFilter(dollarFilter *string) *GetFabricFlavorsParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get fabric flavors params
func (o *GetFabricFlavorsParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithAPIVersion adds the aPIVersion to the get fabric flavors params
func (o *GetFabricFlavorsParams) WithAPIVersion(aPIVersion *string) *GetFabricFlavorsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get fabric flavors params
func (o *GetFabricFlavorsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetFabricFlavorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarFilter != nil {

		// query param $filter
		var qrDollarFilter string

		if o.DollarFilter != nil {
			qrDollarFilter = *o.DollarFilter
		}
		qDollarFilter := qrDollarFilter
		if qDollarFilter != "" {

			if err := r.SetQueryParam("$filter", qDollarFilter); err != nil {
				return err
			}
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
