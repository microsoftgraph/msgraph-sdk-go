package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CancelMediaProcessingOperation 
type CancelMediaProcessingOperation struct {
    CommsOperation
}
// NewCancelMediaProcessingOperation instantiates a new cancelMediaProcessingOperation and sets the default values.
func NewCancelMediaProcessingOperation()(*CancelMediaProcessingOperation) {
    m := &CancelMediaProcessingOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CancelMediaProcessingOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
func (m *CancelMediaProcessingOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CancelMediaProcessingOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
