package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MessageRulePredicates provides operations to manage the drive singleton.
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
    fromAddresses []Recipientable;
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
    sentToAddresses []Recipientable;
    // Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
    sentToMe *bool;
    // Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
    sentToOrCcMe *bool;
    // Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
    subjectContains []string;
    // Represents the minimum and maximum sizes (in kilobytes) that an incoming message must fall in between in order for the condition or exception to apply.
    withinSizeRange SizeRangeable;
}
// NewMessageRulePredicates instantiates a new messageRulePredicates and sets the default values.
func NewMessageRulePredicates()(*MessageRulePredicates) {
    m := &MessageRulePredicates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMessageRulePredicatesFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageRulePredicatesFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMessageRulePredicates(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRulePredicates) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBodyContains gets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetBodyContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.bodyContains
    }
}
// GetBodyOrSubjectContains gets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetBodyOrSubjectContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.bodyOrSubjectContains
    }
}
// GetCategories gets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetFieldDeserializers the deserialization information for the current model
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
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
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
            m.SetImportance(val.(*Importance))
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
            m.SetMessageActionFlag(val.(*MessageActionFlag))
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
            m.SetSensitivity(val.(*Sensitivity))
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
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
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
        val, err := n.GetObjectValue(CreateSizeRangeFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWithinSizeRange(val.(SizeRangeable))
        }
        return nil
    }
    return res
}
// GetFromAddresses gets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetFromAddresses()([]Recipientable) {
    if m == nil {
        return nil
    } else {
        return m.fromAddresses
    }
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// GetHeaderContains gets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetHeaderContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.headerContains
    }
}
// GetImportance gets the importance property value. The importance that is stamped on an incoming message in order for the condition or exception to apply: low, normal, high.
func (m *MessageRulePredicates) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// GetIsApprovalRequest gets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsApprovalRequest()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequest
    }
}
// GetIsAutomaticForward gets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsAutomaticForward()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAutomaticForward
    }
}
// GetIsAutomaticReply gets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsAutomaticReply()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAutomaticReply
    }
}
// GetIsEncrypted gets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
// GetIsMeetingRequest gets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsMeetingRequest()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMeetingRequest
    }
}
// GetIsMeetingResponse gets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsMeetingResponse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMeetingResponse
    }
}
// GetIsNonDeliveryReport gets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsNonDeliveryReport()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isNonDeliveryReport
    }
}
// GetIsPermissionControlled gets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsPermissionControlled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPermissionControlled
    }
}
// GetIsReadReceipt gets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsReadReceipt()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReadReceipt
    }
}
// GetIsSigned gets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsSigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSigned
    }
}
// GetIsVoicemail gets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetIsVoicemail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVoicemail
    }
}
// GetMessageActionFlag gets the messageActionFlag property value. Represents the flag-for-action value that appears on an incoming message in order for the condition or exception to apply. The possible values are: any, call, doNotForward, followUp, fyi, forward, noResponseNecessary, read, reply, replyToAll, review.
func (m *MessageRulePredicates) GetMessageActionFlag()(*MessageActionFlag) {
    if m == nil {
        return nil
    } else {
        return m.messageActionFlag
    }
}
// GetNotSentToMe gets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetNotSentToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notSentToMe
    }
}
// GetRecipientContains gets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetRecipientContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.recipientContains
    }
}
// GetSenderContains gets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSenderContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.senderContains
    }
}
// GetSensitivity gets the sensitivity property value. Represents the sensitivity level that must be stamped on an incoming message in order for the condition or exception to apply. The possible values are: normal, personal, private, confidential.
func (m *MessageRulePredicates) GetSensitivity()(*Sensitivity) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
// GetSentCcMe gets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentCcMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentCcMe
    }
}
// GetSentOnlyToMe gets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentOnlyToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentOnlyToMe
    }
}
// GetSentToAddresses gets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentToAddresses()([]Recipientable) {
    if m == nil {
        return nil
    } else {
        return m.sentToAddresses
    }
}
// GetSentToMe gets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentToMe
    }
}
// GetSentToOrCcMe gets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSentToOrCcMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentToOrCcMe
    }
}
// GetSubjectContains gets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetSubjectContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.subjectContains
    }
}
// GetWithinSizeRange gets the withinSizeRange property value. Represents the minimum and maximum sizes (in kilobytes) that an incoming message must fall in between in order for the condition or exception to apply.
func (m *MessageRulePredicates) GetWithinSizeRange()(SizeRangeable) {
    if m == nil {
        return nil
    } else {
        return m.withinSizeRange
    }
}
func (m *MessageRulePredicates) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MessageRulePredicates) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetBodyContains() != nil {
        err := writer.WriteCollectionOfStringValues("bodyContains", m.GetBodyContains())
        if err != nil {
            return err
        }
    }
    if m.GetBodyOrSubjectContains() != nil {
        err := writer.WriteCollectionOfStringValues("bodyOrSubjectContains", m.GetBodyOrSubjectContains())
        if err != nil {
            return err
        }
    }
    if m.GetCategories() != nil {
        err := writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    if m.GetFromAddresses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetFromAddresses()))
        for i, v := range m.GetFromAddresses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetHeaderContains() != nil {
        err := writer.WriteCollectionOfStringValues("headerContains", m.GetHeaderContains())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
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
        cast := (*m.GetMessageActionFlag()).String()
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
    if m.GetRecipientContains() != nil {
        err := writer.WriteCollectionOfStringValues("recipientContains", m.GetRecipientContains())
        if err != nil {
            return err
        }
    }
    if m.GetSenderContains() != nil {
        err := writer.WriteCollectionOfStringValues("senderContains", m.GetSenderContains())
        if err != nil {
            return err
        }
    }
    if m.GetSensitivity() != nil {
        cast := (*m.GetSensitivity()).String()
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
    if m.GetSentToAddresses() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSentToAddresses()))
        for i, v := range m.GetSentToAddresses() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetSubjectContains() != nil {
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MessageRulePredicates) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetBodyContains sets the bodyContains property value. Represents the strings that should appear in the body of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetBodyContains(value []string)() {
    if m != nil {
        m.bodyContains = value
    }
}
// SetBodyOrSubjectContains sets the bodyOrSubjectContains property value. Represents the strings that should appear in the body or subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetBodyOrSubjectContains(value []string)() {
    if m != nil {
        m.bodyOrSubjectContains = value
    }
}
// SetCategories sets the categories property value. Represents the categories that an incoming message should be labeled with in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetCategories(value []string)() {
    if m != nil {
        m.categories = value
    }
}
// SetFromAddresses sets the fromAddresses property value. Represents the specific sender email addresses of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetFromAddresses(value []Recipientable)() {
    if m != nil {
        m.fromAddresses = value
    }
}
// SetHasAttachments sets the hasAttachments property value. Indicates whether an incoming message must have attachments in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetHasAttachments(value *bool)() {
    if m != nil {
        m.hasAttachments = value
    }
}
// SetHeaderContains sets the headerContains property value. Represents the strings that appear in the headers of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetHeaderContains(value []string)() {
    if m != nil {
        m.headerContains = value
    }
}
// SetImportance sets the importance property value. The importance that is stamped on an incoming message in order for the condition or exception to apply: low, normal, high.
func (m *MessageRulePredicates) SetImportance(value *Importance)() {
    if m != nil {
        m.importance = value
    }
}
// SetIsApprovalRequest sets the isApprovalRequest property value. Indicates whether an incoming message must be an approval request in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsApprovalRequest(value *bool)() {
    if m != nil {
        m.isApprovalRequest = value
    }
}
// SetIsAutomaticForward sets the isAutomaticForward property value. Indicates whether an incoming message must be automatically forwarded in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsAutomaticForward(value *bool)() {
    if m != nil {
        m.isAutomaticForward = value
    }
}
// SetIsAutomaticReply sets the isAutomaticReply property value. Indicates whether an incoming message must be an auto reply in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsAutomaticReply(value *bool)() {
    if m != nil {
        m.isAutomaticReply = value
    }
}
// SetIsEncrypted sets the isEncrypted property value. Indicates whether an incoming message must be encrypted in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsEncrypted(value *bool)() {
    if m != nil {
        m.isEncrypted = value
    }
}
// SetIsMeetingRequest sets the isMeetingRequest property value. Indicates whether an incoming message must be a meeting request in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsMeetingRequest(value *bool)() {
    if m != nil {
        m.isMeetingRequest = value
    }
}
// SetIsMeetingResponse sets the isMeetingResponse property value. Indicates whether an incoming message must be a meeting response in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsMeetingResponse(value *bool)() {
    if m != nil {
        m.isMeetingResponse = value
    }
}
// SetIsNonDeliveryReport sets the isNonDeliveryReport property value. Indicates whether an incoming message must be a non-delivery report in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsNonDeliveryReport(value *bool)() {
    if m != nil {
        m.isNonDeliveryReport = value
    }
}
// SetIsPermissionControlled sets the isPermissionControlled property value. Indicates whether an incoming message must be permission controlled (RMS-protected) in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsPermissionControlled(value *bool)() {
    if m != nil {
        m.isPermissionControlled = value
    }
}
// SetIsReadReceipt sets the isReadReceipt property value. Indicates whether an incoming message must be a read receipt in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsReadReceipt(value *bool)() {
    if m != nil {
        m.isReadReceipt = value
    }
}
// SetIsSigned sets the isSigned property value. Indicates whether an incoming message must be S/MIME-signed in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsSigned(value *bool)() {
    if m != nil {
        m.isSigned = value
    }
}
// SetIsVoicemail sets the isVoicemail property value. Indicates whether an incoming message must be a voice mail in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetIsVoicemail(value *bool)() {
    if m != nil {
        m.isVoicemail = value
    }
}
// SetMessageActionFlag sets the messageActionFlag property value. Represents the flag-for-action value that appears on an incoming message in order for the condition or exception to apply. The possible values are: any, call, doNotForward, followUp, fyi, forward, noResponseNecessary, read, reply, replyToAll, review.
func (m *MessageRulePredicates) SetMessageActionFlag(value *MessageActionFlag)() {
    if m != nil {
        m.messageActionFlag = value
    }
}
// SetNotSentToMe sets the notSentToMe property value. Indicates whether the owner of the mailbox must not be a recipient of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetNotSentToMe(value *bool)() {
    if m != nil {
        m.notSentToMe = value
    }
}
// SetRecipientContains sets the recipientContains property value. Represents the strings that appear in either the toRecipients or ccRecipients properties of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetRecipientContains(value []string)() {
    if m != nil {
        m.recipientContains = value
    }
}
// SetSenderContains sets the senderContains property value. Represents the strings that appear in the from property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSenderContains(value []string)() {
    if m != nil {
        m.senderContains = value
    }
}
// SetSensitivity sets the sensitivity property value. Represents the sensitivity level that must be stamped on an incoming message in order for the condition or exception to apply. The possible values are: normal, personal, private, confidential.
func (m *MessageRulePredicates) SetSensitivity(value *Sensitivity)() {
    if m != nil {
        m.sensitivity = value
    }
}
// SetSentCcMe sets the sentCcMe property value. Indicates whether the owner of the mailbox must be in the ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentCcMe(value *bool)() {
    if m != nil {
        m.sentCcMe = value
    }
}
// SetSentOnlyToMe sets the sentOnlyToMe property value. Indicates whether the owner of the mailbox must be the only recipient in an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentOnlyToMe(value *bool)() {
    if m != nil {
        m.sentOnlyToMe = value
    }
}
// SetSentToAddresses sets the sentToAddresses property value. Represents the email addresses that an incoming message must have been sent to in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentToAddresses(value []Recipientable)() {
    if m != nil {
        m.sentToAddresses = value
    }
}
// SetSentToMe sets the sentToMe property value. Indicates whether the owner of the mailbox must be in the toRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentToMe(value *bool)() {
    if m != nil {
        m.sentToMe = value
    }
}
// SetSentToOrCcMe sets the sentToOrCcMe property value. Indicates whether the owner of the mailbox must be in either a toRecipients or ccRecipients property of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSentToOrCcMe(value *bool)() {
    if m != nil {
        m.sentToOrCcMe = value
    }
}
// SetSubjectContains sets the subjectContains property value. Represents the strings that appear in the subject of an incoming message in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetSubjectContains(value []string)() {
    if m != nil {
        m.subjectContains = value
    }
}
// SetWithinSizeRange sets the withinSizeRange property value. Represents the minimum and maximum sizes (in kilobytes) that an incoming message must fall in between in order for the condition or exception to apply.
func (m *MessageRulePredicates) SetWithinSizeRange(value SizeRangeable)() {
    if m != nil {
        m.withinSizeRange = value
    }
}
