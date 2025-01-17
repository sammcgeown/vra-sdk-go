// Code generated by go-swagger; DO NOT EDIT.

package project

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
	"github.com/go-openapi/swag"
)

// NewGetProjectUsingGETParams creates a new GetProjectUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetProjectUsingGETParams() *GetProjectUsingGETParams {
	return &GetProjectUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectUsingGETParamsWithTimeout creates a new GetProjectUsingGETParams object
// with the ability to set a timeout on a request.
func NewGetProjectUsingGETParamsWithTimeout(timeout time.Duration) *GetProjectUsingGETParams {
	return &GetProjectUsingGETParams{
		timeout: timeout,
	}
}

// NewGetProjectUsingGETParamsWithContext creates a new GetProjectUsingGETParams object
// with the ability to set a context for a request.
func NewGetProjectUsingGETParamsWithContext(ctx context.Context) *GetProjectUsingGETParams {
	return &GetProjectUsingGETParams{
		Context: ctx,
	}
}

// NewGetProjectUsingGETParamsWithHTTPClient creates a new GetProjectUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetProjectUsingGETParamsWithHTTPClient(client *http.Client) *GetProjectUsingGETParams {
	return &GetProjectUsingGETParams{
		HTTPClient: client,
	}
}

/* GetProjectUsingGETParams contains all the parameters to send to the API endpoint
   for the get project using g e t operation.

   Typically these are written to a http.Request.
*/
type GetProjectUsingGETParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format. For versioning information refer to /project-service/api/about.
	*/
	APIVersion *string

	/* ID.

	   The id of the project.
	*/
	ID string

	/* WithAnyPermission.

	   Optional permissions that, if granted to the user, allow him access to the proper set of projects. If the user actually has any of those permissions, the 'excludeViewer' parameter has no effect.
	*/
	WithAnyPermission []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get project using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectUsingGETParams) WithDefaults() *GetProjectUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get project using get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetProjectUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get project using get params
func (o *GetProjectUsingGETParams) WithTimeout(timeout time.Duration) *GetProjectUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project using get params
func (o *GetProjectUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project using get params
func (o *GetProjectUsingGETParams) WithContext(ctx context.Context) *GetProjectUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project using get params
func (o *GetProjectUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project using get params
func (o *GetProjectUsingGETParams) WithHTTPClient(client *http.Client) *GetProjectUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project using get params
func (o *GetProjectUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get project using get params
func (o *GetProjectUsingGETParams) WithAPIVersion(aPIVersion *string) *GetProjectUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get project using get params
func (o *GetProjectUsingGETParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get project using get params
func (o *GetProjectUsingGETParams) WithID(id string) *GetProjectUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get project using get params
func (o *GetProjectUsingGETParams) SetID(id string) {
	o.ID = id
}

// WithWithAnyPermission adds the withAnyPermission to the get project using get params
func (o *GetProjectUsingGETParams) WithWithAnyPermission(withAnyPermission []string) *GetProjectUsingGETParams {
	o.SetWithAnyPermission(withAnyPermission)
	return o
}

// SetWithAnyPermission adds the withAnyPermission to the get project using get params
func (o *GetProjectUsingGETParams) SetWithAnyPermission(withAnyPermission []string) {
	o.WithAnyPermission = withAnyPermission
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.WithAnyPermission != nil {

		// binding items for withAnyPermission
		joinedWithAnyPermission := o.bindParamWithAnyPermission(reg)

		// query array param withAnyPermission
		if err := r.SetQueryParam("withAnyPermission", joinedWithAnyPermission...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetProjectUsingGET binds the parameter withAnyPermission
func (o *GetProjectUsingGETParams) bindParamWithAnyPermission(formats strfmt.Registry) []string {
	withAnyPermissionIR := o.WithAnyPermission

	var withAnyPermissionIC []string
	for _, withAnyPermissionIIR := range withAnyPermissionIR { // explode []string

		withAnyPermissionIIV := withAnyPermissionIIR // string as string
		withAnyPermissionIC = append(withAnyPermissionIC, withAnyPermissionIIV)
	}

	// items.CollectionFormat: "multi"
	withAnyPermissionIS := swag.JoinByFormat(withAnyPermissionIC, "multi")

	return withAnyPermissionIS
}
