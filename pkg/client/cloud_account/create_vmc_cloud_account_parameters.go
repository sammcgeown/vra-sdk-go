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

// NewCreateVmcCloudAccountParams creates a new CreateVmcCloudAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVmcCloudAccountParams() *CreateVmcCloudAccountParams {
	return &CreateVmcCloudAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVmcCloudAccountParamsWithTimeout creates a new CreateVmcCloudAccountParams object
// with the ability to set a timeout on a request.
func NewCreateVmcCloudAccountParamsWithTimeout(timeout time.Duration) *CreateVmcCloudAccountParams {
	return &CreateVmcCloudAccountParams{
		timeout: timeout,
	}
}

// NewCreateVmcCloudAccountParamsWithContext creates a new CreateVmcCloudAccountParams object
// with the ability to set a context for a request.
func NewCreateVmcCloudAccountParamsWithContext(ctx context.Context) *CreateVmcCloudAccountParams {
	return &CreateVmcCloudAccountParams{
		Context: ctx,
	}
}

// NewCreateVmcCloudAccountParamsWithHTTPClient creates a new CreateVmcCloudAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVmcCloudAccountParamsWithHTTPClient(client *http.Client) *CreateVmcCloudAccountParams {
	return &CreateVmcCloudAccountParams{
		HTTPClient: client,
	}
}

/* CreateVmcCloudAccountParams contains all the parameters to send to the API endpoint
   for the create vmc cloud account operation.

   Typically these are written to a http.Request.
*/
type CreateVmcCloudAccountParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   CloudAccountVmc specification
	*/
	Body *models.CloudAccountVmcSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create vmc cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVmcCloudAccountParams) WithDefaults() *CreateVmcCloudAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create vmc cloud account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVmcCloudAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) WithTimeout(timeout time.Duration) *CreateVmcCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) WithContext(ctx context.Context) *CreateVmcCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) WithHTTPClient(client *http.Client) *CreateVmcCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) WithAPIVersion(aPIVersion *string) *CreateVmcCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) WithBody(body *models.CloudAccountVmcSpecification) *CreateVmcCloudAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create vmc cloud account params
func (o *CreateVmcCloudAccountParams) SetBody(body *models.CloudAccountVmcSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVmcCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
