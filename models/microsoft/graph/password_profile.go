package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PasswordProfile struct {
    additionalData map[string]interface{};
    forceChangePasswordNextSignIn *bool;
    forceChangePasswordNextSignInWithMfa *bool;
    password *string;
}
func NewPasswordProfile()(*PasswordProfile) {
    m := &PasswordProfile{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PasswordProfile) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PasswordProfile) GetForceChangePasswordNextSignIn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.forceChangePasswordNextSignIn
    }
}
func (m *PasswordProfile) GetForceChangePasswordNextSignInWithMfa()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.forceChangePasswordNextSignInWithMfa
    }
}
func (m *PasswordProfile) GetPassword()(*string) {
    if m == nil {
        return nil
    } else {
        return m.password
    }
}
func (m *PasswordProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["forceChangePasswordNextSignIn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetForceChangePasswordNextSignIn(val)
        return nil
    }
    res["forceChangePasswordNextSignInWithMfa"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetForceChangePasswordNextSignInWithMfa(val)
        return nil
    }
    res["password"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPassword(val)
        return nil
    }
    return res
}
func (m *PasswordProfile) IsNil()(bool) {
    return m == nil
}
func (m *PasswordProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("forceChangePasswordNextSignIn", m.GetForceChangePasswordNextSignIn())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("forceChangePasswordNextSignInWithMfa", m.GetForceChangePasswordNextSignInWithMfa())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("password", m.GetPassword())
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
func (m *PasswordProfile) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PasswordProfile) SetForceChangePasswordNextSignIn(value *bool)() {
    m.forceChangePasswordNextSignIn = value
}
func (m *PasswordProfile) SetForceChangePasswordNextSignInWithMfa(value *bool)() {
    m.forceChangePasswordNextSignInWithMfa = value
}
func (m *PasswordProfile) SetPassword(value *string)() {
    m.password = value
}
