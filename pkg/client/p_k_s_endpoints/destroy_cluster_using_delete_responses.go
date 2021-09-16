// Code generated by go-swagger; DO NOT EDIT.

package p_k_s_endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DestroyClusterUsingDELETEReader is a Reader for the DestroyClusterUsingDELETE structure.
type DestroyClusterUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DestroyClusterUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDestroyClusterUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDestroyClusterUsingDELETEOK creates a DestroyClusterUsingDELETEOK with default headers values
func NewDestroyClusterUsingDELETEOK() *DestroyClusterUsingDELETEOK {
	return &DestroyClusterUsingDELETEOK{}
}

/* DestroyClusterUsingDELETEOK describes a response with status code 200, with default header values.

OK
*/
type DestroyClusterUsingDELETEOK struct {
}

func (o *DestroyClusterUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /cmx/api/resources/pks/endpoints/{id}/clusters/{clusterId}][%d] destroyClusterUsingDELETEOK ", 200)
}

func (o *DestroyClusterUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
