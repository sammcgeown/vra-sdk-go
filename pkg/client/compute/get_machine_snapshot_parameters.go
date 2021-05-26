// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewGetMachineSnapshotParams creates a new GetMachineSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMachineSnapshotParams() *GetMachineSnapshotParams {
	return &GetMachineSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachineSnapshotParamsWithTimeout creates a new GetMachineSnapshotParams object
// with the ability to set a timeout on a request.
func NewGetMachineSnapshotParamsWithTimeout(timeout time.Duration) *GetMachineSnapshotParams {
	return &GetMachineSnapshotParams{
		timeout: timeout,
	}
}

// NewGetMachineSnapshotParamsWithContext creates a new GetMachineSnapshotParams object
// with the ability to set a context for a request.
func NewGetMachineSnapshotParamsWithContext(ctx context.Context) *GetMachineSnapshotParams {
	return &GetMachineSnapshotParams{
		Context: ctx,
	}
}

// NewGetMachineSnapshotParamsWithHTTPClient creates a new GetMachineSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMachineSnapshotParamsWithHTTPClient(client *http.Client) *GetMachineSnapshotParams {
	return &GetMachineSnapshotParams{
		HTTPClient: client,
	}
}

/* GetMachineSnapshotParams contains all the parameters to send to the API endpoint
   for the get machine snapshot operation.

   Typically these are written to a http.Request.
*/
type GetMachineSnapshotParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the machine.
	*/
	ID string

	/* Id1.

	   The ID of the snapshot.
	*/
	Id1 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get machine snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMachineSnapshotParams) WithDefaults() *GetMachineSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get machine snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMachineSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get machine snapshot params
func (o *GetMachineSnapshotParams) WithTimeout(timeout time.Duration) *GetMachineSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine snapshot params
func (o *GetMachineSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine snapshot params
func (o *GetMachineSnapshotParams) WithContext(ctx context.Context) *GetMachineSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine snapshot params
func (o *GetMachineSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machine snapshot params
func (o *GetMachineSnapshotParams) WithHTTPClient(client *http.Client) *GetMachineSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machine snapshot params
func (o *GetMachineSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get machine snapshot params
func (o *GetMachineSnapshotParams) WithAPIVersion(aPIVersion *string) *GetMachineSnapshotParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get machine snapshot params
func (o *GetMachineSnapshotParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get machine snapshot params
func (o *GetMachineSnapshotParams) WithID(id string) *GetMachineSnapshotParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get machine snapshot params
func (o *GetMachineSnapshotParams) SetID(id string) {
	o.ID = id
}

// WithId1 adds the id1 to the get machine snapshot params
func (o *GetMachineSnapshotParams) WithId1(id1 string) *GetMachineSnapshotParams {
	o.SetId1(id1)
	return o
}

// SetId1 adds the id1 to the get machine snapshot params
func (o *GetMachineSnapshotParams) SetId1(id1 string) {
	o.Id1 = id1
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachineSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id1
	if err := r.SetPathParam("id1", o.Id1); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
