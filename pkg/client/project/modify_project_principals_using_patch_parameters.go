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

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// NewModifyProjectPrincipalsUsingPATCHParams creates a new ModifyProjectPrincipalsUsingPATCHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyProjectPrincipalsUsingPATCHParams() *ModifyProjectPrincipalsUsingPATCHParams {
	return &ModifyProjectPrincipalsUsingPATCHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyProjectPrincipalsUsingPATCHParamsWithTimeout creates a new ModifyProjectPrincipalsUsingPATCHParams object
// with the ability to set a timeout on a request.
func NewModifyProjectPrincipalsUsingPATCHParamsWithTimeout(timeout time.Duration) *ModifyProjectPrincipalsUsingPATCHParams {
	return &ModifyProjectPrincipalsUsingPATCHParams{
		timeout: timeout,
	}
}

// NewModifyProjectPrincipalsUsingPATCHParamsWithContext creates a new ModifyProjectPrincipalsUsingPATCHParams object
// with the ability to set a context for a request.
func NewModifyProjectPrincipalsUsingPATCHParamsWithContext(ctx context.Context) *ModifyProjectPrincipalsUsingPATCHParams {
	return &ModifyProjectPrincipalsUsingPATCHParams{
		Context: ctx,
	}
}

// NewModifyProjectPrincipalsUsingPATCHParamsWithHTTPClient creates a new ModifyProjectPrincipalsUsingPATCHParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyProjectPrincipalsUsingPATCHParamsWithHTTPClient(client *http.Client) *ModifyProjectPrincipalsUsingPATCHParams {
	return &ModifyProjectPrincipalsUsingPATCHParams{
		HTTPClient: client,
	}
}

/* ModifyProjectPrincipalsUsingPATCHParams contains all the parameters to send to the API endpoint
   for the modify project principals using p a t c h operation.

   Typically these are written to a http.Request.
*/
type ModifyProjectPrincipalsUsingPATCHParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format. For versioning information refer to /project-service/api/about.
	*/
	APIVersion *string

	/* ID.

	   The id of the project.
	*/
	ID string

	/* Roles.

	   roles
	*/
	Roles *models.ProjectPrincipalsAssignment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the modify project principals using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyProjectPrincipalsUsingPATCHParams) WithDefaults() *ModifyProjectPrincipalsUsingPATCHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify project principals using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyProjectPrincipalsUsingPATCHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) WithTimeout(timeout time.Duration) *ModifyProjectPrincipalsUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) WithContext(ctx context.Context) *ModifyProjectPrincipalsUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) WithHTTPClient(client *http.Client) *ModifyProjectPrincipalsUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) WithAPIVersion(aPIVersion *string) *ModifyProjectPrincipalsUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) WithID(id string) *ModifyProjectPrincipalsUsingPATCHParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) SetID(id string) {
	o.ID = id
}

// WithRoles adds the roles to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) WithRoles(roles *models.ProjectPrincipalsAssignment) *ModifyProjectPrincipalsUsingPATCHParams {
	o.SetRoles(roles)
	return o
}

// SetRoles adds the roles to the modify project principals using p a t c h params
func (o *ModifyProjectPrincipalsUsingPATCHParams) SetRoles(roles *models.ProjectPrincipalsAssignment) {
	o.Roles = roles
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyProjectPrincipalsUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Roles != nil {
		if err := r.SetBodyParam(o.Roles); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
