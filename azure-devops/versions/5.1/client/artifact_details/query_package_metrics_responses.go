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

// QueryPackageMetricsReader is a Reader for the QueryPackageMetrics structure.
type QueryPackageMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryPackageMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewQueryPackageMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQueryPackageMetricsOK creates a QueryPackageMetricsOK with default headers values
func NewQueryPackageMetricsOK() *QueryPackageMetricsOK {
	return &QueryPackageMetricsOK{}
}

/*QueryPackageMetricsOK handles this case with default header values.

successful operation
*/
type QueryPackageMetricsOK struct {
	Payload []*models.PackageMetrics
}

func (o *QueryPackageMetricsOK) Error() string {
	return fmt.Sprintf("[POST /{organization}/{project}/_apis/packaging/Feeds/{feedId}/packagemetricsbatch][%d] queryPackageMetricsOK  %+v", 200, o.Payload)
}

func (o *QueryPackageMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
