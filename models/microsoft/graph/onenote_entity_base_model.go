package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnenoteEntityBaseModel provides operations to manage the collection of drive entities.
type OnenoteEntityBaseModel struct {
    Entity
    // The endpoint where you can get details about the page. Read-only.
    self *string;
}
// NewOnenoteEntityBaseModel instantiates a new onenoteEntityBaseModel and sets the default values.
func NewOnenoteEntityBaseModel()(*OnenoteEntityBaseModel) {
    m := &OnenoteEntityBaseModel{
        Entity: *NewEntity(),
    }
    return m
}
// CreateOnenoteEntityBaseModelFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenoteEntityBaseModelFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewOnenoteEntityBaseModel(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenoteEntityBaseModel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["self"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSelf(val)
        }
        return nil
    }
    return res
}
// GetSelf gets the self property value. The endpoint where you can get details about the page. Read-only.
func (m *OnenoteEntityBaseModel) GetSelf()(*string) {
    if m == nil {
        return nil
    } else {
        return m.self
    }
}
func (m *OnenoteEntityBaseModel) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetSelf sets the self property value. The endpoint where you can get details about the page. Read-only.
func (m *OnenoteEntityBaseModel) SetSelf(value *string)() {
    if m != nil {
        m.self = value
    }
}
