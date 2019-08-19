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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreateCloudAccountParams creates a new CreateCloudAccountParams object
// with the default values initialized.
func NewCreateCloudAccountParams() *CreateCloudAccountParams {
	var ()
	return &CreateCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCloudAccountParamsWithTimeout creates a new CreateCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCloudAccountParamsWithTimeout(timeout time.Duration) *CreateCloudAccountParams {
	var ()
	return &CreateCloudAccountParams{

		timeout: timeout,
	}
}

// NewCreateCloudAccountParamsWithContext creates a new CreateCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCloudAccountParamsWithContext(ctx context.Context) *CreateCloudAccountParams {
	var ()
	return &CreateCloudAccountParams{

		Context: ctx,
	}
}

// NewCreateCloudAccountParamsWithHTTPClient creates a new CreateCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCloudAccountParamsWithHTTPClient(client *http.Client) *CreateCloudAccountParams {
	var ()
	return &CreateCloudAccountParams{
		HTTPClient: client,
	}
}

/*CreateCloudAccountParams contains all the parameters to send to the API endpoint
for the create cloud account operation typically these are written to a http.Request
*/
type CreateCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  CloudAccount instance

	*/
	Body *models.CloudAccountSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cloud account params
func (o *CreateCloudAccountParams) WithTimeout(timeout time.Duration) *CreateCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cloud account params
func (o *CreateCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cloud account params
func (o *CreateCloudAccountParams) WithContext(ctx context.Context) *CreateCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cloud account params
func (o *CreateCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cloud account params
func (o *CreateCloudAccountParams) WithHTTPClient(client *http.Client) *CreateCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cloud account params
func (o *CreateCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create cloud account params
func (o *CreateCloudAccountParams) WithAPIVersion(aPIVersion *string) *CreateCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create cloud account params
func (o *CreateCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create cloud account params
func (o *CreateCloudAccountParams) WithBody(body *models.CloudAccountSpecification) *CreateCloudAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cloud account params
func (o *CreateCloudAccountParams) SetBody(body *models.CloudAccountSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
