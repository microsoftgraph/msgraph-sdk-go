package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserFlowLanguagePage provides operations to manage the identityContainer singleton.
type UserFlowLanguagePage struct {
    Entity
}
// NewUserFlowLanguagePage instantiates a new userFlowLanguagePage and sets the default values.
func NewUserFlowLanguagePage()(*UserFlowLanguagePage) {
    m := &UserFlowLanguagePage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateUserFlowLanguagePageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUserFlowLanguagePageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUserFlowLanguagePage(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UserFlowLanguagePage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    return res
}
func (m *UserFlowLanguagePage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UserFlowLanguagePage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
