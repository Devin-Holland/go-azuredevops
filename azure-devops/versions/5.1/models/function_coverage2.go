// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// FunctionCoverage2 function coverage2
// swagger:model FunctionCoverage2
type FunctionCoverage2 struct {

	// blocks covered
	BlocksCovered int32 `json:"blocksCovered,omitempty"`

	// blocks not covered
	BlocksNotCovered int32 `json:"blocksNotCovered,omitempty"`

	// class
	Class string `json:"class,omitempty"`

	// coverage Id
	CoverageID int32 `json:"coverageId,omitempty"`

	// function Id
	FunctionID int32 `json:"functionId,omitempty"`

	// lines covered
	LinesCovered int32 `json:"linesCovered,omitempty"`

	// lines not covered
	LinesNotCovered int32 `json:"linesNotCovered,omitempty"`

	// lines partially covered
	LinesPartiallyCovered int32 `json:"linesPartiallyCovered,omitempty"`

	// module Id
	ModuleID int32 `json:"moduleId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// source file
	SourceFile string `json:"sourceFile,omitempty"`
}

// Validate validates this function coverage2
func (m *FunctionCoverage2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FunctionCoverage2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FunctionCoverage2) UnmarshalBinary(b []byte) error {
	var res FunctionCoverage2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
