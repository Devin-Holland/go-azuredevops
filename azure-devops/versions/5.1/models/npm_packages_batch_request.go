// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NpmPackagesBatchRequest A batch of operations to apply to package versions.
// swagger:model NpmPackagesBatchRequest
type NpmPackagesBatchRequest struct {

	// Data required to perform the operation. This is optional based on type of operation. Use BatchPromoteData if performing a promote operation.
	Data BatchOperationData `json:"data,omitempty"`

	// Type of operation that needs to be performed on packages.
	// Enum: [promote deprecate unpublish permanentDelete restoreToFeed delete]
	Operation interface{} `json:"operation,omitempty"`

	// The packages onto which the operation will be performed.
	Packages []*MinimalPackageDetails `json:"packages"`
}

// Validate validates this npm packages batch request
func (m *NpmPackagesBatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NpmPackagesBatchRequest) validatePackages(formats strfmt.Registry) error {

	if swag.IsZero(m.Packages) { // not required
		return nil
	}

	for i := 0; i < len(m.Packages); i++ {
		if swag.IsZero(m.Packages[i]) { // not required
			continue
		}

		if m.Packages[i] != nil {
			if err := m.Packages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NpmPackagesBatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NpmPackagesBatchRequest) UnmarshalBinary(b []byte) error {
	var res NpmPackagesBatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}