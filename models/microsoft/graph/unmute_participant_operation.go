package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type UnmuteParticipantOperation struct {
    CommsOperation
}
func NewUnmuteParticipantOperation()(*UnmuteParticipantOperation) {
    m := &UnmuteParticipantOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
func (m *UnmuteParticipantOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
func (m *UnmuteParticipantOperation) IsNil()(bool) {
    return m == nil
}
func (m *UnmuteParticipantOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
