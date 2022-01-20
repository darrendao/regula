// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateGroupInput create group input
//
// swagger:model CreateGroupInput
type CreateGroupInput struct {

	// environment ids
	EnvironmentIds []string `json:"environment_ids"`

	// name
	Name string `json:"name,omitempty"`

	// policy
	// Enum: [fugue:READONLY fugue:AUDITOR fugue:EDITOR fugue:CONTRIBUTOR fugue:MANAGER fugue:ORGANIZATION_REPORT_VIEWER]
	Policy string `json:"policy,omitempty"`
}

// Validate validates this create group input
func (m *CreateGroupInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createGroupInputTypePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fugue:READONLY","fugue:AUDITOR","fugue:EDITOR","fugue:CONTRIBUTOR","fugue:MANAGER","fugue:ORGANIZATION_REPORT_VIEWER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createGroupInputTypePolicyPropEnum = append(createGroupInputTypePolicyPropEnum, v)
	}
}

const (

	// CreateGroupInputPolicyFugueREADONLY captures enum value "fugue:READONLY"
	CreateGroupInputPolicyFugueREADONLY string = "fugue:READONLY"

	// CreateGroupInputPolicyFugueAUDITOR captures enum value "fugue:AUDITOR"
	CreateGroupInputPolicyFugueAUDITOR string = "fugue:AUDITOR"

	// CreateGroupInputPolicyFugueEDITOR captures enum value "fugue:EDITOR"
	CreateGroupInputPolicyFugueEDITOR string = "fugue:EDITOR"

	// CreateGroupInputPolicyFugueCONTRIBUTOR captures enum value "fugue:CONTRIBUTOR"
	CreateGroupInputPolicyFugueCONTRIBUTOR string = "fugue:CONTRIBUTOR"

	// CreateGroupInputPolicyFugueMANAGER captures enum value "fugue:MANAGER"
	CreateGroupInputPolicyFugueMANAGER string = "fugue:MANAGER"

	// CreateGroupInputPolicyFugueORGANIZATIONREPORTVIEWER captures enum value "fugue:ORGANIZATION_REPORT_VIEWER"
	CreateGroupInputPolicyFugueORGANIZATIONREPORTVIEWER string = "fugue:ORGANIZATION_REPORT_VIEWER"
)

// prop value enum
func (m *CreateGroupInput) validatePolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, createGroupInputTypePolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateGroupInput) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	// value enum
	if err := m.validatePolicyEnum("policy", "body", m.Policy); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateGroupInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGroupInput) UnmarshalBinary(b []byte) error {
	var res CreateGroupInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
