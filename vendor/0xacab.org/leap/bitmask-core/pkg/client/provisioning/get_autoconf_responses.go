// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetAutoconfReader is a Reader for the GetAutoconf structure.
type GetAutoconfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAutoconfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAutoconfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAutoconfBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAutoconfNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAutoconfInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /autoconf] GetAutoconf", response, response.Code())
	}
}

// NewGetAutoconfOK creates a GetAutoconfOK with default headers values
func NewGetAutoconfOK() *GetAutoconfOK {
	return &GetAutoconfOK{}
}

/*
GetAutoconfOK describes a response with status code 200, with default header values.

OK
*/
type GetAutoconfOK struct {
	Payload string
}

// IsSuccess returns true when this get autoconf o k response has a 2xx status code
func (o *GetAutoconfOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get autoconf o k response has a 3xx status code
func (o *GetAutoconfOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get autoconf o k response has a 4xx status code
func (o *GetAutoconfOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get autoconf o k response has a 5xx status code
func (o *GetAutoconfOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get autoconf o k response a status code equal to that given
func (o *GetAutoconfOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get autoconf o k response
func (o *GetAutoconfOK) Code() int {
	return 200
}

func (o *GetAutoconfOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /autoconf][%d] getAutoconfOK %s", 200, payload)
}

func (o *GetAutoconfOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /autoconf][%d] getAutoconfOK %s", 200, payload)
}

func (o *GetAutoconfOK) GetPayload() string {
	return o.Payload
}

func (o *GetAutoconfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAutoconfBadRequest creates a GetAutoconfBadRequest with default headers values
func NewGetAutoconfBadRequest() *GetAutoconfBadRequest {
	return &GetAutoconfBadRequest{}
}

/*
GetAutoconfBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetAutoconfBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this get autoconf bad request response has a 2xx status code
func (o *GetAutoconfBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get autoconf bad request response has a 3xx status code
func (o *GetAutoconfBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get autoconf bad request response has a 4xx status code
func (o *GetAutoconfBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get autoconf bad request response has a 5xx status code
func (o *GetAutoconfBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get autoconf bad request response a status code equal to that given
func (o *GetAutoconfBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get autoconf bad request response
func (o *GetAutoconfBadRequest) Code() int {
	return 400
}

func (o *GetAutoconfBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /autoconf][%d] getAutoconfBadRequest %s", 400, payload)
}

func (o *GetAutoconfBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /autoconf][%d] getAutoconfBadRequest %s", 400, payload)
}

func (o *GetAutoconfBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAutoconfBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAutoconfNotFound creates a GetAutoconfNotFound with default headers values
func NewGetAutoconfNotFound() *GetAutoconfNotFound {
	return &GetAutoconfNotFound{}
}

/*
GetAutoconfNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetAutoconfNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this get autoconf not found response has a 2xx status code
func (o *GetAutoconfNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get autoconf not found response has a 3xx status code
func (o *GetAutoconfNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get autoconf not found response has a 4xx status code
func (o *GetAutoconfNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get autoconf not found response has a 5xx status code
func (o *GetAutoconfNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get autoconf not found response a status code equal to that given
func (o *GetAutoconfNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get autoconf not found response
func (o *GetAutoconfNotFound) Code() int {
	return 404
}

func (o *GetAutoconfNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /autoconf][%d] getAutoconfNotFound %s", 404, payload)
}

func (o *GetAutoconfNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /autoconf][%d] getAutoconfNotFound %s", 404, payload)
}

func (o *GetAutoconfNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAutoconfNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAutoconfInternalServerError creates a GetAutoconfInternalServerError with default headers values
func NewGetAutoconfInternalServerError() *GetAutoconfInternalServerError {
	return &GetAutoconfInternalServerError{}
}

/*
GetAutoconfInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetAutoconfInternalServerError struct {
	Payload interface{}
}

// IsSuccess returns true when this get autoconf internal server error response has a 2xx status code
func (o *GetAutoconfInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get autoconf internal server error response has a 3xx status code
func (o *GetAutoconfInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get autoconf internal server error response has a 4xx status code
func (o *GetAutoconfInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get autoconf internal server error response has a 5xx status code
func (o *GetAutoconfInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get autoconf internal server error response a status code equal to that given
func (o *GetAutoconfInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get autoconf internal server error response
func (o *GetAutoconfInternalServerError) Code() int {
	return 500
}

func (o *GetAutoconfInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /autoconf][%d] getAutoconfInternalServerError %s", 500, payload)
}

func (o *GetAutoconfInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /autoconf][%d] getAutoconfInternalServerError %s", 500, payload)
}

func (o *GetAutoconfInternalServerError) GetPayload() interface{} {
	return o.Payload
}

func (o *GetAutoconfInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
