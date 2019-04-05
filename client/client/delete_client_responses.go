// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gotify/go-api-client/v2/models"
)

// DeleteClientReader is a Reader for the DeleteClient structure.
type DeleteClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteClientBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteClientUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteClientForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteClientNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteClientOK creates a DeleteClientOK with default headers values
func NewDeleteClientOK() *DeleteClientOK {
	return &DeleteClientOK{}
}

/*DeleteClientOK handles this case with default header values.

Ok
*/
type DeleteClientOK struct {
}

func (o *DeleteClientOK) Error() string {
	return fmt.Sprintf("[DELETE /client/{id}][%d] deleteClientOK ", 200)
}

func (o *DeleteClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteClientBadRequest creates a DeleteClientBadRequest with default headers values
func NewDeleteClientBadRequest() *DeleteClientBadRequest {
	return &DeleteClientBadRequest{}
}

/*DeleteClientBadRequest handles this case with default header values.

Bad Request
*/
type DeleteClientBadRequest struct {
	Payload *models.Error
}

func (o *DeleteClientBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /client/{id}][%d] deleteClientBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteClientBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClientUnauthorized creates a DeleteClientUnauthorized with default headers values
func NewDeleteClientUnauthorized() *DeleteClientUnauthorized {
	return &DeleteClientUnauthorized{}
}

/*DeleteClientUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteClientUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteClientUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /client/{id}][%d] deleteClientUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteClientUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClientForbidden creates a DeleteClientForbidden with default headers values
func NewDeleteClientForbidden() *DeleteClientForbidden {
	return &DeleteClientForbidden{}
}

/*DeleteClientForbidden handles this case with default header values.

Forbidden
*/
type DeleteClientForbidden struct {
	Payload *models.Error
}

func (o *DeleteClientForbidden) Error() string {
	return fmt.Sprintf("[DELETE /client/{id}][%d] deleteClientForbidden  %+v", 403, o.Payload)
}

func (o *DeleteClientForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClientNotFound creates a DeleteClientNotFound with default headers values
func NewDeleteClientNotFound() *DeleteClientNotFound {
	return &DeleteClientNotFound{}
}

/*DeleteClientNotFound handles this case with default header values.

Not Found
*/
type DeleteClientNotFound struct {
	Payload *models.Error
}

func (o *DeleteClientNotFound) Error() string {
	return fmt.Sprintf("[DELETE /client/{id}][%d] deleteClientNotFound  %+v", 404, o.Payload)
}

func (o *DeleteClientNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
