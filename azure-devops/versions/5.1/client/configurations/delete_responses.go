// Code generated by go-swagger; DO NOT EDIT.

package configurations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteReader is a Reader for the Delete structure.
type DeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNoContent creates a DeleteNoContent with default headers values
func NewDeleteNoContent() *DeleteNoContent {
	return &DeleteNoContent{}
}

/*DeleteNoContent handles this case with default header values.

NoContent
*/
type DeleteNoContent struct {
}

func (o *DeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/{project}/_apis/policy/configurations/{configurationId}][%d] deleteNoContent ", 204)
}

func (o *DeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
