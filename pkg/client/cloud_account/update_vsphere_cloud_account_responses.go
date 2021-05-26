// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateVSphereCloudAccountReader is a Reader for the UpdateVSphereCloudAccount structure.
type UpdateVSphereCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateVSphereCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateVSphereCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateVSphereCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateVSphereCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateVSphereCloudAccountOK creates a UpdateVSphereCloudAccountOK with default headers values
func NewUpdateVSphereCloudAccountOK() *UpdateVSphereCloudAccountOK {
	return &UpdateVSphereCloudAccountOK{}
}

/* UpdateVSphereCloudAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type UpdateVSphereCloudAccountOK struct {
	Payload *models.CloudAccountVsphere
}

func (o *UpdateVSphereCloudAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vsphere/{id}][%d] updateVSphereCloudAccountOK  %+v", 200, o.Payload)
}
func (o *UpdateVSphereCloudAccountOK) GetPayload() *models.CloudAccountVsphere {
	return o.Payload
}

func (o *UpdateVSphereCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountVsphere)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVSphereCloudAccountForbidden creates a UpdateVSphereCloudAccountForbidden with default headers values
func NewUpdateVSphereCloudAccountForbidden() *UpdateVSphereCloudAccountForbidden {
	return &UpdateVSphereCloudAccountForbidden{}
}

/* UpdateVSphereCloudAccountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateVSphereCloudAccountForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *UpdateVSphereCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vsphere/{id}][%d] updateVSphereCloudAccountForbidden  %+v", 403, o.Payload)
}
func (o *UpdateVSphereCloudAccountForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *UpdateVSphereCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateVSphereCloudAccountNotFound creates a UpdateVSphereCloudAccountNotFound with default headers values
func NewUpdateVSphereCloudAccountNotFound() *UpdateVSphereCloudAccountNotFound {
	return &UpdateVSphereCloudAccountNotFound{}
}

/* UpdateVSphereCloudAccountNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateVSphereCloudAccountNotFound struct {
	Payload *models.Error
}

func (o *UpdateVSphereCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-vsphere/{id}][%d] updateVSphereCloudAccountNotFound  %+v", 404, o.Payload)
}
func (o *UpdateVSphereCloudAccountNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateVSphereCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
