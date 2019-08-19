// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// CreateVSphereStorageProfileReader is a Reader for the CreateVSphereStorageProfile structure.
type CreateVSphereStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVSphereStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateVSphereStorageProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateVSphereStorageProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateVSphereStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateVSphereStorageProfileCreated creates a CreateVSphereStorageProfileCreated with default headers values
func NewCreateVSphereStorageProfileCreated() *CreateVSphereStorageProfileCreated {
	return &CreateVSphereStorageProfileCreated{}
}

/*CreateVSphereStorageProfileCreated handles this case with default header values.

successful operation
*/
type CreateVSphereStorageProfileCreated struct {
	Payload *models.VsphereStorageProfile
}

func (o *CreateVSphereStorageProfileCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-vsphere][%d] createVSphereStorageProfileCreated  %+v", 201, o.Payload)
}

func (o *CreateVSphereStorageProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VsphereStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVSphereStorageProfileBadRequest creates a CreateVSphereStorageProfileBadRequest with default headers values
func NewCreateVSphereStorageProfileBadRequest() *CreateVSphereStorageProfileBadRequest {
	return &CreateVSphereStorageProfileBadRequest{}
}

/*CreateVSphereStorageProfileBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateVSphereStorageProfileBadRequest struct {
}

func (o *CreateVSphereStorageProfileBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-vsphere][%d] createVSphereStorageProfileBadRequest ", 400)
}

func (o *CreateVSphereStorageProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateVSphereStorageProfileForbidden creates a CreateVSphereStorageProfileForbidden with default headers values
func NewCreateVSphereStorageProfileForbidden() *CreateVSphereStorageProfileForbidden {
	return &CreateVSphereStorageProfileForbidden{}
}

/*CreateVSphereStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type CreateVSphereStorageProfileForbidden struct {
}

func (o *CreateVSphereStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles-vsphere][%d] createVSphereStorageProfileForbidden ", 403)
}

func (o *CreateVSphereStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
