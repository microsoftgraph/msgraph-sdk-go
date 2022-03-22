package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChatMessageMention 
type ChatMessageMention struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at id='{index}'> tag in the message body.
    id *int32;
    // The entity (user, application, team, or channel) that was @mentioned.
    mentioned ChatMessageMentionedIdentitySetable;
    // String used to represent the mention. For example, a user's display name, a team name.
    mentionText *string;
}
// NewChatMessageMention instantiates a new chatMessageMention and sets the default values.
func NewChatMessageMention()(*ChatMessageMention) {
    m := &ChatMessageMention{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChatMessageMentionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageMentionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewChatMessageMention(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessageMention) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessageMention) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    res["mentioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessageMentionedIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentioned(val.(ChatMessageMentionedIdentitySetable))
        }
        return nil
    }
    res["mentionText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMentionText(val)
        }
        return nil
    }
    return res
}
// GetId gets the id property value. Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at id='{index}'> tag in the message body.
func (m *ChatMessageMention) GetId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetMentioned gets the mentioned property value. The entity (user, application, team, or channel) that was @mentioned.
func (m *ChatMessageMention) GetMentioned()(ChatMessageMentionedIdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.mentioned
    }
}
// GetMentionText gets the mentionText property value. String used to represent the mention. For example, a user's display name, a team name.
func (m *ChatMessageMention) GetMentionText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mentionText
    }
}
// Serialize serializes information the current object
func (m *ChatMessageMention) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteInt32Value("id", m.GetId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("mentioned", m.GetMentioned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mentionText", m.GetMentionText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessageMention) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetId sets the id property value. Index of an entity being mentioned in the specified chatMessage. Matches the {index} value in the corresponding <at id='{index}'> tag in the message body.
func (m *ChatMessageMention) SetId(value *int32)() {
    if m != nil {
        m.id = value
    }
}
// SetMentioned sets the mentioned property value. The entity (user, application, team, or channel) that was @mentioned.
func (m *ChatMessageMention) SetMentioned(value ChatMessageMentionedIdentitySetable)() {
    if m != nil {
        m.mentioned = value
    }
}
// SetMentionText sets the mentionText property value. String used to represent the mention. For example, a user's display name, a team name.
func (m *ChatMessageMention) SetMentionText(value *string)() {
    if m != nil {
        m.mentionText = value
    }
}
