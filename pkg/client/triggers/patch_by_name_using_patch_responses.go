// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/vmware/vra-sdk-go/pkg/models"
)

// PatchByNameUsingPATCHReader is a Reader for the PatchByNameUsingPATCH structure.
type PatchByNameUsingPATCHReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchByNameUsingPATCHReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchByNameUsingPATCHOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchByNameUsingPATCHUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchByNameUsingPATCHForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchByNameUsingPATCHNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchByNameUsingPATCHInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchByNameUsingPATCHOK creates a PatchByNameUsingPATCHOK with default headers values
func NewPatchByNameUsingPATCHOK() *PatchByNameUsingPATCHOK {
	return &PatchByNameUsingPATCHOK{}
}

/* PatchByNameUsingPATCHOK describes a response with status code 200, with default header values.

'Success' with Docker Registry Webhook patch
*/
type PatchByNameUsingPATCHOK struct {
	Payload models.DockerRegistryWebHook
}

func (o *PatchByNameUsingPATCHOK) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchByNameUsingPATCHOK  %+v", 200, o.Payload)
}
func (o *PatchByNameUsingPATCHOK) GetPayload() models.DockerRegistryWebHook {
	return o.Payload
}

func (o *PatchByNameUsingPATCHOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalDockerRegistryWebHook(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewPatchByNameUsingPATCHUnauthorized creates a PatchByNameUsingPATCHUnauthorized with default headers values
func NewPatchByNameUsingPATCHUnauthorized() *PatchByNameUsingPATCHUnauthorized {
	return &PatchByNameUsingPATCHUnauthorized{}
}

/* PatchByNameUsingPATCHUnauthorized describes a response with status code 401, with default header values.

Unauthorized Request
*/
type PatchByNameUsingPATCHUnauthorized struct {
}

func (o *PatchByNameUsingPATCHUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchByNameUsingPATCHUnauthorized ", 401)
}

func (o *PatchByNameUsingPATCHUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchByNameUsingPATCHForbidden creates a PatchByNameUsingPATCHForbidden with default headers values
func NewPatchByNameUsingPATCHForbidden() *PatchByNameUsingPATCHForbidden {
	return &PatchByNameUsingPATCHForbidden{}
}

/* PatchByNameUsingPATCHForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchByNameUsingPATCHForbidden struct {
}

func (o *PatchByNameUsingPATCHForbidden) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchByNameUsingPATCHForbidden ", 403)
}

func (o *PatchByNameUsingPATCHForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchByNameUsingPATCHNotFound creates a PatchByNameUsingPATCHNotFound with default headers values
func NewPatchByNameUsingPATCHNotFound() *PatchByNameUsingPATCHNotFound {
	return &PatchByNameUsingPATCHNotFound{}
}

/* PatchByNameUsingPATCHNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchByNameUsingPATCHNotFound struct {
	Payload *models.Error
}

func (o *PatchByNameUsingPATCHNotFound) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchByNameUsingPATCHNotFound  %+v", 404, o.Payload)
}
func (o *PatchByNameUsingPATCHNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchByNameUsingPATCHNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchByNameUsingPATCHInternalServerError creates a PatchByNameUsingPATCHInternalServerError with default headers values
func NewPatchByNameUsingPATCHInternalServerError() *PatchByNameUsingPATCHInternalServerError {
	return &PatchByNameUsingPATCHInternalServerError{}
}

/* PatchByNameUsingPATCHInternalServerError describes a response with status code 500, with default header values.

Server Error
*/
type PatchByNameUsingPATCHInternalServerError struct {
}

func (o *PatchByNameUsingPATCHInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /codestream/api/registry-webhooks/{project}/{name}][%d] patchByNameUsingPATCHInternalServerError ", 500)
}

func (o *PatchByNameUsingPATCHInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}