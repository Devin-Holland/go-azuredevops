// Code generated by go-swagger; DO NOT EDIT.

package npm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DownloadScopedPackageReader is a Reader for the DownloadScopedPackage structure.
type DownloadScopedPackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DownloadScopedPackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDownloadScopedPackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDownloadScopedPackageOK creates a DownloadScopedPackageOK with default headers values
func NewDownloadScopedPackageOK() *DownloadScopedPackageOK {
	return &DownloadScopedPackageOK{}
}

/*DownloadScopedPackageOK handles this case with default header values.

successful operation
*/
type DownloadScopedPackageOK struct {
	Payload string
}

func (o *DownloadScopedPackageOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/_apis/packaging/feeds/{feedId}/npm/packages/@{packageScope}/{unscopedPackageName}/versions/{packageVersion}/content][%d] downloadScopedPackageOK  %+v", 200, o.Payload)
}

func (o *DownloadScopedPackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}