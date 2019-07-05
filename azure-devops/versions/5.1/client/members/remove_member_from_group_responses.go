// Code generated by go-swagger; DO NOT EDIT.

package members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveMemberFromGroupReader is a Reader for the RemoveMemberFromGroup structure.
type RemoveMemberFromGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveMemberFromGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveMemberFromGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveMemberFromGroupOK creates a RemoveMemberFromGroupOK with default headers values
func NewRemoveMemberFromGroupOK() *RemoveMemberFromGroupOK {
	return &RemoveMemberFromGroupOK{}
}

/*RemoveMemberFromGroupOK handles this case with default header values.

successful operation
*/
type RemoveMemberFromGroupOK struct {
}

func (o *RemoveMemberFromGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/_apis/GroupEntitlements/{groupId}/members/{memberId}][%d] removeMemberFromGroupOK ", 200)
}

func (o *RemoveMemberFromGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}