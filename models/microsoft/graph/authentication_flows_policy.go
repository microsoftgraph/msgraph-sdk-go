package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type AuthenticationFlowsPolicy struct {
    Entity
    // Inherited property. A description of the policy. Optional. Read-only.
    description *string;
    // Inherited property. The human-readable name of the policy. Optional. Read-only.
    displayName *string;
    // Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only.
    selfServiceSignUp *SelfServiceSignUpAuthenticationFlowConfiguration;
}
// Instantiates a new authenticationFlowsPolicy and sets the default values.
func NewAuthenticationFlowsPolicy()(*AuthenticationFlowsPolicy) {
    m := &AuthenticationFlowsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the description property value. Inherited property. A description of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Inherited property. The human-readable name of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the selfServiceSignUp property value. Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) GetSelfServiceSignUp()(*SelfServiceSignUpAuthenticationFlowConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.selfServiceSignUp
    }
}
// The deserialization information for the current model
func (m *AuthenticationFlowsPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
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
    res["selfServiceSignUp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSelfServiceSignUpAuthenticationFlowConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelfServiceSignUp(val.(*SelfServiceSignUpAuthenticationFlowConfiguration))
        }
        return nil
    }
    return res
}
func (m *AuthenticationFlowsPolicy) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *AuthenticationFlowsPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteObjectValue("selfServiceSignUp", m.GetSelfServiceSignUp())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the description property value. Inherited property. A description of the policy. Optional. Read-only.
// Parameters:
//  - value : Value to set for the description property.
func (m *AuthenticationFlowsPolicy) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Inherited property. The human-readable name of the policy. Optional. Read-only.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *AuthenticationFlowsPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the selfServiceSignUp property value. Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only.
// Parameters:
//  - value : Value to set for the selfServiceSignUp property.
func (m *AuthenticationFlowsPolicy) SetSelfServiceSignUp(value *SelfServiceSignUpAuthenticationFlowConfiguration)() {
    m.selfServiceSignUp = value
}
