// Code generated by go-swagger; DO NOT EDIT.

package kubernetes_zones

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

// NewUpdateZoneProjectsUsingPUTParams creates a new UpdateZoneProjectsUsingPUTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateZoneProjectsUsingPUTParams() *UpdateZoneProjectsUsingPUTParams {
	return &UpdateZoneProjectsUsingPUTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateZoneProjectsUsingPUTParamsWithTimeout creates a new UpdateZoneProjectsUsingPUTParams object
// with the ability to set a timeout on a request.
func NewUpdateZoneProjectsUsingPUTParamsWithTimeout(timeout time.Duration) *UpdateZoneProjectsUsingPUTParams {
	return &UpdateZoneProjectsUsingPUTParams{
		timeout: timeout,
	}
}

// NewUpdateZoneProjectsUsingPUTParamsWithContext creates a new UpdateZoneProjectsUsingPUTParams object
// with the ability to set a context for a request.
func NewUpdateZoneProjectsUsingPUTParamsWithContext(ctx context.Context) *UpdateZoneProjectsUsingPUTParams {
	return &UpdateZoneProjectsUsingPUTParams{
		Context: ctx,
	}
}

// NewUpdateZoneProjectsUsingPUTParamsWithHTTPClient creates a new UpdateZoneProjectsUsingPUTParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateZoneProjectsUsingPUTParamsWithHTTPClient(client *http.Client) *UpdateZoneProjectsUsingPUTParams {
	return &UpdateZoneProjectsUsingPUTParams{
		HTTPClient: client,
	}
}

/* UpdateZoneProjectsUsingPUTParams contains all the parameters to send to the API endpoint
   for the update zone projects using p u t operation.

   Typically these are written to a http.Request.
*/
type UpdateZoneProjectsUsingPUTParams struct {

	/* ID.

	   id

	   Format: uuid
	*/
	ID strfmt.UUID

	/* Projects.

	   projects
	*/
	Projects []*models.K8SZoneProjectAssignment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update zone projects using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateZoneProjectsUsingPUTParams) WithDefaults() *UpdateZoneProjectsUsingPUTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update zone projects using p u t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateZoneProjectsUsingPUTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) WithTimeout(timeout time.Duration) *UpdateZoneProjectsUsingPUTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) WithContext(ctx context.Context) *UpdateZoneProjectsUsingPUTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) WithHTTPClient(client *http.Client) *UpdateZoneProjectsUsingPUTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) WithID(id strfmt.UUID) *UpdateZoneProjectsUsingPUTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithProjects adds the projects to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) WithProjects(projects []*models.K8SZoneProjectAssignment) *UpdateZoneProjectsUsingPUTParams {
	o.SetProjects(projects)
	return o
}

// SetProjects adds the projects to the update zone projects using p u t params
func (o *UpdateZoneProjectsUsingPUTParams) SetProjects(projects []*models.K8SZoneProjectAssignment) {
	o.Projects = projects
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateZoneProjectsUsingPUTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}
	if o.Projects != nil {
		if err := r.SetBodyParam(o.Projects); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}