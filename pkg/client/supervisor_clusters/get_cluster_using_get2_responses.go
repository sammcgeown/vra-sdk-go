// Code generated by go-swagger; DO NOT EDIT.

package supervisor_clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// GetClusterUsingGET2Reader is a Reader for the GetClusterUsingGET2 structure.
type GetClusterUsingGET2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterUsingGET2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterUsingGET2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterUsingGET2OK creates a GetClusterUsingGET2OK with default headers values
func NewGetClusterUsingGET2OK() *GetClusterUsingGET2OK {
	return &GetClusterUsingGET2OK{}
}

/* GetClusterUsingGET2OK describes a response with status code 200, with default header values.

OK
*/
type GetClusterUsingGET2OK struct {
	Payload *models.SupervisorCluster
}

func (o *GetClusterUsingGET2OK) Error() string {
	return fmt.Sprintf("[GET /cmx/api/resources/supervisor-clusters/{selfLinkId}][%d] getClusterUsingGET2OK  %+v", 200, o.Payload)
}
func (o *GetClusterUsingGET2OK) GetPayload() *models.SupervisorCluster {
	return o.Payload
}

func (o *GetClusterUsingGET2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupervisorCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}