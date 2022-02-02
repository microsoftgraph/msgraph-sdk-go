package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Schema 
type Schema struct {
    Entity
    // Must be set to microsoft.graph.externalItem. Required.
    baseType *string;
    // The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
    properties []Property;
}
// NewSchema instantiates a new schema and sets the default values.
func NewSchema()(*Schema) {
    m := &Schema{
        Entity: *NewEntity(),
    }
    return m
}
// GetBaseType gets the baseType property value. Must be set to microsoft.graph.externalItem. Required.
func (m *Schema) GetBaseType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseType
    }
}
// GetProperties gets the properties property value. The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
func (m *Schema) GetProperties()([]Property) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Schema) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["baseType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBaseType(val)
        }
        return nil
    }
    res["properties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Property, len(val))
            for i, v := range val {
                res[i] = *(v.(*Property))
            }
            m.SetProperties(res)
        }
        return nil
    }
    return res
}
func (m *Schema) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
    if m.GetProperties() != nil {
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
// SetBaseType sets the baseType property value. Must be set to microsoft.graph.externalItem. Required.
func (m *Schema) SetBaseType(value *string)() {
    if m != nil {
        m.baseType = value
    }
}
// SetProperties sets the properties property value. The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
func (m *Schema) SetProperties(value []Property)() {
    if m != nil {
        m.properties = value
    }
}
