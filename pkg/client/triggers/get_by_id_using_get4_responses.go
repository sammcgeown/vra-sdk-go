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

// GetByIDUsingGET4Reader is a Reader for the GetByIDUsingGET4 structure.
type GetByIDUsingGET4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetByIDUsingGET4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetByIDUsingGET4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetByIDUsingGET4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetByIDUsingGET4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetByIDUsingGET4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetByIDUsingGET4InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetByIDUsingGET4OK creates a GetByIDUsingGET4OK with default headers values
func NewGetByIDUsingGET4OK() *GetByIDUsingGET4OK {
	return &GetByIDUsingGET4OK{}
}

/* GetByIDUsingGET4OK describes a response with status code 200, with default header values.

'Success' with Git Event
*/
type GetByIDUsingGET4OK struct {
	Payload models.GitEvent
}

func (o *GetByIDUsingGET4OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getByIdUsingGET4OK  %+v", 200, o.Payload)
}
func (o *GetByIDUsingGET4OK) GetPayload() models.GitEvent {
	return o.Payload
}

func (o *GetByIDUsingGET4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGitEvent(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetByIDUsingGET4Unauthorized creates a GetByIDUsingGET4Unauthorized with default headers values
func NewGetByIDUsingGET4Unauthorized() *GetByIDUsingGET4Unauthorized {
	return &GetByIDUsingGET4Unauthorized{}
}

/* GetByIDUsingGET4Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type GetByIDUsingGET4Unauthorized struct {
}

func (o *GetByIDUsingGET4Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getByIdUsingGET4Unauthorized ", 401)
}

func (o *GetByIDUsingGET4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByIDUsingGET4Forbidden creates a GetByIDUsingGET4Forbidden with default headers values
func NewGetByIDUsingGET4Forbidden() *GetByIDUsingGET4Forbidden {
	return &GetByIDUsingGET4Forbidden{}
}

/* GetByIDUsingGET4Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetByIDUsingGET4Forbidden struct {
}

func (o *GetByIDUsingGET4Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getByIdUsingGET4Forbidden ", 403)
}

func (o *GetByIDUsingGET4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByIDUsingGET4NotFound creates a GetByIDUsingGET4NotFound with default headers values
func NewGetByIDUsingGET4NotFound() *GetByIDUsingGET4NotFound {
	return &GetByIDUsingGET4NotFound{}
}

/* GetByIDUsingGET4NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetByIDUsingGET4NotFound struct {
	Payload *models.Error
}

func (o *GetByIDUsingGET4NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getByIdUsingGET4NotFound  %+v", 404, o.Payload)
}
func (o *GetByIDUsingGET4NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetByIDUsingGET4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetByIDUsingGET4InternalServerError creates a GetByIDUsingGET4InternalServerError with default headers values
func NewGetByIDUsingGET4InternalServerError() *GetByIDUsingGET4InternalServerError {
	return &GetByIDUsingGET4InternalServerError{}
}

/* GetByIDUsingGET4InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type GetByIDUsingGET4InternalServerError struct {
}

func (o *GetByIDUsingGET4InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/git-events/{id}][%d] getByIdUsingGET4InternalServerError ", 500)
}

func (o *GetByIDUsingGET4InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
