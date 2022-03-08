package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChatMessagePolicyViolation provides operations to manage the collection of chat entities.
type ChatMessagePolicyViolation struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The action taken by the DLP provider on the message with sensitive content. Supported values are: NoneNotifySender -- Inform the sender of the violation but allow readers to read the message.BlockAccess -- Block readers from reading the message.BlockAccessExternal -- Block users outside the organization from reading the message, while allowing users within the organization to read the message.
    dlpAction *ChatMessagePolicyViolationDlpActionTypes;
    // Justification text provided by the sender of the message when overriding a policy violation.
    justificationText *string;
    // Information to display to the message sender about why the message was flagged as a violation.
    policyTip ChatMessagePolicyViolationPolicyTipable;
    // Indicates the action taken by the user on a message blocked by the DLP provider. Supported values are: NoneOverrideReportFalsePositiveWhen the DLP provider is updating the message for blocking sensitive content, userAction is not required.
    userAction *ChatMessagePolicyViolationUserActionTypes;
    // Indicates what actions the sender may take in response to the policy violation. Supported values are: NoneAllowFalsePositiveOverride -- Allows the sender to declare the policyViolation to be an error in the DLP app and its rules, and allow readers to see the message again if the dlpAction had hidden it.AllowOverrideWithoutJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, without needing to provide an explanation for doing so. AllowOverrideWithJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, after providing an explanation for doing so.AllowOverrideWithoutJustification and AllowOverrideWithJustification are mutually exclusive.
    verdictDetails *ChatMessagePolicyViolationVerdictDetailsTypes;
}
// NewChatMessagePolicyViolation instantiates a new chatMessagePolicyViolation and sets the default values.
func NewChatMessagePolicyViolation()(*ChatMessagePolicyViolation) {
    m := &ChatMessagePolicyViolation{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateChatMessagePolicyViolationFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessagePolicyViolationFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewChatMessagePolicyViolation(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessagePolicyViolation) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDlpAction gets the dlpAction property value. The action taken by the DLP provider on the message with sensitive content. Supported values are: NoneNotifySender -- Inform the sender of the violation but allow readers to read the message.BlockAccess -- Block readers from reading the message.BlockAccessExternal -- Block users outside the organization from reading the message, while allowing users within the organization to read the message.
func (m *ChatMessagePolicyViolation) GetDlpAction()(*ChatMessagePolicyViolationDlpActionTypes) {
    if m == nil {
        return nil
    } else {
        return m.dlpAction
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessagePolicyViolation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["dlpAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationDlpActionTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDlpAction(val.(*ChatMessagePolicyViolationDlpActionTypes))
        }
        return nil
    }
    res["justificationText"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJustificationText(val)
        }
        return nil
    }
    res["policyTip"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessagePolicyViolationPolicyTipFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyTip(val.(ChatMessagePolicyViolationPolicyTipable))
        }
        return nil
    }
    res["userAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationUserActionTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserAction(val.(*ChatMessagePolicyViolationUserActionTypes))
        }
        return nil
    }
    res["verdictDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessagePolicyViolationVerdictDetailsTypes)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVerdictDetails(val.(*ChatMessagePolicyViolationVerdictDetailsTypes))
        }
        return nil
    }
    return res
}
// GetJustificationText gets the justificationText property value. Justification text provided by the sender of the message when overriding a policy violation.
func (m *ChatMessagePolicyViolation) GetJustificationText()(*string) {
    if m == nil {
        return nil
    } else {
        return m.justificationText
    }
}
// GetPolicyTip gets the policyTip property value. Information to display to the message sender about why the message was flagged as a violation.
func (m *ChatMessagePolicyViolation) GetPolicyTip()(ChatMessagePolicyViolationPolicyTipable) {
    if m == nil {
        return nil
    } else {
        return m.policyTip
    }
}
// GetUserAction gets the userAction property value. Indicates the action taken by the user on a message blocked by the DLP provider. Supported values are: NoneOverrideReportFalsePositiveWhen the DLP provider is updating the message for blocking sensitive content, userAction is not required.
func (m *ChatMessagePolicyViolation) GetUserAction()(*ChatMessagePolicyViolationUserActionTypes) {
    if m == nil {
        return nil
    } else {
        return m.userAction
    }
}
// GetVerdictDetails gets the verdictDetails property value. Indicates what actions the sender may take in response to the policy violation. Supported values are: NoneAllowFalsePositiveOverride -- Allows the sender to declare the policyViolation to be an error in the DLP app and its rules, and allow readers to see the message again if the dlpAction had hidden it.AllowOverrideWithoutJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, without needing to provide an explanation for doing so. AllowOverrideWithJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, after providing an explanation for doing so.AllowOverrideWithoutJustification and AllowOverrideWithJustification are mutually exclusive.
func (m *ChatMessagePolicyViolation) GetVerdictDetails()(*ChatMessagePolicyViolationVerdictDetailsTypes) {
    if m == nil {
        return nil
    } else {
        return m.verdictDetails
    }
}
func (m *ChatMessagePolicyViolation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChatMessagePolicyViolation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetDlpAction() != nil {
        cast := (*m.GetDlpAction()).String()
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
        cast := (*m.GetUserAction()).String()
        err := writer.WriteStringValue("userAction", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetVerdictDetails() != nil {
        cast := (*m.GetVerdictDetails()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ChatMessagePolicyViolation) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDlpAction sets the dlpAction property value. The action taken by the DLP provider on the message with sensitive content. Supported values are: NoneNotifySender -- Inform the sender of the violation but allow readers to read the message.BlockAccess -- Block readers from reading the message.BlockAccessExternal -- Block users outside the organization from reading the message, while allowing users within the organization to read the message.
func (m *ChatMessagePolicyViolation) SetDlpAction(value *ChatMessagePolicyViolationDlpActionTypes)() {
    if m != nil {
        m.dlpAction = value
    }
}
// SetJustificationText sets the justificationText property value. Justification text provided by the sender of the message when overriding a policy violation.
func (m *ChatMessagePolicyViolation) SetJustificationText(value *string)() {
    if m != nil {
        m.justificationText = value
    }
}
// SetPolicyTip sets the policyTip property value. Information to display to the message sender about why the message was flagged as a violation.
func (m *ChatMessagePolicyViolation) SetPolicyTip(value ChatMessagePolicyViolationPolicyTipable)() {
    if m != nil {
        m.policyTip = value
    }
}
// SetUserAction sets the userAction property value. Indicates the action taken by the user on a message blocked by the DLP provider. Supported values are: NoneOverrideReportFalsePositiveWhen the DLP provider is updating the message for blocking sensitive content, userAction is not required.
func (m *ChatMessagePolicyViolation) SetUserAction(value *ChatMessagePolicyViolationUserActionTypes)() {
    if m != nil {
        m.userAction = value
    }
}
// SetVerdictDetails sets the verdictDetails property value. Indicates what actions the sender may take in response to the policy violation. Supported values are: NoneAllowFalsePositiveOverride -- Allows the sender to declare the policyViolation to be an error in the DLP app and its rules, and allow readers to see the message again if the dlpAction had hidden it.AllowOverrideWithoutJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, without needing to provide an explanation for doing so. AllowOverrideWithJustification -- Allows the sender to overriide the DLP violation and allow readers to see the message again if the dlpAction had hidden it, after providing an explanation for doing so.AllowOverrideWithoutJustification and AllowOverrideWithJustification are mutually exclusive.
func (m *ChatMessagePolicyViolation) SetVerdictDetails(value *ChatMessagePolicyViolationVerdictDetailsTypes)() {
    if m != nil {
        m.verdictDetails = value
    }
}
