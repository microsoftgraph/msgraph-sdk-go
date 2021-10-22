package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChatMessagePolicyViolation struct {
    additionalData map[string]interface{};
    dlpAction *ChatMessagePolicyViolationDlpActionTypes;
    justificationText *string;
    policyTip *ChatMessagePolicyViolationPolicyTip;
    userAction *ChatMessagePolicyViolationUserActionTypes;
    verdictDetails *ChatMessagePolicyViolationVerdictDetailsTypes;
}
func NewChatMessagePolicyViolation()(*ChatMessagePolicyViolation) {
    m := &ChatMessagePolicyViolation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ChatMessagePolicyViolation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ChatMessagePolicyViolation) GetDlpAction()(*ChatMessagePolicyViolationDlpActionTypes) {
    if m == nil {
        return nil
    } else {
        return m.dlpAction
    }
}
func (m *ChatMessagePolicyViolation) GetJustificationText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justificationText
    }
}
func (m *ChatMessagePolicyViolation) GetPolicyTip()(*ChatMessagePolicyViolationPolicyTip) {
    if m == nil {
        return nil
    } else {
        return m.policyTip
    }
}
func (m *ChatMessagePolicyViolation) GetUserAction()(*ChatMessagePolicyViolationUserActionTypes) {
    if m == nil {
        return nil
    } else {
        return m.userAction
    }
}
func (m *ChatMessagePolicyViolation) GetVerdictDetails()(*ChatMessagePolicyViolationVerdictDetailsTypes) {
    if m == nil {
        return nil
    } else {
        return m.verdictDetails
    }
}
func (m *ChatMessagePolicyViolation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dlpAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationDlpActionTypes)
        if err != nil {
            return err
        }
        cast := val.(ChatMessagePolicyViolationDlpActionTypes)
        m.SetDlpAction(&cast)
        return nil
    }
    res["justificationText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJustificationText(val)
        return nil
    }
    res["policyTip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessagePolicyViolationPolicyTip() })
        if err != nil {
            return err
        }
        m.SetPolicyTip(val.(*ChatMessagePolicyViolationPolicyTip))
        return nil
    }
    res["userAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationUserActionTypes)
        if err != nil {
            return err
        }
        cast := val.(ChatMessagePolicyViolationUserActionTypes)
        m.SetUserAction(&cast)
        return nil
    }
    res["verdictDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationVerdictDetailsTypes)
        if err != nil {
            return err
        }
        cast := val.(ChatMessagePolicyViolationVerdictDetailsTypes)
        m.SetVerdictDetails(&cast)
        return nil
    }
    return res
}
func (m *ChatMessagePolicyViolation) IsNil()(bool) {
    return m == nil
}
func (m *ChatMessagePolicyViolation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDlpAction() != nil {
        cast := m.GetDlpAction().String()
        err := writer.WriteStringValue("dlpAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("justificationText", m.GetJustificationText())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("policyTip", m.GetPolicyTip())
        if err != nil {
            return err
        }
    }
    if m.GetUserAction() != nil {
        cast := m.GetUserAction().String()
        err := writer.WriteStringValue("userAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVerdictDetails() != nil {
        cast := m.GetVerdictDetails().String()
        err := writer.WriteStringValue("verdictDetails", &cast)
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
func (m *ChatMessagePolicyViolation) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ChatMessagePolicyViolation) SetDlpAction(value *ChatMessagePolicyViolationDlpActionTypes)() {
    m.dlpAction = value
}
func (m *ChatMessagePolicyViolation) SetJustificationText(value *string)() {
    m.justificationText = value
}
func (m *ChatMessagePolicyViolation) SetPolicyTip(value *ChatMessagePolicyViolationPolicyTip)() {
    m.policyTip = value
}
func (m *ChatMessagePolicyViolation) SetUserAction(value *ChatMessagePolicyViolationUserActionTypes)() {
    m.userAction = value
}
func (m *ChatMessagePolicyViolation) SetVerdictDetails(value *ChatMessagePolicyViolationVerdictDetailsTypes)() {
    m.verdictDetails = value
}
