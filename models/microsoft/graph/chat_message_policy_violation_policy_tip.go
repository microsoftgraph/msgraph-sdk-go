package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChatMessagePolicyViolationPolicyTip struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn't say in chats)
    complianceUrl *string;
    // Explanatory text shown to the sender of the message.
    generalText *string;
    // The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
    matchedConditionDescriptions []string;
}
// Instantiates a new chatMessagePolicyViolationPolicyTip and sets the default values.
func NewChatMessagePolicyViolationPolicyTip()(*ChatMessagePolicyViolationPolicyTip) {
    m := &ChatMessagePolicyViolationPolicyTip{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessagePolicyViolationPolicyTip) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the complianceUrl property value. The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn't say in chats)
func (m *ChatMessagePolicyViolationPolicyTip) GetComplianceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.complianceUrl
    }
}
// Gets the generalText property value. Explanatory text shown to the sender of the message.
func (m *ChatMessagePolicyViolationPolicyTip) GetGeneralText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.generalText
    }
}
// Gets the matchedConditionDescriptions property value. The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
func (m *ChatMessagePolicyViolationPolicyTip) GetMatchedConditionDescriptions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.matchedConditionDescriptions
    }
}
// The deserialization information for the current model
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
            res[i] = *(v.(*string))
        }
        m.SetMatchedConditionDescriptions(res)
        return nil
    }
    return res
}
func (m *ChatMessagePolicyViolationPolicyTip) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ChatMessagePolicyViolationPolicyTip) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the complianceUrl property value. The URL a user can visit to read about the data loss prevention policies for the organization. (ie, policies about what users shouldn't say in chats)
// Parameters:
//  - value : Value to set for the complianceUrl property.
func (m *ChatMessagePolicyViolationPolicyTip) SetComplianceUrl(value *string)() {
    m.complianceUrl = value
}
// Sets the generalText property value. Explanatory text shown to the sender of the message.
// Parameters:
//  - value : Value to set for the generalText property.
func (m *ChatMessagePolicyViolationPolicyTip) SetGeneralText(value *string)() {
    m.generalText = value
}
// Sets the matchedConditionDescriptions property value. The list of improper data in the message that was detected by the data loss prevention app. Each DLP app defines its own conditions, examples include 'Credit Card Number' and 'Social Security Number'.
// Parameters:
//  - value : Value to set for the matchedConditionDescriptions property.
func (m *ChatMessagePolicyViolationPolicyTip) SetMatchedConditionDescriptions(value []string)() {
    m.matchedConditionDescriptions = value
}
