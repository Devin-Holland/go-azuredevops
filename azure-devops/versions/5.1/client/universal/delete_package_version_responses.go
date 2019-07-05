// Code generated by go-swagger; DO NOT EDIT.

package universal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// DeletePackageVersionReader is a Reader for the DeletePackageVersion structure.
type DeletePackageVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePackageVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeletePackageVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePackageVersionOK creates a DeletePackageVersionOK with default headers values
func NewDeletePackageVersionOK() *DeletePackageVersionOK {
	return &DeletePackageVersionOK{}
}

/*DeletePackageVersionOK handles this case with default header values.

successful operation
*/
type DeletePackageVersionOK struct {
	Payload *models.Package
}

func (o *DeletePackageVersionOK) Error() string {
	return fmt.Sprintf("[DELETE /{organization}/_apis/packaging/feeds/{feedId}/upack/packages/{packageName}/versions/{packageVersion}][%d] deletePackageVersionOK  %+v", 200, o.Payload)
}

func (o *DeletePackageVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Package)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
