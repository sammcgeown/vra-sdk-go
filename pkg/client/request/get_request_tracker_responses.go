// Code generated by go-swagger; DO NOT EDIT.

package request

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetRequestTrackerReader is a Reader for the GetRequestTracker structure.
type GetRequestTrackerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestTrackerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRequestTrackerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetRequestTrackerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRequestTrackerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRequestTrackerOK creates a GetRequestTrackerOK with default headers values
func NewGetRequestTrackerOK() *GetRequestTrackerOK {
	return &GetRequestTrackerOK{}
}

/*GetRequestTrackerOK handles this case with default header values.

successful operation
*/
type GetRequestTrackerOK struct {
	Payload *models.RequestTracker
}

func (o *GetRequestTrackerOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker/{id}][%d] getRequestTrackerOK  %+v", 200, o.Payload)
}

func (o *GetRequestTrackerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestTrackerForbidden creates a GetRequestTrackerForbidden with default headers values
func NewGetRequestTrackerForbidden() *GetRequestTrackerForbidden {
	return &GetRequestTrackerForbidden{}
}

/*GetRequestTrackerForbidden handles this case with default header values.

Forbidden
*/
type GetRequestTrackerForbidden struct {
}

func (o *GetRequestTrackerForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker/{id}][%d] getRequestTrackerForbidden ", 403)
}

func (o *GetRequestTrackerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRequestTrackerNotFound creates a GetRequestTrackerNotFound with default headers values
func NewGetRequestTrackerNotFound() *GetRequestTrackerNotFound {
	return &GetRequestTrackerNotFound{}
}

/*GetRequestTrackerNotFound handles this case with default header values.

Not Found
*/
type GetRequestTrackerNotFound struct {
}

func (o *GetRequestTrackerNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/request-tracker/{id}][%d] getRequestTrackerNotFound ", 404)
}

func (o *GetRequestTrackerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}