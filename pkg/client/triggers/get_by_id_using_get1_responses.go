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

// GetByIDUsingGET1Reader is a Reader for the GetByIDUsingGET1 structure.
type GetByIDUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetByIDUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetByIDUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetByIDUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetByIDUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetByIDUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetByIDUsingGET1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetByIDUsingGET1OK creates a GetByIDUsingGET1OK with default headers values
func NewGetByIDUsingGET1OK() *GetByIDUsingGET1OK {
	return &GetByIDUsingGET1OK{}
}

/* GetByIDUsingGET1OK describes a response with status code 200, with default header values.

'Success' with gerrit trigger
*/
type GetByIDUsingGET1OK struct {
	Payload models.GerritTrigger
}

func (o *GetByIDUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{id}][%d] getByIdUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetByIDUsingGET1OK) GetPayload() models.GerritTrigger {
	return o.Payload
}

func (o *GetByIDUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTrigger(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetByIDUsingGET1Unauthorized creates a GetByIDUsingGET1Unauthorized with default headers values
func NewGetByIDUsingGET1Unauthorized() *GetByIDUsingGET1Unauthorized {
	return &GetByIDUsingGET1Unauthorized{}
}

/* GetByIDUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetByIDUsingGET1Unauthorized struct {
}

func (o *GetByIDUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{id}][%d] getByIdUsingGET1Unauthorized ", 401)
}

func (o *GetByIDUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByIDUsingGET1Forbidden creates a GetByIDUsingGET1Forbidden with default headers values
func NewGetByIDUsingGET1Forbidden() *GetByIDUsingGET1Forbidden {
	return &GetByIDUsingGET1Forbidden{}
}

/* GetByIDUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetByIDUsingGET1Forbidden struct {
}

func (o *GetByIDUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{id}][%d] getByIdUsingGET1Forbidden ", 403)
}

func (o *GetByIDUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByIDUsingGET1NotFound creates a GetByIDUsingGET1NotFound with default headers values
func NewGetByIDUsingGET1NotFound() *GetByIDUsingGET1NotFound {
	return &GetByIDUsingGET1NotFound{}
}

/* GetByIDUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetByIDUsingGET1NotFound struct {
	Payload *models.Error
}

func (o *GetByIDUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{id}][%d] getByIdUsingGET1NotFound  %+v", 404, o.Payload)
}
func (o *GetByIDUsingGET1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetByIDUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByIDUsingGET1InternalServerError creates a GetByIDUsingGET1InternalServerError with default headers values
func NewGetByIDUsingGET1InternalServerError() *GetByIDUsingGET1InternalServerError {
	return &GetByIDUsingGET1InternalServerError{}
}

/* GetByIDUsingGET1InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetByIDUsingGET1InternalServerError struct {
}

func (o *GetByIDUsingGET1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/gerrit-triggers/{id}][%d] getByIdUsingGET1InternalServerError ", 500)
}

func (o *GetByIDUsingGET1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
