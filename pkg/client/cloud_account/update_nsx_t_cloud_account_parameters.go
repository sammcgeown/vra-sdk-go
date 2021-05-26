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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateNsxTCloudAccountParams creates a new UpdateNsxTCloudAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNsxTCloudAccountParams() *UpdateNsxTCloudAccountParams {
	return &UpdateNsxTCloudAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNsxTCloudAccountParamsWithTimeout creates a new UpdateNsxTCloudAccountParams object
// with the ability to set a timeout on a request.
func NewUpdateNsxTCloudAccountParamsWithTimeout(timeout time.Duration) *UpdateNsxTCloudAccountParams {
	return &UpdateNsxTCloudAccountParams{
		timeout: timeout,
	}
}

// NewUpdateNsxTCloudAccountParamsWithContext creates a new UpdateNsxTCloudAccountParams object
// with the ability to set a context for a request.
func NewUpdateNsxTCloudAccountParamsWithContext(ctx context.Context) *UpdateNsxTCloudAccountParams {
	return &UpdateNsxTCloudAccountParams{
		Context: ctx,
	}
}

// NewUpdateNsxTCloudAccountParamsWithHTTPClient creates a new UpdateNsxTCloudAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNsxTCloudAccountParamsWithHTTPClient(client *http.Client) *UpdateNsxTCloudAccountParams {
	return &UpdateNsxTCloudAccountParams{
		HTTPClient: client,
	}
}

/* UpdateNsxTCloudAccountParams contains all the parameters to send to the API endpoint
   for the update nsx t cloud account operation.

   Typically these are written to a http.Request.
*/
type UpdateNsxTCloudAccountParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   NSX-T cloud account details to be updated
	*/
	Body *models.UpdateCloudAccountNsxTSpecification

	/* ID.

	   Cloud account id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update nsx t cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNsxTCloudAccountParams) WithDefaults() *UpdateNsxTCloudAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update nsx t cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNsxTCloudAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) WithTimeout(timeout time.Duration) *UpdateNsxTCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) WithContext(ctx context.Context) *UpdateNsxTCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) WithHTTPClient(client *http.Client) *UpdateNsxTCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) WithAPIVersion(aPIVersion *string) *UpdateNsxTCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) WithBody(body *models.UpdateCloudAccountNsxTSpecification) *UpdateNsxTCloudAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) SetBody(body *models.UpdateCloudAccountNsxTSpecification) {
	o.Body = body
}

// WithID adds the id to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) WithID(id string) *UpdateNsxTCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update nsx t cloud account params
func (o *UpdateNsxTCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNsxTCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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
