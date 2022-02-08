package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedEBookAssignment 
type ManagedEBookAssignment struct {
    Entity
    // The install intent for eBook. Possible values are: available, required, uninstall, availableWithoutEnrollment.
    installIntent *InstallIntent;
    // The assignment target for eBook.
    target *DeviceAndAppManagementAssignmentTarget;
}
// NewManagedEBookAssignment instantiates a new managedEBookAssignment and sets the default values.
func NewManagedEBookAssignment()(*ManagedEBookAssignment) {
    m := &ManagedEBookAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// GetInstallIntent gets the installIntent property value. The install intent for eBook. Possible values are: available, required, uninstall, availableWithoutEnrollment.
func (m *ManagedEBookAssignment) GetInstallIntent()(*InstallIntent) {
    if m == nil {
        return nil
    } else {
        return m.installIntent
    }
}
// GetTarget gets the target property value. The assignment target for eBook.
func (m *ManagedEBookAssignment) GetTarget()(*DeviceAndAppManagementAssignmentTarget) {
    if m == nil {
        return nil
    } else {
        return m.target
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedEBookAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["installIntent"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallIntent)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallIntent(val.(*InstallIntent))
        }
        return nil
    }
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
func (m *ManagedEBookAssignment) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedEBookAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetInstallIntent() != nil {
        cast := (*m.GetInstallIntent()).String()
        err = writer.WriteStringValue("installIntent", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("target", m.GetTarget())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetInstallIntent sets the installIntent property value. The install intent for eBook. Possible values are: available, required, uninstall, availableWithoutEnrollment.
func (m *ManagedEBookAssignment) SetInstallIntent(value *InstallIntent)() {
    if m != nil {
        m.installIntent = value
    }
}
// SetTarget sets the target property value. The assignment target for eBook.
func (m *ManagedEBookAssignment) SetTarget(value *DeviceAndAppManagementAssignmentTarget)() {
    if m != nil {
        m.target = value
    }
}
