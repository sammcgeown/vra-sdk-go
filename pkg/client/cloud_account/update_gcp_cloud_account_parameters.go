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

// NewUpdateGcpCloudAccountParams creates a new UpdateGcpCloudAccountParams object
// with the default values initialized.
func NewUpdateGcpCloudAccountParams() *UpdateGcpCloudAccountParams {
	var ()
	return &UpdateGcpCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateGcpCloudAccountParamsWithTimeout creates a new UpdateGcpCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateGcpCloudAccountParamsWithTimeout(timeout time.Duration) *UpdateGcpCloudAccountParams {
	var ()
	return &UpdateGcpCloudAccountParams{

		timeout: timeout,
	}
}

// NewUpdateGcpCloudAccountParamsWithContext creates a new UpdateGcpCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateGcpCloudAccountParamsWithContext(ctx context.Context) *UpdateGcpCloudAccountParams {
	var ()
	return &UpdateGcpCloudAccountParams{

		Context: ctx,
	}
}

// NewUpdateGcpCloudAccountParamsWithHTTPClient creates a new UpdateGcpCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateGcpCloudAccountParamsWithHTTPClient(client *http.Client) *UpdateGcpCloudAccountParams {
	var ()
	return &UpdateGcpCloudAccountParams{
		HTTPClient: client,
	}
}

/*UpdateGcpCloudAccountParams contains all the parameters to send to the API endpoint
for the update gcp cloud account operation typically these are written to a http.Request
*/
type UpdateGcpCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  GCP cloud account details to be updated

	*/
	Body *models.UpdateCloudAccountGcpSpecification
	/*ID
	  Cloud account id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) WithTimeout(timeout time.Duration) *UpdateGcpCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) WithContext(ctx context.Context) *UpdateGcpCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) WithHTTPClient(client *http.Client) *UpdateGcpCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) WithAPIVersion(aPIVersion *string) *UpdateGcpCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) WithBody(body *models.UpdateCloudAccountGcpSpecification) *UpdateGcpCloudAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) SetBody(body *models.UpdateCloudAccountGcpSpecification) {
	o.Body = body
}

// WithID adds the id to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) WithID(id string) *UpdateGcpCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update gcp cloud account params
func (o *UpdateGcpCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateGcpCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
