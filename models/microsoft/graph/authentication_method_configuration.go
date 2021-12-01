package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AuthenticationMethodConfiguration 
type AuthenticationMethodConfiguration struct {
    Entity
    // The state of the policy. Possible values are: enabled, disabled.
    state *AuthenticationMethodState;
}
// NewAuthenticationMethodConfiguration instantiates a new authenticationMethodConfiguration and sets the default values.
func NewAuthenticationMethodConfiguration()(*AuthenticationMethodConfiguration) {
    m := &AuthenticationMethodConfiguration{
        Entity: *NewEntity(),
    }
    return m
}
// GetState gets the state property value. The state of the policy. Possible values are: enabled, disabled.
func (m *AuthenticationMethodConfiguration) GetState()(*AuthenticationMethodState) {
    if m == nil {
        return nil
    } else {
        return m.state
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationMethodConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["state"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAuthenticationMethodState)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AuthenticationMethodState)
            m.SetState(&cast)
        }
        return nil
    }
    return res
}
func (m *AuthenticationMethodConfiguration) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetState sets the state property value. The state of the policy. Possible values are: enabled, disabled.
func (m *AuthenticationMethodConfiguration) SetState(value *AuthenticationMethodState)() {
    if m != nil {
        m.state = value
    }
}
