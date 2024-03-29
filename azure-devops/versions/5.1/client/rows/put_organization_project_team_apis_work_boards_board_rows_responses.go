// Code generated by go-swagger; DO NOT EDIT.

package rows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "azure-devops/5.1/models"
)

// PutOrganizationProjectTeamApisWorkBoardsBoardRowsReader is a Reader for the PutOrganizationProjectTeamApisWorkBoardsBoardRows structure.
type PutOrganizationProjectTeamApisWorkBoardsBoardRowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrganizationProjectTeamApisWorkBoardsBoardRowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutOrganizationProjectTeamApisWorkBoardsBoardRowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutOrganizationProjectTeamApisWorkBoardsBoardRowsOK creates a PutOrganizationProjectTeamApisWorkBoardsBoardRowsOK with default headers values
func NewPutOrganizationProjectTeamApisWorkBoardsBoardRowsOK() *PutOrganizationProjectTeamApisWorkBoardsBoardRowsOK {
	return &PutOrganizationProjectTeamApisWorkBoardsBoardRowsOK{}
}

/*PutOrganizationProjectTeamApisWorkBoardsBoardRowsOK handles this case with default header values.

successful operation
*/
type PutOrganizationProjectTeamApisWorkBoardsBoardRowsOK struct {
	Payload []*models.BoardRow
}

func (o *PutOrganizationProjectTeamApisWorkBoardsBoardRowsOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/{project}/{team}/_apis/work/boards/{board}/rows][%d] putOrganizationProjectTeamApisWorkBoardsBoardRowsOK  %+v", 200, o.Payload)
}

func (o *PutOrganizationProjectTeamApisWorkBoardsBoardRowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
