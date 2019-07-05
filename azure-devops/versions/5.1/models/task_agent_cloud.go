// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskAgentCloud task agent cloud
// swagger:model TaskAgentCloud
type TaskAgentCloud struct {

	// Gets or sets a AcquireAgentEndpoint using which a request can be made to acquire new agent
	AcquireAgentEndpoint string `json:"acquireAgentEndpoint,omitempty"`

	// acquisition timeout
	AcquisitionTimeout int32 `json:"acquisitionTimeout,omitempty"`

	// agent cloud Id
	AgentCloudID int32 `json:"agentCloudId,omitempty"`

	// get account parallelism endpoint
	GetAccountParallelismEndpoint string `json:"getAccountParallelismEndpoint,omitempty"`

	// get agent definition endpoint
	GetAgentDefinitionEndpoint string `json:"getAgentDefinitionEndpoint,omitempty"`

	// get agent request status endpoint
	GetAgentRequestStatusEndpoint string `json:"getAgentRequestStatusEndpoint,omitempty"`

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// Signifies that this Agent Cloud is internal and should not be user-manageable
	Internal bool `json:"internal,omitempty"`

	// max parallelism
	MaxParallelism int32 `json:"maxParallelism,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// release agent endpoint
	ReleaseAgentEndpoint string `json:"releaseAgentEndpoint,omitempty"`

	// shared secret
	SharedSecret string `json:"sharedSecret,omitempty"`

	// Gets or sets the type of the endpoint.
	Type string `json:"type,omitempty"`
}

// Validate validates this task agent cloud
func (m *TaskAgentCloud) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskAgentCloud) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskAgentCloud) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskAgentCloud) UnmarshalBinary(b []byte) error {
	var res TaskAgentCloud
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
