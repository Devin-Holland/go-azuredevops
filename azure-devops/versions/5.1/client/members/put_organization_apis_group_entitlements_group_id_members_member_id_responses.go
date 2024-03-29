// Code generated by go-swagger; DO NOT EDIT.

package members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDReader is a Reader for the PutOrganizationApisGroupEntitlementsGroupIDMembersMemberID structure.
type PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK creates a PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK with default headers values
func NewPutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK() *PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK {
	return &PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK{}
}

/*PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK handles this case with default header values.

successful operation
*/
type PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK struct {
}

func (o *PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK) Error() string {
	return fmt.Sprintf("[PUT /{organization}/_apis/GroupEntitlements/{groupId}/members/{memberId}][%d] putOrganizationApisGroupEntitlementsGroupIdMembersMemberIdOK ", 200)
}

func (o *PutOrganizationApisGroupEntitlementsGroupIDMembersMemberIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
