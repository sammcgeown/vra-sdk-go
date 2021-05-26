// Code generated by go-swagger; DO NOT EDIT.

package security_group

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

// NewReconfigureSecurityGroupParams creates a new ReconfigureSecurityGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReconfigureSecurityGroupParams() *ReconfigureSecurityGroupParams {
	return &ReconfigureSecurityGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReconfigureSecurityGroupParamsWithTimeout creates a new ReconfigureSecurityGroupParams object
// with the ability to set a timeout on a request.
func NewReconfigureSecurityGroupParamsWithTimeout(timeout time.Duration) *ReconfigureSecurityGroupParams {
	return &ReconfigureSecurityGroupParams{
		timeout: timeout,
	}
}

// NewReconfigureSecurityGroupParamsWithContext creates a new ReconfigureSecurityGroupParams object
// with the ability to set a context for a request.
func NewReconfigureSecurityGroupParamsWithContext(ctx context.Context) *ReconfigureSecurityGroupParams {
	return &ReconfigureSecurityGroupParams{
		Context: ctx,
	}
}

// NewReconfigureSecurityGroupParamsWithHTTPClient creates a new ReconfigureSecurityGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewReconfigureSecurityGroupParamsWithHTTPClient(client *http.Client) *ReconfigureSecurityGroupParams {
	return &ReconfigureSecurityGroupParams{
		HTTPClient: client,
	}
}

/* ReconfigureSecurityGroupParams contains all the parameters to send to the API endpoint
   for the reconfigure security group operation.

   Typically these are written to a http.Request.
*/
type ReconfigureSecurityGroupParams struct {

	/* APIVersion.

	   The version of the API in yyyy-MM-dd format (UTC). For versioning information refer to /iaas/api/about
	*/
	APIVersion *string

	/* Body.

	   Security group Specification instance
	*/
	Body *models.SecurityGroupSpecification

	/* ID.

	   The ID of the security group.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the reconfigure security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReconfigureSecurityGroupParams) WithDefaults() *ReconfigureSecurityGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the reconfigure security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReconfigureSecurityGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) WithTimeout(timeout time.Duration) *ReconfigureSecurityGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) WithContext(ctx context.Context) *ReconfigureSecurityGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) WithHTTPClient(client *http.Client) *ReconfigureSecurityGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) WithAPIVersion(aPIVersion *string) *ReconfigureSecurityGroupParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) WithBody(body *models.SecurityGroupSpecification) *ReconfigureSecurityGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) SetBody(body *models.SecurityGroupSpecification) {
	o.Body = body
}

// WithID adds the id to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) WithID(id string) *ReconfigureSecurityGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the reconfigure security group params
func (o *ReconfigureSecurityGroupParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReconfigureSecurityGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
