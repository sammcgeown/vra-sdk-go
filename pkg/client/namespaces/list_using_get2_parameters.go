// Code generated by go-swagger; DO NOT EDIT.

package namespaces

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
	"github.com/go-openapi/swag"
)

// NewListUsingGET2Params creates a new ListUsingGET2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListUsingGET2Params() *ListUsingGET2Params {
	return &ListUsingGET2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewListUsingGET2ParamsWithTimeout creates a new ListUsingGET2Params object
// with the ability to set a timeout on a request.
func NewListUsingGET2ParamsWithTimeout(timeout time.Duration) *ListUsingGET2Params {
	return &ListUsingGET2Params{
		timeout: timeout,
	}
}

// NewListUsingGET2ParamsWithContext creates a new ListUsingGET2Params object
// with the ability to set a context for a request.
func NewListUsingGET2ParamsWithContext(ctx context.Context) *ListUsingGET2Params {
	return &ListUsingGET2Params{
		Context: ctx,
	}
}

// NewListUsingGET2ParamsWithHTTPClient creates a new ListUsingGET2Params object
// with the ability to set a custom HTTPClient for a request.
func NewListUsingGET2ParamsWithHTTPClient(client *http.Client) *ListUsingGET2Params {
	return &ListUsingGET2Params{
		HTTPClient: client,
	}
}

/* ListUsingGET2Params contains all the parameters to send to the API endpoint
   for the list using g e t 2 operation.

   Typically these are written to a http.Request.
*/
type ListUsingGET2Params struct {

	// Offset.
	//
	// Format: int64
	Offset *int64

	// PageNumber.
	//
	// Format: int32
	PageNumber *int32

	// PageSize.
	//
	// Format: int32
	PageSize *int32

	// Paged.
	Paged *bool

	// SortSorted.
	SortSorted *bool

	// SortUnsorted.
	SortUnsorted *bool

	// Unpaged.
	Unpaged *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUsingGET2Params) WithDefaults() *ListUsingGET2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list using g e t 2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListUsingGET2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list using g e t 2 params
func (o *ListUsingGET2Params) WithTimeout(timeout time.Duration) *ListUsingGET2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list using g e t 2 params
func (o *ListUsingGET2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list using g e t 2 params
func (o *ListUsingGET2Params) WithContext(ctx context.Context) *ListUsingGET2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list using g e t 2 params
func (o *ListUsingGET2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list using g e t 2 params
func (o *ListUsingGET2Params) WithHTTPClient(client *http.Client) *ListUsingGET2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list using g e t 2 params
func (o *ListUsingGET2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOffset adds the offset to the list using g e t 2 params
func (o *ListUsingGET2Params) WithOffset(offset *int64) *ListUsingGET2Params {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the list using g e t 2 params
func (o *ListUsingGET2Params) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPageNumber adds the pageNumber to the list using g e t 2 params
func (o *ListUsingGET2Params) WithPageNumber(pageNumber *int32) *ListUsingGET2Params {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the list using g e t 2 params
func (o *ListUsingGET2Params) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the list using g e t 2 params
func (o *ListUsingGET2Params) WithPageSize(pageSize *int32) *ListUsingGET2Params {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list using g e t 2 params
func (o *ListUsingGET2Params) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPaged adds the paged to the list using g e t 2 params
func (o *ListUsingGET2Params) WithPaged(paged *bool) *ListUsingGET2Params {
	o.SetPaged(paged)
	return o
}

// SetPaged adds the paged to the list using g e t 2 params
func (o *ListUsingGET2Params) SetPaged(paged *bool) {
	o.Paged = paged
}

// WithSortSorted adds the sortSorted to the list using g e t 2 params
func (o *ListUsingGET2Params) WithSortSorted(sortSorted *bool) *ListUsingGET2Params {
	o.SetSortSorted(sortSorted)
	return o
}

// SetSortSorted adds the sortSorted to the list using g e t 2 params
func (o *ListUsingGET2Params) SetSortSorted(sortSorted *bool) {
	o.SortSorted = sortSorted
}

// WithSortUnsorted adds the sortUnsorted to the list using g e t 2 params
func (o *ListUsingGET2Params) WithSortUnsorted(sortUnsorted *bool) *ListUsingGET2Params {
	o.SetSortUnsorted(sortUnsorted)
	return o
}

// SetSortUnsorted adds the sortUnsorted to the list using g e t 2 params
func (o *ListUsingGET2Params) SetSortUnsorted(sortUnsorted *bool) {
	o.SortUnsorted = sortUnsorted
}

// WithUnpaged adds the unpaged to the list using g e t 2 params
func (o *ListUsingGET2Params) WithUnpaged(unpaged *bool) *ListUsingGET2Params {
	o.SetUnpaged(unpaged)
	return o
}

// SetUnpaged adds the unpaged to the list using g e t 2 params
func (o *ListUsingGET2Params) SetUnpaged(unpaged *bool) {
	o.Unpaged = unpaged
}

// WriteToRequest writes these params to a swagger request
func (o *ListUsingGET2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Paged != nil {

		// query param paged
		var qrPaged bool

		if o.Paged != nil {
			qrPaged = *o.Paged
		}
		qPaged := swag.FormatBool(qrPaged)
		if qPaged != "" {

			if err := r.SetQueryParam("paged", qPaged); err != nil {
				return err
			}
		}
	}

	if o.SortSorted != nil {

		// query param sort.sorted
		var qrSortSorted bool

		if o.SortSorted != nil {
			qrSortSorted = *o.SortSorted
		}
		qSortSorted := swag.FormatBool(qrSortSorted)
		if qSortSorted != "" {

			if err := r.SetQueryParam("sort.sorted", qSortSorted); err != nil {
				return err
			}
		}
	}

	if o.SortUnsorted != nil {

		// query param sort.unsorted
		var qrSortUnsorted bool

		if o.SortUnsorted != nil {
			qrSortUnsorted = *o.SortUnsorted
		}
		qSortUnsorted := swag.FormatBool(qrSortUnsorted)
		if qSortUnsorted != "" {

			if err := r.SetQueryParam("sort.unsorted", qSortUnsorted); err != nil {
				return err
			}
		}
	}

	if o.Unpaged != nil {

		// query param unpaged
		var qrUnpaged bool

		if o.Unpaged != nil {
			qrUnpaged = *o.Unpaged
		}
		qUnpaged := swag.FormatBool(qrUnpaged)
		if qUnpaged != "" {

			if err := r.SetQueryParam("unpaged", qUnpaged); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}