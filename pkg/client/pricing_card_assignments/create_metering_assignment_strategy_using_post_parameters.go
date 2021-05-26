// Code generated by go-swagger; DO NOT EDIT.

package pricing_card_assignments

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

// NewCreateMeteringAssignmentStrategyUsingPOSTParams creates a new CreateMeteringAssignmentStrategyUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMeteringAssignmentStrategyUsingPOSTParams() *CreateMeteringAssignmentStrategyUsingPOSTParams {
	return &CreateMeteringAssignmentStrategyUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMeteringAssignmentStrategyUsingPOSTParamsWithTimeout creates a new CreateMeteringAssignmentStrategyUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateMeteringAssignmentStrategyUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateMeteringAssignmentStrategyUsingPOSTParams {
	return &CreateMeteringAssignmentStrategyUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateMeteringAssignmentStrategyUsingPOSTParamsWithContext creates a new CreateMeteringAssignmentStrategyUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateMeteringAssignmentStrategyUsingPOSTParamsWithContext(ctx context.Context) *CreateMeteringAssignmentStrategyUsingPOSTParams {
	return &CreateMeteringAssignmentStrategyUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateMeteringAssignmentStrategyUsingPOSTParamsWithHTTPClient creates a new CreateMeteringAssignmentStrategyUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMeteringAssignmentStrategyUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateMeteringAssignmentStrategyUsingPOSTParams {
	return &CreateMeteringAssignmentStrategyUsingPOSTParams{
		HTTPClient: client,
	}
}

/* CreateMeteringAssignmentStrategyUsingPOSTParams contains all the parameters to send to the API endpoint
   for the create metering assignment strategy using p o s t operation.

   Typically these are written to a http.Request.
*/
type CreateMeteringAssignmentStrategyUsingPOSTParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). If you do not specify explicitly an exact version, you will be calling the latest supported API version.
	*/
	APIVersion *string

	/* MeteringAssignmentStrategy.

	   The pricing card assignment strategy to be created
	*/
	MeteringAssignmentStrategy *models.MeteringAssignmentStrategy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create metering assignment strategy using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) WithDefaults() *CreateMeteringAssignmentStrategyUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create metering assignment strategy using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateMeteringAssignmentStrategyUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) WithContext(ctx context.Context) *CreateMeteringAssignmentStrategyUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateMeteringAssignmentStrategyUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) WithAPIVersion(aPIVersion *string) *CreateMeteringAssignmentStrategyUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithMeteringAssignmentStrategy adds the meteringAssignmentStrategy to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) WithMeteringAssignmentStrategy(meteringAssignmentStrategy *models.MeteringAssignmentStrategy) *CreateMeteringAssignmentStrategyUsingPOSTParams {
	o.SetMeteringAssignmentStrategy(meteringAssignmentStrategy)
	return o
}

// SetMeteringAssignmentStrategy adds the meteringAssignmentStrategy to the create metering assignment strategy using p o s t params
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) SetMeteringAssignmentStrategy(meteringAssignmentStrategy *models.MeteringAssignmentStrategy) {
	o.MeteringAssignmentStrategy = meteringAssignmentStrategy
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMeteringAssignmentStrategyUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.MeteringAssignmentStrategy != nil {
		if err := r.SetBodyParam(o.MeteringAssignmentStrategy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
