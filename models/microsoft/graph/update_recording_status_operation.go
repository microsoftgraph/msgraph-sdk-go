package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdateRecordingStatusOperation provides operations to call the updateRecordingStatus method.
type UpdateRecordingStatusOperation struct {
    CommsOperation
}
// NewUpdateRecordingStatusOperation instantiates a new updateRecordingStatusOperation and sets the default values.
func NewUpdateRecordingStatusOperation()(*UpdateRecordingStatusOperation) {
    m := &UpdateRecordingStatusOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateUpdateRecordingStatusOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateUpdateRecordingStatusOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewUpdateRecordingStatusOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *UpdateRecordingStatusOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
func (m *UpdateRecordingStatusOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *UpdateRecordingStatusOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
