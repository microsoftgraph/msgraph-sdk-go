package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type InvitedUserMessageInfo struct {
    additionalData map[string]interface{};
    ccRecipients []Recipient;
    customizedMessageBody *string;
    messageLanguage *string;
}
func NewInvitedUserMessageInfo()(*InvitedUserMessageInfo) {
    m := &InvitedUserMessageInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *InvitedUserMessageInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *InvitedUserMessageInfo) GetCcRecipients()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.ccRecipients
    }
}
func (m *InvitedUserMessageInfo) GetCustomizedMessageBody()(*string) {
    if m == nil {
        return nil
    } else {
        return m.customizedMessageBody
    }
}
func (m *InvitedUserMessageInfo) GetMessageLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.messageLanguage
    }
}
func (m *InvitedUserMessageInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["ccRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetCcRecipients(res)
        return nil
    }
    res["customizedMessageBody"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCustomizedMessageBody(val)
        return nil
    }
    res["messageLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessageLanguage(val)
        return nil
    }
    return res
}
func (m *InvitedUserMessageInfo) IsNil()(bool) {
    return m == nil
}
func (m *InvitedUserMessageInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCcRecipients()))
        for i, v := range m.GetCcRecipients() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("ccRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("customizedMessageBody", m.GetCustomizedMessageBody())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("messageLanguage", m.GetMessageLanguage())
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
func (m *InvitedUserMessageInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *InvitedUserMessageInfo) SetCcRecipients(value []Recipient)() {
    m.ccRecipients = value
}
func (m *InvitedUserMessageInfo) SetCustomizedMessageBody(value *string)() {
    m.customizedMessageBody = value
}
func (m *InvitedUserMessageInfo) SetMessageLanguage(value *string)() {
    m.messageLanguage = value
}
