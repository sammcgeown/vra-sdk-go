// Code generated by go-swagger; DO NOT EDIT.

package catalog_admin_items

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

// NewGetCatalogItemUsingGETParams creates a new GetCatalogItemUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCatalogItemUsingGETParams() *GetCatalogItemUsingGETParams {
	return &GetCatalogItemUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogItemUsingGETParamsWithTimeout creates a new GetCatalogItemUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetCatalogItemUsingGETParamsWithTimeout(timeout time.Duration) *GetCatalogItemUsingGETParams {
	return &GetCatalogItemUsingGETParams{
		timeout: timeout,
	}
}

// NewGetCatalogItemUsingGETParamsWithContext creates a new GetCatalogItemUsingGETParams object
// with the ability to set a context for a request.
func NewGetCatalogItemUsingGETParamsWithContext(ctx context.Context) *GetCatalogItemUsingGETParams {
	return &GetCatalogItemUsingGETParams{
		Context: ctx,
	}
}

// NewGetCatalogItemUsingGETParamsWithHTTPClient creates a new GetCatalogItemUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCatalogItemUsingGETParamsWithHTTPClient(client *http.Client) *GetCatalogItemUsingGETParams {
	return &GetCatalogItemUsingGETParams{
		HTTPClient: client,
	}
}

/* GetCatalogItemUsingGETParams contains all the parameters to send to the API endpoint
   for the get catalog item using g e t operation.

   Typically these are written to a http.Request.
*/
type GetCatalogItemUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* ID.

	   Catalog item id

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get catalog item using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogItemUsingGETParams) WithDefaults() *GetCatalogItemUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get catalog item using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogItemUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) WithTimeout(timeout time.Duration) *GetCatalogItemUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) WithContext(ctx context.Context) *GetCatalogItemUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) WithHTTPClient(client *http.Client) *GetCatalogItemUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) WithAPIVersion(aPIVersion *string) *GetCatalogItemUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) WithID(id strfmt.UUID) *GetCatalogItemUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get catalog item using get params
func (o *GetCatalogItemUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogItemUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
