// Code generated by go-swagger; DO NOT EDIT.

package policy_types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetTypesUsingGET1Reader is a Reader for the GetTypesUsingGET1 structure.
type GetTypesUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTypesUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTypesUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTypesUsingGET1OK creates a GetTypesUsingGET1OK with default headers values
func NewGetTypesUsingGET1OK() *GetTypesUsingGET1OK {
	return &GetTypesUsingGET1OK{}
}

/* GetTypesUsingGET1OK describes a response with status code 200, with default header values.

OK
*/
type GetTypesUsingGET1OK struct {
	Payload *models.PageOfPolicyType
}

func (o *GetTypesUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /policy/api/policyTypes][%d] getTypesUsingGET1OK  %+v", 200, o.Payload)
}
func (o *GetTypesUsingGET1OK) GetPayload() *models.PageOfPolicyType {
	return o.Payload
}

func (o *GetTypesUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PageOfPolicyType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
