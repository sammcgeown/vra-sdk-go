// Code generated by go-swagger; DO NOT EDIT.

package request

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

// NewDeleteRequestParams creates a new DeleteRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRequestParams() *DeleteRequestParams {
	return &DeleteRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRequestParamsWithTimeout creates a new DeleteRequestParams object
// with the ability to set a timeout on a request.
func NewDeleteRequestParamsWithTimeout(timeout time.Duration) *DeleteRequestParams {
	return &DeleteRequestParams{
		timeout: timeout,
	}
}

// NewDeleteRequestParamsWithContext creates a new DeleteRequestParams object
// with the ability to set a context for a request.
func NewDeleteRequestParamsWithContext(ctx context.Context) *DeleteRequestParams {
	return &DeleteRequestParams{
		Context: ctx,
	}
}

// NewDeleteRequestParamsWithHTTPClient creates a new DeleteRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRequestParamsWithHTTPClient(client *http.Client) *DeleteRequestParams {
	return &DeleteRequestParams{
		HTTPClient: client,
	}
}

/* DeleteRequestParams contains all the parameters to send to the API endpoint
   for the delete request operation.

   Typically these are written to a http.Request.
*/
type DeleteRequestParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the request.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRequestParams) WithDefaults() *DeleteRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete request params
func (o *DeleteRequestParams) WithTimeout(timeout time.Duration) *DeleteRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete request params
func (o *DeleteRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete request params
func (o *DeleteRequestParams) WithContext(ctx context.Context) *DeleteRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete request params
func (o *DeleteRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete request params
func (o *DeleteRequestParams) WithHTTPClient(client *http.Client) *DeleteRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete request params
func (o *DeleteRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete request params
func (o *DeleteRequestParams) WithAPIVersion(aPIVersion *string) *DeleteRequestParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete request params
func (o *DeleteRequestParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete request params
func (o *DeleteRequestParams) WithID(id string) *DeleteRequestParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete request params
func (o *DeleteRequestParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
