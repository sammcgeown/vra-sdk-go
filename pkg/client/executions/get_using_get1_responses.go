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

// GetUsingGET1Reader is a Reader for the GetUsingGET1 structure.
type GetUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsingGET1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsingGET1OK creates a GetUsingGET1OK with default headers values
func NewGetUsingGET1OK() *GetUsingGET1OK {
	return &GetUsingGET1OK{}
}

/* GetUsingGET1OK describes a response with status code 200, with default header values.

'Success' with Execution
*/
type GetUsingGET1OK struct {
	Payload models.Execution
}

func (o *GetUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions/{id}][%d] getUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetUsingGET1OK) GetPayload() models.Execution {
	return o.Payload
}

func (o *GetUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalExecution(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetUsingGET1Unauthorized creates a GetUsingGET1Unauthorized with default headers values
func NewGetUsingGET1Unauthorized() *GetUsingGET1Unauthorized {
	return &GetUsingGET1Unauthorized{}
}

/* GetUsingGET1Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetUsingGET1Unauthorized struct {
}

func (o *GetUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions/{id}][%d] getUsingGET1Unauthorized ", 401)
}

func (o *GetUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET1Forbidden creates a GetUsingGET1Forbidden with default headers values
func NewGetUsingGET1Forbidden() *GetUsingGET1Forbidden {
	return &GetUsingGET1Forbidden{}
}

/* GetUsingGET1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetUsingGET1Forbidden struct {
}

func (o *GetUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions/{id}][%d] getUsingGET1Forbidden ", 403)
}

func (o *GetUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET1NotFound creates a GetUsingGET1NotFound with default headers values
func NewGetUsingGET1NotFound() *GetUsingGET1NotFound {
	return &GetUsingGET1NotFound{}
}

/* GetUsingGET1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetUsingGET1NotFound struct {
	Payload *models.Error
}

func (o *GetUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions/{id}][%d] getUsingGET1NotFound  %+v", 404, o.Payload)
}
func (o *GetUsingGET1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGET1InternalServerError creates a GetUsingGET1InternalServerError with default headers values
func NewGetUsingGET1InternalServerError() *GetUsingGET1InternalServerError {
	return &GetUsingGET1InternalServerError{}
}

/* GetUsingGET1InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetUsingGET1InternalServerError struct {
}

func (o *GetUsingGET1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/executions/{id}][%d] getUsingGET1InternalServerError ", 500)
}

func (o *GetUsingGET1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}