// Code generated by go-swagger; DO NOT EDIT.

package content_source

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteContentSourceUsingDELETEParams creates a new DeleteContentSourceUsingDELETEParams object
// with the default values initialized.
func NewDeleteContentSourceUsingDELETEParams() *DeleteContentSourceUsingDELETEParams {
	var ()
	return &DeleteContentSourceUsingDELETEParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteContentSourceUsingDELETEParamsWithTimeout creates a new DeleteContentSourceUsingDELETEParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteContentSourceUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteContentSourceUsingDELETEParams {
	var ()
	return &DeleteContentSourceUsingDELETEParams{

		timeout: timeout,
	}
}

// NewDeleteContentSourceUsingDELETEParamsWithContext creates a new DeleteContentSourceUsingDELETEParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteContentSourceUsingDELETEParamsWithContext(ctx context.Context) *DeleteContentSourceUsingDELETEParams {
	var ()
	return &DeleteContentSourceUsingDELETEParams{

		Context: ctx,
	}
}

// NewDeleteContentSourceUsingDELETEParamsWithHTTPClient creates a new DeleteContentSourceUsingDELETEParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteContentSourceUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteContentSourceUsingDELETEParams {
	var ()
	return &DeleteContentSourceUsingDELETEParams{
		HTTPClient: client,
	}
}

/*DeleteContentSourceUsingDELETEParams contains all the parameters to send to the API endpoint
for the delete content source using d e l e t e operation typically these are written to a http.Request
*/
type DeleteContentSourceUsingDELETEParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information, please refer to /content/api/about

	*/
	APIVersion *string
	/*ID
	  ID of the content source

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteContentSourceUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) WithContext(ctx context.Context) *DeleteContentSourceUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteContentSourceUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) WithAPIVersion(aPIVersion *string) *DeleteContentSourceUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) WithID(id strfmt.UUID) *DeleteContentSourceUsingDELETEParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete content source using d e l e t e params
func (o *DeleteContentSourceUsingDELETEParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteContentSourceUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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