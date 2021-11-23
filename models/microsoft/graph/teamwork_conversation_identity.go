package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// teamworkConversationIdentity 
type TeamworkConversationIdentity struct {
    Identity
    // Type of conversation. Possible values are: team, channel, chat, and unknownFutureValue.
    conversationIdentityType *TeamworkConversationIdentityType;
}
// NewTeamworkConversationIdentity instantiates a new teamworkConversationIdentity and sets the default values.
func NewTeamworkConversationIdentity()(*TeamworkConversationIdentity) {
    m := &TeamworkConversationIdentity{
        Identity: *NewIdentity(),
    }
    return m
}
// GetConversationIdentityType gets the conversationIdentityType property value. Type of conversation. Possible values are: team, channel, chat, and unknownFutureValue.
func (m *TeamworkConversationIdentity) GetConversationIdentityType()(*TeamworkConversationIdentityType) {
    if m == nil {
        return nil
    } else {
        return m.conversationIdentityType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TeamworkConversationIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Identity.GetFieldDeserializers()
    res["conversationIdentityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTeamworkConversationIdentityType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(TeamworkConversationIdentityType)
            m.SetConversationIdentityType(&cast)
        }
        return nil
    }
    return res
}
func (m *TeamworkConversationIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TeamworkConversationIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Identity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetConversationIdentityType() != nil {
        cast := m.GetConversationIdentityType().String()
        err = writer.WriteStringValue("conversationIdentityType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetConversationIdentityType sets the conversationIdentityType property value. Type of conversation. Possible values are: team, channel, chat, and unknownFutureValue.
func (m *TeamworkConversationIdentity) SetConversationIdentityType(value *TeamworkConversationIdentityType)() {
    m.conversationIdentityType = value
}
