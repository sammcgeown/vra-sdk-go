// Code generated by go-swagger; DO NOT EDIT.

package blueprint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetBlueprintVersionsUsingGETReader is a Reader for the GetBlueprintVersionsUsingGET structure.
type GetBlueprintVersionsUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprintVersionsUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlueprintVersionsUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetBlueprintVersionsUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetBlueprintVersionsUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetBlueprintVersionsUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlueprintVersionsUsingGETOK creates a GetBlueprintVersionsUsingGETOK with default headers values
func NewGetBlueprintVersionsUsingGETOK() *GetBlueprintVersionsUsingGETOK {
	return &GetBlueprintVersionsUsingGETOK{}
}

/*GetBlueprintVersionsUsingGETOK handles this case with default header values.

OK
*/
type GetBlueprintVersionsUsingGETOK struct {
	Payload *models.BlueprintVersionQueryResult
}

func (o *GetBlueprintVersionsUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions][%d] getBlueprintVersionsUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetBlueprintVersionsUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlueprintVersionQueryResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlueprintVersionsUsingGETUnauthorized creates a GetBlueprintVersionsUsingGETUnauthorized with default headers values
func NewGetBlueprintVersionsUsingGETUnauthorized() *GetBlueprintVersionsUsingGETUnauthorized {
	return &GetBlueprintVersionsUsingGETUnauthorized{}
}

/*GetBlueprintVersionsUsingGETUnauthorized handles this case with default header values.

Unauthorized
*/
type GetBlueprintVersionsUsingGETUnauthorized struct {
}

func (o *GetBlueprintVersionsUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions][%d] getBlueprintVersionsUsingGETUnauthorized ", 401)
}

func (o *GetBlueprintVersionsUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintVersionsUsingGETForbidden creates a GetBlueprintVersionsUsingGETForbidden with default headers values
func NewGetBlueprintVersionsUsingGETForbidden() *GetBlueprintVersionsUsingGETForbidden {
	return &GetBlueprintVersionsUsingGETForbidden{}
}

/*GetBlueprintVersionsUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetBlueprintVersionsUsingGETForbidden struct {
}

func (o *GetBlueprintVersionsUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions][%d] getBlueprintVersionsUsingGETForbidden ", 403)
}

func (o *GetBlueprintVersionsUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintVersionsUsingGETNotFound creates a GetBlueprintVersionsUsingGETNotFound with default headers values
func NewGetBlueprintVersionsUsingGETNotFound() *GetBlueprintVersionsUsingGETNotFound {
	return &GetBlueprintVersionsUsingGETNotFound{}
}

/*GetBlueprintVersionsUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetBlueprintVersionsUsingGETNotFound struct {
}

func (o *GetBlueprintVersionsUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}/versions][%d] getBlueprintVersionsUsingGETNotFound ", 404)
}

func (o *GetBlueprintVersionsUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
