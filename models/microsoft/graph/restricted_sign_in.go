package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RestrictedSignIn struct {
    SignIn
    targetTenantId *string;
}
func NewRestrictedSignIn()(*RestrictedSignIn) {
    m := &RestrictedSignIn{
        SignIn: *NewSignIn(),
    }
    return m
}
func (m *RestrictedSignIn) GetTargetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.targetTenantId
    }
}
func (m *RestrictedSignIn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.SignIn.GetFieldDeserializers()
    res["targetTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTargetTenantId(val)
        return nil
    }
    return res
}
func (m *RestrictedSignIn) IsNil()(bool) {
    return m == nil
}
func (m *RestrictedSignIn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.SignIn.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("targetTenantId", m.GetTargetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *RestrictedSignIn) SetTargetTenantId(value *string)() {
    m.targetTenantId = value
}
