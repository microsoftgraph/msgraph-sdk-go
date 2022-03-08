package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ApplicationEnforcedRestrictionsSessionControl provides operations to manage the identityContainer singleton.
type ApplicationEnforcedRestrictionsSessionControl struct {
    ConditionalAccessSessionControl
}
// NewApplicationEnforcedRestrictionsSessionControl instantiates a new applicationEnforcedRestrictionsSessionControl and sets the default values.
func NewApplicationEnforcedRestrictionsSessionControl()(*ApplicationEnforcedRestrictionsSessionControl) {
    m := &ApplicationEnforcedRestrictionsSessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    return m
}
// CreateApplicationEnforcedRestrictionsSessionControlFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateApplicationEnforcedRestrictionsSessionControlFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewApplicationEnforcedRestrictionsSessionControl(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ApplicationEnforcedRestrictionsSessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    return res
}
func (m *ApplicationEnforcedRestrictionsSessionControl) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ApplicationEnforcedRestrictionsSessionControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ConditionalAccessSessionControl.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
