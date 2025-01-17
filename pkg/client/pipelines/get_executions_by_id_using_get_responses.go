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

// GetExecutionsByIDUsingGETReader is a Reader for the GetExecutionsByIDUsingGET structure.
type GetExecutionsByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExecutionsByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExecutionsByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetExecutionsByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetExecutionsByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetExecutionsByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetExecutionsByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetExecutionsByIDUsingGETOK creates a GetExecutionsByIDUsingGETOK with default headers values
func NewGetExecutionsByIDUsingGETOK() *GetExecutionsByIDUsingGETOK {
	return &GetExecutionsByIDUsingGETOK{}
}

/* GetExecutionsByIDUsingGETOK describes a response with status code 200, with default header values.

'Success' with Executions on pages
*/
type GetExecutionsByIDUsingGETOK struct {
	Payload models.Executions
}

func (o *GetExecutionsByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions][%d] getExecutionsByIdUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetExecutionsByIDUsingGETOK) GetPayload() models.Executions {
	return o.Payload
}

func (o *GetExecutionsByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecutions(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetExecutionsByIDUsingGETUnauthorized creates a GetExecutionsByIDUsingGETUnauthorized with default headers values
func NewGetExecutionsByIDUsingGETUnauthorized() *GetExecutionsByIDUsingGETUnauthorized {
	return &GetExecutionsByIDUsingGETUnauthorized{}
}

/* GetExecutionsByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetExecutionsByIDUsingGETUnauthorized struct {
}

func (o *GetExecutionsByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions][%d] getExecutionsByIdUsingGETUnauthorized ", 401)
}

func (o *GetExecutionsByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExecutionsByIDUsingGETForbidden creates a GetExecutionsByIDUsingGETForbidden with default headers values
func NewGetExecutionsByIDUsingGETForbidden() *GetExecutionsByIDUsingGETForbidden {
	return &GetExecutionsByIDUsingGETForbidden{}
}

/* GetExecutionsByIDUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetExecutionsByIDUsingGETForbidden struct {
}

func (o *GetExecutionsByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions][%d] getExecutionsByIdUsingGETForbidden ", 403)
}

func (o *GetExecutionsByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetExecutionsByIDUsingGETNotFound creates a GetExecutionsByIDUsingGETNotFound with default headers values
func NewGetExecutionsByIDUsingGETNotFound() *GetExecutionsByIDUsingGETNotFound {
	return &GetExecutionsByIDUsingGETNotFound{}
}

/* GetExecutionsByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetExecutionsByIDUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetExecutionsByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions][%d] getExecutionsByIdUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetExecutionsByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetExecutionsByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExecutionsByIDUsingGETInternalServerError creates a GetExecutionsByIDUsingGETInternalServerError with default headers values
func NewGetExecutionsByIDUsingGETInternalServerError() *GetExecutionsByIDUsingGETInternalServerError {
	return &GetExecutionsByIDUsingGETInternalServerError{}
}

/* GetExecutionsByIDUsingGETInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetExecutionsByIDUsingGETInternalServerError struct {
}

func (o *GetExecutionsByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/pipelines/{id}/executions][%d] getExecutionsByIdUsingGETInternalServerError ", 500)
}

func (o *GetExecutionsByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
