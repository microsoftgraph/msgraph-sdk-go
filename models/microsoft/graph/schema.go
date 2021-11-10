package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Schema struct {
    Entity
    // Must be set to microsoft.graph.externalConnector.externalItem. Required.
    baseType *string;
    // The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
    properties []Property;
}
// Instantiates a new schema and sets the default values.
func NewSchema()(*Schema) {
    m := &Schema{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the baseType property value. Must be set to microsoft.graph.externalConnector.externalItem. Required.
func (m *Schema) GetBaseType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.baseType
    }
}
// Gets the properties property value. The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
func (m *Schema) GetProperties()([]Property) {
    if m == nil {
        return nil
    } else {
        return m.properties
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the baseType property value. Must be set to microsoft.graph.externalConnector.externalItem. Required.
// Parameters:
//  - value : Value to set for the baseType property.
func (m *Schema) SetBaseType(value *string)() {
    m.baseType = value
}
// Sets the properties property value. The properties defined for the items in the connection. The minimum number of properties is one, the maximum is 128.
// Parameters:
//  - value : Value to set for the properties property.
func (m *Schema) SetProperties(value []Property)() {
    m.properties = value
}
