package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PlayPromptOperation provides operations to call the playPrompt method.
type PlayPromptOperation struct {
    CommsOperation
}
// NewPlayPromptOperation instantiates a new playPromptOperation and sets the default values.
func NewPlayPromptOperation()(*PlayPromptOperation) {
    m := &PlayPromptOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreatePlayPromptOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePlayPromptOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPlayPromptOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PlayPromptOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
func (m *PlayPromptOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PlayPromptOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
