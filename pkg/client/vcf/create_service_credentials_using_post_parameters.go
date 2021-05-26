// Code generated by go-swagger; DO NOT EDIT.

package vcf

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

// NewCreateServiceCredentialsUsingPOSTParams creates a new CreateServiceCredentialsUsingPOSTParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateServiceCredentialsUsingPOSTParams() *CreateServiceCredentialsUsingPOSTParams {
	return &CreateServiceCredentialsUsingPOSTParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateServiceCredentialsUsingPOSTParamsWithTimeout creates a new CreateServiceCredentialsUsingPOSTParams object
// with the ability to set a timeout on a request.
func NewCreateServiceCredentialsUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateServiceCredentialsUsingPOSTParams {
	return &CreateServiceCredentialsUsingPOSTParams{
		timeout: timeout,
	}
}

// NewCreateServiceCredentialsUsingPOSTParamsWithContext creates a new CreateServiceCredentialsUsingPOSTParams object
// with the ability to set a context for a request.
func NewCreateServiceCredentialsUsingPOSTParamsWithContext(ctx context.Context) *CreateServiceCredentialsUsingPOSTParams {
	return &CreateServiceCredentialsUsingPOSTParams{
		Context: ctx,
	}
}

// NewCreateServiceCredentialsUsingPOSTParamsWithHTTPClient creates a new CreateServiceCredentialsUsingPOSTParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateServiceCredentialsUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateServiceCredentialsUsingPOSTParams {
	return &CreateServiceCredentialsUsingPOSTParams{
		HTTPClient: client,
	}
}

/* CreateServiceCredentialsUsingPOSTParams contains all the parameters to send to the API endpoint
   for the create service credentials using p o s t operation.

   Typically these are written to a http.Request.
*/
type CreateServiceCredentialsUsingPOSTParams struct {

	/* DomainID.

	   domainId
	*/
	DomainID string

	/* IntegrationID.

	   integrationId
	*/
	IntegrationID string

	/* Request.

	   request
	*/
	Request *models.VcfCredentialRequests

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create service credentials using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceCredentialsUsingPOSTParams) WithDefaults() *CreateServiceCredentialsUsingPOSTParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create service credentials using p o s t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateServiceCredentialsUsingPOSTParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateServiceCredentialsUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) WithContext(ctx context.Context) *CreateServiceCredentialsUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateServiceCredentialsUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) WithDomainID(domainID string) *CreateServiceCredentialsUsingPOSTParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithIntegrationID adds the integrationID to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) WithIntegrationID(integrationID string) *CreateServiceCredentialsUsingPOSTParams {
	o.SetIntegrationID(integrationID)
	return o
}

// SetIntegrationID adds the integrationId to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) SetIntegrationID(integrationID string) {
	o.IntegrationID = integrationID
}

// WithRequest adds the request to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) WithRequest(request *models.VcfCredentialRequests) *CreateServiceCredentialsUsingPOSTParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the create service credentials using p o s t params
func (o *CreateServiceCredentialsUsingPOSTParams) SetRequest(request *models.VcfCredentialRequests) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *CreateServiceCredentialsUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	// path param integrationId
	if err := r.SetPathParam("integrationId", o.IntegrationID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
