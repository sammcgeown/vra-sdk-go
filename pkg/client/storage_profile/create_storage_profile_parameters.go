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

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateStorageProfileParams creates a new CreateStorageProfileParams object
// with the default values initialized.
func NewCreateStorageProfileParams() *CreateStorageProfileParams {
	var ()
	return &CreateStorageProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateStorageProfileParamsWithTimeout creates a new CreateStorageProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateStorageProfileParamsWithTimeout(timeout time.Duration) *CreateStorageProfileParams {
	var ()
	return &CreateStorageProfileParams{

		timeout: timeout,
	}
}

// NewCreateStorageProfileParamsWithContext creates a new CreateStorageProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateStorageProfileParamsWithContext(ctx context.Context) *CreateStorageProfileParams {
	var ()
	return &CreateStorageProfileParams{

		Context: ctx,
	}
}

// NewCreateStorageProfileParamsWithHTTPClient creates a new CreateStorageProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateStorageProfileParamsWithHTTPClient(client *http.Client) *CreateStorageProfileParams {
	var ()
	return &CreateStorageProfileParams{
		HTTPClient: client,
	}
}

/*CreateStorageProfileParams contains all the parameters to send to the API endpoint
for the create storage profile operation typically these are written to a http.Request
*/
type CreateStorageProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  StorageProfileSpecification instance

	*/
	Body *models.StorageProfileSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create storage profile params
func (o *CreateStorageProfileParams) WithTimeout(timeout time.Duration) *CreateStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create storage profile params
func (o *CreateStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create storage profile params
func (o *CreateStorageProfileParams) WithContext(ctx context.Context) *CreateStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create storage profile params
func (o *CreateStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create storage profile params
func (o *CreateStorageProfileParams) WithHTTPClient(client *http.Client) *CreateStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create storage profile params
func (o *CreateStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create storage profile params
func (o *CreateStorageProfileParams) WithAPIVersion(aPIVersion *string) *CreateStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create storage profile params
func (o *CreateStorageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create storage profile params
func (o *CreateStorageProfileParams) WithBody(body *models.StorageProfileSpecification) *CreateStorageProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create storage profile params
func (o *CreateStorageProfileParams) SetBody(body *models.StorageProfileSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
