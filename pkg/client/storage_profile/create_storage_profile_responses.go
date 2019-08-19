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

// CreateStorageProfileReader is a Reader for the CreateStorageProfile structure.
type CreateStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateStorageProfileCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateStorageProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateStorageProfileCreated creates a CreateStorageProfileCreated with default headers values
func NewCreateStorageProfileCreated() *CreateStorageProfileCreated {
	return &CreateStorageProfileCreated{}
}

/*CreateStorageProfileCreated handles this case with default header values.

successful operation
*/
type CreateStorageProfileCreated struct {
	Payload *models.StorageProfile
}

func (o *CreateStorageProfileCreated) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles][%d] createStorageProfileCreated  %+v", 201, o.Payload)
}

func (o *CreateStorageProfileCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateStorageProfileBadRequest creates a CreateStorageProfileBadRequest with default headers values
func NewCreateStorageProfileBadRequest() *CreateStorageProfileBadRequest {
	return &CreateStorageProfileBadRequest{}
}

/*CreateStorageProfileBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateStorageProfileBadRequest struct {
}

func (o *CreateStorageProfileBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles][%d] createStorageProfileBadRequest ", 400)
}

func (o *CreateStorageProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateStorageProfileForbidden creates a CreateStorageProfileForbidden with default headers values
func NewCreateStorageProfileForbidden() *CreateStorageProfileForbidden {
	return &CreateStorageProfileForbidden{}
}

/*CreateStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type CreateStorageProfileForbidden struct {
}

func (o *CreateStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/storage-profiles][%d] createStorageProfileForbidden ", 403)
}

func (o *CreateStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
