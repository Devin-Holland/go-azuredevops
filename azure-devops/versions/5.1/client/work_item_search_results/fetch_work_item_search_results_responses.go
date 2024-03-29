// Code generated by go-swagger; DO NOT EDIT.

package work_item_search_results

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// FetchWorkItemSearchResultsReader is a Reader for the FetchWorkItemSearchResults structure.
type FetchWorkItemSearchResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FetchWorkItemSearchResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFetchWorkItemSearchResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewFetchWorkItemSearchResultsOK creates a FetchWorkItemSearchResultsOK with default headers values
func NewFetchWorkItemSearchResultsOK() *FetchWorkItemSearchResultsOK {
	return &FetchWorkItemSearchResultsOK{}
}

/*FetchWorkItemSearchResultsOK handles this case with default header values.

successful operation
*/
type FetchWorkItemSearchResultsOK struct {
	Payload *models.WorkItemSearchResponse
}

func (o *FetchWorkItemSearchResultsOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/{project}/_apis/search/workitemsearchresults][%d] fetchWorkItemSearchResultsOK  %+v", 200, o.Payload)
}

func (o *FetchWorkItemSearchResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkItemSearchResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
