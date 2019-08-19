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

// NewUpdateAzureStorageProfileParams creates a new UpdateAzureStorageProfileParams object
// with the default values initialized.
func NewUpdateAzureStorageProfileParams() *UpdateAzureStorageProfileParams {
	var ()
	return &UpdateAzureStorageProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAzureStorageProfileParamsWithTimeout creates a new UpdateAzureStorageProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAzureStorageProfileParamsWithTimeout(timeout time.Duration) *UpdateAzureStorageProfileParams {
	var ()
	return &UpdateAzureStorageProfileParams{

		timeout: timeout,
	}
}

// NewUpdateAzureStorageProfileParamsWithContext creates a new UpdateAzureStorageProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAzureStorageProfileParamsWithContext(ctx context.Context) *UpdateAzureStorageProfileParams {
	var ()
	return &UpdateAzureStorageProfileParams{

		Context: ctx,
	}
}

// NewUpdateAzureStorageProfileParamsWithHTTPClient creates a new UpdateAzureStorageProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAzureStorageProfileParamsWithHTTPClient(client *http.Client) *UpdateAzureStorageProfileParams {
	var ()
	return &UpdateAzureStorageProfileParams{
		HTTPClient: client,
	}
}

/*UpdateAzureStorageProfileParams contains all the parameters to send to the API endpoint
for the update azure storage profile operation typically these are written to a http.Request
*/
type UpdateAzureStorageProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  StorageProfileAzureSpecification instance

	*/
	Body *models.StorageProfileAzureSpecification
	/*ID
	  The ID of the storage profile.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) WithTimeout(timeout time.Duration) *UpdateAzureStorageProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) WithContext(ctx context.Context) *UpdateAzureStorageProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) WithHTTPClient(client *http.Client) *UpdateAzureStorageProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) WithAPIVersion(aPIVersion *string) *UpdateAzureStorageProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) WithBody(body *models.StorageProfileAzureSpecification) *UpdateAzureStorageProfileParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) SetBody(body *models.StorageProfileAzureSpecification) {
	o.Body = body
}

// WithID adds the id to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) WithID(id string) *UpdateAzureStorageProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update azure storage profile params
func (o *UpdateAzureStorageProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAzureStorageProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
