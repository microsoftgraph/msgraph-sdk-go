package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type UpdateRecordingStatusOperation struct {
    CommsOperation
}
// Instantiates a new updateRecordingStatusOperation and sets the default values.
func NewUpdateRecordingStatusOperation()(*UpdateRecordingStatusOperation) {
    m := &UpdateRecordingStatusOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// The deserialization information for the current model
func (m *UpdateRecordingStatusOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
func (m *UpdateRecordingStatusOperation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *UpdateRecordingStatusOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
