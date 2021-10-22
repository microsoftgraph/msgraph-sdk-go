package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConditionalAccessSessionControls struct {
    additionalData map[string]interface{};
    applicationEnforcedRestrictions *ApplicationEnforcedRestrictionsSessionControl;
    cloudAppSecurity *CloudAppSecuritySessionControl;
    persistentBrowser *PersistentBrowserSessionControl;
    signInFrequency *SignInFrequencySessionControl;
}
func NewConditionalAccessSessionControls()(*ConditionalAccessSessionControls) {
    m := &ConditionalAccessSessionControls{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConditionalAccessSessionControls) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConditionalAccessSessionControls) GetApplicationEnforcedRestrictions()(*ApplicationEnforcedRestrictionsSessionControl) {
    if m == nil {
        return nil
    } else {
        return m.applicationEnforcedRestrictions
    }
}
func (m *ConditionalAccessSessionControls) GetCloudAppSecurity()(*CloudAppSecuritySessionControl) {
    if m == nil {
        return nil
    } else {
        return m.cloudAppSecurity
    }
}
func (m *ConditionalAccessSessionControls) GetPersistentBrowser()(*PersistentBrowserSessionControl) {
    if m == nil {
        return nil
    } else {
        return m.persistentBrowser
    }
}
func (m *ConditionalAccessSessionControls) GetSignInFrequency()(*SignInFrequencySessionControl) {
    if m == nil {
        return nil
    } else {
        return m.signInFrequency
    }
}
func (m *ConditionalAccessSessionControls) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["applicationEnforcedRestrictions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApplicationEnforcedRestrictionsSessionControl() })
        if err != nil {
            return err
        }
        m.SetApplicationEnforcedRestrictions(val.(*ApplicationEnforcedRestrictionsSessionControl))
        return nil
    }
    res["cloudAppSecurity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewCloudAppSecuritySessionControl() })
        if err != nil {
            return err
        }
        m.SetCloudAppSecurity(val.(*CloudAppSecuritySessionControl))
        return nil
    }
    res["persistentBrowser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPersistentBrowserSessionControl() })
        if err != nil {
            return err
        }
        m.SetPersistentBrowser(val.(*PersistentBrowserSessionControl))
        return nil
    }
    res["signInFrequency"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSignInFrequencySessionControl() })
        if err != nil {
            return err
        }
        m.SetSignInFrequency(val.(*SignInFrequencySessionControl))
        return nil
    }
    return res
}
func (m *ConditionalAccessSessionControls) IsNil()(bool) {
    return m == nil
}
func (m *ConditionalAccessSessionControls) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("applicationEnforcedRestrictions", m.GetApplicationEnforcedRestrictions())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("cloudAppSecurity", m.GetCloudAppSecurity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("persistentBrowser", m.GetPersistentBrowser())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("signInFrequency", m.GetSignInFrequency())
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
func (m *ConditionalAccessSessionControls) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConditionalAccessSessionControls) SetApplicationEnforcedRestrictions(value *ApplicationEnforcedRestrictionsSessionControl)() {
    m.applicationEnforcedRestrictions = value
}
func (m *ConditionalAccessSessionControls) SetCloudAppSecurity(value *CloudAppSecuritySessionControl)() {
    m.cloudAppSecurity = value
}
func (m *ConditionalAccessSessionControls) SetPersistentBrowser(value *PersistentBrowserSessionControl)() {
    m.persistentBrowser = value
}
func (m *ConditionalAccessSessionControls) SetSignInFrequency(value *SignInFrequencySessionControl)() {
    m.signInFrequency = value
}
