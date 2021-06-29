// Code generated by go-swagger; DO NOT EDIT.

package blueprint_terraform_integrations

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

// NewGetTerraformConfigurationSourceCommitListUsingGET1Params creates a new GetTerraformConfigurationSourceCommitListUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTerraformConfigurationSourceCommitListUsingGET1Params() *GetTerraformConfigurationSourceCommitListUsingGET1Params {
	return &GetTerraformConfigurationSourceCommitListUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTerraformConfigurationSourceCommitListUsingGET1ParamsWithTimeout creates a new GetTerraformConfigurationSourceCommitListUsingGET1Params object
// with the ability to set a timeout on a request.
func NewGetTerraformConfigurationSourceCommitListUsingGET1ParamsWithTimeout(timeout time.Duration) *GetTerraformConfigurationSourceCommitListUsingGET1Params {
	return &GetTerraformConfigurationSourceCommitListUsingGET1Params{
		timeout: timeout,
	}
}

// NewGetTerraformConfigurationSourceCommitListUsingGET1ParamsWithContext creates a new GetTerraformConfigurationSourceCommitListUsingGET1Params object
// with the ability to set a context for a request.
func NewGetTerraformConfigurationSourceCommitListUsingGET1ParamsWithContext(ctx context.Context) *GetTerraformConfigurationSourceCommitListUsingGET1Params {
	return &GetTerraformConfigurationSourceCommitListUsingGET1Params{
		Context: ctx,
	}
}

// NewGetTerraformConfigurationSourceCommitListUsingGET1ParamsWithHTTPClient creates a new GetTerraformConfigurationSourceCommitListUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewGetTerraformConfigurationSourceCommitListUsingGET1ParamsWithHTTPClient(client *http.Client) *GetTerraformConfigurationSourceCommitListUsingGET1Params {
	return &GetTerraformConfigurationSourceCommitListUsingGET1Params{
		HTTPClient: client,
	}
}

/* GetTerraformConfigurationSourceCommitListUsingGET1Params contains all the parameters to send to the API endpoint
   for the get terraform configuration source commit list using get1 operation.

   Typically these are written to a http.Request.
*/
type GetTerraformConfigurationSourceCommitListUsingGET1Params struct {

	/* ConfigurationSourceID.

	   A Configuration Source ID.
	*/
	ConfigurationSourceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get terraform configuration source commit list using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) WithDefaults() *GetTerraformConfigurationSourceCommitListUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get terraform configuration source commit list using get1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get terraform configuration source commit list using get1 params
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) WithTimeout(timeout time.Duration) *GetTerraformConfigurationSourceCommitListUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get terraform configuration source commit list using get1 params
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get terraform configuration source commit list using get1 params
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) WithContext(ctx context.Context) *GetTerraformConfigurationSourceCommitListUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get terraform configuration source commit list using get1 params
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get terraform configuration source commit list using get1 params
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) WithHTTPClient(client *http.Client) *GetTerraformConfigurationSourceCommitListUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get terraform configuration source commit list using get1 params
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConfigurationSourceID adds the configurationSourceID to the get terraform configuration source commit list using get1 params
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) WithConfigurationSourceID(configurationSourceID string) *GetTerraformConfigurationSourceCommitListUsingGET1Params {
	o.SetConfigurationSourceID(configurationSourceID)
	return o
}

// SetConfigurationSourceID adds the configurationSourceId to the get terraform configuration source commit list using get1 params
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) SetConfigurationSourceID(configurationSourceID string) {
	o.ConfigurationSourceID = configurationSourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTerraformConfigurationSourceCommitListUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param configurationSourceId
	qrConfigurationSourceID := o.ConfigurationSourceID
	qConfigurationSourceID := qrConfigurationSourceID
	if qConfigurationSourceID != "" {

		if err := r.SetQueryParam("configurationSourceId", qConfigurationSourceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}