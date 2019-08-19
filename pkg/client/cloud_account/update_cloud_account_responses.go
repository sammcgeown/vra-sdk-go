// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateCloudAccountReader is a Reader for the UpdateCloudAccount structure.
type UpdateCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewUpdateCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCloudAccountOK creates a UpdateCloudAccountOK with default headers values
func NewUpdateCloudAccountOK() *UpdateCloudAccountOK {
	return &UpdateCloudAccountOK{}
}

/*UpdateCloudAccountOK handles this case with default header values.

successful operation
*/
type UpdateCloudAccountOK struct {
	Payload *models.CloudAccount
}

func (o *UpdateCloudAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts/{id}][%d] updateCloudAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudAccountForbidden creates a UpdateCloudAccountForbidden with default headers values
func NewUpdateCloudAccountForbidden() *UpdateCloudAccountForbidden {
	return &UpdateCloudAccountForbidden{}
}

/*UpdateCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type UpdateCloudAccountForbidden struct {
}

func (o *UpdateCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts/{id}][%d] updateCloudAccountForbidden ", 403)
}

func (o *UpdateCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateCloudAccountNotFound creates a UpdateCloudAccountNotFound with default headers values
func NewUpdateCloudAccountNotFound() *UpdateCloudAccountNotFound {
	return &UpdateCloudAccountNotFound{}
}

/*UpdateCloudAccountNotFound handles this case with default header values.

Not Found
*/
type UpdateCloudAccountNotFound struct {
}

func (o *UpdateCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts/{id}][%d] updateCloudAccountNotFound ", 404)
}

func (o *UpdateCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
