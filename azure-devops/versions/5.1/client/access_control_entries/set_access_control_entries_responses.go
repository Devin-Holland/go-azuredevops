// Code generated by go-swagger; DO NOT EDIT.

package access_control_entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// SetAccessControlEntriesReader is a Reader for the SetAccessControlEntries structure.
type SetAccessControlEntriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAccessControlEntriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSetAccessControlEntriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAccessControlEntriesOK creates a SetAccessControlEntriesOK with default headers values
func NewSetAccessControlEntriesOK() *SetAccessControlEntriesOK {
	return &SetAccessControlEntriesOK{}
}

/*SetAccessControlEntriesOK handles this case with default header values.

successful operation
*/
type SetAccessControlEntriesOK struct {
	Payload []*models.AccessControlEntry
}

func (o *SetAccessControlEntriesOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/_apis/accesscontrolentries/{securityNamespaceId}][%d] setAccessControlEntriesOK  %+v", 200, o.Payload)
}

func (o *SetAccessControlEntriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}