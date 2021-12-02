package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppStatus 
type ManagedAppStatus struct {
    Entity
    // Friendly name of the status report.
    displayName *string;
    // Version of the entity.
    version *string;
}
// NewManagedAppStatus instantiates a new managedAppStatus and sets the default values.
func NewManagedAppStatus()(*ManagedAppStatus) {
    m := &ManagedAppStatus{
        Entity: *NewEntity(),
    }
    return m
}
// GetDisplayName gets the displayName property value. Friendly name of the status report.
func (m *ManagedAppStatus) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetVersion gets the version property value. Version of the entity.
func (m *ManagedAppStatus) GetVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.version
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppStatus) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["version"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    return res
}
func (m *ManagedAppStatus) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedAppStatus) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDisplayName sets the displayName property value. Friendly name of the status report.
func (m *ManagedAppStatus) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetVersion sets the version property value. Version of the entity.
func (m *ManagedAppStatus) SetVersion(value *string)() {
    if m != nil {
        m.version = value
    }
}
