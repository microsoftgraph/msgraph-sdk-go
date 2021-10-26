package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ImportedWindowsAutopilotDeviceIdentity struct {
    Entity
    // UPN of the user the device will be assigned
    assignedUserPrincipalName *string;
    // Group Tag of the Windows autopilot device.
    groupTag *string;
    // Hardware Blob of the Windows autopilot device.
    hardwareIdentifier []byte;
    // The Import Id of the Windows autopilot device.
    importId *string;
    // Product Key of the Windows autopilot device.
    productKey *string;
    // Serial number of the Windows autopilot device.
    serialNumber *string;
    // Current state of the imported device.
    state *ImportedWindowsAutopilotDeviceIdentityState;
}
// Instantiates a new importedWindowsAutopilotDeviceIdentity and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentity()(*ImportedWindowsAutopilotDeviceIdentity) {
    m := &ImportedWindowsAutopilotDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the assignedUserPrincipalName property value. UPN of the user the device will be assigned
func (m *ImportedWindowsAutopilotDeviceIdentity) GetAssignedUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedUserPrincipalName
    }
}
// Gets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetGroupTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupTag
    }
}
// Gets the hardwareIdentifier property value. Hardware Blob of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetHardwareIdentifier()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.hardwareIdentifier
    }
}
// Gets the importId property value. The Import Id of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetImportId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.importId
    }
}
// Gets the productKey property value. Product Key of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetProductKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productKey
    }
}
// Gets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// Gets the state property value. Current state of the imported device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetState()(*ImportedWindowsAutopilotDeviceIdentityState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// The deserialization information for the current model
func (m *ImportedWindowsAutopilotDeviceIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAssignedUserPrincipalName(val)
        return nil
    }
    res["groupTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGroupTag(val)
        return nil
    }
    res["hardwareIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        m.SetHardwareIdentifier(val)
        return nil
    }
    res["importId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetImportId(val)
        return nil
    }
    res["productKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetProductKey(val)
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSerialNumber(val)
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImportedWindowsAutopilotDeviceIdentityState() })
        if err != nil {
            return err
        }
        m.SetState(val.(*ImportedWindowsAutopilotDeviceIdentityState))
        return nil
    }
    return res
}
func (m *ImportedWindowsAutopilotDeviceIdentity) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ImportedWindowsAutopilotDeviceIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("assignedUserPrincipalName", m.GetAssignedUserPrincipalName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("groupTag", m.GetGroupTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("hardwareIdentifier", m.GetHardwareIdentifier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("importId", m.GetImportId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("productKey", m.GetProductKey())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("serialNumber", m.GetSerialNumber())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("state", m.GetState())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the assignedUserPrincipalName property value. UPN of the user the device will be assigned
// Parameters:
//  - value : Value to set for the assignedUserPrincipalName property.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetAssignedUserPrincipalName(value *string)() {
    m.assignedUserPrincipalName = value
}
// Sets the groupTag property value. Group Tag of the Windows autopilot device.
// Parameters:
//  - value : Value to set for the groupTag property.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetGroupTag(value *string)() {
    m.groupTag = value
}
// Sets the hardwareIdentifier property value. Hardware Blob of the Windows autopilot device.
// Parameters:
//  - value : Value to set for the hardwareIdentifier property.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetHardwareIdentifier(value []byte)() {
    m.hardwareIdentifier = value
}
// Sets the importId property value. The Import Id of the Windows autopilot device.
// Parameters:
//  - value : Value to set for the importId property.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetImportId(value *string)() {
    m.importId = value
}
// Sets the productKey property value. Product Key of the Windows autopilot device.
// Parameters:
//  - value : Value to set for the productKey property.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetProductKey(value *string)() {
    m.productKey = value
}
// Sets the serialNumber property value. Serial number of the Windows autopilot device.
// Parameters:
//  - value : Value to set for the serialNumber property.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetSerialNumber(value *string)() {
    m.serialNumber = value
}
// Sets the state property value. Current state of the imported device.
// Parameters:
//  - value : Value to set for the state property.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetState(value *ImportedWindowsAutopilotDeviceIdentityState)() {
    m.state = value
}
