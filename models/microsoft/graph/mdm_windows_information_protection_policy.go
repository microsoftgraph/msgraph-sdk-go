package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MdmWindowsInformationProtectionPolicy 
type MdmWindowsInformationProtectionPolicy struct {
    WindowsInformationProtection
}
// NewMdmWindowsInformationProtectionPolicy instantiates a new mdmWindowsInformationProtectionPolicy and sets the default values.
func NewMdmWindowsInformationProtectionPolicy()(*MdmWindowsInformationProtectionPolicy) {
    m := &MdmWindowsInformationProtectionPolicy{
        WindowsInformationProtection: *NewWindowsInformationProtection(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MdmWindowsInformationProtectionPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.WindowsInformationProtection.GetFieldDeserializers()
    return res
}
func (m *MdmWindowsInformationProtectionPolicy) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MdmWindowsInformationProtectionPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.WindowsInformationProtection.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
