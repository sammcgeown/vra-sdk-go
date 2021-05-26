// Code generated by go-swagger; DO NOT EDIT.

package catalog_entitlements

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateEntitlementUsingPOSTReader is a Reader for the CreateEntitlementUsingPOST structure.
type CreateEntitlementUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateEntitlementUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateEntitlementUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateEntitlementUsingPOSTCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateEntitlementUsingPOSTBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateEntitlementUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateEntitlementUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateEntitlementUsingPOSTOK creates a CreateEntitlementUsingPOSTOK with default headers values
func NewCreateEntitlementUsingPOSTOK() *CreateEntitlementUsingPOSTOK {
	return &CreateEntitlementUsingPOSTOK{}
}

/* CreateEntitlementUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type CreateEntitlementUsingPOSTOK struct {
	Payload *models.Entitlement
}

func (o *CreateEntitlementUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/entitlements][%d] createEntitlementUsingPOSTOK  %+v", 200, o.Payload)
}
func (o *CreateEntitlementUsingPOSTOK) GetPayload() *models.Entitlement {
	return o.Payload
}

func (o *CreateEntitlementUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Entitlement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEntitlementUsingPOSTCreated creates a CreateEntitlementUsingPOSTCreated with default headers values
func NewCreateEntitlementUsingPOSTCreated() *CreateEntitlementUsingPOSTCreated {
	return &CreateEntitlementUsingPOSTCreated{}
}

/* CreateEntitlementUsingPOSTCreated describes a response with status code 201, with default header values.

Created
*/
type CreateEntitlementUsingPOSTCreated struct {
	Payload *models.Entitlement
}

func (o *CreateEntitlementUsingPOSTCreated) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/entitlements][%d] createEntitlementUsingPOSTCreated  %+v", 201, o.Payload)
}
func (o *CreateEntitlementUsingPOSTCreated) GetPayload() *models.Entitlement {
	return o.Payload
}

func (o *CreateEntitlementUsingPOSTCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Entitlement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEntitlementUsingPOSTBadRequest creates a CreateEntitlementUsingPOSTBadRequest with default headers values
func NewCreateEntitlementUsingPOSTBadRequest() *CreateEntitlementUsingPOSTBadRequest {
	return &CreateEntitlementUsingPOSTBadRequest{}
}

/* CreateEntitlementUsingPOSTBadRequest describes a response with status code 400, with default header values.

Catalog item or source cannot be entitled to the project
*/
type CreateEntitlementUsingPOSTBadRequest struct {
	Payload *models.Error
}

func (o *CreateEntitlementUsingPOSTBadRequest) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/entitlements][%d] createEntitlementUsingPOSTBadRequest  %+v", 400, o.Payload)
}
func (o *CreateEntitlementUsingPOSTBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateEntitlementUsingPOSTBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateEntitlementUsingPOSTUnauthorized creates a CreateEntitlementUsingPOSTUnauthorized with default headers values
func NewCreateEntitlementUsingPOSTUnauthorized() *CreateEntitlementUsingPOSTUnauthorized {
	return &CreateEntitlementUsingPOSTUnauthorized{}
}

/* CreateEntitlementUsingPOSTUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateEntitlementUsingPOSTUnauthorized struct {
}

func (o *CreateEntitlementUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/entitlements][%d] createEntitlementUsingPOSTUnauthorized ", 401)
}

func (o *CreateEntitlementUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateEntitlementUsingPOSTNotFound creates a CreateEntitlementUsingPOSTNotFound with default headers values
func NewCreateEntitlementUsingPOSTNotFound() *CreateEntitlementUsingPOSTNotFound {
	return &CreateEntitlementUsingPOSTNotFound{}
}

/* CreateEntitlementUsingPOSTNotFound describes a response with status code 404, with default header values.

Catalog item or catalog source not found
*/
type CreateEntitlementUsingPOSTNotFound struct {
	Payload *models.Error
}

func (o *CreateEntitlementUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /catalog/api/admin/entitlements][%d] createEntitlementUsingPOSTNotFound  %+v", 404, o.Payload)
}
func (o *CreateEntitlementUsingPOSTNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateEntitlementUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
