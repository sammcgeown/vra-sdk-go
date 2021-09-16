// Code generated by go-swagger; DO NOT EDIT.

package supervisor_namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateUsingPATCHMixin5Reader is a Reader for the UpdateUsingPATCHMixin5 structure.
type UpdateUsingPATCHMixin5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUsingPATCHMixin5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUsingPATCHMixin5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUsingPATCHMixin5OK creates a UpdateUsingPATCHMixin5OK with default headers values
func NewUpdateUsingPATCHMixin5OK() *UpdateUsingPATCHMixin5OK {
	return &UpdateUsingPATCHMixin5OK{}
}

/* UpdateUsingPATCHMixin5OK describes a response with status code 200, with default header values.

OK
*/
type UpdateUsingPATCHMixin5OK struct {
	Payload *models.SupervisorNamespace
}

func (o *UpdateUsingPATCHMixin5OK) Error() string {
	return fmt.Sprintf("[PATCH /cmx/api/resources/supervisor-namespaces/{selfLinkId}][%d] updateUsingPATCHMixin5OK  %+v", 200, o.Payload)
}
func (o *UpdateUsingPATCHMixin5OK) GetPayload() *models.SupervisorNamespace {
	return o.Payload
}

func (o *UpdateUsingPATCHMixin5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupervisorNamespace)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}