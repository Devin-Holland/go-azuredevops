// Code generated by go-swagger; DO NOT EDIT.

package cardsettings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateTaskboardCardSettingsReader is a Reader for the UpdateTaskboardCardSettings structure.
type UpdateTaskboardCardSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateTaskboardCardSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateTaskboardCardSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateTaskboardCardSettingsOK creates a UpdateTaskboardCardSettingsOK with default headers values
func NewUpdateTaskboardCardSettingsOK() *UpdateTaskboardCardSettingsOK {
	return &UpdateTaskboardCardSettingsOK{}
}

/*UpdateTaskboardCardSettingsOK handles this case with default header values.

successful operation
*/
type UpdateTaskboardCardSettingsOK struct {
}

func (o *UpdateTaskboardCardSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/{project}/{team}/_apis/work/taskboard/cardsettings][%d] updateTaskboardCardSettingsOK ", 200)
}

func (o *UpdateTaskboardCardSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
