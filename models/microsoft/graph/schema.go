package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Schema struct {
    Entity
    baseType *string;
    properties []Property;
}
func NewSchema()(*Schema) {
    m := &Schema{
        Entity: *NewEntity(),
    }
    return m
}
func (m *Schema) GetBaseType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseType
    }
}
func (m *Schema) GetProperties()([]Property) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
func (m *Schema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBaseType(val)
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProperty() })
        if err != nil {
            return err
        }
        res := make([]Property, len(val))
        for i, v := range val {
            res[i] = *(v.(*Property))
        }
        m.SetProperties(res)
        return nil
    }
    return res
}
func (m *Schema) IsNil()(bool) {
    return m == nil
}
func (m *Schema) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("baseType", m.GetBaseType())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProperties()))
        for i, v := range m.GetProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("properties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Schema) SetBaseType(value *string)() {
    m.baseType = value
}
func (m *Schema) SetProperties(value []Property)() {
    m.properties = value
}
