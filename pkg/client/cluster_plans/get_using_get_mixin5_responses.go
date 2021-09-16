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

// GetUsingGETMixin5Reader is a Reader for the GetUsingGETMixin5 structure.
type GetUsingGETMixin5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGETMixin5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGETMixin5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsingGETMixin5OK creates a GetUsingGETMixin5OK with default headers values
func NewGetUsingGETMixin5OK() *GetUsingGETMixin5OK {
	return &GetUsingGETMixin5OK{}
}

/* GetUsingGETMixin5OK describes a response with status code 200, with default header values.

OK
*/
type GetUsingGETMixin5OK struct {
	Payload *models.ClusterPlan
}

func (o *GetUsingGETMixin5OK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/cluster-plans/{id}][%d] getUsingGETMixin5OK  %+v", 200, o.Payload)
}
func (o *GetUsingGETMixin5OK) GetPayload() *models.ClusterPlan {
	return o.Payload
}

func (o *GetUsingGETMixin5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}