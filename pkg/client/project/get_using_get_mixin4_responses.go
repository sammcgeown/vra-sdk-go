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

// GetUsingGETMixin4Reader is a Reader for the GetUsingGETMixin4 structure.
type GetUsingGETMixin4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGETMixin4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGETMixin4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetUsingGETMixin4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsingGETMixin4OK creates a GetUsingGETMixin4OK with default headers values
func NewGetUsingGETMixin4OK() *GetUsingGETMixin4OK {
	return &GetUsingGETMixin4OK{}
}

/* GetUsingGETMixin4OK describes a response with status code 200, with default header values.

'OK' with the ProjectResourceMetadata
*/
type GetUsingGETMixin4OK struct {
	Payload *models.ProjectResourceMetadata
}

func (o *GetUsingGETMixin4OK) Error() string {
	return fmt.Sprintf("[GET /project-service/api/projects/{id}/resource-metadata][%d] getUsingGETMixin4OK  %+v", 200, o.Payload)
}
func (o *GetUsingGETMixin4OK) GetPayload() *models.ProjectResourceMetadata {
	return o.Payload
}

func (o *GetUsingGETMixin4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectResourceMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsingGETMixin4NotFound creates a GetUsingGETMixin4NotFound with default headers values
func NewGetUsingGETMixin4NotFound() *GetUsingGETMixin4NotFound {
	return &GetUsingGETMixin4NotFound{}
}

/* GetUsingGETMixin4NotFound describes a response with status code 404, with default header values.

'Not found' if no project with the provided id
*/
type GetUsingGETMixin4NotFound struct {
	Payload *models.Error
}

func (o *GetUsingGETMixin4NotFound) Error() string {
	return fmt.Sprintf("[GET /project-service/api/projects/{id}/resource-metadata][%d] getUsingGETMixin4NotFound  %+v", 404, o.Payload)
}
func (o *GetUsingGETMixin4NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetUsingGETMixin4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}