// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateUsingPATCHReader is a Reader for the UpdateUsingPATCH structure.
type UpdateUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUsingPATCHBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUsingPATCHOK creates a UpdateUsingPATCHOK with default headers values
func NewUpdateUsingPATCHOK() *UpdateUsingPATCHOK {
	return &UpdateUsingPATCHOK{}
}

/* UpdateUsingPATCHOK describes a response with status code 200, with default header values.

'OK' with the newly updated project resource metadata
*/
type UpdateUsingPATCHOK struct {
	Payload models.ProjectResourceMetadata
}

func (o *UpdateUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /project-service/api/projects/{id}/resource-metadata][%d] updateUsingPATCHOK  %+v", 200, o.Payload)
}
func (o *UpdateUsingPATCHOK) GetPayload() models.ProjectResourceMetadata {
	return o.Payload
}

func (o *UpdateUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalProjectResourceMetadata(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateUsingPATCHBadRequest creates a UpdateUsingPATCHBadRequest with default headers values
func NewUpdateUsingPATCHBadRequest() *UpdateUsingPATCHBadRequest {
	return &UpdateUsingPATCHBadRequest{}
}

/* UpdateUsingPATCHBadRequest describes a response with status code 400, with default header values.

'Bad request' on validation failure
*/
type UpdateUsingPATCHBadRequest struct {
	Payload *models.Error
}

func (o *UpdateUsingPATCHBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /project-service/api/projects/{id}/resource-metadata][%d] updateUsingPATCHBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateUsingPATCHBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUsingPATCHBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUsingPATCHNotFound creates a UpdateUsingPATCHNotFound with default headers values
func NewUpdateUsingPATCHNotFound() *UpdateUsingPATCHNotFound {
	return &UpdateUsingPATCHNotFound{}
}

/* UpdateUsingPATCHNotFound describes a response with status code 404, with default header values.

'Not found' if no project with the provided id
*/
type UpdateUsingPATCHNotFound struct {
	Payload *models.Error
}

func (o *UpdateUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /project-service/api/projects/{id}/resource-metadata][%d] updateUsingPATCHNotFound  %+v", 404, o.Payload)
}
func (o *UpdateUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
