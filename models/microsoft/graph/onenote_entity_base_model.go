package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type OnenoteEntityBaseModel struct {
    Entity
    self *string;
}
func NewOnenoteEntityBaseModel()(*OnenoteEntityBaseModel) {
    m := &OnenoteEntityBaseModel{
        Entity: *NewEntity(),
    }
    return m
}
func (m *OnenoteEntityBaseModel) GetSelf()(*string) {
    if m == nil {
        return nil
    } else {
        return m.self
    }
}
func (m *OnenoteEntityBaseModel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["self"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSelf(val)
        return nil
    }
    return res
}
func (m *OnenoteEntityBaseModel) IsNil()(bool) {
    return m == nil
}
func (m *OnenoteEntityBaseModel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("self", m.GetSelf())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *OnenoteEntityBaseModel) SetSelf(value *string)() {
    m.self = value
}
