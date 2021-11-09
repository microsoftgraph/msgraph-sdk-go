package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceInstallState struct {
    Entity
    // Device Id.
    deviceId *string;
    // Device name.
    deviceName *string;
    // The error code for install failures.
    errorCode *string;
    // The install state of the eBook. Possible values are: notApplicable, installed, failed, notInstalled, uninstallFailed, unknown.
    installState *InstallState;
    // Last sync date and time.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // OS Description.
    osDescription *string;
    // OS Version.
    osVersion *string;
    // Device User Name.
    userName *string;
}
// Instantiates a new deviceInstallState and sets the default values.
func NewDeviceInstallState()(*DeviceInstallState) {
    m := &DeviceInstallState{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the deviceId property value. Device Id.
func (m *DeviceInstallState) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// Gets the deviceName property value. Device name.
func (m *DeviceInstallState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// Gets the errorCode property value. The error code for install failures.
func (m *DeviceInstallState) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// Gets the installState property value. The install state of the eBook. Possible values are: notApplicable, installed, failed, notInstalled, uninstallFailed, unknown.
func (m *DeviceInstallState) GetInstallState()(*InstallState) {
    if m == nil {
        return nil
    } else {
        return m.installState
    }
}
// Gets the lastSyncDateTime property value. Last sync date and time.
func (m *DeviceInstallState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// Gets the osDescription property value. OS Description.
func (m *DeviceInstallState) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// Gets the osVersion property value. OS Version.
func (m *DeviceInstallState) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// Gets the userName property value. Device User Name.
func (m *DeviceInstallState) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// The deserialization information for the current model
func (m *DeviceInstallState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["errorCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["installState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(InstallState)
            m.SetInstallState(&cast)
        }
        return nil
    }
    res["lastSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["osDescription"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["userName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    return res
}
func (m *DeviceInstallState) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceInstallState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceName", m.GetDeviceName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("errorCode", m.GetErrorCode())
        if err != nil {
            return err
        }
    }
    if m.GetInstallState() != nil {
        cast := m.GetInstallState().String()
        err = writer.WriteStringValue("installState", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastSyncDateTime", m.GetLastSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osDescription", m.GetOsDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osVersion", m.GetOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userName", m.GetUserName())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the deviceId property value. Device Id.
// Parameters:
//  - value : Value to set for the deviceId property.
func (m *DeviceInstallState) SetDeviceId(value *string)() {
    m.deviceId = value
}
// Sets the deviceName property value. Device name.
// Parameters:
//  - value : Value to set for the deviceName property.
func (m *DeviceInstallState) SetDeviceName(value *string)() {
    m.deviceName = value
}
// Sets the errorCode property value. The error code for install failures.
// Parameters:
//  - value : Value to set for the errorCode property.
func (m *DeviceInstallState) SetErrorCode(value *string)() {
    m.errorCode = value
}
// Sets the installState property value. The install state of the eBook. Possible values are: notApplicable, installed, failed, notInstalled, uninstallFailed, unknown.
// Parameters:
//  - value : Value to set for the installState property.
func (m *DeviceInstallState) SetInstallState(value *InstallState)() {
    m.installState = value
}
// Sets the lastSyncDateTime property value. Last sync date and time.
// Parameters:
//  - value : Value to set for the lastSyncDateTime property.
func (m *DeviceInstallState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastSyncDateTime = value
}
// Sets the osDescription property value. OS Description.
// Parameters:
//  - value : Value to set for the osDescription property.
func (m *DeviceInstallState) SetOsDescription(value *string)() {
    m.osDescription = value
}
// Sets the osVersion property value. OS Version.
// Parameters:
//  - value : Value to set for the osVersion property.
func (m *DeviceInstallState) SetOsVersion(value *string)() {
    m.osVersion = value
}
// Sets the userName property value. Device User Name.
// Parameters:
//  - value : Value to set for the userName property.
func (m *DeviceInstallState) SetUserName(value *string)() {
    m.userName = value
}
