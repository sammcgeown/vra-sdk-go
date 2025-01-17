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

// DeleteByNameUsingDELETE1Reader is a Reader for the DeleteByNameUsingDELETE1 structure.
type DeleteByNameUsingDELETE1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteByNameUsingDELETE1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteByNameUsingDELETE1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteByNameUsingDELETE1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteByNameUsingDELETE1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteByNameUsingDELETE1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteByNameUsingDELETE1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteByNameUsingDELETE1OK creates a DeleteByNameUsingDELETE1OK with default headers values
func NewDeleteByNameUsingDELETE1OK() *DeleteByNameUsingDELETE1OK {
	return &DeleteByNameUsingDELETE1OK{}
}

/* DeleteByNameUsingDELETE1OK describes a response with status code 200, with default header values.

'Success' with Gerrit Listener Delete
*/
type DeleteByNameUsingDELETE1OK struct {
	Payload models.GerritListener
}

func (o *DeleteByNameUsingDELETE1OK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteByNameUsingDELETE1OK  %+v", 200, o.Payload)
}
func (o *DeleteByNameUsingDELETE1OK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *DeleteByNameUsingDELETE1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteByNameUsingDELETE1Unauthorized creates a DeleteByNameUsingDELETE1Unauthorized with default headers values
func NewDeleteByNameUsingDELETE1Unauthorized() *DeleteByNameUsingDELETE1Unauthorized {
	return &DeleteByNameUsingDELETE1Unauthorized{}
}

/* DeleteByNameUsingDELETE1Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteByNameUsingDELETE1Unauthorized struct {
}

func (o *DeleteByNameUsingDELETE1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteByNameUsingDELETE1Unauthorized ", 401)
}

func (o *DeleteByNameUsingDELETE1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteByNameUsingDELETE1Forbidden creates a DeleteByNameUsingDELETE1Forbidden with default headers values
func NewDeleteByNameUsingDELETE1Forbidden() *DeleteByNameUsingDELETE1Forbidden {
	return &DeleteByNameUsingDELETE1Forbidden{}
}

/* DeleteByNameUsingDELETE1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteByNameUsingDELETE1Forbidden struct {
}

func (o *DeleteByNameUsingDELETE1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteByNameUsingDELETE1Forbidden ", 403)
}

func (o *DeleteByNameUsingDELETE1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteByNameUsingDELETE1NotFound creates a DeleteByNameUsingDELETE1NotFound with default headers values
func NewDeleteByNameUsingDELETE1NotFound() *DeleteByNameUsingDELETE1NotFound {
	return &DeleteByNameUsingDELETE1NotFound{}
}

/* DeleteByNameUsingDELETE1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteByNameUsingDELETE1NotFound struct {
	Payload *models.Error
}

func (o *DeleteByNameUsingDELETE1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteByNameUsingDELETE1NotFound  %+v", 404, o.Payload)
}
func (o *DeleteByNameUsingDELETE1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteByNameUsingDELETE1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteByNameUsingDELETE1InternalServerError creates a DeleteByNameUsingDELETE1InternalServerError with default headers values
func NewDeleteByNameUsingDELETE1InternalServerError() *DeleteByNameUsingDELETE1InternalServerError {
	return &DeleteByNameUsingDELETE1InternalServerError{}
}

/* DeleteByNameUsingDELETE1InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteByNameUsingDELETE1InternalServerError struct {
}

func (o *DeleteByNameUsingDELETE1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{project}/{name}][%d] deleteByNameUsingDELETE1InternalServerError ", 500)
}

func (o *DeleteByNameUsingDELETE1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
