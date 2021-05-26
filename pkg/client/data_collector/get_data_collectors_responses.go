// Code generated by go-swagger; DO NOT EDIT.

package data_collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetDataCollectorsReader is a Reader for the GetDataCollectors structure.
type GetDataCollectorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDataCollectorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDataCollectorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetDataCollectorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDataCollectorsOK creates a GetDataCollectorsOK with default headers values
func NewGetDataCollectorsOK() *GetDataCollectorsOK {
	return &GetDataCollectorsOK{}
}

/* GetDataCollectorsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetDataCollectorsOK struct {
	Payload *models.DataCollectorResult
}

func (o *GetDataCollectorsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/data-collectors][%d] getDataCollectorsOK  %+v", 200, o.Payload)
}
func (o *GetDataCollectorsOK) GetPayload() *models.DataCollectorResult {
	return o.Payload
}

func (o *GetDataCollectorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DataCollectorResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDataCollectorsForbidden creates a GetDataCollectorsForbidden with default headers values
func NewGetDataCollectorsForbidden() *GetDataCollectorsForbidden {
	return &GetDataCollectorsForbidden{}
}

/* GetDataCollectorsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetDataCollectorsForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *GetDataCollectorsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/data-collectors][%d] getDataCollectorsForbidden  %+v", 403, o.Payload)
}
func (o *GetDataCollectorsForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetDataCollectorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
