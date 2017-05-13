package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContainerInstanceVersionInfo container instance version info
// swagger:model ContainerInstanceVersionInfo
type ContainerInstanceVersionInfo struct {

	// agent hash
	AgentHash string `json:"agentHash,omitempty"`

	// agent version
	AgentVersion string `json:"agentVersion,omitempty"`

	// docker version
	DockerVersion string `json:"dockerVersion,omitempty"`
}

// Validate validates this container instance version info
func (m *ContainerInstanceVersionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerInstanceVersionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerInstanceVersionInfo) UnmarshalBinary(b []byte) error {
	var res ContainerInstanceVersionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
