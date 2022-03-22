package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UnmuteParticipantOperation 
type UnmuteParticipantOperation struct {
    CommsOperation
}
// NewUnmuteParticipantOperation instantiates a new unmuteParticipantOperation and sets the default values.
func NewUnmuteParticipantOperation()(*UnmuteParticipantOperation) {
    m := &UnmuteParticipantOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateUnmuteParticipantOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUnmuteParticipantOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUnmuteParticipantOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UnmuteParticipantOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *UnmuteParticipantOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
