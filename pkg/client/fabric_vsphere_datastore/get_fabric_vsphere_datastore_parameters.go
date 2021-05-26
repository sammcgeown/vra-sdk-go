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
	"github.com/go-openapi/strfmt"
)

// NewGetFabricVSphereDatastoreParams creates a new GetFabricVSphereDatastoreParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetFabricVSphereDatastoreParams() *GetFabricVSphereDatastoreParams {
	return &GetFabricVSphereDatastoreParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetFabricVSphereDatastoreParamsWithTimeout creates a new GetFabricVSphereDatastoreParams object
// with the ability to set a timeout on a request.
func NewGetFabricVSphereDatastoreParamsWithTimeout(timeout time.Duration) *GetFabricVSphereDatastoreParams {
	return &GetFabricVSphereDatastoreParams{
		timeout: timeout,
	}
}

// NewGetFabricVSphereDatastoreParamsWithContext creates a new GetFabricVSphereDatastoreParams object
// with the ability to set a context for a request.
func NewGetFabricVSphereDatastoreParamsWithContext(ctx context.Context) *GetFabricVSphereDatastoreParams {
	return &GetFabricVSphereDatastoreParams{
		Context: ctx,
	}
}

// NewGetFabricVSphereDatastoreParamsWithHTTPClient creates a new GetFabricVSphereDatastoreParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetFabricVSphereDatastoreParamsWithHTTPClient(client *http.Client) *GetFabricVSphereDatastoreParams {
	return &GetFabricVSphereDatastoreParams{
		HTTPClient: client,
	}
}

/* GetFabricVSphereDatastoreParams contains all the parameters to send to the API endpoint
   for the get fabric v sphere datastore operation.

   Typically these are written to a http.Request.
*/
type GetFabricVSphereDatastoreParams struct {

	/* DollarSelect.

	   Select a subset of properties to include in the response.
	*/
	DollarSelect *string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the Fabric vSphere Datastore.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get fabric v sphere datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFabricVSphereDatastoreParams) WithDefaults() *GetFabricVSphereDatastoreParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get fabric v sphere datastore params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetFabricVSphereDatastoreParams) SetDefaults() {
	// no default values defined for this parameter
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

// WithDollarSelect adds the dollarSelect to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) WithDollarSelect(dollarSelect *string) *GetFabricVSphereDatastoreParams {
	o.SetDollarSelect(dollarSelect)
	return o
}

// SetDollarSelect adds the dollarSelect to the get fabric v sphere datastore params
func (o *GetFabricVSphereDatastoreParams) SetDollarSelect(dollarSelect *string) {
	o.DollarSelect = dollarSelect
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

	if o.DollarSelect != nil {

		// query param $select
		var qrDollarSelect string

		if o.DollarSelect != nil {
			qrDollarSelect = *o.DollarSelect
		}
		qDollarSelect := qrDollarSelect
		if qDollarSelect != "" {

			if err := r.SetQueryParam("$select", qDollarSelect); err != nil {
				return err
			}
		}
	}

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
