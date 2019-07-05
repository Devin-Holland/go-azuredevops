// Code generated by go-swagger; DO NOT EDIT.

package work_item_types_behaviors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveBehaviorFromWorkItemTypeReader is a Reader for the RemoveBehaviorFromWorkItemType structure.
type RemoveBehaviorFromWorkItemTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveBehaviorFromWorkItemTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRemoveBehaviorFromWorkItemTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveBehaviorFromWorkItemTypeOK creates a RemoveBehaviorFromWorkItemTypeOK with default headers values
func NewRemoveBehaviorFromWorkItemTypeOK() *RemoveBehaviorFromWorkItemTypeOK {
	return &RemoveBehaviorFromWorkItemTypeOK{}
}

/*RemoveBehaviorFromWorkItemTypeOK handles this case with default header values.

successful operation
*/
type RemoveBehaviorFromWorkItemTypeOK struct {
}

func (o *RemoveBehaviorFromWorkItemTypeOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/_apis/work/processes/{processId}/workitemtypesbehaviors/{witRefNameForBehaviors}/behaviors/{behaviorRefName}][%d] removeBehaviorFromWorkItemTypeOK ", 200)
}

func (o *RemoveBehaviorFromWorkItemTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
