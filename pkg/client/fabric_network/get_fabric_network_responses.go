// Code generated by go-swagger; DO NOT EDIT.

package fabric_network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricNetworkReader is a Reader for the GetFabricNetwork structure.
type GetFabricNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFabricNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetFabricNetworkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFabricNetworkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFabricNetworkOK creates a GetFabricNetworkOK with default headers values
func NewGetFabricNetworkOK() *GetFabricNetworkOK {
	return &GetFabricNetworkOK{}
}

/*GetFabricNetworkOK handles this case with default header values.

successful operation
*/
type GetFabricNetworkOK struct {
	Payload *models.FabricNetwork
}

func (o *GetFabricNetworkOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-networks/{id}][%d] getFabricNetworkOK  %+v", 200, o.Payload)
}

func (o *GetFabricNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricNetworkForbidden creates a GetFabricNetworkForbidden with default headers values
func NewGetFabricNetworkForbidden() *GetFabricNetworkForbidden {
	return &GetFabricNetworkForbidden{}
}

/*GetFabricNetworkForbidden handles this case with default header values.

Forbidden
*/
type GetFabricNetworkForbidden struct {
}

func (o *GetFabricNetworkForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-networks/{id}][%d] getFabricNetworkForbidden ", 403)
}

func (o *GetFabricNetworkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFabricNetworkNotFound creates a GetFabricNetworkNotFound with default headers values
func NewGetFabricNetworkNotFound() *GetFabricNetworkNotFound {
	return &GetFabricNetworkNotFound{}
}

/*GetFabricNetworkNotFound handles this case with default header values.

Not Found
*/
type GetFabricNetworkNotFound struct {
}

func (o *GetFabricNetworkNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-networks/{id}][%d] getFabricNetworkNotFound ", 404)
}

func (o *GetFabricNetworkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
