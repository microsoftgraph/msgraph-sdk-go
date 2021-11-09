package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MessageRulePredicates struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
    bodyContains []string;
    // Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
    bodyOrSubjectContains []string;
    // Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
    categories []string;
    // Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
    fromAddresses []Recipient;
    // Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
    hasAttachments *bool;
    // Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
    headerContains []string;
    // The importance that is stamped on an incoming message in order for the condition or exception to apply: low, normal, high.
    importance *Importance;
    // Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
    isApprovalRequest *bool;
    // Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
    isAutomaticForward *bool;
    // Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
    isAutomaticReply *bool;
    // Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
    isEncrypted *bool;
    // Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
    isMeetingRequest *bool;
    // Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
    isMeetingResponse *bool;
    // Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
    isNonDeliveryReport *bool;
    // Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
    isPermissionControlled *bool;
    // Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
    isReadReceipt *bool;
    // Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
    isSigned *bool;
    // Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
    isVoicemail *bool;
    // Represents the flag-for-action value that appears on an incoming message in order for the condition or exception to apply. The possible values are: any, call, doNotForward, followUp, fyi, forward, noResponseNecessary, read, reply, replyToAll, review.
    messageActionFlag *MessageActionFlag;
    // Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
    notSentToMe *bool;
    // Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
    recipientContains []string;
    // Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
    senderContains []string;
    // Represents the sensitivity level that must be stamped on an incoming message in order for the condition or exception to apply. The possible values are: normal, personal, private, confidential.
    sensitivity *Sensitivity;
    // Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
    sentCcMe *bool;
    // Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
    sentOnlyToMe *bool;
    // Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
    sentToAddresses []Recipient;
    // Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
    sentToMe *bool;
    // Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
    sentToOrCcMe *bool;
    // Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
    subjectContains []string;
    // Represents the minimum and maximum sizes (in kilobytes) that an incoming message must fall in between in order for the condition or exception to apply.
    withinSizeRange *SizeRange;
}
// Instantiates a new messageRulePredicates and sets the default values.
func NewMessageRulePredicates()(*MessageRulePredicates) {
    m := &MessageRulePredicates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRulePredicates) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetBodyContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.bodyContains
    }
}
// Gets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetBodyOrSubjectContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.bodyOrSubjectContains
    }
}
// Gets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// Gets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetFromAddresses()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.fromAddresses
    }
}
// Gets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// Gets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetHeaderContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.headerContains
    }
}
// Gets the importance property value. The importance that is stamped on an incoming message in order for the condition or exception to apply: low, normal, high.
func (m *MessageRulePredicates) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// Gets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsApprovalRequest()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequest
    }
}
// Gets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsAutomaticForward()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAutomaticForward
    }
}
// Gets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsAutomaticReply()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAutomaticReply
    }
}
// Gets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
// Gets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsMeetingRequest()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMeetingRequest
    }
}
// Gets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsMeetingResponse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMeetingResponse
    }
}
// Gets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsNonDeliveryReport()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isNonDeliveryReport
    }
}
// Gets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsPermissionControlled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPermissionControlled
    }
}
// Gets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsReadReceipt()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReadReceipt
    }
}
// Gets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsSigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSigned
    }
}
// Gets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsVoicemail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVoicemail
    }
}
// Gets the messageActionFlag property value. Represents the flag-for-action value that appears on an incoming message in order for the condition or exception to apply. The possible values are: any, call, doNotForward, followUp, fyi, forward, noResponseNecessary, read, reply, replyToAll, review.
func (m *MessageRulePredicates) GetMessageActionFlag()(*MessageActionFlag) {
    if m == nil {
        return nil
    } else {
        return m.messageActionFlag
    }
}
// Gets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetNotSentToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notSentToMe
    }
}
// Gets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetRecipientContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.recipientContains
    }
}
// Gets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSenderContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.senderContains
    }
}
// Gets the sensitivity property value. Represents the sensitivity level that must be stamped on an incoming message in order for the condition or exception to apply. The possible values are: normal, personal, private, confidential.
func (m *MessageRulePredicates) GetSensitivity()(*Sensitivity) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// Gets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentCcMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentCcMe
    }
}
// Gets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentOnlyToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentOnlyToMe
    }
}
// Gets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentToAddresses()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.sentToAddresses
    }
}
// Gets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentToMe
    }
}
// Gets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentToOrCcMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentToOrCcMe
    }
}
// Gets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSubjectContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.subjectContains
    }
}
// Gets the withinSizeRange property value. Represents the minimum and maximum sizes (in kilobytes) that an incoming message must fall in between in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetWithinSizeRange()(*SizeRange) {
    if m == nil {
        return nil
    } else {
        return m.withinSizeRange
    }
}
// The deserialization information for the current model
func (m *MessageRulePredicates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bodyContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBodyContains(res)
        }
        return nil
    }
    res["bodyOrSubjectContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetBodyOrSubjectContains(res)
        }
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["fromAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetFromAddresses(res)
        }
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasAttachments(val)
        }
        return nil
    }
    res["headerContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetHeaderContains(res)
        }
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(Importance)
            m.SetImportance(&cast)
        }
        return nil
    }
    res["isApprovalRequest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsApprovalRequest(val)
        }
        return nil
    }
    res["isAutomaticForward"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutomaticForward(val)
        }
        return nil
    }
    res["isAutomaticReply"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAutomaticReply(val)
        }
        return nil
    }
    res["isEncrypted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEncrypted(val)
        }
        return nil
    }
    res["isMeetingRequest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMeetingRequest(val)
        }
        return nil
    }
    res["isMeetingResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsMeetingResponse(val)
        }
        return nil
    }
    res["isNonDeliveryReport"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsNonDeliveryReport(val)
        }
        return nil
    }
    res["isPermissionControlled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsPermissionControlled(val)
        }
        return nil
    }
    res["isReadReceipt"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReadReceipt(val)
        }
        return nil
    }
    res["isSigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsSigned(val)
        }
        return nil
    }
    res["isVoicemail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVoicemail(val)
        }
        return nil
    }
    res["messageActionFlag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMessageActionFlag)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MessageActionFlag)
            m.SetMessageActionFlag(&cast)
        }
        return nil
    }
    res["notSentToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotSentToMe(val)
        }
        return nil
    }
    res["recipientContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRecipientContains(res)
        }
        return nil
    }
    res["senderContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSenderContains(res)
        }
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivity)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(Sensitivity)
            m.SetSensitivity(&cast)
        }
        return nil
    }
    res["sentCcMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentCcMe(val)
        }
        return nil
    }
    res["sentOnlyToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentOnlyToMe(val)
        }
        return nil
    }
    res["sentToAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetSentToAddresses(res)
        }
        return nil
    }
    res["sentToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentToMe(val)
        }
        return nil
    }
    res["sentToOrCcMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentToOrCcMe(val)
        }
        return nil
    }
    res["subjectContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetSubjectContains(res)
        }
        return nil
    }
    res["withinSizeRange"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSizeRange() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWithinSizeRange(val.(*SizeRange))
        }
        return nil
    }
    return res
}
func (m *MessageRulePredicates) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MessageRulePredicates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("bodyContains", m.GetBodyContains())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("bodyOrSubjectContains", m.GetBodyOrSubjectContains())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFromAddresses()))
        for i, v := range m.GetFromAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("fromAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("headerContains", m.GetHeaderContains())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := m.GetImportance().String()
        err := writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isApprovalRequest", m.GetIsApprovalRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAutomaticForward", m.GetIsAutomaticForward())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAutomaticReply", m.GetIsAutomaticReply())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isEncrypted", m.GetIsEncrypted())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMeetingRequest", m.GetIsMeetingRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isMeetingResponse", m.GetIsMeetingResponse())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isNonDeliveryReport", m.GetIsNonDeliveryReport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isPermissionControlled", m.GetIsPermissionControlled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isReadReceipt", m.GetIsReadReceipt())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isSigned", m.GetIsSigned())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVoicemail", m.GetIsVoicemail())
        if err != nil {
            return err
        }
    }
    if m.GetMessageActionFlag() != nil {
        cast := m.GetMessageActionFlag().String()
        err := writer.WriteStringValue("messageActionFlag", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("notSentToMe", m.GetNotSentToMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("recipientContains", m.GetRecipientContains())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("senderContains", m.GetSenderContains())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivity() != nil {
        cast := m.GetSensitivity().String()
        err := writer.WriteStringValue("sensitivity", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sentCcMe", m.GetSentCcMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sentOnlyToMe", m.GetSentOnlyToMe())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSentToAddresses()))
        for i, v := range m.GetSentToAddresses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("sentToAddresses", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sentToMe", m.GetSentToMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("sentToOrCcMe", m.GetSentToOrCcMe())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("subjectContains", m.GetSubjectContains())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("withinSizeRange", m.GetWithinSizeRange())
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
func (m *MessageRulePredicates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the bodyContains property.
func (m *MessageRulePredicates) SetBodyContains(value []string)() {
    m.bodyContains = value
}
// Sets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the bodyOrSubjectContains property.
func (m *MessageRulePredicates) SetBodyOrSubjectContains(value []string)() {
    m.bodyOrSubjectContains = value
}
// Sets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the categories property.
func (m *MessageRulePredicates) SetCategories(value []string)() {
    m.categories = value
}
// Sets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the fromAddresses property.
func (m *MessageRulePredicates) SetFromAddresses(value []Recipient)() {
    m.fromAddresses = value
}
// Sets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the hasAttachments property.
func (m *MessageRulePredicates) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// Sets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the headerContains property.
func (m *MessageRulePredicates) SetHeaderContains(value []string)() {
    m.headerContains = value
}
// Sets the importance property value. The importance that is stamped on an incoming message in order for the condition or exception to apply: low, normal, high.
// Parameters:
//  - value : Value to set for the importance property.
func (m *MessageRulePredicates) SetImportance(value *Importance)() {
    m.importance = value
}
// Sets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isApprovalRequest property.
func (m *MessageRulePredicates) SetIsApprovalRequest(value *bool)() {
    m.isApprovalRequest = value
}
// Sets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isAutomaticForward property.
func (m *MessageRulePredicates) SetIsAutomaticForward(value *bool)() {
    m.isAutomaticForward = value
}
// Sets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isAutomaticReply property.
func (m *MessageRulePredicates) SetIsAutomaticReply(value *bool)() {
    m.isAutomaticReply = value
}
// Sets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isEncrypted property.
func (m *MessageRulePredicates) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
// Sets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isMeetingRequest property.
func (m *MessageRulePredicates) SetIsMeetingRequest(value *bool)() {
    m.isMeetingRequest = value
}
// Sets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isMeetingResponse property.
func (m *MessageRulePredicates) SetIsMeetingResponse(value *bool)() {
    m.isMeetingResponse = value
}
// Sets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isNonDeliveryReport property.
func (m *MessageRulePredicates) SetIsNonDeliveryReport(value *bool)() {
    m.isNonDeliveryReport = value
}
// Sets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isPermissionControlled property.
func (m *MessageRulePredicates) SetIsPermissionControlled(value *bool)() {
    m.isPermissionControlled = value
}
// Sets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isReadReceipt property.
func (m *MessageRulePredicates) SetIsReadReceipt(value *bool)() {
    m.isReadReceipt = value
}
// Sets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isSigned property.
func (m *MessageRulePredicates) SetIsSigned(value *bool)() {
    m.isSigned = value
}
// Sets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the isVoicemail property.
func (m *MessageRulePredicates) SetIsVoicemail(value *bool)() {
    m.isVoicemail = value
}
// Sets the messageActionFlag property value. Represents the flag-for-action value that appears on an incoming message in order for the condition or exception to apply. The possible values are: any, call, doNotForward, followUp, fyi, forward, noResponseNecessary, read, reply, replyToAll, review.
// Parameters:
//  - value : Value to set for the messageActionFlag property.
func (m *MessageRulePredicates) SetMessageActionFlag(value *MessageActionFlag)() {
    m.messageActionFlag = value
}
// Sets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the notSentToMe property.
func (m *MessageRulePredicates) SetNotSentToMe(value *bool)() {
    m.notSentToMe = value
}
// Sets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the recipientContains property.
func (m *MessageRulePredicates) SetRecipientContains(value []string)() {
    m.recipientContains = value
}
// Sets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the senderContains property.
func (m *MessageRulePredicates) SetSenderContains(value []string)() {
    m.senderContains = value
}
// Sets the sensitivity property value. Represents the sensitivity level that must be stamped on an incoming message in order for the condition or exception to apply. The possible values are: normal, personal, private, confidential.
// Parameters:
//  - value : Value to set for the sensitivity property.
func (m *MessageRulePredicates) SetSensitivity(value *Sensitivity)() {
    m.sensitivity = value
}
// Sets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the sentCcMe property.
func (m *MessageRulePredicates) SetSentCcMe(value *bool)() {
    m.sentCcMe = value
}
// Sets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the sentOnlyToMe property.
func (m *MessageRulePredicates) SetSentOnlyToMe(value *bool)() {
    m.sentOnlyToMe = value
}
// Sets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the sentToAddresses property.
func (m *MessageRulePredicates) SetSentToAddresses(value []Recipient)() {
    m.sentToAddresses = value
}
// Sets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the sentToMe property.
func (m *MessageRulePredicates) SetSentToMe(value *bool)() {
    m.sentToMe = value
}
// Sets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the sentToOrCcMe property.
func (m *MessageRulePredicates) SetSentToOrCcMe(value *bool)() {
    m.sentToOrCcMe = value
}
// Sets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the subjectContains property.
func (m *MessageRulePredicates) SetSubjectContains(value []string)() {
    m.subjectContains = value
}
// Sets the withinSizeRange property value. Represents the minimum and maximum sizes (in kilobytes) that an incoming message must fall in between in order for the condition or exception to apply.
// Parameters:
//  - value : Value to set for the withinSizeRange property.
func (m *MessageRulePredicates) SetWithinSizeRange(value *SizeRange)() {
    m.withinSizeRange = value
}
