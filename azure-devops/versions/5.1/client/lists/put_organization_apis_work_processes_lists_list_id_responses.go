// Code generated by go-swagger; DO NOT EDIT.

package lists

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// PutOrganizationApisWorkProcessesListsListIDReader is a Reader for the PutOrganizationApisWorkProcessesListsListID structure.
type PutOrganizationApisWorkProcessesListsListIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrganizationApisWorkProcessesListsListIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutOrganizationApisWorkProcessesListsListIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutOrganizationApisWorkProcessesListsListIDOK creates a PutOrganizationApisWorkProcessesListsListIDOK with default headers values
func NewPutOrganizationApisWorkProcessesListsListIDOK() *PutOrganizationApisWorkProcessesListsListIDOK {
	return &PutOrganizationApisWorkProcessesListsListIDOK{}
}

/*PutOrganizationApisWorkProcessesListsListIDOK handles this case with default header values.

successful operation
*/
type PutOrganizationApisWorkProcessesListsListIDOK struct {
	Payload *models.PickList
}

func (o *PutOrganizationApisWorkProcessesListsListIDOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/_apis/work/processes/lists/{listId}][%d] putOrganizationApisWorkProcessesListsListIdOK  %+v", 200, o.Payload)
}

func (o *PutOrganizationApisWorkProcessesListsListIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PickList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
