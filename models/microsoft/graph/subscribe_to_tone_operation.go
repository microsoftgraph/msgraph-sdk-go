package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubscribeToToneOperation provides operations to call the subscribeToTone method.
type SubscribeToToneOperation struct {
    CommsOperation
}
// NewSubscribeToToneOperation instantiates a new subscribeToToneOperation and sets the default values.
func NewSubscribeToToneOperation()(*SubscribeToToneOperation) {
    m := &SubscribeToToneOperation{
        CommsOperation: *NewCommsOperation(),
    }
    return m
}
// CreateSubscribeToToneOperationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubscribeToToneOperationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSubscribeToToneOperation(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubscribeToToneOperation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CommsOperation.GetFieldDeserializers()
    return res
}
func (m *SubscribeToToneOperation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SubscribeToToneOperation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CommsOperation.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
