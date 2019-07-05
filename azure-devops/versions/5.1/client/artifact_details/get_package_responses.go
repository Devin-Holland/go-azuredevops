// Code generated by go-swagger; DO NOT EDIT.

package artifact_details

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetPackageReader is a Reader for the GetPackage structure.
type GetPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPackageOK creates a GetPackageOK with default headers values
func NewGetPackageOK() *GetPackageOK {
	return &GetPackageOK{}
}

/*GetPackageOK handles this case with default header values.

successful operation
*/
type GetPackageOK struct {
	Payload *models.Package
}

func (o *GetPackageOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/_apis/packaging/Feeds/{feedId}/packages/{packageId}][%d] getPackageOK  %+v", 200, o.Payload)
}

func (o *GetPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Package)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}