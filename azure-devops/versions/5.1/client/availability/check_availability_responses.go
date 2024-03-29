// Code generated by go-swagger; DO NOT EDIT.

package availability

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CheckAvailabilityReader is a Reader for the CheckAvailability structure.
type CheckAvailabilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckAvailabilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCheckAvailabilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckAvailabilityOK creates a CheckAvailabilityOK with default headers values
func NewCheckAvailabilityOK() *CheckAvailabilityOK {
	return &CheckAvailabilityOK{}
}

/*CheckAvailabilityOK handles this case with default header values.

The symbol service is available
*/
type CheckAvailabilityOK struct {
}

func (o *CheckAvailabilityOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/_apis/symbol/availability][%d] checkAvailabilityOK ", 200)
}

func (o *CheckAvailabilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
