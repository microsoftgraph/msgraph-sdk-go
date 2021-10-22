package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuthenticationFlowsPolicy struct {
    Entity
    description *string;
    displayName *string;
    selfServiceSignUp *SelfServiceSignUpAuthenticationFlowConfiguration;
}
func NewAuthenticationFlowsPolicy()(*AuthenticationFlowsPolicy) {
    m := &AuthenticationFlowsPolicy{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AuthenticationFlowsPolicy) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *AuthenticationFlowsPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AuthenticationFlowsPolicy) GetSelfServiceSignUp()(*SelfServiceSignUpAuthenticationFlowConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.selfServiceSignUp
    }
}
func (m *AuthenticationFlowsPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["selfServiceSignUp"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSelfServiceSignUpAuthenticationFlowConfiguration() })
        if err != nil {
            return err
        }
        m.SetSelfServiceSignUp(val.(*SelfServiceSignUpAuthenticationFlowConfiguration))
        return nil
    }
    return res
}
func (m *AuthenticationFlowsPolicy) IsNil()(bool) {
    return m == nil
}
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
func (m *AuthenticationFlowsPolicy) SetDescription(value *string)() {
    m.description = value
}
func (m *AuthenticationFlowsPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AuthenticationFlowsPolicy) SetSelfServiceSignUp(value *SelfServiceSignUpAuthenticationFlowConfiguration)() {
    m.selfServiceSignUp = value
}
