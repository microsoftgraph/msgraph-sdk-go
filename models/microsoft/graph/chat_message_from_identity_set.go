package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChatMessageFromIdentitySet struct {
    IdentitySet
}
// Instantiates a new chatMessageFromIdentitySet and sets the default values.
func NewChatMessageFromIdentitySet()(*ChatMessageFromIdentitySet) {
    m := &ChatMessageFromIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// The deserialization information for the current model
func (m *ChatMessageFromIdentitySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    return res
}
func (m *ChatMessageFromIdentitySet) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ChatMessageFromIdentitySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
