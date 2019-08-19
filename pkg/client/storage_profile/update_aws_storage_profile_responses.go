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

// UpdateAwsStorageProfileReader is a Reader for the UpdateAwsStorageProfile structure.
type UpdateAwsStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAwsStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAwsStorageProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateAwsStorageProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewUpdateAwsStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAwsStorageProfileOK creates a UpdateAwsStorageProfileOK with default headers values
func NewUpdateAwsStorageProfileOK() *UpdateAwsStorageProfileOK {
	return &UpdateAwsStorageProfileOK{}
}

/*UpdateAwsStorageProfileOK handles this case with default header values.

successful operation
*/
type UpdateAwsStorageProfileOK struct {
	Payload *models.AwsStorageProfile
}

func (o *UpdateAwsStorageProfileOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/storage-profiles-aws/{id}][%d] updateAwsStorageProfileOK  %+v", 200, o.Payload)
}

func (o *UpdateAwsStorageProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AwsStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAwsStorageProfileBadRequest creates a UpdateAwsStorageProfileBadRequest with default headers values
func NewUpdateAwsStorageProfileBadRequest() *UpdateAwsStorageProfileBadRequest {
	return &UpdateAwsStorageProfileBadRequest{}
}

/*UpdateAwsStorageProfileBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type UpdateAwsStorageProfileBadRequest struct {
}

func (o *UpdateAwsStorageProfileBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/storage-profiles-aws/{id}][%d] updateAwsStorageProfileBadRequest ", 400)
}

func (o *UpdateAwsStorageProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAwsStorageProfileForbidden creates a UpdateAwsStorageProfileForbidden with default headers values
func NewUpdateAwsStorageProfileForbidden() *UpdateAwsStorageProfileForbidden {
	return &UpdateAwsStorageProfileForbidden{}
}

/*UpdateAwsStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type UpdateAwsStorageProfileForbidden struct {
}

func (o *UpdateAwsStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/storage-profiles-aws/{id}][%d] updateAwsStorageProfileForbidden ", 403)
}

func (o *UpdateAwsStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
