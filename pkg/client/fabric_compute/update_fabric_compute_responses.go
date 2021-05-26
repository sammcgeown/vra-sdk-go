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

// UpdateFabricComputeReader is a Reader for the UpdateFabricCompute structure.
type UpdateFabricComputeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFabricComputeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFabricComputeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateFabricComputeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateFabricComputeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateFabricComputeOK creates a UpdateFabricComputeOK with default headers values
func NewUpdateFabricComputeOK() *UpdateFabricComputeOK {
	return &UpdateFabricComputeOK{}
}

/* UpdateFabricComputeOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateFabricComputeOK struct {
	Payload *models.FabricCompute
}

func (o *UpdateFabricComputeOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/fabric-computes/{id}][%d] updateFabricComputeOK  %+v", 200, o.Payload)
}
func (o *UpdateFabricComputeOK) GetPayload() *models.FabricCompute {
	return o.Payload
}

func (o *UpdateFabricComputeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricCompute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFabricComputeForbidden creates a UpdateFabricComputeForbidden with default headers values
func NewUpdateFabricComputeForbidden() *UpdateFabricComputeForbidden {
	return &UpdateFabricComputeForbidden{}
}

/* UpdateFabricComputeForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateFabricComputeForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *UpdateFabricComputeForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/fabric-computes/{id}][%d] updateFabricComputeForbidden  %+v", 403, o.Payload)
}
func (o *UpdateFabricComputeForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateFabricComputeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFabricComputeNotFound creates a UpdateFabricComputeNotFound with default headers values
func NewUpdateFabricComputeNotFound() *UpdateFabricComputeNotFound {
	return &UpdateFabricComputeNotFound{}
}

/* UpdateFabricComputeNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateFabricComputeNotFound struct {
	Payload *models.Error
}

func (o *UpdateFabricComputeNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/fabric-computes/{id}][%d] updateFabricComputeNotFound  %+v", 404, o.Payload)
}
func (o *UpdateFabricComputeNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateFabricComputeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
