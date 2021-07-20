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

// UpdateUsingPUT5Reader is a Reader for the UpdateUsingPUT5 structure.
type UpdateUsingPUT5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUsingPUT5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUsingPUT5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUsingPUT5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUsingPUT5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUsingPUT5NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUsingPUT5InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUsingPUT5OK creates a UpdateUsingPUT5OK with default headers values
func NewUpdateUsingPUT5OK() *UpdateUsingPUT5OK {
	return &UpdateUsingPUT5OK{}
}

/* UpdateUsingPUT5OK describes a response with status code 200, with default header values.

'Success' with updated Variable
*/
type UpdateUsingPUT5OK struct {
	Payload *models.Variable
}

func (o *UpdateUsingPUT5OK) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateUsingPUT5OK  %+v", 200, o.Payload)
}
func (o *UpdateUsingPUT5OK) GetPayload() *models.Variable {
	return o.Payload
}

func (o *UpdateUsingPUT5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Variable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUsingPUT5Unauthorized creates a UpdateUsingPUT5Unauthorized with default headers values
func NewUpdateUsingPUT5Unauthorized() *UpdateUsingPUT5Unauthorized {
	return &UpdateUsingPUT5Unauthorized{}
}

/* UpdateUsingPUT5Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type UpdateUsingPUT5Unauthorized struct {
}

func (o *UpdateUsingPUT5Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateUsingPUT5Unauthorized ", 401)
}

func (o *UpdateUsingPUT5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUsingPUT5Forbidden creates a UpdateUsingPUT5Forbidden with default headers values
func NewUpdateUsingPUT5Forbidden() *UpdateUsingPUT5Forbidden {
	return &UpdateUsingPUT5Forbidden{}
}

/* UpdateUsingPUT5Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateUsingPUT5Forbidden struct {
}

func (o *UpdateUsingPUT5Forbidden) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateUsingPUT5Forbidden ", 403)
}

func (o *UpdateUsingPUT5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUsingPUT5NotFound creates a UpdateUsingPUT5NotFound with default headers values
func NewUpdateUsingPUT5NotFound() *UpdateUsingPUT5NotFound {
	return &UpdateUsingPUT5NotFound{}
}

/* UpdateUsingPUT5NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateUsingPUT5NotFound struct {
	Payload *models.Error
}

func (o *UpdateUsingPUT5NotFound) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateUsingPUT5NotFound  %+v", 404, o.Payload)
}
func (o *UpdateUsingPUT5NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUsingPUT5NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUsingPUT5InternalServerError creates a UpdateUsingPUT5InternalServerError with default headers values
func NewUpdateUsingPUT5InternalServerError() *UpdateUsingPUT5InternalServerError {
	return &UpdateUsingPUT5InternalServerError{}
}

/* UpdateUsingPUT5InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateUsingPUT5InternalServerError struct {
}

func (o *UpdateUsingPUT5InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{project}/{name}][%d] updateUsingPUT5InternalServerError ", 500)
}

func (o *UpdateUsingPUT5InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}