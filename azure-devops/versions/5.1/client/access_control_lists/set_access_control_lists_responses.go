// Code generated by go-swagger; DO NOT EDIT.

package access_control_lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// SetAccessControlListsReader is a Reader for the SetAccessControlLists structure.
type SetAccessControlListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAccessControlListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetAccessControlListsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAccessControlListsOK creates a SetAccessControlListsOK with default headers values
func NewSetAccessControlListsOK() *SetAccessControlListsOK {
	return &SetAccessControlListsOK{}
}

/*SetAccessControlListsOK handles this case with default header values.

successful operation
*/
type SetAccessControlListsOK struct {
}

func (o *SetAccessControlListsOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/_apis/accesscontrollists/{securityNamespaceId}][%d] setAccessControlListsOK ", 200)
}

func (o *SetAccessControlListsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}