// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetProjectPropertiesReader is a Reader for the SetProjectProperties structure.
type SetProjectPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetProjectPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetProjectPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetProjectPropertiesOK creates a SetProjectPropertiesOK with default headers values
func NewSetProjectPropertiesOK() *SetProjectPropertiesOK {
	return &SetProjectPropertiesOK{}
}

/*SetProjectPropertiesOK handles this case with default header values.

successful operation
*/
type SetProjectPropertiesOK struct {
}

func (o *SetProjectPropertiesOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/_apis/projects/{projectId}/properties][%d] setProjectPropertiesOK ", 200)
}

func (o *SetProjectPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
