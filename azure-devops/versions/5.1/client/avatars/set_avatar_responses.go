// Code generated by go-swagger; DO NOT EDIT.

package avatars

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetAvatarReader is a Reader for the SetAvatar structure.
type SetAvatarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAvatarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetAvatarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAvatarOK creates a SetAvatarOK with default headers values
func NewSetAvatarOK() *SetAvatarOK {
	return &SetAvatarOK{}
}

/*SetAvatarOK handles this case with default header values.

successful operation
*/
type SetAvatarOK struct {
}

func (o *SetAvatarOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/_apis/graph/Subjects/{subjectDescriptor}/avatars][%d] setAvatarOK ", 200)
}

func (o *SetAvatarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
