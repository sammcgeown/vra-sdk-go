// Code generated by go-swagger; DO NOT EDIT.

package triggers

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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateByIDUsingPUT3Params creates a new UpdateByIDUsingPUT3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateByIDUsingPUT3Params() *UpdateByIDUsingPUT3Params {
	return &UpdateByIDUsingPUT3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateByIDUsingPUT3ParamsWithTimeout creates a new UpdateByIDUsingPUT3Params object
// with the ability to set a timeout on a request.
func NewUpdateByIDUsingPUT3ParamsWithTimeout(timeout time.Duration) *UpdateByIDUsingPUT3Params {
	return &UpdateByIDUsingPUT3Params{
		timeout: timeout,
	}
}

// NewUpdateByIDUsingPUT3ParamsWithContext creates a new UpdateByIDUsingPUT3Params object
// with the ability to set a context for a request.
func NewUpdateByIDUsingPUT3ParamsWithContext(ctx context.Context) *UpdateByIDUsingPUT3Params {
	return &UpdateByIDUsingPUT3Params{
		Context: ctx,
	}
}

// NewUpdateByIDUsingPUT3ParamsWithHTTPClient creates a new UpdateByIDUsingPUT3Params object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateByIDUsingPUT3ParamsWithHTTPClient(client *http.Client) *UpdateByIDUsingPUT3Params {
	return &UpdateByIDUsingPUT3Params{
		HTTPClient: client,
	}
}

/* UpdateByIDUsingPUT3Params contains all the parameters to send to the API endpoint
   for the update by Id using p u t 3 operation.

   Typically these are written to a http.Request.
*/
type UpdateByIDUsingPUT3Params struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* GerritTriggerSpec.

	   gerritTriggerSpec
	*/
	GerritTriggerSpec models.GerritTriggerSpec

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update by Id using p u t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateByIDUsingPUT3Params) WithDefaults() *UpdateByIDUsingPUT3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update by Id using p u t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateByIDUsingPUT3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) WithTimeout(timeout time.Duration) *UpdateByIDUsingPUT3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) WithContext(ctx context.Context) *UpdateByIDUsingPUT3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) WithHTTPClient(client *http.Client) *UpdateByIDUsingPUT3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) WithAuthorization(authorization string) *UpdateByIDUsingPUT3Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAPIVersion adds the aPIVersion to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) WithAPIVersion(aPIVersion string) *UpdateByIDUsingPUT3Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGerritTriggerSpec adds the gerritTriggerSpec to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) WithGerritTriggerSpec(gerritTriggerSpec models.GerritTriggerSpec) *UpdateByIDUsingPUT3Params {
	o.SetGerritTriggerSpec(gerritTriggerSpec)
	return o
}

// SetGerritTriggerSpec adds the gerritTriggerSpec to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) SetGerritTriggerSpec(gerritTriggerSpec models.GerritTriggerSpec) {
	o.GerritTriggerSpec = gerritTriggerSpec
}

// WithID adds the id to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) WithID(id string) *UpdateByIDUsingPUT3Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the update by Id using p u t 3 params
func (o *UpdateByIDUsingPUT3Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateByIDUsingPUT3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.GerritTriggerSpec); err != nil {
		return err
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