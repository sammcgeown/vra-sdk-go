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

// DeleteByNameUsingDELETE2Reader is a Reader for the DeleteByNameUsingDELETE2 structure.
type DeleteByNameUsingDELETE2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteByNameUsingDELETE2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteByNameUsingDELETE2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteByNameUsingDELETE2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteByNameUsingDELETE2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteByNameUsingDELETE2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteByNameUsingDELETE2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteByNameUsingDELETE2OK creates a DeleteByNameUsingDELETE2OK with default headers values
func NewDeleteByNameUsingDELETE2OK() *DeleteByNameUsingDELETE2OK {
	return &DeleteByNameUsingDELETE2OK{}
}

/* DeleteByNameUsingDELETE2OK describes a response with status code 200, with default header values.

'Success' with Gerrit Trigger Delete
*/
type DeleteByNameUsingDELETE2OK struct {
	Payload models.GerritTrigger
}

func (o *DeleteByNameUsingDELETE2OK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-triggers/{project}/{name}][%d] deleteByNameUsingDELETE2OK  %+v", 200, o.Payload)
}
func (o *DeleteByNameUsingDELETE2OK) GetPayload() models.GerritTrigger {
	return o.Payload
}

func (o *DeleteByNameUsingDELETE2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTrigger(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteByNameUsingDELETE2Unauthorized creates a DeleteByNameUsingDELETE2Unauthorized with default headers values
func NewDeleteByNameUsingDELETE2Unauthorized() *DeleteByNameUsingDELETE2Unauthorized {
	return &DeleteByNameUsingDELETE2Unauthorized{}
}

/* DeleteByNameUsingDELETE2Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteByNameUsingDELETE2Unauthorized struct {
}

func (o *DeleteByNameUsingDELETE2Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-triggers/{project}/{name}][%d] deleteByNameUsingDELETE2Unauthorized ", 401)
}

func (o *DeleteByNameUsingDELETE2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteByNameUsingDELETE2Forbidden creates a DeleteByNameUsingDELETE2Forbidden with default headers values
func NewDeleteByNameUsingDELETE2Forbidden() *DeleteByNameUsingDELETE2Forbidden {
	return &DeleteByNameUsingDELETE2Forbidden{}
}

/* DeleteByNameUsingDELETE2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteByNameUsingDELETE2Forbidden struct {
}

func (o *DeleteByNameUsingDELETE2Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-triggers/{project}/{name}][%d] deleteByNameUsingDELETE2Forbidden ", 403)
}

func (o *DeleteByNameUsingDELETE2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteByNameUsingDELETE2NotFound creates a DeleteByNameUsingDELETE2NotFound with default headers values
func NewDeleteByNameUsingDELETE2NotFound() *DeleteByNameUsingDELETE2NotFound {
	return &DeleteByNameUsingDELETE2NotFound{}
}

/* DeleteByNameUsingDELETE2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteByNameUsingDELETE2NotFound struct {
	Payload *models.Error
}

func (o *DeleteByNameUsingDELETE2NotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-triggers/{project}/{name}][%d] deleteByNameUsingDELETE2NotFound  %+v", 404, o.Payload)
}
func (o *DeleteByNameUsingDELETE2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteByNameUsingDELETE2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteByNameUsingDELETE2InternalServerError creates a DeleteByNameUsingDELETE2InternalServerError with default headers values
func NewDeleteByNameUsingDELETE2InternalServerError() *DeleteByNameUsingDELETE2InternalServerError {
	return &DeleteByNameUsingDELETE2InternalServerError{}
}

/* DeleteByNameUsingDELETE2InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteByNameUsingDELETE2InternalServerError struct {
}

func (o *DeleteByNameUsingDELETE2InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-triggers/{project}/{name}][%d] deleteByNameUsingDELETE2InternalServerError ", 500)
}

func (o *DeleteByNameUsingDELETE2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
