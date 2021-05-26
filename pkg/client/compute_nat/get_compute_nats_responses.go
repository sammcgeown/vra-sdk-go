// Code generated by go-swagger; DO NOT EDIT.

package compute_nat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetComputeNatsReader is a Reader for the GetComputeNats structure.
type GetComputeNatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComputeNatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComputeNatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetComputeNatsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComputeNatsOK creates a GetComputeNatsOK with default headers values
func NewGetComputeNatsOK() *GetComputeNatsOK {
	return &GetComputeNatsOK{}
}

/* GetComputeNatsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetComputeNatsOK struct {
	Payload *models.ComputeNatResult
}

func (o *GetComputeNatsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-nats][%d] getComputeNatsOK  %+v", 200, o.Payload)
}
func (o *GetComputeNatsOK) GetPayload() *models.ComputeNatResult {
	return o.Payload
}

func (o *GetComputeNatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputeNatResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeNatsForbidden creates a GetComputeNatsForbidden with default headers values
func NewGetComputeNatsForbidden() *GetComputeNatsForbidden {
	return &GetComputeNatsForbidden{}
}

/* GetComputeNatsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetComputeNatsForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *GetComputeNatsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/compute-nats][%d] getComputeNatsForbidden  %+v", 403, o.Payload)
}
func (o *GetComputeNatsForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetComputeNatsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
