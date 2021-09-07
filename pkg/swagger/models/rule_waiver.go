// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RuleWaiver Rule waiver
//
// swagger:model RuleWaiver
type RuleWaiver struct {

	// comment
	Comment string `json:"comment,omitempty"`

	// The date and time when the rule waiver was created.
	CreatedAt int64 `json:"created_at,omitempty"`

	// Principal that created the rule waiver.
	CreatedBy string `json:"created_by,omitempty"`

	// Display name of the user that created the rule waiver.
	CreatedByDisplayName string `json:"created_by_display_name,omitempty"`

	// environment id
	// Required: true
	EnvironmentID *string `json:"environment_id"`

	// environment name
	EnvironmentName string `json:"environment_name,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// name
	// Required: true
	Name *string `json:"name"`

	// resource id
	// Required: true
	ResourceID *string `json:"resource_id"`

	// resource provider
	// Required: true
	ResourceProvider *string `json:"resource_provider"`

	// resource tag
	ResourceTag string `json:"resource_tag,omitempty"`

	// resource type
	// Required: true
	ResourceType *string `json:"resource_type"`

	// Mapping of this rule in compliance families and their controls which are enabled in a given environment.
	RuleComplianceMapping interface{} `json:"rule_compliance_mapping,omitempty"`

	// Description of the rule.
	RuleDescription string `json:"rule_description,omitempty"`

	// rule id
	// Required: true
	RuleID *string `json:"rule_id"`

	// The date and time when the rule waiver was last updated.
	UpdatedAt int64 `json:"updated_at,omitempty"`

	// Principal that last updated the rule waiver.
	UpdatedBy string `json:"updated_by,omitempty"`

	// Display name of the user that last updated the rule waiver.
	UpdatedByDisplayName string `json:"updated_by_display_name,omitempty"`
}

// Validate validates this rule waiver
func (m *RuleWaiver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnvironmentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RuleWaiver) validateEnvironmentID(formats strfmt.Registry) error {

	if err := validate.Required("environment_id", "body", m.EnvironmentID); err != nil {
		return err
	}

	return nil
}

func (m *RuleWaiver) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *RuleWaiver) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RuleWaiver) validateResourceID(formats strfmt.Registry) error {

	if err := validate.Required("resource_id", "body", m.ResourceID); err != nil {
		return err
	}

	return nil
}

func (m *RuleWaiver) validateResourceProvider(formats strfmt.Registry) error {

	if err := validate.Required("resource_provider", "body", m.ResourceProvider); err != nil {
		return err
	}

	return nil
}

func (m *RuleWaiver) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("resource_type", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *RuleWaiver) validateRuleID(formats strfmt.Registry) error {

	if err := validate.Required("rule_id", "body", m.RuleID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RuleWaiver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleWaiver) UnmarshalBinary(b []byte) error {
	var res RuleWaiver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}