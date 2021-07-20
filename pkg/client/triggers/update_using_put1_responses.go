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

// UpdateUsingPUT1Reader is a Reader for the UpdateUsingPUT1 structure.
type UpdateUsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUsingPUT1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUsingPUT1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUsingPUT1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUsingPUT1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUsingPUT1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUsingPUT1OK creates a UpdateUsingPUT1OK with default headers values
func NewUpdateUsingPUT1OK() *UpdateUsingPUT1OK {
	return &UpdateUsingPUT1OK{}
}

/* UpdateUsingPUT1OK describes a response with status code 200, with default header values.

'Success' with Git Webhook Update
*/
type UpdateUsingPUT1OK struct {
	Payload models.GitWebhook
}

func (o *UpdateUsingPUT1OK) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/git-webhooks/{id}][%d] updateUsingPUT1OK  %+v", 200, o.Payload)
}
func (o *UpdateUsingPUT1OK) GetPayload() models.GitWebhook {
	return o.Payload
}

func (o *UpdateUsingPUT1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGitWebhook(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateUsingPUT1Unauthorized creates a UpdateUsingPUT1Unauthorized with default headers values
func NewUpdateUsingPUT1Unauthorized() *UpdateUsingPUT1Unauthorized {
	return &UpdateUsingPUT1Unauthorized{}
}

/* UpdateUsingPUT1Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type UpdateUsingPUT1Unauthorized struct {
}

func (o *UpdateUsingPUT1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/git-webhooks/{id}][%d] updateUsingPUT1Unauthorized ", 401)
}

func (o *UpdateUsingPUT1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUsingPUT1Forbidden creates a UpdateUsingPUT1Forbidden with default headers values
func NewUpdateUsingPUT1Forbidden() *UpdateUsingPUT1Forbidden {
	return &UpdateUsingPUT1Forbidden{}
}

/* UpdateUsingPUT1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateUsingPUT1Forbidden struct {
}

func (o *UpdateUsingPUT1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/git-webhooks/{id}][%d] updateUsingPUT1Forbidden ", 403)
}

func (o *UpdateUsingPUT1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUsingPUT1NotFound creates a UpdateUsingPUT1NotFound with default headers values
func NewUpdateUsingPUT1NotFound() *UpdateUsingPUT1NotFound {
	return &UpdateUsingPUT1NotFound{}
}

/* UpdateUsingPUT1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateUsingPUT1NotFound struct {
	Payload *models.Error
}

func (o *UpdateUsingPUT1NotFound) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/git-webhooks/{id}][%d] updateUsingPUT1NotFound  %+v", 404, o.Payload)
}
func (o *UpdateUsingPUT1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUsingPUT1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUsingPUT1InternalServerError creates a UpdateUsingPUT1InternalServerError with default headers values
func NewUpdateUsingPUT1InternalServerError() *UpdateUsingPUT1InternalServerError {
	return &UpdateUsingPUT1InternalServerError{}
}

/* UpdateUsingPUT1InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateUsingPUT1InternalServerError struct {
}

func (o *UpdateUsingPUT1InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/git-webhooks/{id}][%d] updateUsingPUT1InternalServerError ", 500)
}

func (o *UpdateUsingPUT1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}