// Code generated by go-swagger; DO NOT EDIT.

package disk

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

// NewGetDiskSnapshotsParams creates a new GetDiskSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDiskSnapshotsParams() *GetDiskSnapshotsParams {
	return &GetDiskSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDiskSnapshotsParamsWithTimeout creates a new GetDiskSnapshotsParams object
// with the ability to set a timeout on a request.
func NewGetDiskSnapshotsParamsWithTimeout(timeout time.Duration) *GetDiskSnapshotsParams {
	return &GetDiskSnapshotsParams{
		timeout: timeout,
	}
}

// NewGetDiskSnapshotsParamsWithContext creates a new GetDiskSnapshotsParams object
// with the ability to set a context for a request.
func NewGetDiskSnapshotsParamsWithContext(ctx context.Context) *GetDiskSnapshotsParams {
	return &GetDiskSnapshotsParams{
		Context: ctx,
	}
}

// NewGetDiskSnapshotsParamsWithHTTPClient creates a new GetDiskSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDiskSnapshotsParamsWithHTTPClient(client *http.Client) *GetDiskSnapshotsParams {
	return &GetDiskSnapshotsParams{
		HTTPClient: client,
	}
}

/* GetDiskSnapshotsParams contains all the parameters to send to the API endpoint
   for the get disk snapshots operation.

   Typically these are written to a http.Request.
*/
type GetDiskSnapshotsParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* ID.

	   The ID of the disk.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get disk snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiskSnapshotsParams) WithDefaults() *GetDiskSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get disk snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDiskSnapshotsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get disk snapshots params
func (o *GetDiskSnapshotsParams) WithTimeout(timeout time.Duration) *GetDiskSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get disk snapshots params
func (o *GetDiskSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get disk snapshots params
func (o *GetDiskSnapshotsParams) WithContext(ctx context.Context) *GetDiskSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get disk snapshots params
func (o *GetDiskSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get disk snapshots params
func (o *GetDiskSnapshotsParams) WithHTTPClient(client *http.Client) *GetDiskSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get disk snapshots params
func (o *GetDiskSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get disk snapshots params
func (o *GetDiskSnapshotsParams) WithAPIVersion(aPIVersion *string) *GetDiskSnapshotsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get disk snapshots params
func (o *GetDiskSnapshotsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get disk snapshots params
func (o *GetDiskSnapshotsParams) WithID(id string) *GetDiskSnapshotsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get disk snapshots params
func (o *GetDiskSnapshotsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetDiskSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
