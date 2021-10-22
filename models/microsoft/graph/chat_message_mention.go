package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChatMessageMention struct {
    additionalData map[string]interface{};
    id *int32;
    mentioned *ChatMessageMentionedIdentitySet;
    mentionText *string;
}
func NewChatMessageMention()(*ChatMessageMention) {
    m := &ChatMessageMention{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChatMessageMention) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChatMessageMention) GetId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *ChatMessageMention) GetMentioned()(*ChatMessageMentionedIdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.mentioned
    }
}
func (m *ChatMessageMention) GetMentionText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mentionText
    }
}
func (m *ChatMessageMention) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["mentioned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageMentionedIdentitySet() })
        if err != nil {
            return err
        }
        m.SetMentioned(val.(*ChatMessageMentionedIdentitySet))
        return nil
    }
    res["mentionText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMentionText(val)
        return nil
    }
    return res
}
func (m *ChatMessageMention) IsNil()(bool) {
    return m == nil
}
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
func (m *ChatMessageMention) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChatMessageMention) SetId(value *int32)() {
    m.id = value
}
func (m *ChatMessageMention) SetMentioned(value *ChatMessageMentionedIdentitySet)() {
    m.mentioned = value
}
func (m *ChatMessageMention) SetMentionText(value *string)() {
    m.mentionText = value
}
