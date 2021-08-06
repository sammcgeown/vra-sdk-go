// Code generated by go-swagger; DO NOT EDIT.

package project

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateUsingPATCHParams creates a new UpdateUsingPATCHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateUsingPATCHParams() *UpdateUsingPATCHParams {
	return &UpdateUsingPATCHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUsingPATCHParamsWithTimeout creates a new UpdateUsingPATCHParams object
// with the ability to set a timeout on a request.
func NewUpdateUsingPATCHParamsWithTimeout(timeout time.Duration) *UpdateUsingPATCHParams {
	return &UpdateUsingPATCHParams{
		timeout: timeout,
	}
}

// NewUpdateUsingPATCHParamsWithContext creates a new UpdateUsingPATCHParams object
// with the ability to set a context for a request.
func NewUpdateUsingPATCHParamsWithContext(ctx context.Context) *UpdateUsingPATCHParams {
	return &UpdateUsingPATCHParams{
		Context: ctx,
	}
}

// NewUpdateUsingPATCHParamsWithHTTPClient creates a new UpdateUsingPATCHParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateUsingPATCHParamsWithHTTPClient(client *http.Client) *UpdateUsingPATCHParams {
	return &UpdateUsingPATCHParams{
		HTTPClient: client,
	}
}

/* UpdateUsingPATCHParams contains all the parameters to send to the API endpoint
   for the update using p a t c h operation.

   Typically these are written to a http.Request.
*/
type UpdateUsingPATCHParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format. For versioning information refer to /project-service/api/about.
	*/
	APIVersion *string

	/* ID.

	   Тhe id of the project.
	*/
	ID string

	/* Request.

	   request
	*/
	Request *models.ProjectResourceMetadata

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUsingPATCHParams) WithDefaults() *UpdateUsingPATCHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateUsingPATCHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update using p a t c h params
func (o *UpdateUsingPATCHParams) WithTimeout(timeout time.Duration) *UpdateUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update using p a t c h params
func (o *UpdateUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update using p a t c h params
func (o *UpdateUsingPATCHParams) WithContext(ctx context.Context) *UpdateUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update using p a t c h params
func (o *UpdateUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update using p a t c h params
func (o *UpdateUsingPATCHParams) WithHTTPClient(client *http.Client) *UpdateUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update using p a t c h params
func (o *UpdateUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update using p a t c h params
func (o *UpdateUsingPATCHParams) WithAPIVersion(aPIVersion *string) *UpdateUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update using p a t c h params
func (o *UpdateUsingPATCHParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the update using p a t c h params
func (o *UpdateUsingPATCHParams) WithID(id string) *UpdateUsingPATCHParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update using p a t c h params
func (o *UpdateUsingPATCHParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the update using p a t c h params
func (o *UpdateUsingPATCHParams) WithRequest(request *models.ProjectResourceMetadata) *UpdateUsingPATCHParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the update using p a t c h params
func (o *UpdateUsingPATCHParams) SetRequest(request *models.ProjectResourceMetadata) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}