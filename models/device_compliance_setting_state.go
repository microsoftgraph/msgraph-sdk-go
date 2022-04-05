package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceComplianceSettingState 
type DeviceComplianceSettingState struct {
    Entity
    // The DateTime when device compliance grace period expires
    complianceGracePeriodExpirationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Device Id that is being reported
    deviceId *string;
    // The device model that is being reported
    deviceModel *string;
    // The Device Name that is being reported
    deviceName *string;
    // The setting class name and property name.
    setting *string;
    // The Setting Name that is being reported
    settingName *string;
    // The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
    state *ComplianceStatus;
    // The User email address that is being reported
    userEmail *string;
    // The user Id that is being reported
    userId *string;
    // The User Name that is being reported
    userName *string;
    // The User PrincipalName that is being reported
    userPrincipalName *string;
}
// NewDeviceComplianceSettingState instantiates a new deviceComplianceSettingState and sets the default values.
func NewDeviceComplianceSettingState()(*DeviceComplianceSettingState) {
    m := &DeviceComplianceSettingState{
        Entity: *NewEntity(),
    }
    return m
}
// CreateDeviceComplianceSettingStateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceComplianceSettingStateFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceComplianceSettingState(), nil
}
// GetComplianceGracePeriodExpirationDateTime gets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
func (m *DeviceComplianceSettingState) GetComplianceGracePeriodExpirationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.complianceGracePeriodExpirationDateTime
    }
}
// GetDeviceId gets the deviceId property value. The Device Id that is being reported
func (m *DeviceComplianceSettingState) GetDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceId
    }
}
// GetDeviceModel gets the deviceModel property value. The device model that is being reported
func (m *DeviceComplianceSettingState) GetDeviceModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceModel
    }
}
// GetDeviceName gets the deviceName property value. The Device Name that is being reported
func (m *DeviceComplianceSettingState) GetDeviceName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceComplianceSettingState) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["complianceGracePeriodExpirationDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComplianceGracePeriodExpirationDateTime(val)
        }
        return nil
    }
    res["deviceId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceId(val)
        }
        return nil
    }
    res["deviceModel"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceModel(val)
        }
        return nil
    }
    res["deviceName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceName(val)
        }
        return nil
    }
    res["setting"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSetting(val)
        }
        return nil
    }
    res["settingName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettingName(val)
        }
        return nil
    }
    res["state"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseComplianceStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetState(val.(*ComplianceStatus))
        }
        return nil
    }
    res["userEmail"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserEmail(val)
        }
        return nil
    }
    res["userId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserId(val)
        }
        return nil
    }
    res["userName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserName(val)
        }
        return nil
    }
    res["userPrincipalName"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserPrincipalName(val)
        }
        return nil
    }
    return res
}
// GetSetting gets the setting property value. The setting class name and property name.
func (m *DeviceComplianceSettingState) GetSetting()(*string) {
    if m == nil {
        return nil
    } else {
        return m.setting
    }
}
// GetSettingName gets the settingName property value. The Setting Name that is being reported
func (m *DeviceComplianceSettingState) GetSettingName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.settingName
    }
}
// GetState gets the state property value. The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *DeviceComplianceSettingState) GetState()(*ComplianceStatus) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetUserEmail gets the userEmail property value. The User email address that is being reported
func (m *DeviceComplianceSettingState) GetUserEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userEmail
    }
}
// GetUserId gets the userId property value. The user Id that is being reported
func (m *DeviceComplianceSettingState) GetUserId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userId
    }
}
// GetUserName gets the userName property value. The User Name that is being reported
func (m *DeviceComplianceSettingState) GetUserName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userName
    }
}
// GetUserPrincipalName gets the userPrincipalName property value. The User PrincipalName that is being reported
func (m *DeviceComplianceSettingState) GetUserPrincipalName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userPrincipalName
    }
}
// Serialize serializes information the current object
func (m *DeviceComplianceSettingState) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("complianceGracePeriodExpirationDateTime", m.GetComplianceGracePeriodExpirationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceId", m.GetDeviceId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceModel", m.GetDeviceModel())
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
        err = writer.WriteStringValue("setting", m.GetSetting())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("settingName", m.GetSettingName())
        if err != nil {
            return err
        }
    }
    if m.GetState() != nil {
        cast := (*m.GetState()).String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userEmail", m.GetUserEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userId", m.GetUserId())
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
    {
        err = writer.WriteStringValue("userPrincipalName", m.GetUserPrincipalName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetComplianceGracePeriodExpirationDateTime sets the complianceGracePeriodExpirationDateTime property value. The DateTime when device compliance grace period expires
func (m *DeviceComplianceSettingState) SetComplianceGracePeriodExpirationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.complianceGracePeriodExpirationDateTime = value
    }
}
// SetDeviceId sets the deviceId property value. The Device Id that is being reported
func (m *DeviceComplianceSettingState) SetDeviceId(value *string)() {
    if m != nil {
        m.deviceId = value
    }
}
// SetDeviceModel sets the deviceModel property value. The device model that is being reported
func (m *DeviceComplianceSettingState) SetDeviceModel(value *string)() {
    if m != nil {
        m.deviceModel = value
    }
}
// SetDeviceName sets the deviceName property value. The Device Name that is being reported
func (m *DeviceComplianceSettingState) SetDeviceName(value *string)() {
    if m != nil {
        m.deviceName = value
    }
}
// SetSetting sets the setting property value. The setting class name and property name.
func (m *DeviceComplianceSettingState) SetSetting(value *string)() {
    if m != nil {
        m.setting = value
    }
}
// SetSettingName sets the settingName property value. The Setting Name that is being reported
func (m *DeviceComplianceSettingState) SetSettingName(value *string)() {
    if m != nil {
        m.settingName = value
    }
}
// SetState sets the state property value. The compliance state of the setting. Possible values are: unknown, notApplicable, compliant, remediated, nonCompliant, error, conflict, notAssigned.
func (m *DeviceComplianceSettingState) SetState(value *ComplianceStatus)() {
    if m != nil {
        m.state = value
    }
}
// SetUserEmail sets the userEmail property value. The User email address that is being reported
func (m *DeviceComplianceSettingState) SetUserEmail(value *string)() {
    if m != nil {
        m.userEmail = value
    }
}
// SetUserId sets the userId property value. The user Id that is being reported
func (m *DeviceComplianceSettingState) SetUserId(value *string)() {
    if m != nil {
        m.userId = value
    }
}
// SetUserName sets the userName property value. The User Name that is being reported
func (m *DeviceComplianceSettingState) SetUserName(value *string)() {
    if m != nil {
        m.userName = value
    }
}
// SetUserPrincipalName sets the userPrincipalName property value. The User PrincipalName that is being reported
func (m *DeviceComplianceSettingState) SetUserPrincipalName(value *string)() {
    if m != nil {
        m.userPrincipalName = value
    }
}
