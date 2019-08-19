// Code generated by go-swagger; DO NOT EDIT.

package catalog_admin_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetCatalogItemUsingGETReader is a Reader for the GetCatalogItemUsingGET structure.
type GetCatalogItemUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogItemUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCatalogItemUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetCatalogItemUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCatalogItemUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCatalogItemUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCatalogItemUsingGETOK creates a GetCatalogItemUsingGETOK with default headers values
func NewGetCatalogItemUsingGETOK() *GetCatalogItemUsingGETOK {
	return &GetCatalogItemUsingGETOK{}
}

/*GetCatalogItemUsingGETOK handles this case with default header values.

OK
*/
type GetCatalogItemUsingGETOK struct {
	Payload *models.CatalogItem
}

func (o *GetCatalogItemUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/items/{id}][%d] getCatalogItemUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetCatalogItemUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CatalogItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogItemUsingGETUnauthorized creates a GetCatalogItemUsingGETUnauthorized with default headers values
func NewGetCatalogItemUsingGETUnauthorized() *GetCatalogItemUsingGETUnauthorized {
	return &GetCatalogItemUsingGETUnauthorized{}
}

/*GetCatalogItemUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCatalogItemUsingGETUnauthorized struct {
}

func (o *GetCatalogItemUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/items/{id}][%d] getCatalogItemUsingGETUnauthorized ", 401)
}

func (o *GetCatalogItemUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCatalogItemUsingGETForbidden creates a GetCatalogItemUsingGETForbidden with default headers values
func NewGetCatalogItemUsingGETForbidden() *GetCatalogItemUsingGETForbidden {
	return &GetCatalogItemUsingGETForbidden{}
}

/*GetCatalogItemUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetCatalogItemUsingGETForbidden struct {
}

func (o *GetCatalogItemUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/items/{id}][%d] getCatalogItemUsingGETForbidden ", 403)
}

func (o *GetCatalogItemUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCatalogItemUsingGETNotFound creates a GetCatalogItemUsingGETNotFound with default headers values
func NewGetCatalogItemUsingGETNotFound() *GetCatalogItemUsingGETNotFound {
	return &GetCatalogItemUsingGETNotFound{}
}

/*GetCatalogItemUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetCatalogItemUsingGETNotFound struct {
}

func (o *GetCatalogItemUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /catalog/api/admin/items/{id}][%d] getCatalogItemUsingGETNotFound ", 404)
}

func (o *GetCatalogItemUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
