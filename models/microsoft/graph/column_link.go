package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ColumnLink provides operations to manage the collection of drive entities.
type ColumnLink struct {
    Entity
    // The name of the column  in this content type.
    name *string;
}
// NewColumnLink instantiates a new columnLink and sets the default values.
func NewColumnLink()(*ColumnLink) {
    m := &ColumnLink{
        Entity: *NewEntity(),
    }
    return m
}
// CreateColumnLinkFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateColumnLinkFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewColumnLink(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ColumnLink) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetName gets the name property value. The name of the column  in this content type.
func (m *ColumnLink) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *ColumnLink) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ColumnLink) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetName sets the name property value. The name of the column  in this content type.
func (m *ColumnLink) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
