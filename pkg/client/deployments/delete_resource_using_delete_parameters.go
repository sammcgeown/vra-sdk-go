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
)

// NewDeleteResourceUsingDELETEParams creates a new DeleteResourceUsingDELETEParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteResourceUsingDELETEParams() *DeleteResourceUsingDELETEParams {
	return &DeleteResourceUsingDELETEParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResourceUsingDELETEParamsWithTimeout creates a new DeleteResourceUsingDELETEParams object
// with the ability to set a timeout on a request.
func NewDeleteResourceUsingDELETEParamsWithTimeout(timeout time.Duration) *DeleteResourceUsingDELETEParams {
	return &DeleteResourceUsingDELETEParams{
		timeout: timeout,
	}
}

// NewDeleteResourceUsingDELETEParamsWithContext creates a new DeleteResourceUsingDELETEParams object
// with the ability to set a context for a request.
func NewDeleteResourceUsingDELETEParamsWithContext(ctx context.Context) *DeleteResourceUsingDELETEParams {
	return &DeleteResourceUsingDELETEParams{
		Context: ctx,
	}
}

// NewDeleteResourceUsingDELETEParamsWithHTTPClient creates a new DeleteResourceUsingDELETEParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteResourceUsingDELETEParamsWithHTTPClient(client *http.Client) *DeleteResourceUsingDELETEParams {
	return &DeleteResourceUsingDELETEParams{
		HTTPClient: client,
	}
}

/* DeleteResourceUsingDELETEParams contains all the parameters to send to the API endpoint
   for the delete resource using d e l e t e operation.

   Typically these are written to a http.Request.
*/
type DeleteResourceUsingDELETEParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* DeploymentID.

	   Deployment ID

	   Format: uuid
	*/
	DeploymentID strfmt.UUID

	/* ResourceID.

	   Resource ID

	   Format: uuid
	*/
	ResourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete resource using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceUsingDELETEParams) WithDefaults() *DeleteResourceUsingDELETEParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete resource using d e l e t e params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteResourceUsingDELETEParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) WithTimeout(timeout time.Duration) *DeleteResourceUsingDELETEParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) WithContext(ctx context.Context) *DeleteResourceUsingDELETEParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) WithHTTPClient(client *http.Client) *DeleteResourceUsingDELETEParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) WithAPIVersion(aPIVersion *string) *DeleteResourceUsingDELETEParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithDeploymentID adds the deploymentID to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) WithDeploymentID(deploymentID strfmt.UUID) *DeleteResourceUsingDELETEParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) SetDeploymentID(deploymentID strfmt.UUID) {
	o.DeploymentID = deploymentID
}

// WithResourceID adds the resourceID to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) WithResourceID(resourceID strfmt.UUID) *DeleteResourceUsingDELETEParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the delete resource using d e l e t e params
func (o *DeleteResourceUsingDELETEParams) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResourceUsingDELETEParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param deploymentId
	if err := r.SetPathParam("deploymentId", o.DeploymentID.String()); err != nil {
		return err
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
