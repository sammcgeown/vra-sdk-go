// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewPatchDeploymentUsingPATCHParams creates a new PatchDeploymentUsingPATCHParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchDeploymentUsingPATCHParams() *PatchDeploymentUsingPATCHParams {
	return &PatchDeploymentUsingPATCHParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDeploymentUsingPATCHParamsWithTimeout creates a new PatchDeploymentUsingPATCHParams object
// with the ability to set a timeout on a request.
func NewPatchDeploymentUsingPATCHParamsWithTimeout(timeout time.Duration) *PatchDeploymentUsingPATCHParams {
	return &PatchDeploymentUsingPATCHParams{
		timeout: timeout,
	}
}

// NewPatchDeploymentUsingPATCHParamsWithContext creates a new PatchDeploymentUsingPATCHParams object
// with the ability to set a context for a request.
func NewPatchDeploymentUsingPATCHParamsWithContext(ctx context.Context) *PatchDeploymentUsingPATCHParams {
	return &PatchDeploymentUsingPATCHParams{
		Context: ctx,
	}
}

// NewPatchDeploymentUsingPATCHParamsWithHTTPClient creates a new PatchDeploymentUsingPATCHParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchDeploymentUsingPATCHParamsWithHTTPClient(client *http.Client) *PatchDeploymentUsingPATCHParams {
	return &PatchDeploymentUsingPATCHParams{
		HTTPClient: client,
	}
}

/* PatchDeploymentUsingPATCHParams contains all the parameters to send to the API endpoint
   for the patch deployment using p a t c h operation.

   Typically these are written to a http.Request.
*/
type PatchDeploymentUsingPATCHParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* DepID.

	   Deployment ID

	   Format: uuid
	*/
	DepID strfmt.UUID

	/* Update.

	   A set of fields to overwrite the corresponding fields in the deployment
	*/
	Update *models.DeploymentUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch deployment using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeploymentUsingPATCHParams) WithDefaults() *PatchDeploymentUsingPATCHParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch deployment using p a t c h params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchDeploymentUsingPATCHParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) WithTimeout(timeout time.Duration) *PatchDeploymentUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) WithContext(ctx context.Context) *PatchDeploymentUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) WithHTTPClient(client *http.Client) *PatchDeploymentUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) WithAPIVersion(aPIVersion *string) *PatchDeploymentUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDepID adds the depID to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) WithDepID(depID strfmt.UUID) *PatchDeploymentUsingPATCHParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WithUpdate adds the update to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) WithUpdate(update *models.DeploymentUpdate) *PatchDeploymentUsingPATCHParams {
	o.SetUpdate(update)
	return o
}

// SetUpdate adds the update to the patch deployment using p a t c h params
func (o *PatchDeploymentUsingPATCHParams) SetUpdate(update *models.DeploymentUpdate) {
	o.Update = update
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDeploymentUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param depId
	if err := r.SetPathParam("depId", o.DepID.String()); err != nil {
		return err
	}
	if o.Update != nil {
		if err := r.SetBodyParam(o.Update); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
