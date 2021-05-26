// Code generated by go-swagger; DO NOT EDIT.

package deployment

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
	"github.com/go-openapi/swag"
)

// NewGetDeploymentsParams creates a new GetDeploymentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDeploymentsParams() *GetDeploymentsParams {
	return &GetDeploymentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentsParamsWithTimeout creates a new GetDeploymentsParams object
// with the ability to set a timeout on a request.
func NewGetDeploymentsParamsWithTimeout(timeout time.Duration) *GetDeploymentsParams {
	return &GetDeploymentsParams{
		timeout: timeout,
	}
}

// NewGetDeploymentsParamsWithContext creates a new GetDeploymentsParams object
// with the ability to set a context for a request.
func NewGetDeploymentsParamsWithContext(ctx context.Context) *GetDeploymentsParams {
	return &GetDeploymentsParams{
		Context: ctx,
	}
}

// NewGetDeploymentsParamsWithHTTPClient creates a new GetDeploymentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDeploymentsParamsWithHTTPClient(client *http.Client) *GetDeploymentsParams {
	return &GetDeploymentsParams{
		HTTPClient: client,
	}
}

/* GetDeploymentsParams contains all the parameters to send to the API endpoint
   for the get deployments operation.

   Typically these are written to a http.Request.
*/
type GetDeploymentsParams struct {

	/* DollarCount.

	   Flag which when specified shows the total number of records. If the collection has a filter it shows the number of records matching the filter.
	*/
	DollarCount *bool

	/* DollarFilter.

	   Filter the results by a specified predicate expression. Operators: eq, ne, and, or.
	*/
	DollarFilter *string

	/* DollarSkip.

	   Number of records you want to skip.
	*/
	DollarSkip *int64

	/* DollarTop.

	   Number of records you want to get.
	*/
	DollarTop *int64

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentsParams) WithDefaults() *GetDeploymentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get deployments params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDeploymentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get deployments params
func (o *GetDeploymentsParams) WithTimeout(timeout time.Duration) *GetDeploymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployments params
func (o *GetDeploymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployments params
func (o *GetDeploymentsParams) WithContext(ctx context.Context) *GetDeploymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployments params
func (o *GetDeploymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployments params
func (o *GetDeploymentsParams) WithHTTPClient(client *http.Client) *GetDeploymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployments params
func (o *GetDeploymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDollarCount adds the dollarCount to the get deployments params
func (o *GetDeploymentsParams) WithDollarCount(dollarCount *bool) *GetDeploymentsParams {
	o.SetDollarCount(dollarCount)
	return o
}

// SetDollarCount adds the dollarCount to the get deployments params
func (o *GetDeploymentsParams) SetDollarCount(dollarCount *bool) {
	o.DollarCount = dollarCount
}

// WithDollarFilter adds the dollarFilter to the get deployments params
func (o *GetDeploymentsParams) WithDollarFilter(dollarFilter *string) *GetDeploymentsParams {
	o.SetDollarFilter(dollarFilter)
	return o
}

// SetDollarFilter adds the dollarFilter to the get deployments params
func (o *GetDeploymentsParams) SetDollarFilter(dollarFilter *string) {
	o.DollarFilter = dollarFilter
}

// WithDollarSkip adds the dollarSkip to the get deployments params
func (o *GetDeploymentsParams) WithDollarSkip(dollarSkip *int64) *GetDeploymentsParams {
	o.SetDollarSkip(dollarSkip)
	return o
}

// SetDollarSkip adds the dollarSkip to the get deployments params
func (o *GetDeploymentsParams) SetDollarSkip(dollarSkip *int64) {
	o.DollarSkip = dollarSkip
}

// WithDollarTop adds the dollarTop to the get deployments params
func (o *GetDeploymentsParams) WithDollarTop(dollarTop *int64) *GetDeploymentsParams {
	o.SetDollarTop(dollarTop)
	return o
}

// SetDollarTop adds the dollarTop to the get deployments params
func (o *GetDeploymentsParams) SetDollarTop(dollarTop *int64) {
	o.DollarTop = dollarTop
}

// WithAPIVersion adds the aPIVersion to the get deployments params
func (o *GetDeploymentsParams) WithAPIVersion(aPIVersion *string) *GetDeploymentsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get deployments params
func (o *GetDeploymentsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DollarCount != nil {

		// query param $count
		var qrDollarCount bool

		if o.DollarCount != nil {
			qrDollarCount = *o.DollarCount
		}
		qDollarCount := swag.FormatBool(qrDollarCount)
		if qDollarCount != "" {

			if err := r.SetQueryParam("$count", qDollarCount); err != nil {
				return err
			}
		}
	}

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

	if o.DollarSkip != nil {

		// query param $skip
		var qrDollarSkip int64

		if o.DollarSkip != nil {
			qrDollarSkip = *o.DollarSkip
		}
		qDollarSkip := swag.FormatInt64(qrDollarSkip)
		if qDollarSkip != "" {

			if err := r.SetQueryParam("$skip", qDollarSkip); err != nil {
				return err
			}
		}
	}

	if o.DollarTop != nil {

		// query param $top
		var qrDollarTop int64

		if o.DollarTop != nil {
			qrDollarTop = *o.DollarTop
		}
		qDollarTop := swag.FormatInt64(qrDollarTop)
		if qDollarTop != "" {

			if err := r.SetQueryParam("$top", qDollarTop); err != nil {
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
