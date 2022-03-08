package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Extension provides operations to manage the collection of device entities.
type Extension struct {
    Entity
}
// NewExtension instantiates a new extension and sets the default values.
func NewExtension()(*Extension) {
    m := &Extension{
        Entity: *NewEntity(),
    }
    return m
}
// CreateExtensionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateExtensionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewExtension(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Extension) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *Extension) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Extension) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
