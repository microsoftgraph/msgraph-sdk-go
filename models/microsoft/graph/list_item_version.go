package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ListItemVersion 
type ListItemVersion struct {
    BaseItemVersion
    // A collection of the fields and values for this version of the list item.
    fields FieldValueSetable;
}
// NewListItemVersion instantiates a new listItemVersion and sets the default values.
func NewListItemVersion()(*ListItemVersion) {
    m := &ListItemVersion{
        BaseItemVersion: *NewBaseItemVersion(),
    }
    return m
}
// CreateListItemVersionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateListItemVersionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewListItemVersion(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ListItemVersion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.BaseItemVersion.GetFieldDeserializers()
    res["fields"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateFieldValueSetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFields(val.(FieldValueSetable))
        }
        return nil
    }
    return res
}
// GetFields gets the fields property value. A collection of the fields and values for this version of the list item.
func (m *ListItemVersion) GetFields()(FieldValueSetable) {
    if m == nil {
        return nil
    } else {
        return m.fields
    }
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
func (m *ListItemVersion) SetFields(value FieldValueSetable)() {
    if m != nil {
        m.fields = value
    }
}
