package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChatInfo struct {
    additionalData map[string]interface{};
    messageId *string;
    replyChainMessageId *string;
    threadId *string;
}
func NewChatInfo()(*ChatInfo) {
    m := &ChatInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChatInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChatInfo) GetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageId
    }
}
func (m *ChatInfo) GetReplyChainMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.replyChainMessageId
    }
}
func (m *ChatInfo) GetThreadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.threadId
    }
}
func (m *ChatInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["messageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessageId(val)
        return nil
    }
    res["replyChainMessageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReplyChainMessageId(val)
        return nil
    }
    res["threadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetThreadId(val)
        return nil
    }
    return res
}
func (m *ChatInfo) IsNil()(bool) {
    return m == nil
}
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
func (m *ChatInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChatInfo) SetMessageId(value *string)() {
    m.messageId = value
}
func (m *ChatInfo) SetReplyChainMessageId(value *string)() {
    m.replyChainMessageId = value
}
func (m *ChatInfo) SetThreadId(value *string)() {
    m.threadId = value
}
