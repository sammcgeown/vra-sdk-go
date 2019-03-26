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

// NewGetFabricVSphereDatastoreParams creates a new GetFabricVSphereDatastoreParams object
// with the default values initialized.
func NewGetFabricVSphereDatastoreParams() *GetFabricVSphereDatastoreParams {
	var ()
	return &GetFabricVSphereDatastoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricVSphereDatastoreParamsWithTimeout creates a new GetFabricVSphereDatastoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFabricVSphereDatastoreParamsWithTimeout(timeout time.Duration) *GetFabricVSphereDatastoreParams {
	var ()
	return &GetFabricVSphereDatastoreParams{

		timeout: timeout,
	}
}

// NewGetFabricVSphereDatastoreParamsWithContext creates a new GetFabricVSphereDatastoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFabricVSphereDatastoreParamsWithContext(ctx context.Context) *GetFabricVSphereDatastoreParams {
	var ()
	return &GetFabricVSphereDatastoreParams{

		Context: ctx,
	}
}

// NewGetFabricVSphereDatastoreParamsWithHTTPClient creates a new GetFabricVSphereDatastoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFabricVSphereDatastoreParamsWithHTTPClient(client *http.Client) *GetFabricVSphereDatastoreParams {
	var ()
	return &GetFabricVSphereDatastoreParams{
		HTTPClient: client,
	}
}

/*GetFabricVSphereDatastoreParams contains all the parameters to send to the API endpoint
for the get fabric v sphere datastore operation typically these are written to a http.Request
*/
type GetFabricVSphereDatastoreParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the Fabric vSphere Datastore.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) WithTimeout(timeout time.Duration) *GetFabricVSphereDatastoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) WithContext(ctx context.Context) *GetFabricVSphereDatastoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) WithHTTPClient(client *http.Client) *GetFabricVSphereDatastoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) WithAPIVersion(aPIVersion *string) *GetFabricVSphereDatastoreParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) WithID(id string) *GetFabricVSphereDatastoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetFabricVSphereDatastoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
