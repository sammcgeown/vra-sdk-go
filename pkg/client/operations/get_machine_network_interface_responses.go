// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetMachineNetworkInterfaceReader is a Reader for the GetMachineNetworkInterface structure.
type GetMachineNetworkInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineNetworkInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMachineNetworkInterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetMachineNetworkInterfaceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetMachineNetworkInterfaceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMachineNetworkInterfaceOK creates a GetMachineNetworkInterfaceOK with default headers values
func NewGetMachineNetworkInterfaceOK() *GetMachineNetworkInterfaceOK {
	return &GetMachineNetworkInterfaceOK{}
}

/*GetMachineNetworkInterfaceOK handles this case with default header values.

successful operation
*/
type GetMachineNetworkInterfaceOK struct {
	Payload *models.NetworkInterface
}

func (o *GetMachineNetworkInterfaceOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}/network-interfaces/{id1}][%d] getMachineNetworkInterfaceOK  %+v", 200, o.Payload)
}

func (o *GetMachineNetworkInterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineNetworkInterfaceForbidden creates a GetMachineNetworkInterfaceForbidden with default headers values
func NewGetMachineNetworkInterfaceForbidden() *GetMachineNetworkInterfaceForbidden {
	return &GetMachineNetworkInterfaceForbidden{}
}

/*GetMachineNetworkInterfaceForbidden handles this case with default header values.

Forbidden
*/
type GetMachineNetworkInterfaceForbidden struct {
}

func (o *GetMachineNetworkInterfaceForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}/network-interfaces/{id1}][%d] getMachineNetworkInterfaceForbidden ", 403)
}

func (o *GetMachineNetworkInterfaceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMachineNetworkInterfaceNotFound creates a GetMachineNetworkInterfaceNotFound with default headers values
func NewGetMachineNetworkInterfaceNotFound() *GetMachineNetworkInterfaceNotFound {
	return &GetMachineNetworkInterfaceNotFound{}
}

/*GetMachineNetworkInterfaceNotFound handles this case with default header values.

Not Found
*/
type GetMachineNetworkInterfaceNotFound struct {
}

func (o *GetMachineNetworkInterfaceNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}/network-interfaces/{id1}][%d] getMachineNetworkInterfaceNotFound ", 404)
}

func (o *GetMachineNetworkInterfaceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
