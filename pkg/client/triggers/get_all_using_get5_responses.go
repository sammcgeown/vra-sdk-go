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

// GetAllUsingGET5Reader is a Reader for the GetAllUsingGET5 structure.
type GetAllUsingGET5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUsingGET5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllUsingGET5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllUsingGET5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllUsingGET5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllUsingGET5NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllUsingGET5InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllUsingGET5OK creates a GetAllUsingGET5OK with default headers values
func NewGetAllUsingGET5OK() *GetAllUsingGET5OK {
	return &GetAllUsingGET5OK{}
}

/* GetAllUsingGET5OK describes a response with status code 200, with default header values.

'Success' with get of gerrit listeners
*/
type GetAllUsingGET5OK struct {
	Payload models.GerritListeners
}

func (o *GetAllUsingGET5OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners][%d] getAllUsingGET5OK  %+v", 200, o.Payload)
}
func (o *GetAllUsingGET5OK) GetPayload() models.GerritListeners {
	return o.Payload
}

func (o *GetAllUsingGET5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListeners(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetAllUsingGET5Unauthorized creates a GetAllUsingGET5Unauthorized with default headers values
func NewGetAllUsingGET5Unauthorized() *GetAllUsingGET5Unauthorized {
	return &GetAllUsingGET5Unauthorized{}
}

/* GetAllUsingGET5Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetAllUsingGET5Unauthorized struct {
}

func (o *GetAllUsingGET5Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners][%d] getAllUsingGET5Unauthorized ", 401)
}

func (o *GetAllUsingGET5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsingGET5Forbidden creates a GetAllUsingGET5Forbidden with default headers values
func NewGetAllUsingGET5Forbidden() *GetAllUsingGET5Forbidden {
	return &GetAllUsingGET5Forbidden{}
}

/* GetAllUsingGET5Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAllUsingGET5Forbidden struct {
}

func (o *GetAllUsingGET5Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners][%d] getAllUsingGET5Forbidden ", 403)
}

func (o *GetAllUsingGET5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsingGET5NotFound creates a GetAllUsingGET5NotFound with default headers values
func NewGetAllUsingGET5NotFound() *GetAllUsingGET5NotFound {
	return &GetAllUsingGET5NotFound{}
}

/* GetAllUsingGET5NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAllUsingGET5NotFound struct {
	Payload *models.Error
}

func (o *GetAllUsingGET5NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners][%d] getAllUsingGET5NotFound  %+v", 404, o.Payload)
}
func (o *GetAllUsingGET5NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAllUsingGET5NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUsingGET5InternalServerError creates a GetAllUsingGET5InternalServerError with default headers values
func NewGetAllUsingGET5InternalServerError() *GetAllUsingGET5InternalServerError {
	return &GetAllUsingGET5InternalServerError{}
}

/* GetAllUsingGET5InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetAllUsingGET5InternalServerError struct {
}

func (o *GetAllUsingGET5InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-listeners][%d] getAllUsingGET5InternalServerError ", 500)
}

func (o *GetAllUsingGET5InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}