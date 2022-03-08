package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// StopHoldMusicOperation provides operations to call the stopHoldMusic method.
type StopHoldMusicOperation struct {
    CommsOperation
}
// NewStopHoldMusicOperation instantiates a new stopHoldMusicOperation and sets the default values.
func NewStopHoldMusicOperation()(*StopHoldMusicOperation) {
    m := &StopHoldMusicOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateStopHoldMusicOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateStopHoldMusicOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewStopHoldMusicOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *StopHoldMusicOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
func (m *StopHoldMusicOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *StopHoldMusicOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
