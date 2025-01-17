// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateUsingPOST7Reader is a Reader for the CreateUsingPOST7 structure.
type CreateUsingPOST7Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUsingPOST7Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUsingPOST7OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateUsingPOST7Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUsingPOST7Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUsingPOST7NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUsingPOST7InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUsingPOST7OK creates a CreateUsingPOST7OK with default headers values
func NewCreateUsingPOST7OK() *CreateUsingPOST7OK {
	return &CreateUsingPOST7OK{}
}

/* CreateUsingPOST7OK describes a response with status code 200, with default header values.

'Success' with Variable creation
*/
type CreateUsingPOST7OK struct {
	Payload *models.Variable
}

func (o *CreateUsingPOST7OK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/variables][%d] createUsingPOST7OK  %+v", 200, o.Payload)
}
func (o *CreateUsingPOST7OK) GetPayload() *models.Variable {
	return o.Payload
}

func (o *CreateUsingPOST7OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Variable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUsingPOST7Unauthorized creates a CreateUsingPOST7Unauthorized with default headers values
func NewCreateUsingPOST7Unauthorized() *CreateUsingPOST7Unauthorized {
	return &CreateUsingPOST7Unauthorized{}
}

/* CreateUsingPOST7Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CreateUsingPOST7Unauthorized struct {
}

func (o *CreateUsingPOST7Unauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/variables][%d] createUsingPOST7Unauthorized ", 401)
}

func (o *CreateUsingPOST7Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOST7Forbidden creates a CreateUsingPOST7Forbidden with default headers values
func NewCreateUsingPOST7Forbidden() *CreateUsingPOST7Forbidden {
	return &CreateUsingPOST7Forbidden{}
}

/* CreateUsingPOST7Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateUsingPOST7Forbidden struct {
}

func (o *CreateUsingPOST7Forbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/variables][%d] createUsingPOST7Forbidden ", 403)
}

func (o *CreateUsingPOST7Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOST7NotFound creates a CreateUsingPOST7NotFound with default headers values
func NewCreateUsingPOST7NotFound() *CreateUsingPOST7NotFound {
	return &CreateUsingPOST7NotFound{}
}

/* CreateUsingPOST7NotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateUsingPOST7NotFound struct {
	Payload *models.Error
}

func (o *CreateUsingPOST7NotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/variables][%d] createUsingPOST7NotFound  %+v", 404, o.Payload)
}
func (o *CreateUsingPOST7NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUsingPOST7NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUsingPOST7InternalServerError creates a CreateUsingPOST7InternalServerError with default headers values
func NewCreateUsingPOST7InternalServerError() *CreateUsingPOST7InternalServerError {
	return &CreateUsingPOST7InternalServerError{}
}

/* CreateUsingPOST7InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CreateUsingPOST7InternalServerError struct {
}

func (o *CreateUsingPOST7InternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/variables][%d] createUsingPOST7InternalServerError ", 500)
}

func (o *CreateUsingPOST7InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
