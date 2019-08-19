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

// GetVSphereStorageProfilesReader is a Reader for the GetVSphereStorageProfiles structure.
type GetVSphereStorageProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVSphereStorageProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVSphereStorageProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetVSphereStorageProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVSphereStorageProfilesOK creates a GetVSphereStorageProfilesOK with default headers values
func NewGetVSphereStorageProfilesOK() *GetVSphereStorageProfilesOK {
	return &GetVSphereStorageProfilesOK{}
}

/*GetVSphereStorageProfilesOK handles this case with default header values.

successful operation
*/
type GetVSphereStorageProfilesOK struct {
	Payload *models.StorageProfileVsphereResult
}

func (o *GetVSphereStorageProfilesOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-vsphere][%d] getVSphereStorageProfilesOK  %+v", 200, o.Payload)
}

func (o *GetVSphereStorageProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageProfileVsphereResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVSphereStorageProfilesForbidden creates a GetVSphereStorageProfilesForbidden with default headers values
func NewGetVSphereStorageProfilesForbidden() *GetVSphereStorageProfilesForbidden {
	return &GetVSphereStorageProfilesForbidden{}
}

/*GetVSphereStorageProfilesForbidden handles this case with default header values.

Forbidden
*/
type GetVSphereStorageProfilesForbidden struct {
}

func (o *GetVSphereStorageProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-vsphere][%d] getVSphereStorageProfilesForbidden ", 403)
}

func (o *GetVSphereStorageProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
