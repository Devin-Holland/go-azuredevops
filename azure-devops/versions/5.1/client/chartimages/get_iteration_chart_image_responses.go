// Code generated by go-swagger; DO NOT EDIT.

package chartimages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetIterationChartImageReader is a Reader for the GetIterationChartImage structure.
type GetIterationChartImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIterationChartImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIterationChartImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIterationChartImageOK creates a GetIterationChartImageOK with default headers values
func NewGetIterationChartImageOK() *GetIterationChartImageOK {
	return &GetIterationChartImageOK{}
}

/*GetIterationChartImageOK handles this case with default header values.

successful operation
*/
type GetIterationChartImageOK struct {
	Payload string
}

func (o *GetIterationChartImageOK) Error() string {
	return fmt.Sprintf("[GET /{organization}/{project}/{team}/_apis/work/iterations/{iterationId}/chartimages/{name}][%d] getIterationChartImageOK  %+v", 200, o.Payload)
}

func (o *GetIterationChartImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
