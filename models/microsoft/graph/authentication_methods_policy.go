package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuthenticationMethodsPolicy struct {
    Entity
    // Represents the settings for each authentication method.
    authenticationMethodConfigurations []AuthenticationMethodConfiguration;
    // A description of the policy. Read-only.
    description *string;
    // The name of the policy. Read-only.
    displayName *string;
    // The date and time of the last update to the policy. Read-only.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The version of the policy in use. Read-only.
    policyVersion *string;
    // 
    reconfirmationInDays *int32;
}
// Instantiates a new authenticationMethodsPolicy and sets the default values.
func NewAuthenticationMethodsPolicy()(*AuthenticationMethodsPolicy) {
    m := &AuthenticationMethodsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the authenticationMethodConfigurations property value. Represents the settings for each authentication method.
func (m *AuthenticationMethodsPolicy) GetAuthenticationMethodConfigurations()([]AuthenticationMethodConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.authenticationMethodConfigurations
    }
}
// Gets the description property value. A description of the policy. Read-only.
func (m *AuthenticationMethodsPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. The name of the policy. Read-only.
func (m *AuthenticationMethodsPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the lastModifiedDateTime property value. The date and time of the last update to the policy. Read-only.
func (m *AuthenticationMethodsPolicy) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the policyVersion property value. The version of the policy in use. Read-only.
func (m *AuthenticationMethodsPolicy) GetPolicyVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.policyVersion
    }
}
// Gets the reconfirmationInDays property value. 
func (m *AuthenticationMethodsPolicy) GetReconfirmationInDays()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.reconfirmationInDays
    }
}
// The deserialization information for the current model
func (m *AuthenticationMethodsPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["authenticationMethodConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuthenticationMethodConfiguration() })
        if err != nil {
            return err
        }
        res := make([]AuthenticationMethodConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*AuthenticationMethodConfiguration))
        }
        m.SetAuthenticationMethodConfigurations(res)
        return nil
    }
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["policyVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPolicyVersion(val)
        return nil
    }
    res["reconfirmationInDays"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetReconfirmationInDays(val)
        return nil
    }
    return res
}
func (m *AuthenticationMethodsPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AuthenticationMethodsPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAuthenticationMethodConfigurations()))
        for i, v := range m.GetAuthenticationMethodConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("authenticationMethodConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
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
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("policyVersion", m.GetPolicyVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("reconfirmationInDays", m.GetReconfirmationInDays())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the authenticationMethodConfigurations property value. Represents the settings for each authentication method.
// Parameters:
//  - value : Value to set for the authenticationMethodConfigurations property.
func (m *AuthenticationMethodsPolicy) SetAuthenticationMethodConfigurations(value []AuthenticationMethodConfiguration)() {
    m.authenticationMethodConfigurations = value
}
// Sets the description property value. A description of the policy. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *AuthenticationMethodsPolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. The name of the policy. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AuthenticationMethodsPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the lastModifiedDateTime property value. The date and time of the last update to the policy. Read-only.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *AuthenticationMethodsPolicy) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the policyVersion property value. The version of the policy in use. Read-only.
// Parameters:
//  - value : Value to set for the policyVersion property.
func (m *AuthenticationMethodsPolicy) SetPolicyVersion(value *string)() {
    m.policyVersion = value
}
// Sets the reconfirmationInDays property value. 
// Parameters:
//  - value : Value to set for the reconfirmationInDays property.
func (m *AuthenticationMethodsPolicy) SetReconfirmationInDays(value *int32)() {
    m.reconfirmationInDays = value
}
