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

// NewCreateNsxTCloudAccountParams creates a new CreateNsxTCloudAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNsxTCloudAccountParams() *CreateNsxTCloudAccountParams {
	return &CreateNsxTCloudAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNsxTCloudAccountParamsWithTimeout creates a new CreateNsxTCloudAccountParams object
// with the ability to set a timeout on a request.
func NewCreateNsxTCloudAccountParamsWithTimeout(timeout time.Duration) *CreateNsxTCloudAccountParams {
	return &CreateNsxTCloudAccountParams{
		timeout: timeout,
	}
}

// NewCreateNsxTCloudAccountParamsWithContext creates a new CreateNsxTCloudAccountParams object
// with the ability to set a context for a request.
func NewCreateNsxTCloudAccountParamsWithContext(ctx context.Context) *CreateNsxTCloudAccountParams {
	return &CreateNsxTCloudAccountParams{
		Context: ctx,
	}
}

// NewCreateNsxTCloudAccountParamsWithHTTPClient creates a new CreateNsxTCloudAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNsxTCloudAccountParamsWithHTTPClient(client *http.Client) *CreateNsxTCloudAccountParams {
	return &CreateNsxTCloudAccountParams{
		HTTPClient: client,
	}
}

/* CreateNsxTCloudAccountParams contains all the parameters to send to the API endpoint
   for the create nsx t cloud account operation.

   Typically these are written to a http.Request.
*/
type CreateNsxTCloudAccountParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   CloudAccountNsxT specification
	*/
	Body *models.CloudAccountNsxTSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create nsx t cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNsxTCloudAccountParams) WithDefaults() *CreateNsxTCloudAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create nsx t cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNsxTCloudAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) WithTimeout(timeout time.Duration) *CreateNsxTCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) WithContext(ctx context.Context) *CreateNsxTCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) WithHTTPClient(client *http.Client) *CreateNsxTCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) WithAPIVersion(aPIVersion *string) *CreateNsxTCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) WithBody(body *models.CloudAccountNsxTSpecification) *CreateNsxTCloudAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create nsx t cloud account params
func (o *CreateNsxTCloudAccountParams) SetBody(body *models.CloudAccountNsxTSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNsxTCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
