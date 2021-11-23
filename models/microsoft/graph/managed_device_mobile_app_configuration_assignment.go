package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// managedDeviceMobileAppConfigurationAssignment 
type ManagedDeviceMobileAppConfigurationAssignment struct {
    Entity
    // Assignment target that the T&C policy is assigned to.
    target *DeviceAndAppManagementAssignmentTarget;
}
// NewManagedDeviceMobileAppConfigurationAssignment instantiates a new managedDeviceMobileAppConfigurationAssignment and sets the default values.
func NewManagedDeviceMobileAppConfigurationAssignment()(*ManagedDeviceMobileAppConfigurationAssignment) {
    m := &ManagedDeviceMobileAppConfigurationAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetTarget gets the target property value. Assignment target that the T&C policy is assigned to.
func (m *ManagedDeviceMobileAppConfigurationAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceMobileAppConfigurationAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
func (m *ManagedDeviceMobileAppConfigurationAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceMobileAppConfigurationAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
func (m *ManagedDeviceMobileAppConfigurationAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    m.target = value
}
