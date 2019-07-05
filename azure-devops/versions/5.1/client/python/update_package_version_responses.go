// Code generated by go-swagger; DO NOT EDIT.

package python

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdatePackageVersionReader is a Reader for the UpdatePackageVersion structure.
type UpdatePackageVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePackageVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdatePackageVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePackageVersionOK creates a UpdatePackageVersionOK with default headers values
func NewUpdatePackageVersionOK() *UpdatePackageVersionOK {
	return &UpdatePackageVersionOK{}
}

/*UpdatePackageVersionOK handles this case with default header values.

successful operation
*/
type UpdatePackageVersionOK struct {
}

func (o *UpdatePackageVersionOK) Error() string {
	return fmt.Sprintf("[PATCH /{organization}/{project}/_apis/packaging/feeds/{feedId}/pypi/packages/{packageName}/versions/{packageVersion}][%d] updatePackageVersionOK ", 200)
}

func (o *UpdatePackageVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
