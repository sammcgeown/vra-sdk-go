// Code generated by go-swagger; DO NOT EDIT.

package blueprint_deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetRequestEventsUsingGETReader is a Reader for the GetRequestEventsUsingGET structure.
type GetRequestEventsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestEventsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRequestEventsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetRequestEventsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRequestEventsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRequestEventsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRequestEventsUsingGETOK creates a GetRequestEventsUsingGETOK with default headers values
func NewGetRequestEventsUsingGETOK() *GetRequestEventsUsingGETOK {
	return &GetRequestEventsUsingGETOK{}
}

/*GetRequestEventsUsingGETOK handles this case with default header values.

OK
*/
type GetRequestEventsUsingGETOK struct {
	Payload *models.BlueprintRequestEventQueryResult
}

func (o *GetRequestEventsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-deployments/{deploymentId}/events][%d] getRequestEventsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetRequestEventsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlueprintRequestEventQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRequestEventsUsingGETUnauthorized creates a GetRequestEventsUsingGETUnauthorized with default headers values
func NewGetRequestEventsUsingGETUnauthorized() *GetRequestEventsUsingGETUnauthorized {
	return &GetRequestEventsUsingGETUnauthorized{}
}

/*GetRequestEventsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetRequestEventsUsingGETUnauthorized struct {
}

func (o *GetRequestEventsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-deployments/{deploymentId}/events][%d] getRequestEventsUsingGETUnauthorized ", 401)
}

func (o *GetRequestEventsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRequestEventsUsingGETForbidden creates a GetRequestEventsUsingGETForbidden with default headers values
func NewGetRequestEventsUsingGETForbidden() *GetRequestEventsUsingGETForbidden {
	return &GetRequestEventsUsingGETForbidden{}
}

/*GetRequestEventsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetRequestEventsUsingGETForbidden struct {
}

func (o *GetRequestEventsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-deployments/{deploymentId}/events][%d] getRequestEventsUsingGETForbidden ", 403)
}

func (o *GetRequestEventsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRequestEventsUsingGETNotFound creates a GetRequestEventsUsingGETNotFound with default headers values
func NewGetRequestEventsUsingGETNotFound() *GetRequestEventsUsingGETNotFound {
	return &GetRequestEventsUsingGETNotFound{}
}

/*GetRequestEventsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetRequestEventsUsingGETNotFound struct {
}

func (o *GetRequestEventsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprint-deployments/{deploymentId}/events][%d] getRequestEventsUsingGETNotFound ", 404)
}

func (o *GetRequestEventsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
