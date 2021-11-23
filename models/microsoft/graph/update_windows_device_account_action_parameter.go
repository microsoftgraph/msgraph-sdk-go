package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// updateWindowsDeviceAccountActionParameter 
type UpdateWindowsDeviceAccountActionParameter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Not yet documented
    calendarSyncEnabled *bool;
    // Not yet documented
    deviceAccount *WindowsDeviceAccount;
    // Not yet documented
    deviceAccountEmail *string;
    // Not yet documented
    exchangeServer *string;
    // Not yet documented
    passwordRotationEnabled *bool;
    // Not yet documented
    sessionInitiationProtocalAddress *string;
}
// NewUpdateWindowsDeviceAccountActionParameter instantiates a new updateWindowsDeviceAccountActionParameter and sets the default values.
func NewUpdateWindowsDeviceAccountActionParameter()(*UpdateWindowsDeviceAccountActionParameter) {
    m := &UpdateWindowsDeviceAccountActionParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateWindowsDeviceAccountActionParameter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCalendarSyncEnabled gets the calendarSyncEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.calendarSyncEnabled
    }
}
// GetDeviceAccount gets the deviceAccount property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetDeviceAccount()(*WindowsDeviceAccount) {
    if m == nil {
        return nil
    } else {
        return m.deviceAccount
    }
}
// GetDeviceAccountEmail gets the deviceAccountEmail property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceAccountEmail
    }
}
// GetExchangeServer gets the exchangeServer property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetExchangeServer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeServer
    }
}
// GetPasswordRotationEnabled gets the passwordRotationEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordRotationEnabled
    }
}
// GetSessionInitiationProtocalAddress gets the sessionInitiationProtocalAddress property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sessionInitiationProtocalAddress
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateWindowsDeviceAccountActionParameter) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["calendarSyncEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCalendarSyncEnabled(val)
        }
        return nil
    }
    res["deviceAccount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsDeviceAccount() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccount(val.(*WindowsDeviceAccount))
        }
        return nil
    }
    res["deviceAccountEmail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceAccountEmail(val)
        }
        return nil
    }
    res["exchangeServer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExchangeServer(val)
        }
        return nil
    }
    res["passwordRotationEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPasswordRotationEnabled(val)
        }
        return nil
    }
    res["sessionInitiationProtocalAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSessionInitiationProtocalAddress(val)
        }
        return nil
    }
    return res
}
func (m *UpdateWindowsDeviceAccountActionParameter) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateWindowsDeviceAccountActionParameter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCalendarSyncEnabled sets the calendarSyncEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabled(value *bool)() {
    m.calendarSyncEnabled = value
}
// SetDeviceAccount sets the deviceAccount property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetDeviceAccount(value *WindowsDeviceAccount)() {
    m.deviceAccount = value
}
// SetDeviceAccountEmail sets the deviceAccountEmail property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmail(value *string)() {
    m.deviceAccountEmail = value
}
// SetExchangeServer sets the exchangeServer property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetExchangeServer(value *string)() {
    m.exchangeServer = value
}
// SetPasswordRotationEnabled sets the passwordRotationEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabled(value *bool)() {
    m.passwordRotationEnabled = value
}
// SetSessionInitiationProtocalAddress sets the sessionInitiationProtocalAddress property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddress(value *string)() {
    m.sessionInitiationProtocalAddress = value
}
