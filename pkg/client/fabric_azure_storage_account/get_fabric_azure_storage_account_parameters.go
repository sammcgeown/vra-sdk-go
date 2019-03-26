// Code generated by go-swagger; DO NOT EDIT.

package fabric_azure_storage_account

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
)

// NewGetFabricAzureStorageAccountParams creates a new GetFabricAzureStorageAccountParams object
// with the default values initialized.
func NewGetFabricAzureStorageAccountParams() *GetFabricAzureStorageAccountParams {
	var ()
	return &GetFabricAzureStorageAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricAzureStorageAccountParamsWithTimeout creates a new GetFabricAzureStorageAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFabricAzureStorageAccountParamsWithTimeout(timeout time.Duration) *GetFabricAzureStorageAccountParams {
	var ()
	return &GetFabricAzureStorageAccountParams{

		timeout: timeout,
	}
}

// NewGetFabricAzureStorageAccountParamsWithContext creates a new GetFabricAzureStorageAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFabricAzureStorageAccountParamsWithContext(ctx context.Context) *GetFabricAzureStorageAccountParams {
	var ()
	return &GetFabricAzureStorageAccountParams{

		Context: ctx,
	}
}

// NewGetFabricAzureStorageAccountParamsWithHTTPClient creates a new GetFabricAzureStorageAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFabricAzureStorageAccountParamsWithHTTPClient(client *http.Client) *GetFabricAzureStorageAccountParams {
	var ()
	return &GetFabricAzureStorageAccountParams{
		HTTPClient: client,
	}
}

/*GetFabricAzureStorageAccountParams contains all the parameters to send to the API endpoint
for the get fabric azure storage account operation typically these are written to a http.Request
*/
type GetFabricAzureStorageAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the Fabric Azure Storage Account.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) WithTimeout(timeout time.Duration) *GetFabricAzureStorageAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) WithContext(ctx context.Context) *GetFabricAzureStorageAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) WithHTTPClient(client *http.Client) *GetFabricAzureStorageAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) WithAPIVersion(aPIVersion *string) *GetFabricAzureStorageAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) WithID(id string) *GetFabricAzureStorageAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get fabric azure storage account params
func (o *GetFabricAzureStorageAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetFabricAzureStorageAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
