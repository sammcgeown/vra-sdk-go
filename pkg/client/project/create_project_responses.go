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

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProjectCreated creates a CreateProjectCreated with default headers values
func NewCreateProjectCreated() *CreateProjectCreated {
	return &CreateProjectCreated{}
}

/* CreateProjectCreated describes a response with status code 201, with default header values.

successful operation
*/
type CreateProjectCreated struct {
	Payload *models.IaaSProject
}

func (o *CreateProjectCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/projects][%d] createProjectCreated  %+v", 201, o.Payload)
}
func (o *CreateProjectCreated) GetPayload() *models.IaaSProject {
	return o.Payload
}

func (o *CreateProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IaaSProject)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectBadRequest creates a CreateProjectBadRequest with default headers values
func NewCreateProjectBadRequest() *CreateProjectBadRequest {
	return &CreateProjectBadRequest{}
}

/* CreateProjectBadRequest describes a response with status code 400, with default header values.

Invalid Request - bad data
*/
type CreateProjectBadRequest struct {
	Payload *models.Error
}

func (o *CreateProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/projects][%d] createProjectBadRequest  %+v", 400, o.Payload)
}
func (o *CreateProjectBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectForbidden creates a CreateProjectForbidden with default headers values
func NewCreateProjectForbidden() *CreateProjectForbidden {
	return &CreateProjectForbidden{}
}

/* CreateProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateProjectForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *CreateProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/projects][%d] createProjectForbidden  %+v", 403, o.Payload)
}
func (o *CreateProjectForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *CreateProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
