package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessSessionControls struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session control.
    applicationEnforcedRestrictions *ApplicationEnforcedRestrictionsSessionControl;
    // Session control to apply cloud app security.
    cloudAppSecurity *CloudAppSecuritySessionControl;
    // Session control to define whether to persist cookies or not. All apps should be selected for this session control to work correctly.
    persistentBrowser *PersistentBrowserSessionControl;
    // Session control to enforce signin frequency.
    signInFrequency *SignInFrequencySessionControl;
}
// Instantiates a new conditionalAccessSessionControls and sets the default values.
func NewConditionalAccessSessionControls()(*ConditionalAccessSessionControls) {
    m := &ConditionalAccessSessionControls{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessSessionControls) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the applicationEnforcedRestrictions property value. Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session control.
func (m *ConditionalAccessSessionControls) GetApplicationEnforcedRestrictions()(*ApplicationEnforcedRestrictionsSessionControl) {
    if m == nil {
        return nil
    } else {
        return m.applicationEnforcedRestrictions
    }
}
// Gets the cloudAppSecurity property value. Session control to apply cloud app security.
func (m *ConditionalAccessSessionControls) GetCloudAppSecurity()(*CloudAppSecuritySessionControl) {
    if m == nil {
        return nil
    } else {
        return m.cloudAppSecurity
    }
}
// Gets the persistentBrowser property value. Session control to define whether to persist cookies or not. All apps should be selected for this session control to work correctly.
func (m *ConditionalAccessSessionControls) GetPersistentBrowser()(*PersistentBrowserSessionControl) {
    if m == nil {
        return nil
    } else {
        return m.persistentBrowser
    }
}
// Gets the signInFrequency property value. Session control to enforce signin frequency.
func (m *ConditionalAccessSessionControls) GetSignInFrequency()(*SignInFrequencySessionControl) {
    if m == nil {
        return nil
    } else {
        return m.signInFrequency
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ConditionalAccessSessionControls) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the applicationEnforcedRestrictions property value. Session control to enforce application restrictions. Only Exchange Online and Sharepoint Online support this session control.
// Parameters:
//  - value : Value to set for the applicationEnforcedRestrictions property.
func (m *ConditionalAccessSessionControls) SetApplicationEnforcedRestrictions(value *ApplicationEnforcedRestrictionsSessionControl)() {
    m.applicationEnforcedRestrictions = value
}
// Sets the cloudAppSecurity property value. Session control to apply cloud app security.
// Parameters:
//  - value : Value to set for the cloudAppSecurity property.
func (m *ConditionalAccessSessionControls) SetCloudAppSecurity(value *CloudAppSecuritySessionControl)() {
    m.cloudAppSecurity = value
}
// Sets the persistentBrowser property value. Session control to define whether to persist cookies or not. All apps should be selected for this session control to work correctly.
// Parameters:
//  - value : Value to set for the persistentBrowser property.
func (m *ConditionalAccessSessionControls) SetPersistentBrowser(value *PersistentBrowserSessionControl)() {
    m.persistentBrowser = value
}
// Sets the signInFrequency property value. Session control to enforce signin frequency.
// Parameters:
//  - value : Value to set for the signInFrequency property.
func (m *ConditionalAccessSessionControls) SetSignInFrequency(value *SignInFrequencySessionControl)() {
    m.signInFrequency = value
}
