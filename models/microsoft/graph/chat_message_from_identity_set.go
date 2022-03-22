package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChatMessageFromIdentitySet 
type ChatMessageFromIdentitySet struct {
    IdentitySet
}
// NewChatMessageFromIdentitySet instantiates a new chatMessageFromIdentitySet and sets the default values.
func NewChatMessageFromIdentitySet()(*ChatMessageFromIdentitySet) {
    m := &ChatMessageFromIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// CreateChatMessageFromIdentitySetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageFromIdentitySetFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewChatMessageFromIdentitySet(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageFromIdentitySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ChatMessageFromIdentitySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
