package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/segmentio/ecs-blox/blox/models"
)

// GetInstanceReader is a Reader for the GetInstance structure.
type GetInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstanceOK creates a GetInstanceOK with default headers values
func NewGetInstanceOK() *GetInstanceOK {
	return &GetInstanceOK{}
}

/*GetInstanceOK handles this case with default header values.

Get instance using cluster name and instance ARN - success
*/
type GetInstanceOK struct {
	Payload *models.ContainerInstance
}

func (o *GetInstanceOK) Error() string {
	return fmt.Sprintf("[GET /instances/{cluster}/{arn}][%d] getInstanceOK  %+v", 200, o.Payload)
}

func (o *GetInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceNotFound creates a GetInstanceNotFound with default headers values
func NewGetInstanceNotFound() *GetInstanceNotFound {
	return &GetInstanceNotFound{}
}

/*GetInstanceNotFound handles this case with default header values.

Get instance using cluster name and instance ARN - instance not found
*/
type GetInstanceNotFound struct {
	Payload string
}

func (o *GetInstanceNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{cluster}/{arn}][%d] getInstanceNotFound  %+v", 404, o.Payload)
}

func (o *GetInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstanceInternalServerError creates a GetInstanceInternalServerError with default headers values
func NewGetInstanceInternalServerError() *GetInstanceInternalServerError {
	return &GetInstanceInternalServerError{}
}

/*GetInstanceInternalServerError handles this case with default header values.

Get instance using cluster name and instance ARN - unexpected error
*/
type GetInstanceInternalServerError struct {
	Payload string
}

func (o *GetInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /instances/{cluster}/{arn}][%d] getInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
