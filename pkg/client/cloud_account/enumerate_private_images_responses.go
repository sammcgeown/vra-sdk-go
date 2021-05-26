// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EnumeratePrivateImagesReader is a Reader for the EnumeratePrivateImages structure.
type EnumeratePrivateImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EnumeratePrivateImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewEnumeratePrivateImagesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewEnumeratePrivateImagesDefault creates a EnumeratePrivateImagesDefault with default headers values
func NewEnumeratePrivateImagesDefault(code int) *EnumeratePrivateImagesDefault {
	return &EnumeratePrivateImagesDefault{
		_statusCode: code,
	}
}

/* EnumeratePrivateImagesDefault describes a response with status code -1, with default header values.

successful operation
*/
type EnumeratePrivateImagesDefault struct {
	_statusCode int
}

// Code gets the status code for the enumerate private images default response
func (o *EnumeratePrivateImagesDefault) Code() int {
	return o._statusCode
}

func (o *EnumeratePrivateImagesDefault) Error() string {
	return fmt.Sprintf("[POST /iaas/api/cloud-accounts/{id}/private-image-enumeration][%d] enumeratePrivateImages default ", o._statusCode)
}

func (o *EnumeratePrivateImagesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
