package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ListItemVersion 
type ListItemVersion struct {
    BaseItemVersion
    // A collection of the fields and values for this version of the list item.
    fields *FieldValueSet;
}
// NewListItemVersion instantiates a new listItemVersion and sets the default values.
func NewListItemVersion()(*ListItemVersion) {
    m := &ListItemVersion{
        BaseItemVersion: *NewBaseItemVersion(),
    }
    return m
}
// GetFields gets the fields property value. A collection of the fields and values for this version of the list item.
func (m *ListItemVersion) GetFields()(*FieldValueSet) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListItemVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItemVersion.GetFieldDeserializers()
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFieldValueSet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFields(val.(*FieldValueSet))
        }
        return nil
    }
    return res
}
func (m *ListItemVersion) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ListItemVersion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.BaseItemVersion.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("fields", m.GetFields())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetFields sets the fields property value. A collection of the fields and values for this version of the list item.
func (m *ListItemVersion) SetFields(value *FieldValueSet)() {
    m.fields = value
}
