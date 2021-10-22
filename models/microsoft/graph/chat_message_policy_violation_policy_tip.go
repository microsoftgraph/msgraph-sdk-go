package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChatMessagePolicyViolationPolicyTip struct {
    additionalData map[string]interface{};
    complianceUrl *string;
    generalText *string;
    matchedConditionDescriptions []string;
}
func NewChatMessagePolicyViolationPolicyTip()(*ChatMessagePolicyViolationPolicyTip) {
    m := &ChatMessagePolicyViolationPolicyTip{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChatMessagePolicyViolationPolicyTip) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChatMessagePolicyViolationPolicyTip) GetComplianceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.complianceUrl
    }
}
func (m *ChatMessagePolicyViolationPolicyTip) GetGeneralText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.generalText
    }
}
func (m *ChatMessagePolicyViolationPolicyTip) GetMatchedConditionDescriptions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.matchedConditionDescriptions
    }
}
func (m *ChatMessagePolicyViolationPolicyTip) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["complianceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetComplianceUrl(val)
        return nil
    }
    res["generalText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetGeneralText(val)
        return nil
    }
    res["matchedConditionDescriptions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetMatchedConditionDescriptions(res)
        return nil
    }
    return res
}
func (m *ChatMessagePolicyViolationPolicyTip) IsNil()(bool) {
    return m == nil
}
func (m *ChatMessagePolicyViolationPolicyTip) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("complianceUrl", m.GetComplianceUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("generalText", m.GetGeneralText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("matchedConditionDescriptions", m.GetMatchedConditionDescriptions())
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
func (m *ChatMessagePolicyViolationPolicyTip) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChatMessagePolicyViolationPolicyTip) SetComplianceUrl(value *string)() {
    m.complianceUrl = value
}
func (m *ChatMessagePolicyViolationPolicyTip) SetGeneralText(value *string)() {
    m.generalText = value
}
func (m *ChatMessagePolicyViolationPolicyTip) SetMatchedConditionDescriptions(value []string)() {
    m.matchedConditionDescriptions = value
}
