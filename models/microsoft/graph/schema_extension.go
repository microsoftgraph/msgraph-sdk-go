package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SchemaExtension struct {
    Entity
    description *string;
    owner *string;
    properties []ExtensionSchemaProperty;
    status *string;
    targetTypes []string;
}
func NewSchemaExtension()(*SchemaExtension) {
    m := &SchemaExtension{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SchemaExtension) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
func (m *SchemaExtension) GetOwner()(*string) {
    if m == nil {
        return nil
    } else {
        return m.owner
    }
}
func (m *SchemaExtension) GetProperties()([]ExtensionSchemaProperty) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
func (m *SchemaExtension) GetStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *SchemaExtension) GetTargetTypes()([]string) {
    if m == nil {
        return nil
    } else {
        return m.targetTypes
    }
}
func (m *SchemaExtension) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDescription(val)
        return nil
    }
    res["owner"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOwner(val)
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtensionSchemaProperty() })
        if err != nil {
            return err
        }
        res := make([]ExtensionSchemaProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*ExtensionSchemaProperty))
        }
        m.SetProperties(res)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStatus(val)
        return nil
    }
    res["targetTypes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetTargetTypes(res)
        return nil
    }
    return res
}
func (m *SchemaExtension) IsNil()(bool) {
    return m == nil
}
func (m *SchemaExtension) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("owner", m.GetOwner())
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
    {
        err = writer.WriteStringValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("targetTypes", m.GetTargetTypes())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *SchemaExtension) SetDescription(value *string)() {
    m.description = value
}
func (m *SchemaExtension) SetOwner(value *string)() {
    m.owner = value
}
func (m *SchemaExtension) SetProperties(value []ExtensionSchemaProperty)() {
    m.properties = value
}
func (m *SchemaExtension) SetStatus(value *string)() {
    m.status = value
}
func (m *SchemaExtension) SetTargetTypes(value []string)() {
    m.targetTypes = value
}
