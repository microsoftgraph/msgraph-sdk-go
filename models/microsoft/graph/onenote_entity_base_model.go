package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnenoteEntityBaseModel struct {
    Entity
    // The endpoint where you can get details about the page. Read-only.
    self *string;
}
// Instantiates a new onenoteEntityBaseModel and sets the default values.
func NewOnenoteEntityBaseModel()(*OnenoteEntityBaseModel) {
    m := &OnenoteEntityBaseModel{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the self property value. The endpoint where you can get details about the page. Read-only.
func (m *OnenoteEntityBaseModel) GetSelf()(*string) {
    if m == nil {
        return nil
    } else {
        return m.self
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the self property value. The endpoint where you can get details about the page. Read-only.
// Parameters:
//  - value : Value to set for the self property.
func (m *OnenoteEntityBaseModel) SetSelf(value *string)() {
    m.self = value
}
