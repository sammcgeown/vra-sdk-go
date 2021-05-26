// Code generated by go-swagger; DO NOT EDIT.

package source_control_sync

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

// NewGetSyncRequestUsingGETParams creates a new GetSyncRequestUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSyncRequestUsingGETParams() *GetSyncRequestUsingGETParams {
	return &GetSyncRequestUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSyncRequestUsingGETParamsWithTimeout creates a new GetSyncRequestUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetSyncRequestUsingGETParamsWithTimeout(timeout time.Duration) *GetSyncRequestUsingGETParams {
	return &GetSyncRequestUsingGETParams{
		timeout: timeout,
	}
}

// NewGetSyncRequestUsingGETParamsWithContext creates a new GetSyncRequestUsingGETParams object
// with the ability to set a context for a request.
func NewGetSyncRequestUsingGETParamsWithContext(ctx context.Context) *GetSyncRequestUsingGETParams {
	return &GetSyncRequestUsingGETParams{
		Context: ctx,
	}
}

// NewGetSyncRequestUsingGETParamsWithHTTPClient creates a new GetSyncRequestUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSyncRequestUsingGETParamsWithHTTPClient(client *http.Client) *GetSyncRequestUsingGETParams {
	return &GetSyncRequestUsingGETParams{
		HTTPClient: client,
	}
}

/* GetSyncRequestUsingGETParams contains all the parameters to send to the API endpoint
   for the get sync request using g e t operation.

   Typically these are written to a http.Request.
*/
type GetSyncRequestUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information, please refer to /content/api/about
	*/
	APIVersion *string

	/* ID.

	   ID of the sync request

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get sync request using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSyncRequestUsingGETParams) WithDefaults() *GetSyncRequestUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get sync request using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSyncRequestUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) WithTimeout(timeout time.Duration) *GetSyncRequestUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) WithContext(ctx context.Context) *GetSyncRequestUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) WithHTTPClient(client *http.Client) *GetSyncRequestUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) WithAPIVersion(aPIVersion *string) *GetSyncRequestUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) WithID(id strfmt.UUID) *GetSyncRequestUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get sync request using get params
func (o *GetSyncRequestUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSyncRequestUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
