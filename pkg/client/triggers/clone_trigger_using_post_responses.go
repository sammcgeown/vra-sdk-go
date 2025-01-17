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

// CloneTriggerUsingPOSTReader is a Reader for the CloneTriggerUsingPOST structure.
type CloneTriggerUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneTriggerUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneTriggerUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCloneTriggerUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCloneTriggerUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloneTriggerUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCloneTriggerUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneTriggerUsingPOSTOK creates a CloneTriggerUsingPOSTOK with default headers values
func NewCloneTriggerUsingPOSTOK() *CloneTriggerUsingPOSTOK {
	return &CloneTriggerUsingPOSTOK{}
}

/* CloneTriggerUsingPOSTOK describes a response with status code 200, with default header values.

'Success' with Gerrit Trigger Clone
*/
type CloneTriggerUsingPOSTOK struct {
	Payload models.GerritTrigger
}

func (o *CloneTriggerUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{id}][%d] cloneTriggerUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *CloneTriggerUsingPOSTOK) GetPayload() models.GerritTrigger {
	return o.Payload
}

func (o *CloneTriggerUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritTrigger(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCloneTriggerUsingPOSTUnauthorized creates a CloneTriggerUsingPOSTUnauthorized with default headers values
func NewCloneTriggerUsingPOSTUnauthorized() *CloneTriggerUsingPOSTUnauthorized {
	return &CloneTriggerUsingPOSTUnauthorized{}
}

/* CloneTriggerUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type CloneTriggerUsingPOSTUnauthorized struct {
}

func (o *CloneTriggerUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{id}][%d] cloneTriggerUsingPOSTUnauthorized ", 401)
}

func (o *CloneTriggerUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneTriggerUsingPOSTForbidden creates a CloneTriggerUsingPOSTForbidden with default headers values
func NewCloneTriggerUsingPOSTForbidden() *CloneTriggerUsingPOSTForbidden {
	return &CloneTriggerUsingPOSTForbidden{}
}

/* CloneTriggerUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CloneTriggerUsingPOSTForbidden struct {
}

func (o *CloneTriggerUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{id}][%d] cloneTriggerUsingPOSTForbidden ", 403)
}

func (o *CloneTriggerUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneTriggerUsingPOSTNotFound creates a CloneTriggerUsingPOSTNotFound with default headers values
func NewCloneTriggerUsingPOSTNotFound() *CloneTriggerUsingPOSTNotFound {
	return &CloneTriggerUsingPOSTNotFound{}
}

/* CloneTriggerUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CloneTriggerUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *CloneTriggerUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{id}][%d] cloneTriggerUsingPOSTNotFound  %+v", 404, o.Payload)
}
func (o *CloneTriggerUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CloneTriggerUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneTriggerUsingPOSTInternalServerError creates a CloneTriggerUsingPOSTInternalServerError with default headers values
func NewCloneTriggerUsingPOSTInternalServerError() *CloneTriggerUsingPOSTInternalServerError {
	return &CloneTriggerUsingPOSTInternalServerError{}
}

/* CloneTriggerUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type CloneTriggerUsingPOSTInternalServerError struct {
}

func (o *CloneTriggerUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-triggers/{id}][%d] cloneTriggerUsingPOSTInternalServerError ", 500)
}

func (o *CloneTriggerUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
