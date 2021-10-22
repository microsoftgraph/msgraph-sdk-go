package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UpdateWindowsDeviceAccountActionParameter struct {
    additionalData map[string]interface{};
    calendarSyncEnabled *bool;
    deviceAccount *WindowsDeviceAccount;
    deviceAccountEmail *string;
    exchangeServer *string;
    passwordRotationEnabled *bool;
    sessionInitiationProtocalAddress *string;
}
func NewUpdateWindowsDeviceAccountActionParameter()(*UpdateWindowsDeviceAccountActionParameter) {
    m := &UpdateWindowsDeviceAccountActionParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *UpdateWindowsDeviceAccountActionParameter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *UpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.calendarSyncEnabled
    }
}
func (m *UpdateWindowsDeviceAccountActionParameter) GetDeviceAccount()(*WindowsDeviceAccount) {
    if m == nil {
        return nil
    } else {
        return m.deviceAccount
    }
}
func (m *UpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceAccountEmail
    }
}
func (m *UpdateWindowsDeviceAccountActionParameter) GetExchangeServer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeServer
    }
}
func (m *UpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordRotationEnabled
    }
}
func (m *UpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sessionInitiationProtocalAddress
    }
}
func (m *UpdateWindowsDeviceAccountActionParameter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["calendarSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetCalendarSyncEnabled(val)
        return nil
    }
    res["deviceAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDeviceAccount() })
        if err != nil {
            return err
        }
        m.SetDeviceAccount(val.(*WindowsDeviceAccount))
        return nil
    }
    res["deviceAccountEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceAccountEmail(val)
        return nil
    }
    res["exchangeServer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExchangeServer(val)
        return nil
    }
    res["passwordRotationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPasswordRotationEnabled(val)
        return nil
    }
    res["sessionInitiationProtocalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSessionInitiationProtocalAddress(val)
        return nil
    }
    return res
}
func (m *UpdateWindowsDeviceAccountActionParameter) IsNil()(bool) {
    return m == nil
}
func (m *UpdateWindowsDeviceAccountActionParameter) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("calendarSyncEnabled", m.GetCalendarSyncEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("deviceAccount", m.GetDeviceAccount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceAccountEmail", m.GetDeviceAccountEmail())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("exchangeServer", m.GetExchangeServer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("passwordRotationEnabled", m.GetPasswordRotationEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sessionInitiationProtocalAddress", m.GetSessionInitiationProtocalAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *UpdateWindowsDeviceAccountActionParameter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *UpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabled(value *bool)() {
    m.calendarSyncEnabled = value
}
func (m *UpdateWindowsDeviceAccountActionParameter) SetDeviceAccount(value *WindowsDeviceAccount)() {
    m.deviceAccount = value
}
func (m *UpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmail(value *string)() {
    m.deviceAccountEmail = value
}
func (m *UpdateWindowsDeviceAccountActionParameter) SetExchangeServer(value *string)() {
    m.exchangeServer = value
}
func (m *UpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabled(value *bool)() {
    m.passwordRotationEnabled = value
}
func (m *UpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddress(value *string)() {
    m.sessionInitiationProtocalAddress = value
}
