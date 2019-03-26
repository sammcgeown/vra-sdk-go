// Code generated by go-swagger; DO NOT EDIT.

package fabric_vsphere_datastore

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

// NewGetFabricVSphereDatastoresParams creates a new GetFabricVSphereDatastoresParams object
// with the default values initialized.
func NewGetFabricVSphereDatastoresParams() *GetFabricVSphereDatastoresParams {
	var ()
	return &GetFabricVSphereDatastoresParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricVSphereDatastoresParamsWithTimeout creates a new GetFabricVSphereDatastoresParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFabricVSphereDatastoresParamsWithTimeout(timeout time.Duration) *GetFabricVSphereDatastoresParams {
	var ()
	return &GetFabricVSphereDatastoresParams{

		timeout: timeout,
	}
}

// NewGetFabricVSphereDatastoresParamsWithContext creates a new GetFabricVSphereDatastoresParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFabricVSphereDatastoresParamsWithContext(ctx context.Context) *GetFabricVSphereDatastoresParams {
	var ()
	return &GetFabricVSphereDatastoresParams{

		Context: ctx,
	}
}

// NewGetFabricVSphereDatastoresParamsWithHTTPClient creates a new GetFabricVSphereDatastoresParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFabricVSphereDatastoresParamsWithHTTPClient(client *http.Client) *GetFabricVSphereDatastoresParams {
	var ()
	return &GetFabricVSphereDatastoresParams{
		HTTPClient: client,
	}
}

/*GetFabricVSphereDatastoresParams contains all the parameters to send to the API endpoint
for the get fabric v sphere datastores operation typically these are written to a http.Request
*/
type GetFabricVSphereDatastoresParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fabric v sphere datastores params
func (o *GetFabricVSphereDatastoresParams) WithTimeout(timeout time.Duration) *GetFabricVSphereDatastoresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fabric v sphere datastores params
func (o *GetFabricVSphereDatastoresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fabric v sphere datastores params
func (o *GetFabricVSphereDatastoresParams) WithContext(ctx context.Context) *GetFabricVSphereDatastoresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fabric v sphere datastores params
func (o *GetFabricVSphereDatastoresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fabric v sphere datastores params
func (o *GetFabricVSphereDatastoresParams) WithHTTPClient(client *http.Client) *GetFabricVSphereDatastoresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fabric v sphere datastores params
func (o *GetFabricVSphereDatastoresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get fabric v sphere datastores params
func (o *GetFabricVSphereDatastoresParams) WithAPIVersion(aPIVersion *string) *GetFabricVSphereDatastoresParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get fabric v sphere datastores params
func (o *GetFabricVSphereDatastoresParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *GetFabricVSphereDatastoresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
