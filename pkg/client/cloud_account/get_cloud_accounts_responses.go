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

// GetCloudAccountsReader is a Reader for the GetCloudAccounts structure.
type GetCloudAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCloudAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCloudAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetCloudAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCloudAccountsOK creates a GetCloudAccountsOK with default headers values
func NewGetCloudAccountsOK() *GetCloudAccountsOK {
	return &GetCloudAccountsOK{}
}

/*GetCloudAccountsOK handles this case with default header values.

successful operation
*/
type GetCloudAccountsOK struct {
	Payload *models.CloudAccountResult
}

func (o *GetCloudAccountsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts][%d] getCloudAccountsOK  %+v", 200, o.Payload)
}

func (o *GetCloudAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCloudAccountsForbidden creates a GetCloudAccountsForbidden with default headers values
func NewGetCloudAccountsForbidden() *GetCloudAccountsForbidden {
	return &GetCloudAccountsForbidden{}
}

/*GetCloudAccountsForbidden handles this case with default header values.

Forbidden
*/
type GetCloudAccountsForbidden struct {
}

func (o *GetCloudAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts][%d] getCloudAccountsForbidden ", 403)
}

func (o *GetCloudAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
