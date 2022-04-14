package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceInstallState 
type DeviceInstallState struct {
    Entity
    // Device Id.
    deviceId *string
    // Device name.
    deviceName *string
    // The error code for install failures.
    errorCode *string
    // The install state of the eBook. Possible values are: notApplicable, installed, failed, notInstalled, uninstallFailed, unknown.
    installState *InstallState
    // Last sync date and time.
    lastSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // OS Description.
    osDescription *string
    // OS Version.
    osVersion *string
    // Device User Name.
    userName *string
}
// NewDeviceInstallState instantiates a new deviceInstallState and sets the default values.
func NewDeviceInstallState()(*DeviceInstallState) {
    m := &DeviceInstallState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceInstallStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceInstallStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceInstallState(), nil
}
// GetDeviceId gets the deviceId property value. Device Id.
func (m *DeviceInstallState) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDeviceName gets the deviceName property value. Device name.
func (m *DeviceInstallState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetErrorCode gets the errorCode property value. The error code for install failures.
func (m *DeviceInstallState) GetErrorCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.errorCode
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceInstallState) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["deviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["errorCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorCode(val)
        }
        return nil
    }
    res["installState"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInstallState)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInstallState(val.(*InstallState))
        }
        return nil
    }
    res["lastSyncDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastSyncDateTime(val)
        }
        return nil
    }
    res["osDescription"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsDescription(val)
        }
        return nil
    }
    res["osVersion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsVersion(val)
        }
        return nil
    }
    res["userName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetInstallState gets the installState property value. The install state of the eBook. Possible values are: notApplicable, installed, failed, notInstalled, uninstallFailed, unknown.
func (m *DeviceInstallState) GetInstallState()(*InstallState) {
    if m == nil {
        return nil
    } else {
        return m.installState
    }
}
// GetLastSyncDateTime gets the lastSyncDateTime property value. Last sync date and time.
func (m *DeviceInstallState) GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastSyncDateTime
    }
}
// GetOsDescription gets the osDescription property value. OS Description.
func (m *DeviceInstallState) GetOsDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osDescription
    }
}
// GetOsVersion gets the osVersion property value. OS Version.
func (m *DeviceInstallState) GetOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.osVersion
    }
}
// GetUserName gets the userName property value. Device User Name.
func (m *DeviceInstallState) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// Serialize serializes information the current object
func (m *DeviceInstallState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
        cast := (*m.GetInstallState()).String()
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
// SetDeviceId sets the deviceId property value. Device Id.
func (m *DeviceInstallState) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDeviceName sets the deviceName property value. Device name.
func (m *DeviceInstallState) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetErrorCode sets the errorCode property value. The error code for install failures.
func (m *DeviceInstallState) SetErrorCode(value *string)() {
    if m != nil {
        m.errorCode = value
    }
}
// SetInstallState sets the installState property value. The install state of the eBook. Possible values are: notApplicable, installed, failed, notInstalled, uninstallFailed, unknown.
func (m *DeviceInstallState) SetInstallState(value *InstallState)() {
    if m != nil {
        m.installState = value
    }
}
// SetLastSyncDateTime sets the lastSyncDateTime property value. Last sync date and time.
func (m *DeviceInstallState) SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastSyncDateTime = value
    }
}
// SetOsDescription sets the osDescription property value. OS Description.
func (m *DeviceInstallState) SetOsDescription(value *string)() {
    if m != nil {
        m.osDescription = value
    }
}
// SetOsVersion sets the osVersion property value. OS Version.
func (m *DeviceInstallState) SetOsVersion(value *string)() {
    if m != nil {
        m.osVersion = value
    }
}
// SetUserName sets the userName property value. Device User Name.
func (m *DeviceInstallState) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
