package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuthenticationFlowsPolicy 
type AuthenticationFlowsPolicy struct {
    Entity
    // Inherited property. A description of the policy. Optional. Read-only.
    description *string;
    // Inherited property. The human-readable name of the policy. Optional. Read-only.
    displayName *string;
    // Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only.
    selfServiceSignUp *SelfServiceSignUpAuthenticationFlowConfiguration;
}
// NewAuthenticationFlowsPolicy instantiates a new authenticationFlowsPolicy and sets the default values.
func NewAuthenticationFlowsPolicy()(*AuthenticationFlowsPolicy) {
    m := &AuthenticationFlowsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
// GetDescription gets the description property value. Inherited property. A description of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// GetDisplayName gets the displayName property value. Inherited property. The human-readable name of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetSelfServiceSignUp gets the selfServiceSignUp property value. Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) GetSelfServiceSignUp()(*SelfServiceSignUpAuthenticationFlowConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.selfServiceSignUp
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
// Serialize serializes information the current object
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
// SetDescription sets the description property value. Inherited property. A description of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) SetDescription(value *string)() {
    if m != nil {
        m.description = value
    }
}
// SetDisplayName sets the displayName property value. Inherited property. The human-readable name of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetSelfServiceSignUp sets the selfServiceSignUp property value. Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) SetSelfServiceSignUp(value *SelfServiceSignUpAuthenticationFlowConfiguration)() {
    if m != nil {
        m.selfServiceSignUp = value
    }
}
