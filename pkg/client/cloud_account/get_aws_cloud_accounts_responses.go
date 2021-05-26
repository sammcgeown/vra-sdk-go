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

// GetAwsCloudAccountsReader is a Reader for the GetAwsCloudAccounts structure.
type GetAwsCloudAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAwsCloudAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAwsCloudAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetAwsCloudAccountsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAwsCloudAccountsOK creates a GetAwsCloudAccountsOK with default headers values
func NewGetAwsCloudAccountsOK() *GetAwsCloudAccountsOK {
	return &GetAwsCloudAccountsOK{}
}

/* GetAwsCloudAccountsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAwsCloudAccountsOK struct {
	Payload *models.CloudAccountAwsResult
}

func (o *GetAwsCloudAccountsOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-aws][%d] getAwsCloudAccountsOK  %+v", 200, o.Payload)
}
func (o *GetAwsCloudAccountsOK) GetPayload() *models.CloudAccountAwsResult {
	return o.Payload
}

func (o *GetAwsCloudAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountAwsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAwsCloudAccountsForbidden creates a GetAwsCloudAccountsForbidden with default headers values
func NewGetAwsCloudAccountsForbidden() *GetAwsCloudAccountsForbidden {
	return &GetAwsCloudAccountsForbidden{}
}

/* GetAwsCloudAccountsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAwsCloudAccountsForbidden struct {
	Payload *models.ServiceErrorResponse
}

func (o *GetAwsCloudAccountsForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/cloud-accounts-aws][%d] getAwsCloudAccountsForbidden  %+v", 403, o.Payload)
}
func (o *GetAwsCloudAccountsForbidden) GetPayload() *models.ServiceErrorResponse {
	return o.Payload
}

func (o *GetAwsCloudAccountsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
