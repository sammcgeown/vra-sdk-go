// Code generated by go-swagger; DO NOT EDIT.

package load_balancer

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

// NewGetLoadBalancerParams creates a new GetLoadBalancerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLoadBalancerParams() *GetLoadBalancerParams {
	return &GetLoadBalancerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoadBalancerParamsWithTimeout creates a new GetLoadBalancerParams object
// with the ability to set a timeout on a request.
func NewGetLoadBalancerParamsWithTimeout(timeout time.Duration) *GetLoadBalancerParams {
	return &GetLoadBalancerParams{
		timeout: timeout,
	}
}

// NewGetLoadBalancerParamsWithContext creates a new GetLoadBalancerParams object
// with the ability to set a context for a request.
func NewGetLoadBalancerParamsWithContext(ctx context.Context) *GetLoadBalancerParams {
	return &GetLoadBalancerParams{
		Context: ctx,
	}
}

// NewGetLoadBalancerParamsWithHTTPClient creates a new GetLoadBalancerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLoadBalancerParamsWithHTTPClient(client *http.Client) *GetLoadBalancerParams {
	return &GetLoadBalancerParams{
		HTTPClient: client,
	}
}

/* GetLoadBalancerParams contains all the parameters to send to the API endpoint
   for the get load balancer operation.

   Typically these are written to a http.Request.
*/
type GetLoadBalancerParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the load balancer.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get load balancer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoadBalancerParams) WithDefaults() *GetLoadBalancerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get load balancer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLoadBalancerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get load balancer params
func (o *GetLoadBalancerParams) WithTimeout(timeout time.Duration) *GetLoadBalancerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get load balancer params
func (o *GetLoadBalancerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get load balancer params
func (o *GetLoadBalancerParams) WithContext(ctx context.Context) *GetLoadBalancerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get load balancer params
func (o *GetLoadBalancerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get load balancer params
func (o *GetLoadBalancerParams) WithHTTPClient(client *http.Client) *GetLoadBalancerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get load balancer params
func (o *GetLoadBalancerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get load balancer params
func (o *GetLoadBalancerParams) WithAPIVersion(aPIVersion *string) *GetLoadBalancerParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get load balancer params
func (o *GetLoadBalancerParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get load balancer params
func (o *GetLoadBalancerParams) WithID(id string) *GetLoadBalancerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get load balancer params
func (o *GetLoadBalancerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoadBalancerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
