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

// GetGcpCloudAccountReader is a Reader for the GetGcpCloudAccount structure.
type GetGcpCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGcpCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetGcpCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetGcpCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetGcpCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGcpCloudAccountOK creates a GetGcpCloudAccountOK with default headers values
func NewGetGcpCloudAccountOK() *GetGcpCloudAccountOK {
	return &GetGcpCloudAccountOK{}
}

/*GetGcpCloudAccountOK handles this case with default header values.

successful operation
*/
type GetGcpCloudAccountOK struct {
	Payload *models.CloudAccountGcp
}

func (o *GetGcpCloudAccountOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-gcp/{id}][%d] getGcpCloudAccountOK  %+v", 200, o.Payload)
}

func (o *GetGcpCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountGcp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGcpCloudAccountForbidden creates a GetGcpCloudAccountForbidden with default headers values
func NewGetGcpCloudAccountForbidden() *GetGcpCloudAccountForbidden {
	return &GetGcpCloudAccountForbidden{}
}

/*GetGcpCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type GetGcpCloudAccountForbidden struct {
}

func (o *GetGcpCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-gcp/{id}][%d] getGcpCloudAccountForbidden ", 403)
}

func (o *GetGcpCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGcpCloudAccountNotFound creates a GetGcpCloudAccountNotFound with default headers values
func NewGetGcpCloudAccountNotFound() *GetGcpCloudAccountNotFound {
	return &GetGcpCloudAccountNotFound{}
}

/*GetGcpCloudAccountNotFound handles this case with default header values.

Not Found
*/
type GetGcpCloudAccountNotFound struct {
}

func (o *GetGcpCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-gcp/{id}][%d] getGcpCloudAccountNotFound ", 404)
}

func (o *GetGcpCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
