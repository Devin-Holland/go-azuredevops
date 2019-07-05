// Code generated by go-swagger; DO NOT EDIT.

package workitemsorder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// ReorderBacklogWorkItemsReader is a Reader for the ReorderBacklogWorkItems structure.
type ReorderBacklogWorkItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReorderBacklogWorkItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewReorderBacklogWorkItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReorderBacklogWorkItemsOK creates a ReorderBacklogWorkItemsOK with default headers values
func NewReorderBacklogWorkItemsOK() *ReorderBacklogWorkItemsOK {
	return &ReorderBacklogWorkItemsOK{}
}

/*ReorderBacklogWorkItemsOK handles this case with default header values.

successful operation
*/
type ReorderBacklogWorkItemsOK struct {
	Payload []*models.ReorderResult
}

func (o *ReorderBacklogWorkItemsOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/{project}/{team}/_apis/work/workitemsorder][%d] reorderBacklogWorkItemsOK  %+v", 200, o.Payload)
}

func (o *ReorderBacklogWorkItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}