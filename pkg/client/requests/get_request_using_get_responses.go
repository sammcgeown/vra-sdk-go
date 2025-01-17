// Code generated by go-swagger; DO NOT EDIT.

package requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetRequestUsingGETReader is a Reader for the GetRequestUsingGET structure.
type GetRequestUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRequestUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRequestUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRequestUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRequestUsingGETOK creates a GetRequestUsingGETOK with default headers values
func NewGetRequestUsingGETOK() *GetRequestUsingGETOK {
	return &GetRequestUsingGETOK{}
}

/* GetRequestUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetRequestUsingGETOK struct {
	Payload *models.Request
}

func (o *GetRequestUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/requests/{requestId}][%d] getRequestUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetRequestUsingGETOK) GetPayload() *models.Request {
	return o.Payload
}

func (o *GetRequestUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestUsingGETUnauthorized creates a GetRequestUsingGETUnauthorized with default headers values
func NewGetRequestUsingGETUnauthorized() *GetRequestUsingGETUnauthorized {
	return &GetRequestUsingGETUnauthorized{}
}

/* GetRequestUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRequestUsingGETUnauthorized struct {
}

func (o *GetRequestUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/requests/{requestId}][%d] getRequestUsingGETUnauthorized ", 401)
}

func (o *GetRequestUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRequestUsingGETNotFound creates a GetRequestUsingGETNotFound with default headers values
func NewGetRequestUsingGETNotFound() *GetRequestUsingGETNotFound {
	return &GetRequestUsingGETNotFound{}
}

/* GetRequestUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRequestUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetRequestUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/requests/{requestId}][%d] getRequestUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetRequestUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetRequestUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
