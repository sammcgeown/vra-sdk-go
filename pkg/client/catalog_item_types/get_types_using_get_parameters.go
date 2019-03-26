// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_types

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

// NewGetTypesUsingGETParams creates a new GetTypesUsingGETParams object
// with the default values initialized.
func NewGetTypesUsingGETParams() *GetTypesUsingGETParams {
	var ()
	return &GetTypesUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTypesUsingGETParamsWithTimeout creates a new GetTypesUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTypesUsingGETParamsWithTimeout(timeout time.Duration) *GetTypesUsingGETParams {
	var ()
	return &GetTypesUsingGETParams{

		timeout: timeout,
	}
}

// NewGetTypesUsingGETParamsWithContext creates a new GetTypesUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTypesUsingGETParamsWithContext(ctx context.Context) *GetTypesUsingGETParams {
	var ()
	return &GetTypesUsingGETParams{

		Context: ctx,
	}
}

// NewGetTypesUsingGETParamsWithHTTPClient creates a new GetTypesUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTypesUsingGETParamsWithHTTPClient(client *http.Client) *GetTypesUsingGETParams {
	var ()
	return &GetTypesUsingGETParams{
		HTTPClient: client,
	}
}

/*GetTypesUsingGETParams contains all the parameters to send to the API endpoint
for the get types using g e t operation typically these are written to a http.Request
*/
type GetTypesUsingGETParams struct {

	/*Page
	  Results page you want to retrieve (0..N)

	*/
	Page *int32
	/*Size
	  Number of records per page.

	*/
	Size *int32
	/*Sort
	  Sorting criteria in the format: property(,asc|desc). Default sort order is ascending. Multiple sort criteria are supported.

	*/
	Sort []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get types using get params
func (o *GetTypesUsingGETParams) WithTimeout(timeout time.Duration) *GetTypesUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get types using get params
func (o *GetTypesUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get types using get params
func (o *GetTypesUsingGETParams) WithContext(ctx context.Context) *GetTypesUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get types using get params
func (o *GetTypesUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get types using get params
func (o *GetTypesUsingGETParams) WithHTTPClient(client *http.Client) *GetTypesUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get types using get params
func (o *GetTypesUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get types using get params
func (o *GetTypesUsingGETParams) WithPage(page *int32) *GetTypesUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get types using get params
func (o *GetTypesUsingGETParams) SetPage(page *int32) {
	o.Page = page
}

// WithSize adds the size to the get types using get params
func (o *GetTypesUsingGETParams) WithSize(size *int32) *GetTypesUsingGETParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get types using get params
func (o *GetTypesUsingGETParams) SetSize(size *int32) {
	o.Size = size
}

// WithSort adds the sort to the get types using get params
func (o *GetTypesUsingGETParams) WithSort(sort []string) *GetTypesUsingGETParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get types using get params
func (o *GetTypesUsingGETParams) SetSort(sort []string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetTypesUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "multi")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
