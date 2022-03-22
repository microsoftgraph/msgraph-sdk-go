package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// StartHoldMusicOperation 
type StartHoldMusicOperation struct {
    CommsOperation
}
// NewStartHoldMusicOperation instantiates a new startHoldMusicOperation and sets the default values.
func NewStartHoldMusicOperation()(*StartHoldMusicOperation) {
    m := &StartHoldMusicOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateStartHoldMusicOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStartHoldMusicOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewStartHoldMusicOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StartHoldMusicOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *StartHoldMusicOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
