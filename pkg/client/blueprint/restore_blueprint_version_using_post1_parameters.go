// Code generated by go-swagger; DO NOT EDIT.

package blueprint

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

// NewRestoreBlueprintVersionUsingPOST1Params creates a new RestoreBlueprintVersionUsingPOST1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestoreBlueprintVersionUsingPOST1Params() *RestoreBlueprintVersionUsingPOST1Params {
	return &RestoreBlueprintVersionUsingPOST1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestoreBlueprintVersionUsingPOST1ParamsWithTimeout creates a new RestoreBlueprintVersionUsingPOST1Params object
// with the ability to set a timeout on a request.
func NewRestoreBlueprintVersionUsingPOST1ParamsWithTimeout(timeout time.Duration) *RestoreBlueprintVersionUsingPOST1Params {
	return &RestoreBlueprintVersionUsingPOST1Params{
		timeout: timeout,
	}
}

// NewRestoreBlueprintVersionUsingPOST1ParamsWithContext creates a new RestoreBlueprintVersionUsingPOST1Params object
// with the ability to set a context for a request.
func NewRestoreBlueprintVersionUsingPOST1ParamsWithContext(ctx context.Context) *RestoreBlueprintVersionUsingPOST1Params {
	return &RestoreBlueprintVersionUsingPOST1Params{
		Context: ctx,
	}
}

// NewRestoreBlueprintVersionUsingPOST1ParamsWithHTTPClient creates a new RestoreBlueprintVersionUsingPOST1Params object
// with the ability to set a custom HTTPClient for a request.
func NewRestoreBlueprintVersionUsingPOST1ParamsWithHTTPClient(client *http.Client) *RestoreBlueprintVersionUsingPOST1Params {
	return &RestoreBlueprintVersionUsingPOST1Params{
		HTTPClient: client,
	}
}

/* RestoreBlueprintVersionUsingPOST1Params contains all the parameters to send to the API endpoint
   for the restore blueprint version using p o s t 1 operation.

   Typically these are written to a http.Request.
*/
type RestoreBlueprintVersionUsingPOST1Params struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /blueprint/api/about
	*/
	APIVersion *string

	/* BlueprintID.

	   blueprintId

	   Format: uuid
	*/
	BlueprintID strfmt.UUID

	/* Version.

	   version
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the restore blueprint version using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreBlueprintVersionUsingPOST1Params) WithDefaults() *RestoreBlueprintVersionUsingPOST1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the restore blueprint version using p o s t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestoreBlueprintVersionUsingPOST1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) WithTimeout(timeout time.Duration) *RestoreBlueprintVersionUsingPOST1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) WithContext(ctx context.Context) *RestoreBlueprintVersionUsingPOST1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) WithHTTPClient(client *http.Client) *RestoreBlueprintVersionUsingPOST1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) WithAPIVersion(aPIVersion *string) *RestoreBlueprintVersionUsingPOST1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBlueprintID adds the blueprintID to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) WithBlueprintID(blueprintID strfmt.UUID) *RestoreBlueprintVersionUsingPOST1Params {
	o.SetBlueprintID(blueprintID)
	return o
}

// SetBlueprintID adds the blueprintId to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) SetBlueprintID(blueprintID strfmt.UUID) {
	o.BlueprintID = blueprintID
}

// WithVersion adds the version to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) WithVersion(version string) *RestoreBlueprintVersionUsingPOST1Params {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the restore blueprint version using p o s t 1 params
func (o *RestoreBlueprintVersionUsingPOST1Params) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *RestoreBlueprintVersionUsingPOST1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param blueprintId
	if err := r.SetPathParam("blueprintId", o.BlueprintID.String()); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
