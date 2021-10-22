package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChatMessageReaction struct {
    additionalData map[string]interface{};
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    reactionType *string;
    user *ChatMessageReactionIdentitySet;
}
func NewChatMessageReaction()(*ChatMessageReaction) {
    m := &ChatMessageReaction{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChatMessageReaction) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChatMessageReaction) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *ChatMessageReaction) GetReactionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reactionType
    }
}
func (m *ChatMessageReaction) GetUser()(*ChatMessageReactionIdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.user
    }
}
func (m *ChatMessageReaction) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["reactionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReactionType(val)
        return nil
    }
    res["user"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageReactionIdentitySet() })
        if err != nil {
            return err
        }
        m.SetUser(val.(*ChatMessageReactionIdentitySet))
        return nil
    }
    return res
}
func (m *ChatMessageReaction) IsNil()(bool) {
    return m == nil
}
func (m *ChatMessageReaction) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reactionType", m.GetReactionType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("user", m.GetUser())
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
func (m *ChatMessageReaction) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChatMessageReaction) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *ChatMessageReaction) SetReactionType(value *string)() {
    m.reactionType = value
}
func (m *ChatMessageReaction) SetUser(value *ChatMessageReactionIdentitySet)() {
    m.user = value
}
