// Code generated by go-swagger; DO NOT EDIT.

package load_balancer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteLoadBalancerReader is a Reader for the DeleteLoadBalancer structure.
type DeleteLoadBalancerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLoadBalancerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteLoadBalancerAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteLoadBalancerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLoadBalancerAccepted creates a DeleteLoadBalancerAccepted with default headers values
func NewDeleteLoadBalancerAccepted() *DeleteLoadBalancerAccepted {
	return &DeleteLoadBalancerAccepted{}
}

/*DeleteLoadBalancerAccepted handles this case with default header values.

successful operation
*/
type DeleteLoadBalancerAccepted struct {
	Payload *models.RequestTracker
}

func (o *DeleteLoadBalancerAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/load-balancers/{id}][%d] deleteLoadBalancerAccepted  %+v", 202, o.Payload)
}

func (o *DeleteLoadBalancerAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteLoadBalancerForbidden creates a DeleteLoadBalancerForbidden with default headers values
func NewDeleteLoadBalancerForbidden() *DeleteLoadBalancerForbidden {
	return &DeleteLoadBalancerForbidden{}
}

/*DeleteLoadBalancerForbidden handles this case with default header values.

Forbidden
*/
type DeleteLoadBalancerForbidden struct {
}

func (o *DeleteLoadBalancerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/load-balancers/{id}][%d] deleteLoadBalancerForbidden ", 403)
}

func (o *DeleteLoadBalancerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
