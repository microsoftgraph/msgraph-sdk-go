package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IncomingCallOptions provides operations to call the answer method.
type IncomingCallOptions struct {
    CallOptions
}
// NewIncomingCallOptions instantiates a new incomingCallOptions and sets the default values.
func NewIncomingCallOptions()(*IncomingCallOptions) {
    m := &IncomingCallOptions{
        CallOptions: *NewCallOptions(),
    }
    return m
}
// CreateIncomingCallOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIncomingCallOptionsFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewIncomingCallOptions(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IncomingCallOptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.CallOptions.GetFieldDeserializers()
    return res
}
func (m *IncomingCallOptions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *IncomingCallOptions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.CallOptions.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
