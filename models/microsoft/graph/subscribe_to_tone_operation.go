package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SubscribeToToneOperation struct {
    CommsOperation
}
func NewSubscribeToToneOperation()(*SubscribeToToneOperation) {
    m := &SubscribeToToneOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
func (m *SubscribeToToneOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
func (m *SubscribeToToneOperation) IsNil()(bool) {
    return m == nil
}
func (m *SubscribeToToneOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
