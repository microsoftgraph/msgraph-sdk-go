package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MultiValueLegacyExtendedProperty provides operations to manage the drive singleton.
type MultiValueLegacyExtendedProperty struct {
    Entity
    // A collection of property values.
    value []string;
}
// NewMultiValueLegacyExtendedProperty instantiates a new multiValueLegacyExtendedProperty and sets the default values.
func NewMultiValueLegacyExtendedProperty()(*MultiValueLegacyExtendedProperty) {
    m := &MultiValueLegacyExtendedProperty{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMultiValueLegacyExtendedProperty(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MultiValueLegacyExtendedProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. A collection of property values.
func (m *MultiValueLegacyExtendedProperty) GetValue()([]string) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *MultiValueLegacyExtendedProperty) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MultiValueLegacyExtendedProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        err = writer.WriteCollectionOfStringValues("value", m.GetValue())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. A collection of property values.
func (m *MultiValueLegacyExtendedProperty) SetValue(value []string)() {
    if m != nil {
        m.value = value
    }
}
