// Code generated by go-swagger; DO NOT EDIT.

package controls

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveControlFromGroupReader is a Reader for the RemoveControlFromGroup structure.
type RemoveControlFromGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveControlFromGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveControlFromGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveControlFromGroupOK creates a RemoveControlFromGroupOK with default headers values
func NewRemoveControlFromGroupOK() *RemoveControlFromGroupOK {
	return &RemoveControlFromGroupOK{}
}

/*RemoveControlFromGroupOK handles this case with default header values.

successful operation
*/
type RemoveControlFromGroupOK struct {
}

func (o *RemoveControlFromGroupOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/_apis/work/processes/{processId}/workItemTypes/{witRefName}/layout/groups/{groupId}/controls/{controlId}][%d] removeControlFromGroupOK ", 200)
}

func (o *RemoveControlFromGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
