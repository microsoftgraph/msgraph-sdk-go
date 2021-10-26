package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MultiValueLegacyExtendedProperty struct {
    Entity
    // A collection of property values.
    value []string;
}
// Instantiates a new multiValueLegacyExtendedProperty and sets the default values.
func NewMultiValueLegacyExtendedProperty()(*MultiValueLegacyExtendedProperty) {
    m := &MultiValueLegacyExtendedProperty{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the value property value. A collection of property values.
func (m *MultiValueLegacyExtendedProperty) GetValue()([]string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
// The deserialization information for the current model
func (m *MultiValueLegacyExtendedProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetValue(res)
        return nil
    }
    return res
}
func (m *MultiValueLegacyExtendedProperty) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MultiValueLegacyExtendedProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the value property value. A collection of property values.
// Parameters:
//  - value : Value to set for the value property.
func (m *MultiValueLegacyExtendedProperty) SetValue(value []string)() {
    m.value = value
}
