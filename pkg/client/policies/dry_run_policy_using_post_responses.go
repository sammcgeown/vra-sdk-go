// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DryRunPolicyUsingPOSTReader is a Reader for the DryRunPolicyUsingPOST structure.
type DryRunPolicyUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DryRunPolicyUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDryRunPolicyUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDryRunPolicyUsingPOSTOK creates a DryRunPolicyUsingPOSTOK with default headers values
func NewDryRunPolicyUsingPOSTOK() *DryRunPolicyUsingPOSTOK {
	return &DryRunPolicyUsingPOSTOK{}
}

/* DryRunPolicyUsingPOSTOK describes a response with status code 200, with default header values.

OK
*/
type DryRunPolicyUsingPOSTOK struct {
}

func (o *DryRunPolicyUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /policy/api/policies][%d] dryRunPolicyUsingPOSTOK ", 200)
}

func (o *DryRunPolicyUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
