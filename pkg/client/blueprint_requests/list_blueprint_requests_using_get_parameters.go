// Code generated by go-swagger; DO NOT EDIT.

package blueprint_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListBlueprintRequestsUsingGETParams creates a new ListBlueprintRequestsUsingGETParams object
// with the default values initialized.
func NewListBlueprintRequestsUsingGETParams() *ListBlueprintRequestsUsingGETParams {
	var (
		expandDefault  = bool(true)
		orderByDefault = string("updatedAt DESC")
		pageDefault    = int32(1)
		sizeDefault    = int32(20)
	)
	return &ListBlueprintRequestsUsingGETParams{
		Expand:  &expandDefault,
		OrderBy: &orderByDefault,
		Page:    &pageDefault,
		Size:    &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListBlueprintRequestsUsingGETParamsWithTimeout creates a new ListBlueprintRequestsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBlueprintRequestsUsingGETParamsWithTimeout(timeout time.Duration) *ListBlueprintRequestsUsingGETParams {
	var (
		expandDefault  = bool(true)
		orderByDefault = string("updatedAt DESC")
		pageDefault    = int32(1)
		sizeDefault    = int32(20)
	)
	return &ListBlueprintRequestsUsingGETParams{
		Expand:  &expandDefault,
		OrderBy: &orderByDefault,
		Page:    &pageDefault,
		Size:    &sizeDefault,

		timeout: timeout,
	}
}

// NewListBlueprintRequestsUsingGETParamsWithContext creates a new ListBlueprintRequestsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBlueprintRequestsUsingGETParamsWithContext(ctx context.Context) *ListBlueprintRequestsUsingGETParams {
	var (
		expandDefault  = bool(true)
		orderByDefault = string("updatedAt DESC")
		pageDefault    = int32(1)
		sizeDefault    = int32(20)
	)
	return &ListBlueprintRequestsUsingGETParams{
		Expand:  &expandDefault,
		OrderBy: &orderByDefault,
		Page:    &pageDefault,
		Size:    &sizeDefault,

		Context: ctx,
	}
}

// NewListBlueprintRequestsUsingGETParamsWithHTTPClient creates a new ListBlueprintRequestsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBlueprintRequestsUsingGETParamsWithHTTPClient(client *http.Client) *ListBlueprintRequestsUsingGETParams {
	var (
		expandDefault  = bool(true)
		orderByDefault = string("updatedAt DESC")
		pageDefault    = int32(1)
		sizeDefault    = int32(20)
	)
	return &ListBlueprintRequestsUsingGETParams{
		Expand:     &expandDefault,
		OrderBy:    &orderByDefault,
		Page:       &pageDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*ListBlueprintRequestsUsingGETParams contains all the parameters to send to the API endpoint
for the list blueprint requests using g e t operation typically these are written to a http.Request
*/
type ListBlueprintRequestsUsingGETParams struct {

	/*DeploymentID
	  Deployment Id filter

	*/
	DeploymentID *string
	/*Expand
	  Expand with content

	*/
	Expand *bool
	/*Fields
	  Fields to include in content

	*/
	Fields []string
	/*IncludePlan
	  Plan Requests filter

	*/
	IncludePlan *bool
	/*OrderBy
	  Sorts blueprint requests (e.g. 'updatedAt DESC').

	*/
	OrderBy *string
	/*Page
	  Page index

	*/
	Page *int32
	/*Size
	  Page size

	*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithTimeout(timeout time.Duration) *ListBlueprintRequestsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithContext(ctx context.Context) *ListBlueprintRequestsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithHTTPClient(client *http.Client) *ListBlueprintRequestsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeploymentID adds the deploymentID to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithDeploymentID(deploymentID *string) *ListBlueprintRequestsUsingGETParams {
	o.SetDeploymentID(deploymentID)
	return o
}

// SetDeploymentID adds the deploymentId to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetDeploymentID(deploymentID *string) {
	o.DeploymentID = deploymentID
}

// WithExpand adds the expand to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithExpand(expand *bool) *ListBlueprintRequestsUsingGETParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetExpand(expand *bool) {
	o.Expand = expand
}

// WithFields adds the fields to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithFields(fields []string) *ListBlueprintRequestsUsingGETParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIncludePlan adds the includePlan to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithIncludePlan(includePlan *bool) *ListBlueprintRequestsUsingGETParams {
	o.SetIncludePlan(includePlan)
	return o
}

// SetIncludePlan adds the includePlan to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetIncludePlan(includePlan *bool) {
	o.IncludePlan = includePlan
}

// WithOrderBy adds the orderBy to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithOrderBy(orderBy *string) *ListBlueprintRequestsUsingGETParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithPage adds the page to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithPage(page *int32) *ListBlueprintRequestsUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetPage(page *int32) {
	o.Page = page
}

// WithSize adds the size to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) WithSize(size *int32) *ListBlueprintRequestsUsingGETParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the list blueprint requests using get params
func (o *ListBlueprintRequestsUsingGETParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *ListBlueprintRequestsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeploymentID != nil {

		// query param deploymentId
		var qrDeploymentID string
		if o.DeploymentID != nil {
			qrDeploymentID = *o.DeploymentID
		}
		qDeploymentID := qrDeploymentID
		if qDeploymentID != "" {
			if err := r.SetQueryParam("deploymentId", qDeploymentID); err != nil {
				return err
			}
		}

	}

	if o.Expand != nil {

		// query param expand
		var qrExpand bool
		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := swag.FormatBool(qrExpand)
		if qExpand != "" {
			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}

	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "multi")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	if o.IncludePlan != nil {

		// query param includePlan
		var qrIncludePlan bool
		if o.IncludePlan != nil {
			qrIncludePlan = *o.IncludePlan
		}
		qIncludePlan := swag.FormatBool(qrIncludePlan)
		if qIncludePlan != "" {
			if err := r.SetQueryParam("includePlan", qIncludePlan); err != nil {
				return err
			}
		}

	}

	if o.OrderBy != nil {

		// query param orderBy
		var qrOrderBy string
		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {
			if err := r.SetQueryParam("orderBy", qOrderBy); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
