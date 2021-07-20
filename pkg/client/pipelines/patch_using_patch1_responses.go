// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PatchUsingPATCH1Reader is a Reader for the PatchUsingPATCH1 structure.
type PatchUsingPATCH1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchUsingPATCH1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchUsingPATCH1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchUsingPATCH1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchUsingPATCH1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchUsingPATCH1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchUsingPATCH1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchUsingPATCH1OK creates a PatchUsingPATCH1OK with default headers values
func NewPatchUsingPATCH1OK() *PatchUsingPATCH1OK {
	return &PatchUsingPATCH1OK{}
}

/* PatchUsingPATCH1OK describes a response with status code 200, with default header values.

'Success' with the updated Pipeline
*/
type PatchUsingPATCH1OK struct {
	Payload models.Pipeline
}

func (o *PatchUsingPATCH1OK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}][%d] patchUsingPATCH1OK  %+v", 200, o.Payload)
}
func (o *PatchUsingPATCH1OK) GetPayload() models.Pipeline {
	return o.Payload
}

func (o *PatchUsingPATCH1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalPipeline(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchUsingPATCH1Unauthorized creates a PatchUsingPATCH1Unauthorized with default headers values
func NewPatchUsingPATCH1Unauthorized() *PatchUsingPATCH1Unauthorized {
	return &PatchUsingPATCH1Unauthorized{}
}

/* PatchUsingPATCH1Unauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type PatchUsingPATCH1Unauthorized struct {
}

func (o *PatchUsingPATCH1Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}][%d] patchUsingPATCH1Unauthorized ", 401)
}

func (o *PatchUsingPATCH1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchUsingPATCH1Forbidden creates a PatchUsingPATCH1Forbidden with default headers values
func NewPatchUsingPATCH1Forbidden() *PatchUsingPATCH1Forbidden {
	return &PatchUsingPATCH1Forbidden{}
}

/* PatchUsingPATCH1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchUsingPATCH1Forbidden struct {
}

func (o *PatchUsingPATCH1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}][%d] patchUsingPATCH1Forbidden ", 403)
}

func (o *PatchUsingPATCH1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchUsingPATCH1NotFound creates a PatchUsingPATCH1NotFound with default headers values
func NewPatchUsingPATCH1NotFound() *PatchUsingPATCH1NotFound {
	return &PatchUsingPATCH1NotFound{}
}

/* PatchUsingPATCH1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchUsingPATCH1NotFound struct {
	Payload *models.Error
}

func (o *PatchUsingPATCH1NotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}][%d] patchUsingPATCH1NotFound  %+v", 404, o.Payload)
}
func (o *PatchUsingPATCH1NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchUsingPATCH1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchUsingPATCH1InternalServerError creates a PatchUsingPATCH1InternalServerError with default headers values
func NewPatchUsingPATCH1InternalServerError() *PatchUsingPATCH1InternalServerError {
	return &PatchUsingPATCH1InternalServerError{}
}

/* PatchUsingPATCH1InternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PatchUsingPATCH1InternalServerError struct {
}

func (o *PatchUsingPATCH1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/pipelines/{project}/{name}][%d] patchUsingPATCH1InternalServerError ", 500)
}

func (o *PatchUsingPATCH1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}