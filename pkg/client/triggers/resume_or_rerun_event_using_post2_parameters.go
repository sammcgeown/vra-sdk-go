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
)

// NewResumeOrRerunEventUsingPOST2Params creates a new ResumeOrRerunEventUsingPOST2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResumeOrRerunEventUsingPOST2Params() *ResumeOrRerunEventUsingPOST2Params {
	return &ResumeOrRerunEventUsingPOST2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewResumeOrRerunEventUsingPOST2ParamsWithTimeout creates a new ResumeOrRerunEventUsingPOST2Params object
// with the ability to set a timeout on a request.
func NewResumeOrRerunEventUsingPOST2ParamsWithTimeout(timeout time.Duration) *ResumeOrRerunEventUsingPOST2Params {
	return &ResumeOrRerunEventUsingPOST2Params{
		timeout: timeout,
	}
}

// NewResumeOrRerunEventUsingPOST2ParamsWithContext creates a new ResumeOrRerunEventUsingPOST2Params object
// with the ability to set a context for a request.
func NewResumeOrRerunEventUsingPOST2ParamsWithContext(ctx context.Context) *ResumeOrRerunEventUsingPOST2Params {
	return &ResumeOrRerunEventUsingPOST2Params{
		Context: ctx,
	}
}

// NewResumeOrRerunEventUsingPOST2ParamsWithHTTPClient creates a new ResumeOrRerunEventUsingPOST2Params object
// with the ability to set a custom HTTPClient for a request.
func NewResumeOrRerunEventUsingPOST2ParamsWithHTTPClient(client *http.Client) *ResumeOrRerunEventUsingPOST2Params {
	return &ResumeOrRerunEventUsingPOST2Params{
		HTTPClient: client,
	}
}

/* ResumeOrRerunEventUsingPOST2Params contains all the parameters to send to the API endpoint
   for the resume or rerun event using p o s t 2 operation.

   Typically these are written to a http.Request.
*/
type ResumeOrRerunEventUsingPOST2Params struct {

	/* Authorization.

	   Bearer token
	*/
	Authorization string

	/* Action.

	   Resume/Rerun
	*/
	Action string

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about
	*/
	APIVersion string

	/* ID.

	   id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resume or rerun event using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeOrRerunEventUsingPOST2Params) WithDefaults() *ResumeOrRerunEventUsingPOST2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resume or rerun event using p o s t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeOrRerunEventUsingPOST2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) WithTimeout(timeout time.Duration) *ResumeOrRerunEventUsingPOST2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) WithContext(ctx context.Context) *ResumeOrRerunEventUsingPOST2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) WithHTTPClient(client *http.Client) *ResumeOrRerunEventUsingPOST2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) WithAuthorization(authorization string) *ResumeOrRerunEventUsingPOST2Params {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithAction adds the action to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) WithAction(action string) *ResumeOrRerunEventUsingPOST2Params {
	o.SetAction(action)
	return o
}

// SetAction adds the action to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) SetAction(action string) {
	o.Action = action
}

// WithAPIVersion adds the aPIVersion to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) WithAPIVersion(aPIVersion string) *ResumeOrRerunEventUsingPOST2Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) WithID(id string) *ResumeOrRerunEventUsingPOST2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the resume or rerun event using p o s t 2 params
func (o *ResumeOrRerunEventUsingPOST2Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ResumeOrRerunEventUsingPOST2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// query param action
	qrAction := o.Action
	qAction := qrAction
	if qAction != "" {

		if err := r.SetQueryParam("action", qAction); err != nil {
			return err
		}
	}

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion

	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
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
