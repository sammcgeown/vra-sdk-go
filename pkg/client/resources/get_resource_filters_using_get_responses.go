// Code generated by go-swagger; DO NOT EDIT.

package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetResourceFiltersUsingGETReader is a Reader for the GetResourceFiltersUsingGET structure.
type GetResourceFiltersUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResourceFiltersUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResourceFiltersUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetResourceFiltersUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetResourceFiltersUsingGETOK creates a GetResourceFiltersUsingGETOK with default headers values
func NewGetResourceFiltersUsingGETOK() *GetResourceFiltersUsingGETOK {
	return &GetResourceFiltersUsingGETOK{}
}

/* GetResourceFiltersUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type GetResourceFiltersUsingGETOK struct {
	Payload *models.DeploymentFilterSchema
}

func (o *GetResourceFiltersUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resources/filters][%d] getResourceFiltersUsingGETOK  %+v", 200, o.Payload)
}
func (o *GetResourceFiltersUsingGETOK) GetPayload() *models.DeploymentFilterSchema {
	return o.Payload
}

func (o *GetResourceFiltersUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentFilterSchema)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResourceFiltersUsingGETUnauthorized creates a GetResourceFiltersUsingGETUnauthorized with default headers values
func NewGetResourceFiltersUsingGETUnauthorized() *GetResourceFiltersUsingGETUnauthorized {
	return &GetResourceFiltersUsingGETUnauthorized{}
}

/* GetResourceFiltersUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetResourceFiltersUsingGETUnauthorized struct {
}

func (o *GetResourceFiltersUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /deployment/api/resources/filters][%d] getResourceFiltersUsingGETUnauthorized ", 401)
}

func (o *GetResourceFiltersUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
