// Code generated by go-swagger; DO NOT EDIT.

package vcf

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

// NewGetDomainsUsingGETParams creates a new GetDomainsUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDomainsUsingGETParams() *GetDomainsUsingGETParams {
	return &GetDomainsUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDomainsUsingGETParamsWithTimeout creates a new GetDomainsUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetDomainsUsingGETParamsWithTimeout(timeout time.Duration) *GetDomainsUsingGETParams {
	return &GetDomainsUsingGETParams{
		timeout: timeout,
	}
}

// NewGetDomainsUsingGETParamsWithContext creates a new GetDomainsUsingGETParams object
// with the ability to set a context for a request.
func NewGetDomainsUsingGETParamsWithContext(ctx context.Context) *GetDomainsUsingGETParams {
	return &GetDomainsUsingGETParams{
		Context: ctx,
	}
}

// NewGetDomainsUsingGETParamsWithHTTPClient creates a new GetDomainsUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDomainsUsingGETParamsWithHTTPClient(client *http.Client) *GetDomainsUsingGETParams {
	return &GetDomainsUsingGETParams{
		HTTPClient: client,
	}
}

/* GetDomainsUsingGETParams contains all the parameters to send to the API endpoint
   for the get domains using g e t operation.

   Typically these are written to a http.Request.
*/
type GetDomainsUsingGETParams struct {

	/* IntegrationID.

	   integrationId
	*/
	IntegrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get domains using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainsUsingGETParams) WithDefaults() *GetDomainsUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get domains using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDomainsUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get domains using get params
func (o *GetDomainsUsingGETParams) WithTimeout(timeout time.Duration) *GetDomainsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get domains using get params
func (o *GetDomainsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get domains using get params
func (o *GetDomainsUsingGETParams) WithContext(ctx context.Context) *GetDomainsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get domains using get params
func (o *GetDomainsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get domains using get params
func (o *GetDomainsUsingGETParams) WithHTTPClient(client *http.Client) *GetDomainsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get domains using get params
func (o *GetDomainsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIntegrationID adds the integrationID to the get domains using get params
func (o *GetDomainsUsingGETParams) WithIntegrationID(integrationID string) *GetDomainsUsingGETParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the get domains using get params
func (o *GetDomainsUsingGETParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDomainsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param integrationId
	if err := r.SetPathParam("integrationId", o.IntegrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
