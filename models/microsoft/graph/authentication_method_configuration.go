package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AuthenticationMethodConfiguration struct {
    Entity
    state *AuthenticationMethodState;
}
func NewAuthenticationMethodConfiguration()(*AuthenticationMethodConfiguration) {
    m := &AuthenticationMethodConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
func (m *AuthenticationMethodConfiguration) GetState()(*AuthenticationMethodState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
func (m *AuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodState)
        if err != nil {
            return err
        }
        cast := val.(AuthenticationMethodState)
        m.SetState(&cast)
        return nil
    }
    return res
}
func (m *AuthenticationMethodConfiguration) IsNil()(bool) {
    return m == nil
}
func (m *AuthenticationMethodConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetState() != nil {
        cast := m.GetState().String()
        err = writer.WriteStringValue("state", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *AuthenticationMethodConfiguration) SetState(value *AuthenticationMethodState)() {
    m.state = value
}
