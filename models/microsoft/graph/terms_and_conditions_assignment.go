package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TermsAndConditionsAssignment 
type TermsAndConditionsAssignment struct {
    Entity
    // Assignment target that the T&C policy is assigned to.
    target *DeviceAndAppManagementAssignmentTarget;
}
// NewTermsAndConditionsAssignment instantiates a new termsAndConditionsAssignment and sets the default values.
func NewTermsAndConditionsAssignment()(*TermsAndConditionsAssignment) {
    m := &TermsAndConditionsAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetTarget gets the target property value. Assignment target that the T&C policy is assigned to.
func (m *TermsAndConditionsAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TermsAndConditionsAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["target"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementAssignmentTarget() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTarget(val.(*DeviceAndAppManagementAssignmentTarget))
        }
        return nil
    }
    return res
}
func (m *TermsAndConditionsAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TermsAndConditionsAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetTarget sets the target property value. Assignment target that the T&C policy is assigned to.
func (m *TermsAndConditionsAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    m.target = value
}
