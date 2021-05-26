// Code generated by go-swagger; DO NOT EDIT.

package fabric_compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricComputeReader is a Reader for the GetFabricCompute structure.
type GetFabricComputeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricComputeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFabricComputeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFabricComputeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFabricComputeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetFabricComputeOK creates a GetFabricComputeOK with default headers values
func NewGetFabricComputeOK() *GetFabricComputeOK {
	return &GetFabricComputeOK{}
}

/* GetFabricComputeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetFabricComputeOK struct {
	Payload *models.FabricCompute
}

func (o *GetFabricComputeOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-computes/{id}][%d] getFabricComputeOK  %+v", 200, o.Payload)
}
func (o *GetFabricComputeOK) GetPayload() *models.FabricCompute {
	return o.Payload
}

func (o *GetFabricComputeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricCompute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricComputeForbidden creates a GetFabricComputeForbidden with default headers values
func NewGetFabricComputeForbidden() *GetFabricComputeForbidden {
	return &GetFabricComputeForbidden{}
}

/* GetFabricComputeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetFabricComputeForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *GetFabricComputeForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-computes/{id}][%d] getFabricComputeForbidden  %+v", 403, o.Payload)
}
func (o *GetFabricComputeForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetFabricComputeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricComputeNotFound creates a GetFabricComputeNotFound with default headers values
func NewGetFabricComputeNotFound() *GetFabricComputeNotFound {
	return &GetFabricComputeNotFound{}
}

/* GetFabricComputeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetFabricComputeNotFound struct {
	Payload *models.Error
}

func (o *GetFabricComputeNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-computes/{id}][%d] getFabricComputeNotFound  %+v", 404, o.Payload)
}
func (o *GetFabricComputeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFabricComputeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
