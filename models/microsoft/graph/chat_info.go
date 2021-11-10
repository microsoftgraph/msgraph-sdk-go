package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChatInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The unique identifier of a message in a Microsoft Teams channel.
    messageId *string;
    // The ID of the reply message.
    replyChainMessageId *string;
    // The unique identifier for a thread in Microsoft Teams.
    threadId *string;
}
// Instantiates a new chatInfo and sets the default values.
func NewChatInfo()(*ChatInfo) {
    m := &ChatInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the messageId property value. The unique identifier of a message in a Microsoft Teams channel.
func (m *ChatInfo) GetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageId
    }
}
// Gets the replyChainMessageId property value. The ID of the reply message.
func (m *ChatInfo) GetReplyChainMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.replyChainMessageId
    }
}
// Gets the threadId property value. The unique identifier for a thread in Microsoft Teams.
func (m *ChatInfo) GetThreadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threadId
    }
}
// The deserialization information for the current model
func (m *ChatInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["messageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageId(val)
        }
        return nil
    }
    res["replyChainMessageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplyChainMessageId(val)
        }
        return nil
    }
    res["threadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetThreadId(val)
        }
        return nil
    }
    return res
}
func (m *ChatInfo) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ChatInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("messageId", m.GetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("replyChainMessageId", m.GetReplyChainMessageId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("threadId", m.GetThreadId())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ChatInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the messageId property value. The unique identifier of a message in a Microsoft Teams channel.
// Parameters:
//  - value : Value to set for the messageId property.
func (m *ChatInfo) SetMessageId(value *string)() {
    m.messageId = value
}
// Sets the replyChainMessageId property value. The ID of the reply message.
// Parameters:
//  - value : Value to set for the replyChainMessageId property.
func (m *ChatInfo) SetReplyChainMessageId(value *string)() {
    m.replyChainMessageId = value
}
// Sets the threadId property value. The unique identifier for a thread in Microsoft Teams.
// Parameters:
//  - value : Value to set for the threadId property.
func (m *ChatInfo) SetThreadId(value *string)() {
    m.threadId = value
}
