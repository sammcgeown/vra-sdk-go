// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetResourceByIDUsingGETReader is a Reader for the GetResourceByIDUsingGET structure.
type GetResourceByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetResourceByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceByIDUsingGETOK creates a GetResourceByIDUsingGETOK with default headers values
func NewGetResourceByIDUsingGETOK() *GetResourceByIDUsingGETOK {
	return &GetResourceByIDUsingGETOK{}
}

/* GetResourceByIDUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetResourceByIDUsingGETOK struct {
	Payload *models.DeploymentResource
}

func (o *GetResourceByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] getResourceByIdUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetResourceByIDUsingGETOK) GetPayload() *models.DeploymentResource {
	return o.Payload
}

func (o *GetResourceByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceByIDUsingGETUnauthorized creates a GetResourceByIDUsingGETUnauthorized with default headers values
func NewGetResourceByIDUsingGETUnauthorized() *GetResourceByIDUsingGETUnauthorized {
	return &GetResourceByIDUsingGETUnauthorized{}
}

/* GetResourceByIDUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetResourceByIDUsingGETUnauthorized struct {
}

func (o *GetResourceByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] getResourceByIdUsingGETUnauthorized ", 401)
}

func (o *GetResourceByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetResourceByIDUsingGETNotFound creates a GetResourceByIDUsingGETNotFound with default headers values
func NewGetResourceByIDUsingGETNotFound() *GetResourceByIDUsingGETNotFound {
	return &GetResourceByIDUsingGETNotFound{}
}

/* GetResourceByIDUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetResourceByIDUsingGETNotFound struct {
	Payload *models.Error
}

func (o *GetResourceByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] getResourceByIdUsingGETNotFound  %+v", 404, o.Payload)
}
func (o *GetResourceByIDUsingGETNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetResourceByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
