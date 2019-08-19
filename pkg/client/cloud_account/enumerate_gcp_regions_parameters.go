// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewEnumerateGcpRegionsParams creates a new EnumerateGcpRegionsParams object
// with the default values initialized.
func NewEnumerateGcpRegionsParams() *EnumerateGcpRegionsParams {
	var ()
	return &EnumerateGcpRegionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnumerateGcpRegionsParamsWithTimeout creates a new EnumerateGcpRegionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnumerateGcpRegionsParamsWithTimeout(timeout time.Duration) *EnumerateGcpRegionsParams {
	var ()
	return &EnumerateGcpRegionsParams{

		timeout: timeout,
	}
}

// NewEnumerateGcpRegionsParamsWithContext creates a new EnumerateGcpRegionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnumerateGcpRegionsParamsWithContext(ctx context.Context) *EnumerateGcpRegionsParams {
	var ()
	return &EnumerateGcpRegionsParams{

		Context: ctx,
	}
}

// NewEnumerateGcpRegionsParamsWithHTTPClient creates a new EnumerateGcpRegionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnumerateGcpRegionsParamsWithHTTPClient(client *http.Client) *EnumerateGcpRegionsParams {
	var ()
	return &EnumerateGcpRegionsParams{
		HTTPClient: client,
	}
}

/*EnumerateGcpRegionsParams contains all the parameters to send to the API endpoint
for the enumerate gcp regions operation typically these are written to a http.Request
*/
type EnumerateGcpRegionsParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  CloudAccount specification

	*/
	Body *models.CloudAccountGcpSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) WithTimeout(timeout time.Duration) *EnumerateGcpRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) WithContext(ctx context.Context) *EnumerateGcpRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) WithHTTPClient(client *http.Client) *EnumerateGcpRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) WithAPIVersion(aPIVersion *string) *EnumerateGcpRegionsParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) WithBody(body *models.CloudAccountGcpSpecification) *EnumerateGcpRegionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the enumerate gcp regions params
func (o *EnumerateGcpRegionsParams) SetBody(body *models.CloudAccountGcpSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *EnumerateGcpRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
