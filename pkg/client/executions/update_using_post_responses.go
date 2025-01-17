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

// UpdateUsingPOSTReader is a Reader for the UpdateUsingPOST structure.
type UpdateUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUsingPOSTOK creates a UpdateUsingPOSTOK with default headers values
func NewUpdateUsingPOSTOK() *UpdateUsingPOSTOK {
	return &UpdateUsingPOSTOK{}
}

/* UpdateUsingPOSTOK describes a response with status code 200, with default header values.

Approved/Rejected all selected user operations successfully.
*/
type UpdateUsingPOSTOK struct {
	Payload models.BatchUserOperationResponse
}

func (o *UpdateUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/batch-user-operations][%d] updateUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *UpdateUsingPOSTOK) GetPayload() models.BatchUserOperationResponse {
	return o.Payload
}

func (o *UpdateUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalBatchUserOperationResponse(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateUsingPOSTUnauthorized creates a UpdateUsingPOSTUnauthorized with default headers values
func NewUpdateUsingPOSTUnauthorized() *UpdateUsingPOSTUnauthorized {
	return &UpdateUsingPOSTUnauthorized{}
}

/* UpdateUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type UpdateUsingPOSTUnauthorized struct {
}

func (o *UpdateUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/batch-user-operations][%d] updateUsingPOSTUnauthorized ", 401)
}

func (o *UpdateUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUsingPOSTForbidden creates a UpdateUsingPOSTForbidden with default headers values
func NewUpdateUsingPOSTForbidden() *UpdateUsingPOSTForbidden {
	return &UpdateUsingPOSTForbidden{}
}

/* UpdateUsingPOSTForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateUsingPOSTForbidden struct {
}

func (o *UpdateUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/batch-user-operations][%d] updateUsingPOSTForbidden ", 403)
}

func (o *UpdateUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUsingPOSTNotFound creates a UpdateUsingPOSTNotFound with default headers values
func NewUpdateUsingPOSTNotFound() *UpdateUsingPOSTNotFound {
	return &UpdateUsingPOSTNotFound{}
}

/* UpdateUsingPOSTNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *UpdateUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/batch-user-operations][%d] updateUsingPOSTNotFound  %+v", 404, o.Payload)
}
func (o *UpdateUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUsingPOSTInternalServerError creates a UpdateUsingPOSTInternalServerError with default headers values
func NewUpdateUsingPOSTInternalServerError() *UpdateUsingPOSTInternalServerError {
	return &UpdateUsingPOSTInternalServerError{}
}

/* UpdateUsingPOSTInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type UpdateUsingPOSTInternalServerError struct {
}

func (o *UpdateUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/batch-user-operations][%d] updateUsingPOSTInternalServerError ", 500)
}

func (o *UpdateUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
