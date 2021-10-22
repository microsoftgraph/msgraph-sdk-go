package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChatMessageMentionedIdentitySet struct {
    IdentitySet
    conversation *TeamworkConversationIdentity;
}
func NewChatMessageMentionedIdentitySet()(*ChatMessageMentionedIdentitySet) {
    m := &ChatMessageMentionedIdentitySet{
        IdentitySet: *NewIdentitySet(),
    }
    return m
}
func (m *ChatMessageMentionedIdentitySet) GetConversation()(*TeamworkConversationIdentity) {
    if m == nil {
        return nil
    } else {
        return m.conversation
    }
}
func (m *ChatMessageMentionedIdentitySet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.IdentitySet.GetFieldDeserializers()
    res["conversation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamworkConversationIdentity() })
        if err != nil {
            return err
        }
        m.SetConversation(val.(*TeamworkConversationIdentity))
        return nil
    }
    return res
}
func (m *ChatMessageMentionedIdentitySet) IsNil()(bool) {
    return m == nil
}
func (m *ChatMessageMentionedIdentitySet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.IdentitySet.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("conversation", m.GetConversation())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ChatMessageMentionedIdentitySet) SetConversation(value *TeamworkConversationIdentity)() {
    m.conversation = value
}
