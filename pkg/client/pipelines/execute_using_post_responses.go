// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// ExecuteUsingPOSTReader is a Reader for the ExecuteUsingPOST structure.
type ExecuteUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExecuteUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewExecuteUsingPOSTAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewExecuteUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewExecuteUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewExecuteUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewExecuteUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExecuteUsingPOSTOK creates a ExecuteUsingPOSTOK with default headers values
func NewExecuteUsingPOSTOK() *ExecuteUsingPOSTOK {
	return &ExecuteUsingPOSTOK{}
}

/* ExecuteUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with the created execution response
*/
type ExecuteUsingPOSTOK struct {
	Payload models.ExecutionResponse
}

func (o *ExecuteUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executeUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *ExecuteUsingPOSTOK) GetPayload() models.ExecutionResponse {
	return o.Payload
}

func (o *ExecuteUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutionResponse(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewExecuteUsingPOSTAccepted creates a ExecuteUsingPOSTAccepted with default headers values
func NewExecuteUsingPOSTAccepted() *ExecuteUsingPOSTAccepted {
	return &ExecuteUsingPOSTAccepted{}
}

/* ExecuteUsingPOSTAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ExecuteUsingPOSTAccepted struct {
	Payload models.ExecutionResponse
}

func (o *ExecuteUsingPOSTAccepted) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executeUsingPOSTAccepted  %+v", 202, o.Payload)
}
func (o *ExecuteUsingPOSTAccepted) GetPayload() models.ExecutionResponse {
	return o.Payload
}

func (o *ExecuteUsingPOSTAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutionResponse(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewExecuteUsingPOSTUnauthorized creates a ExecuteUsingPOSTUnauthorized with default headers values
func NewExecuteUsingPOSTUnauthorized() *ExecuteUsingPOSTUnauthorized {
	return &ExecuteUsingPOSTUnauthorized{}
}

/* ExecuteUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type ExecuteUsingPOSTUnauthorized struct {
}

func (o *ExecuteUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executeUsingPOSTUnauthorized ", 401)
}

func (o *ExecuteUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecuteUsingPOSTForbidden creates a ExecuteUsingPOSTForbidden with default headers values
func NewExecuteUsingPOSTForbidden() *ExecuteUsingPOSTForbidden {
	return &ExecuteUsingPOSTForbidden{}
}

/* ExecuteUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ExecuteUsingPOSTForbidden struct {
}

func (o *ExecuteUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executeUsingPOSTForbidden ", 403)
}

func (o *ExecuteUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExecuteUsingPOSTNotFound creates a ExecuteUsingPOSTNotFound with default headers values
func NewExecuteUsingPOSTNotFound() *ExecuteUsingPOSTNotFound {
	return &ExecuteUsingPOSTNotFound{}
}

/* ExecuteUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ExecuteUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *ExecuteUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executeUsingPOSTNotFound  %+v", 404, o.Payload)
}
func (o *ExecuteUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ExecuteUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteUsingPOSTInternalServerError creates a ExecuteUsingPOSTInternalServerError with default headers values
func NewExecuteUsingPOSTInternalServerError() *ExecuteUsingPOSTInternalServerError {
	return &ExecuteUsingPOSTInternalServerError{}
}

/* ExecuteUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type ExecuteUsingPOSTInternalServerError struct {
}

func (o *ExecuteUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/pipelines/{id}/executions][%d] executeUsingPOSTInternalServerError ", 500)
}

func (o *ExecuteUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}