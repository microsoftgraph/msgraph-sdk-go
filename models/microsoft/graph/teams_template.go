package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TeamsTemplate provides operations to manage the collection of group entities.
type TeamsTemplate struct {
    Entity
}
// NewTeamsTemplate instantiates a new teamsTemplate and sets the default values.
func NewTeamsTemplate()(*TeamsTemplate) {
    m := &TeamsTemplate{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTeamsTemplateFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTeamsTemplateFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTeamsTemplate(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamsTemplate) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *TeamsTemplate) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamsTemplate) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
