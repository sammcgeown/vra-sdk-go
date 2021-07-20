// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetAllUsingGETReader is a Reader for the GetAllUsingGET structure.
type GetAllUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllUsingGETOK creates a GetAllUsingGETOK with default headers values
func NewGetAllUsingGETOK() *GetAllUsingGETOK {
	return &GetAllUsingGETOK{}
}

/* GetAllUsingGETOK describes a response with status code 200, with default header values.

'Success' with get of Docker Registry Events
*/
type GetAllUsingGETOK struct {
	Payload models.DockerRegistryEvents
}

func (o *GetAllUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-events][%d] getAllUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetAllUsingGETOK) GetPayload() models.DockerRegistryEvents {
	return o.Payload
}

func (o *GetAllUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalDockerRegistryEvents(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetAllUsingGETUnauthorized creates a GetAllUsingGETUnauthorized with default headers values
func NewGetAllUsingGETUnauthorized() *GetAllUsingGETUnauthorized {
	return &GetAllUsingGETUnauthorized{}
}

/* GetAllUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetAllUsingGETUnauthorized struct {
}

func (o *GetAllUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-events][%d] getAllUsingGETUnauthorized ", 401)
}

func (o *GetAllUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsingGETForbidden creates a GetAllUsingGETForbidden with default headers values
func NewGetAllUsingGETForbidden() *GetAllUsingGETForbidden {
	return &GetAllUsingGETForbidden{}
}

/* GetAllUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllUsingGETForbidden struct {
}

func (o *GetAllUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-events][%d] getAllUsingGETForbidden ", 403)
}

func (o *GetAllUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsingGETNotFound creates a GetAllUsingGETNotFound with default headers values
func NewGetAllUsingGETNotFound() *GetAllUsingGETNotFound {
	return &GetAllUsingGETNotFound{}
}

/* GetAllUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetAllUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-events][%d] getAllUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetAllUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUsingGETInternalServerError creates a GetAllUsingGETInternalServerError with default headers values
func NewGetAllUsingGETInternalServerError() *GetAllUsingGETInternalServerError {
	return &GetAllUsingGETInternalServerError{}
}

/* GetAllUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAllUsingGETInternalServerError struct {
}

func (o *GetAllUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/registry-events][%d] getAllUsingGETInternalServerError ", 500)
}

func (o *GetAllUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}