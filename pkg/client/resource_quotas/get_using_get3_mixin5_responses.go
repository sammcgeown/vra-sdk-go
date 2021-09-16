// Code generated by go-swagger; DO NOT EDIT.

package resource_quotas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetUsingGET3Mixin5Reader is a Reader for the GetUsingGET3Mixin5 structure.
type GetUsingGET3Mixin5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGET3Mixin5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGET3Mixin5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsingGET3Mixin5OK creates a GetUsingGET3Mixin5OK with default headers values
func NewGetUsingGET3Mixin5OK() *GetUsingGET3Mixin5OK {
	return &GetUsingGET3Mixin5OK{}
}

/* GetUsingGET3Mixin5OK describes a response with status code 200, with default header values.

OK
*/
type GetUsingGET3Mixin5OK struct {
	Payload *models.K8SResourceQuota
}

func (o *GetUsingGET3Mixin5OK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/resource-quotas/{id}][%d] getUsingGET3Mixin5OK  %+v", 200, o.Payload)
}
func (o *GetUsingGET3Mixin5OK) GetPayload() *models.K8SResourceQuota {
	return o.Payload
}

func (o *GetUsingGET3Mixin5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.K8SResourceQuota)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}