package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// muteParticipantOperation 
type MuteParticipantOperation struct {
    CommsOperation
}
// NewMuteParticipantOperation instantiates a new muteParticipantOperation and sets the default values.
func NewMuteParticipantOperation()(*MuteParticipantOperation) {
    m := &MuteParticipantOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MuteParticipantOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
func (m *MuteParticipantOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MuteParticipantOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
