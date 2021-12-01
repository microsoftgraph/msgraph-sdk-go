package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportedWindowsAutopilotDeviceIdentity 
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
// NewImportedWindowsAutopilotDeviceIdentity instantiates a new importedWindowsAutopilotDeviceIdentity and sets the default values.
func NewImportedWindowsAutopilotDeviceIdentity()(*ImportedWindowsAutopilotDeviceIdentity) {
    m := &ImportedWindowsAutopilotDeviceIdentity{
        Entity: *NewEntity(),
    }
    return m
}
// GetAssignedUserPrincipalName gets the assignedUserPrincipalName property value. UPN of the user the device will be assigned
func (m *ImportedWindowsAutopilotDeviceIdentity) GetAssignedUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.assignedUserPrincipalName
    }
}
// GetGroupTag gets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetGroupTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.groupTag
    }
}
// GetHardwareIdentifier gets the hardwareIdentifier property value. Hardware Blob of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetHardwareIdentifier()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.hardwareIdentifier
    }
}
// GetImportId gets the importId property value. The Import Id of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetImportId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.importId
    }
}
// GetProductKey gets the productKey property value. Product Key of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetProductKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.productKey
    }
}
// GetSerialNumber gets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetSerialNumber()(*string) {
    if m == nil {
        return nil
    } else {
        return m.serialNumber
    }
}
// GetState gets the state property value. Current state of the imported device.
func (m *ImportedWindowsAutopilotDeviceIdentity) GetState()(*ImportedWindowsAutopilotDeviceIdentityState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ImportedWindowsAutopilotDeviceIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["assignedUserPrincipalName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAssignedUserPrincipalName(val)
        }
        return nil
    }
    res["groupTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetGroupTag(val)
        }
        return nil
    }
    res["hardwareIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHardwareIdentifier(val)
        }
        return nil
    }
    res["importId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportId(val)
        }
        return nil
    }
    res["productKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductKey(val)
        }
        return nil
    }
    res["serialNumber"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSerialNumber(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImportedWindowsAutopilotDeviceIdentityState() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ImportedWindowsAutopilotDeviceIdentityState))
        }
        return nil
    }
    return res
}
func (m *ImportedWindowsAutopilotDeviceIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAssignedUserPrincipalName sets the assignedUserPrincipalName property value. UPN of the user the device will be assigned
func (m *ImportedWindowsAutopilotDeviceIdentity) SetAssignedUserPrincipalName(value *string)() {
    if m != nil {
        m.assignedUserPrincipalName = value
    }
}
// SetGroupTag sets the groupTag property value. Group Tag of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetGroupTag(value *string)() {
    if m != nil {
        m.groupTag = value
    }
}
// SetHardwareIdentifier sets the hardwareIdentifier property value. Hardware Blob of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetHardwareIdentifier(value []byte)() {
    if m != nil {
        m.hardwareIdentifier = value
    }
}
// SetImportId sets the importId property value. The Import Id of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetImportId(value *string)() {
    if m != nil {
        m.importId = value
    }
}
// SetProductKey sets the productKey property value. Product Key of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetProductKey(value *string)() {
    if m != nil {
        m.productKey = value
    }
}
// SetSerialNumber sets the serialNumber property value. Serial number of the Windows autopilot device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetSerialNumber(value *string)() {
    if m != nil {
        m.serialNumber = value
    }
}
// SetState sets the state property value. Current state of the imported device.
func (m *ImportedWindowsAutopilotDeviceIdentity) SetState(value *ImportedWindowsAutopilotDeviceIdentityState)() {
    if m != nil {
        m.state = value
    }
}
