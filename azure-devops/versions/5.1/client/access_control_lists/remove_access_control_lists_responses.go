// Code generated by go-swagger; DO NOT EDIT.

package access_control_lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveAccessControlListsReader is a Reader for the RemoveAccessControlLists structure.
type RemoveAccessControlListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveAccessControlListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveAccessControlListsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveAccessControlListsOK creates a RemoveAccessControlListsOK with default headers values
func NewRemoveAccessControlListsOK() *RemoveAccessControlListsOK {
	return &RemoveAccessControlListsOK{}
}

/*RemoveAccessControlListsOK handles this case with default header values.

successful operation
*/
type RemoveAccessControlListsOK struct {
	Payload bool
}

func (o *RemoveAccessControlListsOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/_apis/accesscontrollists/{securityNamespaceId}][%d] removeAccessControlListsOK  %+v", 200, o.Payload)
}

func (o *RemoveAccessControlListsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
