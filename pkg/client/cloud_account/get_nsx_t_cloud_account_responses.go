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

// GetNsxTCloudAccountReader is a Reader for the GetNsxTCloudAccount structure.
type GetNsxTCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNsxTCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNsxTCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetNsxTCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetNsxTCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNsxTCloudAccountOK creates a GetNsxTCloudAccountOK with default headers values
func NewGetNsxTCloudAccountOK() *GetNsxTCloudAccountOK {
	return &GetNsxTCloudAccountOK{}
}

/*GetNsxTCloudAccountOK handles this case with default header values.

successful operation
*/
type GetNsxTCloudAccountOK struct {
	Payload *models.CloudAccountNsxT
}

func (o *GetNsxTCloudAccountOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-t/{id}][%d] getNsxTCloudAccountOK  %+v", 200, o.Payload)
}

func (o *GetNsxTCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountNsxT)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNsxTCloudAccountForbidden creates a GetNsxTCloudAccountForbidden with default headers values
func NewGetNsxTCloudAccountForbidden() *GetNsxTCloudAccountForbidden {
	return &GetNsxTCloudAccountForbidden{}
}

/*GetNsxTCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type GetNsxTCloudAccountForbidden struct {
}

func (o *GetNsxTCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-t/{id}][%d] getNsxTCloudAccountForbidden ", 403)
}

func (o *GetNsxTCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNsxTCloudAccountNotFound creates a GetNsxTCloudAccountNotFound with default headers values
func NewGetNsxTCloudAccountNotFound() *GetNsxTCloudAccountNotFound {
	return &GetNsxTCloudAccountNotFound{}
}

/*GetNsxTCloudAccountNotFound handles this case with default header values.

Not Found
*/
type GetNsxTCloudAccountNotFound struct {
}

func (o *GetNsxTCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-nsx-t/{id}][%d] getNsxTCloudAccountNotFound ", 404)
}

func (o *GetNsxTCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
