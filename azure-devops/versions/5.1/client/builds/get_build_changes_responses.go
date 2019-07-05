// Code generated by go-swagger; DO NOT EDIT.

package builds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetBuildChangesReader is a Reader for the GetBuildChanges structure.
type GetBuildChangesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildChangesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBuildChangesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBuildChangesOK creates a GetBuildChangesOK with default headers values
func NewGetBuildChangesOK() *GetBuildChangesOK {
	return &GetBuildChangesOK{}
}

/*GetBuildChangesOK handles this case with default header values.

successful operation
*/
type GetBuildChangesOK struct {
	Payload []*models.Change
}

func (o *GetBuildChangesOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/build/builds/{buildId}/changes][%d] getBuildChangesOK  %+v", 200, o.Payload)
}

func (o *GetBuildChangesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}