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

// DeleteResourceUsingDELETEReader is a Reader for the DeleteResourceUsingDELETE structure.
type DeleteResourceUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResourceUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteResourceUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteResourceUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteResourceUsingDELETEForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteResourceUsingDELETENotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteResourceUsingDELETEOK creates a DeleteResourceUsingDELETEOK with default headers values
func NewDeleteResourceUsingDELETEOK() *DeleteResourceUsingDELETEOK {
	return &DeleteResourceUsingDELETEOK{}
}

/* DeleteResourceUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type DeleteResourceUsingDELETEOK struct {
	Payload *models.Request
}

func (o *DeleteResourceUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETEOK  %+v", 200, o.Payload)
}
func (o *DeleteResourceUsingDELETEOK) GetPayload() *models.Request {
	return o.Payload
}

func (o *DeleteResourceUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Request)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResourceUsingDELETEUnauthorized creates a DeleteResourceUsingDELETEUnauthorized with default headers values
func NewDeleteResourceUsingDELETEUnauthorized() *DeleteResourceUsingDELETEUnauthorized {
	return &DeleteResourceUsingDELETEUnauthorized{}
}

/* DeleteResourceUsingDELETEUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteResourceUsingDELETEUnauthorized struct {
}

func (o *DeleteResourceUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETEUnauthorized ", 401)
}

func (o *DeleteResourceUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourceUsingDELETEForbidden creates a DeleteResourceUsingDELETEForbidden with default headers values
func NewDeleteResourceUsingDELETEForbidden() *DeleteResourceUsingDELETEForbidden {
	return &DeleteResourceUsingDELETEForbidden{}
}

/* DeleteResourceUsingDELETEForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteResourceUsingDELETEForbidden struct {
}

func (o *DeleteResourceUsingDELETEForbidden) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETEForbidden ", 403)
}

func (o *DeleteResourceUsingDELETEForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteResourceUsingDELETENotFound creates a DeleteResourceUsingDELETENotFound with default headers values
func NewDeleteResourceUsingDELETENotFound() *DeleteResourceUsingDELETENotFound {
	return &DeleteResourceUsingDELETENotFound{}
}

/* DeleteResourceUsingDELETENotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteResourceUsingDELETENotFound struct {
	Payload *models.Error
}

func (o *DeleteResourceUsingDELETENotFound) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{deploymentId}/resources/{resourceId}][%d] deleteResourceUsingDELETENotFound  %+v", 404, o.Payload)
}
func (o *DeleteResourceUsingDELETENotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteResourceUsingDELETENotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
