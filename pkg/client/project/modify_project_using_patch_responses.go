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

// ModifyProjectUsingPATCHReader is a Reader for the ModifyProjectUsingPATCH structure.
type ModifyProjectUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyProjectUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModifyProjectUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewModifyProjectUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewModifyProjectUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewModifyProjectUsingPATCHOK creates a ModifyProjectUsingPATCHOK with default headers values
func NewModifyProjectUsingPATCHOK() *ModifyProjectUsingPATCHOK {
	return &ModifyProjectUsingPATCHOK{}
}

/* ModifyProjectUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with the Project
*/
type ModifyProjectUsingPATCHOK struct {
	Payload models.Project
}

func (o *ModifyProjectUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /project-service/api/projects/{id}][%d] modifyProjectUsingPATCHOK  %+v", 200, o.Payload)
}
func (o *ModifyProjectUsingPATCHOK) GetPayload() models.Project {
	return o.Payload
}

func (o *ModifyProjectUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalProject(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewModifyProjectUsingPATCHForbidden creates a ModifyProjectUsingPATCHForbidden with default headers values
func NewModifyProjectUsingPATCHForbidden() *ModifyProjectUsingPATCHForbidden {
	return &ModifyProjectUsingPATCHForbidden{}
}

/* ModifyProjectUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden, the user lacks permissions
*/
type ModifyProjectUsingPATCHForbidden struct {
}

func (o *ModifyProjectUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /project-service/api/projects/{id}][%d] modifyProjectUsingPATCHForbidden ", 403)
}

func (o *ModifyProjectUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyProjectUsingPATCHNotFound creates a ModifyProjectUsingPATCHNotFound with default headers values
func NewModifyProjectUsingPATCHNotFound() *ModifyProjectUsingPATCHNotFound {
	return &ModifyProjectUsingPATCHNotFound{}
}

/* ModifyProjectUsingPATCHNotFound describes a response with status code 404, with default header values.

'Not found' if no project with the provided id
*/
type ModifyProjectUsingPATCHNotFound struct {
	Payload *models.Error
}

func (o *ModifyProjectUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /project-service/api/projects/{id}][%d] modifyProjectUsingPATCHNotFound  %+v", 404, o.Payload)
}
func (o *ModifyProjectUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ModifyProjectUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
