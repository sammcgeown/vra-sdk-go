// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

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

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// NewCreateAwsStorageProfileParams creates a new CreateAwsStorageProfileParams object
// with the default values initialized.
func NewCreateAwsStorageProfileParams() *CreateAwsStorageProfileParams {
	var ()
	return &CreateAwsStorageProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAwsStorageProfileParamsWithTimeout creates a new CreateAwsStorageProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAwsStorageProfileParamsWithTimeout(timeout time.Duration) *CreateAwsStorageProfileParams {
	var ()
	return &CreateAwsStorageProfileParams{

		timeout: timeout,
	}
}

// NewCreateAwsStorageProfileParamsWithContext creates a new CreateAwsStorageProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAwsStorageProfileParamsWithContext(ctx context.Context) *CreateAwsStorageProfileParams {
	var ()
	return &CreateAwsStorageProfileParams{

		Context: ctx,
	}
}

// NewCreateAwsStorageProfileParamsWithHTTPClient creates a new CreateAwsStorageProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAwsStorageProfileParamsWithHTTPClient(client *http.Client) *CreateAwsStorageProfileParams {
	var ()
	return &CreateAwsStorageProfileParams{
		HTTPClient: client,
	}
}

/*CreateAwsStorageProfileParams contains all the parameters to send to the API endpoint
for the create aws storage profile operation typically these are written to a http.Request
*/
type CreateAwsStorageProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  StorageProfileAwsSpecification instance

	*/
	Body *models.StorageProfileAwsSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) WithTimeout(timeout time.Duration) *CreateAwsStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) WithContext(ctx context.Context) *CreateAwsStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) WithHTTPClient(client *http.Client) *CreateAwsStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) WithAPIVersion(aPIVersion *string) *CreateAwsStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) WithBody(body *models.StorageProfileAwsSpecification) *CreateAwsStorageProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create aws storage profile params
func (o *CreateAwsStorageProfileParams) SetBody(body *models.StorageProfileAwsSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAwsStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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