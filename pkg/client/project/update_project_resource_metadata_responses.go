// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// UpdateProjectResourceMetadataReader is a Reader for the UpdateProjectResourceMetadata structure.
type UpdateProjectResourceMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectResourceMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateProjectResourceMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateProjectResourceMetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateProjectResourceMetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateProjectResourceMetadataOK creates a UpdateProjectResourceMetadataOK with default headers values
func NewUpdateProjectResourceMetadataOK() *UpdateProjectResourceMetadataOK {
	return &UpdateProjectResourceMetadataOK{}
}

/*UpdateProjectResourceMetadataOK handles this case with default header values.

successful operation
*/
type UpdateProjectResourceMetadataOK struct {
	Payload *models.Project
}

func (o *UpdateProjectResourceMetadataOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}/resource-metadata][%d] updateProjectResourceMetadataOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectResourceMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectResourceMetadataBadRequest creates a UpdateProjectResourceMetadataBadRequest with default headers values
func NewUpdateProjectResourceMetadataBadRequest() *UpdateProjectResourceMetadataBadRequest {
	return &UpdateProjectResourceMetadataBadRequest{}
}

/*UpdateProjectResourceMetadataBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type UpdateProjectResourceMetadataBadRequest struct {
}

func (o *UpdateProjectResourceMetadataBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}/resource-metadata][%d] updateProjectResourceMetadataBadRequest ", 400)
}

func (o *UpdateProjectResourceMetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateProjectResourceMetadataForbidden creates a UpdateProjectResourceMetadataForbidden with default headers values
func NewUpdateProjectResourceMetadataForbidden() *UpdateProjectResourceMetadataForbidden {
	return &UpdateProjectResourceMetadataForbidden{}
}

/*UpdateProjectResourceMetadataForbidden handles this case with default header values.

Forbidden
*/
type UpdateProjectResourceMetadataForbidden struct {
}

func (o *UpdateProjectResourceMetadataForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/projects/{id}/resource-metadata][%d] updateProjectResourceMetadataForbidden ", 403)
}

func (o *UpdateProjectResourceMetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
