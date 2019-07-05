// Code generated by go-swagger; DO NOT EDIT.

package avatar

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveProjectAvatarReader is a Reader for the RemoveProjectAvatar structure.
type RemoveProjectAvatarReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveProjectAvatarReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveProjectAvatarOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveProjectAvatarOK creates a RemoveProjectAvatarOK with default headers values
func NewRemoveProjectAvatarOK() *RemoveProjectAvatarOK {
	return &RemoveProjectAvatarOK{}
}

/*RemoveProjectAvatarOK handles this case with default header values.

successful operation
*/
type RemoveProjectAvatarOK struct {
}

func (o *RemoveProjectAvatarOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/_apis/projects/{projectId}/avatar][%d] removeProjectAvatarOK ", 200)
}

func (o *RemoveProjectAvatarOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}