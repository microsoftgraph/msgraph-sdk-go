package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// chatMessageReactionIdentitySet 
type ChatMessageReactionIdentitySet struct {
    IdentitySet
}
// NewChatMessageReactionIdentitySet instantiates a new chatMessageReactionIdentitySet and sets the default values.
func NewChatMessageReactionIdentitySet()(*ChatMessageReactionIdentitySet) {
    m := &ChatMessageReactionIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageReactionIdentitySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    return res
}
func (m *ChatMessageReactionIdentitySet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChatMessageReactionIdentitySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
