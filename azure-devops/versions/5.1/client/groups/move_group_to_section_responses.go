// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// MoveGroupToSectionReader is a Reader for the MoveGroupToSection structure.
type MoveGroupToSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveGroupToSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMoveGroupToSectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMoveGroupToSectionOK creates a MoveGroupToSectionOK with default headers values
func NewMoveGroupToSectionOK() *MoveGroupToSectionOK {
	return &MoveGroupToSectionOK{}
}

/*MoveGroupToSectionOK handles this case with default header values.

successful operation
*/
type MoveGroupToSectionOK struct {
	Payload *models.Group
}

func (o *MoveGroupToSectionOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/_apis/work/processes/{processId}/workItemTypes/{witRefName}/layout/pages/{pageId}/sections/{sectionId}/groups/{groupId}][%d] moveGroupToSectionOK  %+v", 200, o.Payload)
}

func (o *MoveGroupToSectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
