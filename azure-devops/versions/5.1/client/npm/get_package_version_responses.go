// Code generated by go-swagger; DO NOT EDIT.

package npm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// GetPackageVersionReader is a Reader for the GetPackageVersion structure.
type GetPackageVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackageVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPackageVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPackageVersionOK creates a GetPackageVersionOK with default headers values
func NewGetPackageVersionOK() *GetPackageVersionOK {
	return &GetPackageVersionOK{}
}

/*GetPackageVersionOK handles this case with default header values.

successful operation
*/
type GetPackageVersionOK struct {
	Payload *models.Package
}

func (o *GetPackageVersionOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/_apis/packaging/feeds/{feedId}/npm/{packageName}/versions/{packageVersion}][%d] getPackageVersionOK  %+v", 200, o.Payload)
}

func (o *GetPackageVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Package)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
