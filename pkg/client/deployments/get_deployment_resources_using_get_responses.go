// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDeploymentResourcesUsingGETReader is a Reader for the GetDeploymentResourcesUsingGET structure.
type GetDeploymentResourcesUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentResourcesUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeploymentResourcesUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetDeploymentResourcesUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetDeploymentResourcesUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDeploymentResourcesUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDeploymentResourcesUsingGETOK creates a GetDeploymentResourcesUsingGETOK with default headers values
func NewGetDeploymentResourcesUsingGETOK() *GetDeploymentResourcesUsingGETOK {
	return &GetDeploymentResourcesUsingGETOK{}
}

/*GetDeploymentResourcesUsingGETOK handles this case with default header values.

OK
*/
type GetDeploymentResourcesUsingGETOK struct {
	Payload *models.PageOfResource
}

func (o *GetDeploymentResourcesUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{depId}/resources][%d] getDeploymentResourcesUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentResourcesUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentResourcesUsingGETUnauthorized creates a GetDeploymentResourcesUsingGETUnauthorized with default headers values
func NewGetDeploymentResourcesUsingGETUnauthorized() *GetDeploymentResourcesUsingGETUnauthorized {
	return &GetDeploymentResourcesUsingGETUnauthorized{}
}

/*GetDeploymentResourcesUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetDeploymentResourcesUsingGETUnauthorized struct {
}

func (o *GetDeploymentResourcesUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{depId}/resources][%d] getDeploymentResourcesUsingGETUnauthorized ", 401)
}

func (o *GetDeploymentResourcesUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentResourcesUsingGETForbidden creates a GetDeploymentResourcesUsingGETForbidden with default headers values
func NewGetDeploymentResourcesUsingGETForbidden() *GetDeploymentResourcesUsingGETForbidden {
	return &GetDeploymentResourcesUsingGETForbidden{}
}

/*GetDeploymentResourcesUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetDeploymentResourcesUsingGETForbidden struct {
}

func (o *GetDeploymentResourcesUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{depId}/resources][%d] getDeploymentResourcesUsingGETForbidden ", 403)
}

func (o *GetDeploymentResourcesUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDeploymentResourcesUsingGETNotFound creates a GetDeploymentResourcesUsingGETNotFound with default headers values
func NewGetDeploymentResourcesUsingGETNotFound() *GetDeploymentResourcesUsingGETNotFound {
	return &GetDeploymentResourcesUsingGETNotFound{}
}

/*GetDeploymentResourcesUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetDeploymentResourcesUsingGETNotFound struct {
}

func (o *GetDeploymentResourcesUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /deployment/api/deployments/{depId}/resources][%d] getDeploymentResourcesUsingGETNotFound ", 404)
}

func (o *GetDeploymentResourcesUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
