// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetProjectsReader is a Reader for the GetProjects structure.
type GetProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetProjectsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProjectsOK creates a GetProjectsOK with default headers values
func NewGetProjectsOK() *GetProjectsOK {
	return &GetProjectsOK{}
}

/*GetProjectsOK handles this case with default header values.

successful operation
*/
type GetProjectsOK struct {
	Payload *models.ProjectResult
}

func (o *GetProjectsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/projects][%d] getProjectsOK  %+v", 200, o.Payload)
}

func (o *GetProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectsForbidden creates a GetProjectsForbidden with default headers values
func NewGetProjectsForbidden() *GetProjectsForbidden {
	return &GetProjectsForbidden{}
}

/*GetProjectsForbidden handles this case with default header values.

Forbidden
*/
type GetProjectsForbidden struct {
}

func (o *GetProjectsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/projects][%d] getProjectsForbidden ", 403)
}

func (o *GetProjectsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
