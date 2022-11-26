// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EndpointChangeRequest Structure which contains the mutable elements of an Endpoint.
//
// swagger:model EndpointChangeRequest
type EndpointChangeRequest struct {

	// addressing
	Addressing *AddressPair `json:"addressing,omitempty"`

	// ID assigned by container runtime
	ContainerID string `json:"container-id,omitempty"`

	// Name assigned to container
	ContainerName string `json:"container-name,omitempty"`

	// datapath configuration
	DatapathConfiguration *EndpointDatapathConfiguration `json:"datapath-configuration,omitempty"`

	// ID of datapath tail call map
	DatapathMapID int64 `json:"datapath-map-id,omitempty"`

	// Docker endpoint ID
	DockerEndpointID string `json:"docker-endpoint-id,omitempty"`

	// Docker network ID
	DockerNetworkID string `json:"docker-network-id,omitempty"`

	// MAC address
	HostMac string `json:"host-mac,omitempty"`

	// Local endpoint ID
	ID int64 `json:"id,omitempty"`

	// Index of network device
	InterfaceIndex int64 `json:"interface-index,omitempty"`

	// Name of network device
	InterfaceName string `json:"interface-name,omitempty"`

	// Kubernetes namespace name
	K8sNamespace string `json:"k8s-namespace,omitempty"`

	// Kubernetes pod name
	K8sPodName string `json:"k8s-pod-name,omitempty"`

	// Labels describing the identity
	Labels Labels `json:"labels,omitempty"`

	// MAC address
	Mac string `json:"mac,omitempty"`

	// Process ID of the workload belonging to this endpoint
	Pid int64 `json:"pid,omitempty"`

	// Whether policy enforcement is enabled or not
	PolicyEnabled bool `json:"policy-enabled,omitempty"`

	// Current state of endpoint
	// Required: true
	State *EndpointState `json:"state"`

	// Whether to build an endpoint synchronously
	//
	SyncBuildEndpoint bool `json:"sync-build-endpoint,omitempty"`
}

// Validate validates this endpoint change request
func (m *EndpointChangeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatapathConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointChangeRequest) validateAddressing(formats strfmt.Registry) error {
	if swag.IsZero(m.Addressing) { // not required
		return nil
	}

	if m.Addressing != nil {
		if err := m.Addressing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addressing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addressing")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointChangeRequest) validateDatapathConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.DatapathConfiguration) { // not required
		return nil
	}

	if m.DatapathConfiguration != nil {
		if err := m.DatapathConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datapath-configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datapath-configuration")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointChangeRequest) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	if err := m.Labels.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labels")
		}
		return err
	}

	return nil
}

func (m *EndpointChangeRequest) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this endpoint change request based on the context it is used
func (m *EndpointChangeRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAddressing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDatapathConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointChangeRequest) contextValidateAddressing(ctx context.Context, formats strfmt.Registry) error {

	if m.Addressing != nil {
		if err := m.Addressing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addressing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("addressing")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointChangeRequest) contextValidateDatapathConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.DatapathConfiguration != nil {
		if err := m.DatapathConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("datapath-configuration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("datapath-configuration")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointChangeRequest) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Labels.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("labels")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("labels")
		}
		return err
	}

	return nil
}

func (m *EndpointChangeRequest) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointChangeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointChangeRequest) UnmarshalBinary(b []byte) error {
	var res EndpointChangeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
