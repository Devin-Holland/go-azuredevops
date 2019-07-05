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

// KubernetesResource kubernetes resource
// swagger:model KubernetesResource
type KubernetesResource struct {
	EnvironmentResource

	// cluster name
	ClusterName string `json:"clusterName,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// service endpoint Id
	// Format: uuid
	ServiceEndpointID strfmt.UUID `json:"serviceEndpointId,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *KubernetesResource) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 EnvironmentResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.EnvironmentResource = aO0

	// now for regular properties
	var propsKubernetesResource struct {
		ClusterName string `json:"clusterName,omitempty"`

		Namespace string `json:"namespace,omitempty"`

		ServiceEndpointID strfmt.UUID `json:"serviceEndpointId,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsKubernetesResource); err != nil {
		return err
	}
	m.ClusterName = propsKubernetesResource.ClusterName

	m.Namespace = propsKubernetesResource.Namespace

	m.ServiceEndpointID = propsKubernetesResource.ServiceEndpointID

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m KubernetesResource) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.EnvironmentResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsKubernetesResource struct {
		ClusterName string `json:"clusterName,omitempty"`

		Namespace string `json:"namespace,omitempty"`

		ServiceEndpointID strfmt.UUID `json:"serviceEndpointId,omitempty"`
	}
	propsKubernetesResource.ClusterName = m.ClusterName

	propsKubernetesResource.Namespace = m.Namespace

	propsKubernetesResource.ServiceEndpointID = m.ServiceEndpointID

	jsonDataPropsKubernetesResource, errKubernetesResource := swag.WriteJSON(propsKubernetesResource)
	if errKubernetesResource != nil {
		return nil, errKubernetesResource
	}
	_parts = append(_parts, jsonDataPropsKubernetesResource)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this kubernetes resource
func (m *KubernetesResource) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with EnvironmentResource
	if err := m.EnvironmentResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceEndpointID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesResource) validateServiceEndpointID(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceEndpointID) { // not required
		return nil
	}

	if err := validate.FormatOf("serviceEndpointId", "body", "uuid", m.ServiceEndpointID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesResource) UnmarshalBinary(b []byte) error {
	var res KubernetesResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
