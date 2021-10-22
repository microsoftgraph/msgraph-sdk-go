package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MessageRulePredicates struct {
    additionalData map[string]interface{};
    bodyContains []string;
    bodyOrSubjectContains []string;
    categories []string;
    fromAddresses []Recipient;
    hasAttachments *bool;
    headerContains []string;
    importance *Importance;
    isApprovalRequest *bool;
    isAutomaticForward *bool;
    isAutomaticReply *bool;
    isEncrypted *bool;
    isMeetingRequest *bool;
    isMeetingResponse *bool;
    isNonDeliveryReport *bool;
    isPermissionControlled *bool;
    isReadReceipt *bool;
    isSigned *bool;
    isVoicemail *bool;
    messageActionFlag *MessageActionFlag;
    notSentToMe *bool;
    recipientContains []string;
    senderContains []string;
    sensitivity *Sensitivity;
    sentCcMe *bool;
    sentOnlyToMe *bool;
    sentToAddresses []Recipient;
    sentToMe *bool;
    sentToOrCcMe *bool;
    subjectContains []string;
    withinSizeRange *SizeRange;
}
func NewMessageRulePredicates()(*MessageRulePredicates) {
    m := &MessageRulePredicates{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MessageRulePredicates) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MessageRulePredicates) GetBodyContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.bodyContains
    }
}
func (m *MessageRulePredicates) GetBodyOrSubjectContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.bodyOrSubjectContains
    }
}
func (m *MessageRulePredicates) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
func (m *MessageRulePredicates) GetFromAddresses()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.fromAddresses
    }
}
func (m *MessageRulePredicates) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
func (m *MessageRulePredicates) GetHeaderContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.headerContains
    }
}
func (m *MessageRulePredicates) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
func (m *MessageRulePredicates) GetIsApprovalRequest()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isApprovalRequest
    }
}
func (m *MessageRulePredicates) GetIsAutomaticForward()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAutomaticForward
    }
}
func (m *MessageRulePredicates) GetIsAutomaticReply()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAutomaticReply
    }
}
func (m *MessageRulePredicates) GetIsEncrypted()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEncrypted
    }
}
func (m *MessageRulePredicates) GetIsMeetingRequest()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMeetingRequest
    }
}
func (m *MessageRulePredicates) GetIsMeetingResponse()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isMeetingResponse
    }
}
func (m *MessageRulePredicates) GetIsNonDeliveryReport()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isNonDeliveryReport
    }
}
func (m *MessageRulePredicates) GetIsPermissionControlled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isPermissionControlled
    }
}
func (m *MessageRulePredicates) GetIsReadReceipt()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReadReceipt
    }
}
func (m *MessageRulePredicates) GetIsSigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isSigned
    }
}
func (m *MessageRulePredicates) GetIsVoicemail()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVoicemail
    }
}
func (m *MessageRulePredicates) GetMessageActionFlag()(*MessageActionFlag) {
    if m == nil {
        return nil
    } else {
        return m.messageActionFlag
    }
}
func (m *MessageRulePredicates) GetNotSentToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.notSentToMe
    }
}
func (m *MessageRulePredicates) GetRecipientContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.recipientContains
    }
}
func (m *MessageRulePredicates) GetSenderContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.senderContains
    }
}
func (m *MessageRulePredicates) GetSensitivity()(*Sensitivity) {
    if m == nil {
        return nil
    } else {
        return m.sensitivity
    }
}
func (m *MessageRulePredicates) GetSentCcMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentCcMe
    }
}
func (m *MessageRulePredicates) GetSentOnlyToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentOnlyToMe
    }
}
func (m *MessageRulePredicates) GetSentToAddresses()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.sentToAddresses
    }
}
func (m *MessageRulePredicates) GetSentToMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentToMe
    }
}
func (m *MessageRulePredicates) GetSentToOrCcMe()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.sentToOrCcMe
    }
}
func (m *MessageRulePredicates) GetSubjectContains()([]string) {
    if m == nil {
        return nil
    } else {
        return m.subjectContains
    }
}
func (m *MessageRulePredicates) GetWithinSizeRange()(*SizeRange) {
    if m == nil {
        return nil
    } else {
        return m.withinSizeRange
    }
}
func (m *MessageRulePredicates) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bodyContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetBodyContains(res)
        return nil
    }
    res["bodyOrSubjectContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetBodyOrSubjectContains(res)
        return nil
    }
    res["categories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetCategories(res)
        return nil
    }
    res["fromAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetFromAddresses(res)
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasAttachments(val)
        return nil
    }
    res["headerContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetHeaderContains(res)
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        cast := val.(Importance)
        m.SetImportance(&cast)
        return nil
    }
    res["isApprovalRequest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsApprovalRequest(val)
        return nil
    }
    res["isAutomaticForward"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAutomaticForward(val)
        return nil
    }
    res["isAutomaticReply"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAutomaticReply(val)
        return nil
    }
    res["isEncrypted"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEncrypted(val)
        return nil
    }
    res["isMeetingRequest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMeetingRequest(val)
        return nil
    }
    res["isMeetingResponse"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsMeetingResponse(val)
        return nil
    }
    res["isNonDeliveryReport"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsNonDeliveryReport(val)
        return nil
    }
    res["isPermissionControlled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsPermissionControlled(val)
        return nil
    }
    res["isReadReceipt"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsReadReceipt(val)
        return nil
    }
    res["isSigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsSigned(val)
        return nil
    }
    res["isVoicemail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsVoicemail(val)
        return nil
    }
    res["messageActionFlag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMessageActionFlag)
        if err != nil {
            return err
        }
        cast := val.(MessageActionFlag)
        m.SetMessageActionFlag(&cast)
        return nil
    }
    res["notSentToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetNotSentToMe(val)
        return nil
    }
    res["recipientContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetRecipientContains(res)
        return nil
    }
    res["senderContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSenderContains(res)
        return nil
    }
    res["sensitivity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSensitivity)
        if err != nil {
            return err
        }
        cast := val.(Sensitivity)
        m.SetSensitivity(&cast)
        return nil
    }
    res["sentCcMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSentCcMe(val)
        return nil
    }
    res["sentOnlyToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSentOnlyToMe(val)
        return nil
    }
    res["sentToAddresses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        res := make([]Recipient, len(val))
        for i, v := range val {
            res[i] = *(v.(*Recipient))
        }
        m.SetSentToAddresses(res)
        return nil
    }
    res["sentToMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSentToMe(val)
        return nil
    }
    res["sentToOrCcMe"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSentToOrCcMe(val)
        return nil
    }
    res["subjectContains"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetSubjectContains(res)
        return nil
    }
    res["withinSizeRange"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSizeRange() })
        if err != nil {
            return err
        }
        m.SetWithinSizeRange(val.(*SizeRange))
        return nil
    }
    return res
}
func (m *MessageRulePredicates) IsNil()(bool) {
    return m == nil
}
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
func (m *MessageRulePredicates) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MessageRulePredicates) SetBodyContains(value []string)() {
    m.bodyContains = value
}
func (m *MessageRulePredicates) SetBodyOrSubjectContains(value []string)() {
    m.bodyOrSubjectContains = value
}
func (m *MessageRulePredicates) SetCategories(value []string)() {
    m.categories = value
}
func (m *MessageRulePredicates) SetFromAddresses(value []Recipient)() {
    m.fromAddresses = value
}
func (m *MessageRulePredicates) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
func (m *MessageRulePredicates) SetHeaderContains(value []string)() {
    m.headerContains = value
}
func (m *MessageRulePredicates) SetImportance(value *Importance)() {
    m.importance = value
}
func (m *MessageRulePredicates) SetIsApprovalRequest(value *bool)() {
    m.isApprovalRequest = value
}
func (m *MessageRulePredicates) SetIsAutomaticForward(value *bool)() {
    m.isAutomaticForward = value
}
func (m *MessageRulePredicates) SetIsAutomaticReply(value *bool)() {
    m.isAutomaticReply = value
}
func (m *MessageRulePredicates) SetIsEncrypted(value *bool)() {
    m.isEncrypted = value
}
func (m *MessageRulePredicates) SetIsMeetingRequest(value *bool)() {
    m.isMeetingRequest = value
}
func (m *MessageRulePredicates) SetIsMeetingResponse(value *bool)() {
    m.isMeetingResponse = value
}
func (m *MessageRulePredicates) SetIsNonDeliveryReport(value *bool)() {
    m.isNonDeliveryReport = value
}
func (m *MessageRulePredicates) SetIsPermissionControlled(value *bool)() {
    m.isPermissionControlled = value
}
func (m *MessageRulePredicates) SetIsReadReceipt(value *bool)() {
    m.isReadReceipt = value
}
func (m *MessageRulePredicates) SetIsSigned(value *bool)() {
    m.isSigned = value
}
func (m *MessageRulePredicates) SetIsVoicemail(value *bool)() {
    m.isVoicemail = value
}
func (m *MessageRulePredicates) SetMessageActionFlag(value *MessageActionFlag)() {
    m.messageActionFlag = value
}
func (m *MessageRulePredicates) SetNotSentToMe(value *bool)() {
    m.notSentToMe = value
}
func (m *MessageRulePredicates) SetRecipientContains(value []string)() {
    m.recipientContains = value
}
func (m *MessageRulePredicates) SetSenderContains(value []string)() {
    m.senderContains = value
}
func (m *MessageRulePredicates) SetSensitivity(value *Sensitivity)() {
    m.sensitivity = value
}
func (m *MessageRulePredicates) SetSentCcMe(value *bool)() {
    m.sentCcMe = value
}
func (m *MessageRulePredicates) SetSentOnlyToMe(value *bool)() {
    m.sentOnlyToMe = value
}
func (m *MessageRulePredicates) SetSentToAddresses(value []Recipient)() {
    m.sentToAddresses = value
}
func (m *MessageRulePredicates) SetSentToMe(value *bool)() {
    m.sentToMe = value
}
func (m *MessageRulePredicates) SetSentToOrCcMe(value *bool)() {
    m.sentToOrCcMe = value
}
func (m *MessageRulePredicates) SetSubjectContains(value []string)() {
    m.subjectContains = value
}
func (m *MessageRulePredicates) SetWithinSizeRange(value *SizeRange)() {
    m.withinSizeRange = value
}
