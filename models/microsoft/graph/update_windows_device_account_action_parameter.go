package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new updateWindowsDeviceAccountActionParameter and sets the default values.
func NewUpdateWindowsDeviceAccountActionParameter()(*UpdateWindowsDeviceAccountActionParameter) {
    m := &UpdateWindowsDeviceAccountActionParameter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *UpdateWindowsDeviceAccountActionParameter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the calendarSyncEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetCalendarSyncEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.calendarSyncEnabled
    }
}
// Gets the deviceAccount property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetDeviceAccount()(*WindowsDeviceAccount) {
    if m == nil {
        return nil
    } else {
        return m.deviceAccount
    }
}
// Gets the deviceAccountEmail property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetDeviceAccountEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceAccountEmail
    }
}
// Gets the exchangeServer property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetExchangeServer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.exchangeServer
    }
}
// Gets the passwordRotationEnabled property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetPasswordRotationEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.passwordRotationEnabled
    }
}
// Gets the sessionInitiationProtocalAddress property value. Not yet documented
func (m *UpdateWindowsDeviceAccountActionParameter) GetSessionInitiationProtocalAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sessionInitiationProtocalAddress
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *UpdateWindowsDeviceAccountActionParameter) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the calendarSyncEnabled property value. Not yet documented
// Parameters:
//  - value : Value to set for the calendarSyncEnabled property.
func (m *UpdateWindowsDeviceAccountActionParameter) SetCalendarSyncEnabled(value *bool)() {
    m.calendarSyncEnabled = value
}
// Sets the deviceAccount property value. Not yet documented
// Parameters:
//  - value : Value to set for the deviceAccount property.
func (m *UpdateWindowsDeviceAccountActionParameter) SetDeviceAccount(value *WindowsDeviceAccount)() {
    m.deviceAccount = value
}
// Sets the deviceAccountEmail property value. Not yet documented
// Parameters:
//  - value : Value to set for the deviceAccountEmail property.
func (m *UpdateWindowsDeviceAccountActionParameter) SetDeviceAccountEmail(value *string)() {
    m.deviceAccountEmail = value
}
// Sets the exchangeServer property value. Not yet documented
// Parameters:
//  - value : Value to set for the exchangeServer property.
func (m *UpdateWindowsDeviceAccountActionParameter) SetExchangeServer(value *string)() {
    m.exchangeServer = value
}
// Sets the passwordRotationEnabled property value. Not yet documented
// Parameters:
//  - value : Value to set for the passwordRotationEnabled property.
func (m *UpdateWindowsDeviceAccountActionParameter) SetPasswordRotationEnabled(value *bool)() {
    m.passwordRotationEnabled = value
}
// Sets the sessionInitiationProtocalAddress property value. Not yet documented
// Parameters:
//  - value : Value to set for the sessionInitiationProtocalAddress property.
func (m *UpdateWindowsDeviceAccountActionParameter) SetSessionInitiationProtocalAddress(value *string)() {
    m.sessionInitiationProtocalAddress = value
}
