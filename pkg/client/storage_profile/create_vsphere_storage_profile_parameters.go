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

// NewCreateVSphereStorageProfileParams creates a new CreateVSphereStorageProfileParams object
// with the default values initialized.
func NewCreateVSphereStorageProfileParams() *CreateVSphereStorageProfileParams {
	var ()
	return &CreateVSphereStorageProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVSphereStorageProfileParamsWithTimeout creates a new CreateVSphereStorageProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateVSphereStorageProfileParamsWithTimeout(timeout time.Duration) *CreateVSphereStorageProfileParams {
	var ()
	return &CreateVSphereStorageProfileParams{

		timeout: timeout,
	}
}

// NewCreateVSphereStorageProfileParamsWithContext creates a new CreateVSphereStorageProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateVSphereStorageProfileParamsWithContext(ctx context.Context) *CreateVSphereStorageProfileParams {
	var ()
	return &CreateVSphereStorageProfileParams{

		Context: ctx,
	}
}

// NewCreateVSphereStorageProfileParamsWithHTTPClient creates a new CreateVSphereStorageProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateVSphereStorageProfileParamsWithHTTPClient(client *http.Client) *CreateVSphereStorageProfileParams {
	var ()
	return &CreateVSphereStorageProfileParams{
		HTTPClient: client,
	}
}

/*CreateVSphereStorageProfileParams contains all the parameters to send to the API endpoint
for the create v sphere storage profile operation typically these are written to a http.Request
*/
type CreateVSphereStorageProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  StorageProfileVsphereSpecification instance

	*/
	Body *models.StorageProfileVsphereSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) WithTimeout(timeout time.Duration) *CreateVSphereStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) WithContext(ctx context.Context) *CreateVSphereStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) WithHTTPClient(client *http.Client) *CreateVSphereStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) WithAPIVersion(aPIVersion *string) *CreateVSphereStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) WithBody(body *models.StorageProfileVsphereSpecification) *CreateVSphereStorageProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create v sphere storage profile params
func (o *CreateVSphereStorageProfileParams) SetBody(body *models.StorageProfileVsphereSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVSphereStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
