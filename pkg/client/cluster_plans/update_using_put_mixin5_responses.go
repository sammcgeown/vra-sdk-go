// Code generated by go-swagger; DO NOT EDIT.

package cluster_plans

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateUsingPUTMixin5Reader is a Reader for the UpdateUsingPUTMixin5 structure.
type UpdateUsingPUTMixin5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUsingPUTMixin5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUsingPUTMixin5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUsingPUTMixin5OK creates a UpdateUsingPUTMixin5OK with default headers values
func NewUpdateUsingPUTMixin5OK() *UpdateUsingPUTMixin5OK {
	return &UpdateUsingPUTMixin5OK{}
}

/* UpdateUsingPUTMixin5OK describes a response with status code 200, with default header values.

OK
*/
type UpdateUsingPUTMixin5OK struct {
	Payload *models.ClusterPlan
}

func (o *UpdateUsingPUTMixin5OK) Error() string {
	return fmt.Sprintf("[PUT /cmx/api/resources/cluster-plans/{id}][%d] updateUsingPUTMixin5OK  %+v", 200, o.Payload)
}
func (o *UpdateUsingPUTMixin5OK) GetPayload() *models.ClusterPlan {
	return o.Payload
}

func (o *UpdateUsingPUTMixin5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
