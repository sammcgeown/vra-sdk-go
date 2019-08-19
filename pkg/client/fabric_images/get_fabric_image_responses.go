// Code generated by go-swagger; DO NOT EDIT.

package fabric_images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetFabricImageReader is a Reader for the GetFabricImage structure.
type GetFabricImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFabricImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetFabricImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFabricImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFabricImageOK creates a GetFabricImageOK with default headers values
func NewGetFabricImageOK() *GetFabricImageOK {
	return &GetFabricImageOK{}
}

/*GetFabricImageOK handles this case with default header values.

successful operation
*/
type GetFabricImageOK struct {
	Payload *models.FabricImage
}

func (o *GetFabricImageOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-images/{id}][%d] getFabricImageOK  %+v", 200, o.Payload)
}

func (o *GetFabricImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricImage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricImageForbidden creates a GetFabricImageForbidden with default headers values
func NewGetFabricImageForbidden() *GetFabricImageForbidden {
	return &GetFabricImageForbidden{}
}

/*GetFabricImageForbidden handles this case with default header values.

Forbidden
*/
type GetFabricImageForbidden struct {
}

func (o *GetFabricImageForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-images/{id}][%d] getFabricImageForbidden ", 403)
}

func (o *GetFabricImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFabricImageNotFound creates a GetFabricImageNotFound with default headers values
func NewGetFabricImageNotFound() *GetFabricImageNotFound {
	return &GetFabricImageNotFound{}
}

/*GetFabricImageNotFound handles this case with default header values.

Not Found
*/
type GetFabricImageNotFound struct {
}

func (o *GetFabricImageNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-images/{id}][%d] getFabricImageNotFound ", 404)
}

func (o *GetFabricImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
