package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ApplicationEnforcedRestrictionsSessionControl struct {
    ConditionalAccessSessionControl
}
func NewApplicationEnforcedRestrictionsSessionControl()(*ApplicationEnforcedRestrictionsSessionControl) {
    m := &ApplicationEnforcedRestrictionsSessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    return m
}
func (m *ApplicationEnforcedRestrictionsSessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    return res
}
func (m *ApplicationEnforcedRestrictionsSessionControl) IsNil()(bool) {
    return m == nil
}
func (m *ApplicationEnforcedRestrictionsSessionControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ConditionalAccessSessionControl.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
