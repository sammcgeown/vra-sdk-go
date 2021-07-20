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

// CreateUsingPOST5Reader is a Reader for the CreateUsingPOST5 structure.
type CreateUsingPOST5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUsingPOST5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUsingPOST5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateUsingPOST5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUsingPOST5Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUsingPOST5NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUsingPOST5InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUsingPOST5OK creates a CreateUsingPOST5OK with default headers values
func NewCreateUsingPOST5OK() *CreateUsingPOST5OK {
	return &CreateUsingPOST5OK{}
}

/* CreateUsingPOST5OK describes a response with status code 200, with default header values.

'Success' with Git Webhook Creation
*/
type CreateUsingPOST5OK struct {
	Payload models.GitWebhook
}

func (o *CreateUsingPOST5OK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createUsingPOST5OK  %+v", 200, o.Payload)
}
func (o *CreateUsingPOST5OK) GetPayload() models.GitWebhook {
	return o.Payload
}

func (o *CreateUsingPOST5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGitWebhook(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCreateUsingPOST5Unauthorized creates a CreateUsingPOST5Unauthorized with default headers values
func NewCreateUsingPOST5Unauthorized() *CreateUsingPOST5Unauthorized {
	return &CreateUsingPOST5Unauthorized{}
}

/* CreateUsingPOST5Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CreateUsingPOST5Unauthorized struct {
}

func (o *CreateUsingPOST5Unauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createUsingPOST5Unauthorized ", 401)
}

func (o *CreateUsingPOST5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOST5Forbidden creates a CreateUsingPOST5Forbidden with default headers values
func NewCreateUsingPOST5Forbidden() *CreateUsingPOST5Forbidden {
	return &CreateUsingPOST5Forbidden{}
}

/* CreateUsingPOST5Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateUsingPOST5Forbidden struct {
}

func (o *CreateUsingPOST5Forbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createUsingPOST5Forbidden ", 403)
}

func (o *CreateUsingPOST5Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUsingPOST5NotFound creates a CreateUsingPOST5NotFound with default headers values
func NewCreateUsingPOST5NotFound() *CreateUsingPOST5NotFound {
	return &CreateUsingPOST5NotFound{}
}

/* CreateUsingPOST5NotFound describes a response with status code 404, with default header values.

Not Found
*/
type CreateUsingPOST5NotFound struct {
	Payload *models.Error
}

func (o *CreateUsingPOST5NotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createUsingPOST5NotFound  %+v", 404, o.Payload)
}
func (o *CreateUsingPOST5NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUsingPOST5NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUsingPOST5InternalServerError creates a CreateUsingPOST5InternalServerError with default headers values
func NewCreateUsingPOST5InternalServerError() *CreateUsingPOST5InternalServerError {
	return &CreateUsingPOST5InternalServerError{}
}

/* CreateUsingPOST5InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CreateUsingPOST5InternalServerError struct {
}

func (o *CreateUsingPOST5InternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/git-webhooks][%d] createUsingPOST5InternalServerError ", 500)
}

func (o *CreateUsingPOST5InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}