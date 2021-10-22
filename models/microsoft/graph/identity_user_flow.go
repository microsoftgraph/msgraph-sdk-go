package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type IdentityUserFlow struct {
    Entity
    userFlowType *UserFlowType;
    userFlowTypeVersion *float32;
}
func NewIdentityUserFlow()(*IdentityUserFlow) {
    m := &IdentityUserFlow{
        Entity: *NewEntity(),
    }
    return m
}
func (m *IdentityUserFlow) GetUserFlowType()(*UserFlowType) {
    if m == nil {
        return nil
    } else {
        return m.userFlowType
    }
}
func (m *IdentityUserFlow) GetUserFlowTypeVersion()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.userFlowTypeVersion
    }
}
func (m *IdentityUserFlow) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["userFlowType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseUserFlowType)
        if err != nil {
            return err
        }
        cast := val.(UserFlowType)
        m.SetUserFlowType(&cast)
        return nil
    }
    res["userFlowTypeVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        m.SetUserFlowTypeVersion(val)
        return nil
    }
    return res
}
func (m *IdentityUserFlow) IsNil()(bool) {
    return m == nil
}
func (m *IdentityUserFlow) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetUserFlowType() != nil {
        cast := m.GetUserFlowType().String()
        err = writer.WriteStringValue("userFlowType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteFloat32Value("userFlowTypeVersion", m.GetUserFlowTypeVersion())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *IdentityUserFlow) SetUserFlowType(value *UserFlowType)() {
    m.userFlowType = value
}
func (m *IdentityUserFlow) SetUserFlowTypeVersion(value *float32)() {
    m.userFlowTypeVersion = value
}
