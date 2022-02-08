package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsHelloForBusinessAuthenticationMethod 
type WindowsHelloForBusinessAuthenticationMethod struct {
    AuthenticationMethod
    // The date and time that this Windows Hello for Business key was registered.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The registered device on which this Windows Hello for Business key resides.
    device *Device;
    // The name of the device on which Windows Hello for Business is registered
    displayName *string;
    // Key strength of this Windows Hello for Business key. Possible values are: normal, weak, unknown.
    keyStrength *AuthenticationMethodKeyStrength;
}
// NewWindowsHelloForBusinessAuthenticationMethod instantiates a new windowsHelloForBusinessAuthenticationMethod and sets the default values.
func NewWindowsHelloForBusinessAuthenticationMethod()(*WindowsHelloForBusinessAuthenticationMethod) {
    m := &WindowsHelloForBusinessAuthenticationMethod{
        AuthenticationMethod: *NewAuthenticationMethod(),
    }
    return m
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time that this Windows Hello for Business key was registered.
func (m *WindowsHelloForBusinessAuthenticationMethod) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDevice gets the device property value. The registered device on which this Windows Hello for Business key resides.
func (m *WindowsHelloForBusinessAuthenticationMethod) GetDevice()(*Device) {
    if m == nil {
        return nil
    } else {
        return m.device
    }
}
// GetDisplayName gets the displayName property value. The name of the device on which Windows Hello for Business is registered
func (m *WindowsHelloForBusinessAuthenticationMethod) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetKeyStrength gets the keyStrength property value. Key strength of this Windows Hello for Business key. Possible values are: normal, weak, unknown.
func (m *WindowsHelloForBusinessAuthenticationMethod) GetKeyStrength()(*AuthenticationMethodKeyStrength) {
    if m == nil {
        return nil
    } else {
        return m.keyStrength
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WindowsHelloForBusinessAuthenticationMethod) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["keyStrength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodKeyStrength)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKeyStrength(val.(*AuthenticationMethodKeyStrength))
        }
        return nil
    }
    return res
}
func (m *WindowsHelloForBusinessAuthenticationMethod) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *WindowsHelloForBusinessAuthenticationMethod) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    if m.GetKeyStrength() != nil {
        cast := (*m.GetKeyStrength()).String()
        err = writer.WriteStringValue("keyStrength", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time that this Windows Hello for Business key was registered.
func (m *WindowsHelloForBusinessAuthenticationMethod) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDevice sets the device property value. The registered device on which this Windows Hello for Business key resides.
func (m *WindowsHelloForBusinessAuthenticationMethod) SetDevice(value *Device)() {
    if m != nil {
        m.device = value
    }
}
// SetDisplayName sets the displayName property value. The name of the device on which Windows Hello for Business is registered
func (m *WindowsHelloForBusinessAuthenticationMethod) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetKeyStrength sets the keyStrength property value. Key strength of this Windows Hello for Business key. Possible values are: normal, weak, unknown.
func (m *WindowsHelloForBusinessAuthenticationMethod) SetKeyStrength(value *AuthenticationMethodKeyStrength)() {
    if m != nil {
        m.keyStrength = value
    }
}
