// Code generated by go-swagger; DO NOT EDIT.

package executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteAllUsingDELETEReader is a Reader for the DeleteAllUsingDELETE structure.
type DeleteAllUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAllUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewDeleteAllUsingDELETEAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAllUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteAllUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAllUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteAllUsingDELETEInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAllUsingDELETEAccepted creates a DeleteAllUsingDELETEAccepted with default headers values
func NewDeleteAllUsingDELETEAccepted() *DeleteAllUsingDELETEAccepted {
	return &DeleteAllUsingDELETEAccepted{}
}

/* DeleteAllUsingDELETEAccepted describes a response with status code 202, with default header values.

Count of executions to be deleted
*/
type DeleteAllUsingDELETEAccepted struct {
	Payload models.Executions
}

func (o *DeleteAllUsingDELETEAccepted) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions][%d] deleteAllUsingDELETEAccepted  %+v", 202, o.Payload)
}
func (o *DeleteAllUsingDELETEAccepted) GetPayload() models.Executions {
	return o.Payload
}

func (o *DeleteAllUsingDELETEAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutions(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteAllUsingDELETEUnauthorized creates a DeleteAllUsingDELETEUnauthorized with default headers values
func NewDeleteAllUsingDELETEUnauthorized() *DeleteAllUsingDELETEUnauthorized {
	return &DeleteAllUsingDELETEUnauthorized{}
}

/* DeleteAllUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type DeleteAllUsingDELETEUnauthorized struct {
}

func (o *DeleteAllUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions][%d] deleteAllUsingDELETEUnauthorized ", 401)
}

func (o *DeleteAllUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAllUsingDELETEForbidden creates a DeleteAllUsingDELETEForbidden with default headers values
func NewDeleteAllUsingDELETEForbidden() *DeleteAllUsingDELETEForbidden {
	return &DeleteAllUsingDELETEForbidden{}
}

/* DeleteAllUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteAllUsingDELETEForbidden struct {
}

func (o *DeleteAllUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions][%d] deleteAllUsingDELETEForbidden ", 403)
}

func (o *DeleteAllUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAllUsingDELETENotFound creates a DeleteAllUsingDELETENotFound with default headers values
func NewDeleteAllUsingDELETENotFound() *DeleteAllUsingDELETENotFound {
	return &DeleteAllUsingDELETENotFound{}
}

/* DeleteAllUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteAllUsingDELETENotFound struct {
	Payload *models.Error
}

func (o *DeleteAllUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions][%d] deleteAllUsingDELETENotFound  %+v", 404, o.Payload)
}
func (o *DeleteAllUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteAllUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllUsingDELETEInternalServerError creates a DeleteAllUsingDELETEInternalServerError with default headers values
func NewDeleteAllUsingDELETEInternalServerError() *DeleteAllUsingDELETEInternalServerError {
	return &DeleteAllUsingDELETEInternalServerError{}
}

/* DeleteAllUsingDELETEInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type DeleteAllUsingDELETEInternalServerError struct {
}

func (o *DeleteAllUsingDELETEInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/executions][%d] deleteAllUsingDELETEInternalServerError ", 500)
}

func (o *DeleteAllUsingDELETEInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
