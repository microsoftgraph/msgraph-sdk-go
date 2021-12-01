package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MicrosoftAuthenticatorAuthenticationMethod 
type MicrosoftAuthenticatorAuthenticationMethod struct {
    AuthenticationMethod
    // The date and time that this app was registered. This property is null if the device is not registered for passwordless Phone Sign-In.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
    device *Device;
    // Tags containing app metadata.
    deviceTag *string;
    // The name of the device on which this app is registered.
    displayName *string;
    // Numerical version of this instance of the Authenticator app.
    phoneAppVersion *string;
}
// NewMicrosoftAuthenticatorAuthenticationMethod instantiates a new microsoftAuthenticatorAuthenticationMethod and sets the default values.
func NewMicrosoftAuthenticatorAuthenticationMethod()(*MicrosoftAuthenticatorAuthenticationMethod) {
    m := &MicrosoftAuthenticatorAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time that this app was registered. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDevice gets the device property value. The registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetDevice()(*Device) {
    if m == nil {
        return nil
    } else {
        return m.device
    }
}
// GetDeviceTag gets the deviceTag property value. Tags containing app metadata.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetDeviceTag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceTag
    }
}
// GetDisplayName gets the displayName property value. The name of the device on which this app is registered.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetPhoneAppVersion gets the phoneAppVersion property value. Numerical version of this instance of the Authenticator app.
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetPhoneAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.phoneAppVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MicrosoftAuthenticatorAuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AuthenticationMethod.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["device"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDevice() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDevice(val.(*Device))
        }
        return nil
    }
    res["deviceTag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceTag(val)
        }
        return nil
    }
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
    res["phoneAppVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhoneAppVersion(val)
        }
        return nil
    }
    return res
}
func (m *MicrosoftAuthenticatorAuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MicrosoftAuthenticatorAuthenticationMethod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AuthenticationMethod.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("device", m.GetDevice())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceTag", m.GetDeviceTag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("phoneAppVersion", m.GetPhoneAppVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time that this app was registered. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDevice sets the device property value. The registered device on which Microsoft Authenticator resides. This property is null if the device is not registered for passwordless Phone Sign-In.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetDevice(value *Device)() {
    if m != nil {
        m.device = value
    }
}
// SetDeviceTag sets the deviceTag property value. Tags containing app metadata.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetDeviceTag(value *string)() {
    if m != nil {
        m.deviceTag = value
    }
}
// SetDisplayName sets the displayName property value. The name of the device on which this app is registered.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetPhoneAppVersion sets the phoneAppVersion property value. Numerical version of this instance of the Authenticator app.
func (m *MicrosoftAuthenticatorAuthenticationMethod) SetPhoneAppVersion(value *string)() {
    if m != nil {
        m.phoneAppVersion = value
    }
}
