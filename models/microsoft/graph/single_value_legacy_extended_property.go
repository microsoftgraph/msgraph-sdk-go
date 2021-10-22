package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SingleValueLegacyExtendedProperty struct {
    Entity
    value *string;
}
func NewSingleValueLegacyExtendedProperty()(*SingleValueLegacyExtendedProperty) {
    m := &SingleValueLegacyExtendedProperty{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SingleValueLegacyExtendedProperty) GetValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *SingleValueLegacyExtendedProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValue(val)
        return nil
    }
    return res
}
func (m *SingleValueLegacyExtendedProperty) IsNil()(bool) {
    return m == nil
}
func (m *SingleValueLegacyExtendedProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SingleValueLegacyExtendedProperty) SetValue(value *string)() {
    m.value = value
}
