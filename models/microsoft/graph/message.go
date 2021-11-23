package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// message 
type Message struct {
    OutlookItem
    // The fileAttachment and itemAttachment attachments for the message.
    attachments []Attachment;
    // The Bcc: recipients for the message.
    bccRecipients []Recipient;
    // The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
    body *ItemBody;
    // The first 255 characters of the message body. It is in text format.
    bodyPreview *string;
    // The Cc: recipients for the message.
    ccRecipients []Recipient;
    // The ID of the conversation the email belongs to.
    conversationId *string;
    // Indicates the position of the message within the conversation.
    conversationIndex []byte;
    // The collection of open extensions defined for the message. Nullable.
    extensions []Extension;
    // The flag value that indicates the status, start date, due date, or completion date for the message.
    flag *FollowupFlag;
    // The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
    from *Recipient;
    // Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
    hasAttachments *bool;
    // 
    importance *Importance;
    // 
    inferenceClassification *InferenceClassificationType;
    // 
    internetMessageHeaders []InternetMessageHeader;
    // 
    internetMessageId *string;
    // 
    isDeliveryReceiptRequested *bool;
    // 
    isDraft *bool;
    // 
    isRead *bool;
    // 
    isReadReceiptRequested *bool;
    // The collection of multi-value extended properties defined for the message. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    // 
    parentFolderId *string;
    // 
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // 
    replyTo []Recipient;
    // 
    sender *Recipient;
    // 
    sentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of single-value extended properties defined for the message. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
    // 
    subject *string;
    // 
    toRecipients []Recipient;
    // 
    uniqueBody *ItemBody;
    // 
    webLink *string;
}
// NewMessage instantiates a new message and sets the default values.
func NewMessage()(*Message) {
    m := &Message{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
// GetAttachments gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
func (m *Message) GetAttachments()([]Attachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// GetBccRecipients gets the bccRecipients property value. The Bcc: recipients for the message.
func (m *Message) GetBccRecipients()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.bccRecipients
    }
}
// GetBody gets the body property value. The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
func (m *Message) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// GetBodyPreview gets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
func (m *Message) GetBodyPreview()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bodyPreview
    }
}
// GetCcRecipients gets the ccRecipients property value. The Cc: recipients for the message.
func (m *Message) GetCcRecipients()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.ccRecipients
    }
}
// GetConversationId gets the conversationId property value. The ID of the conversation the email belongs to.
func (m *Message) GetConversationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conversationId
    }
}
// GetConversationIndex gets the conversationIndex property value. Indicates the position of the message within the conversation.
func (m *Message) GetConversationIndex()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.conversationIndex
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the message. Nullable.
func (m *Message) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFlag gets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the message.
func (m *Message) GetFlag()(*FollowupFlag) {
    if m == nil {
        return nil
    } else {
        return m.flag
    }
}
// GetFrom gets the from property value. The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
func (m *Message) GetFrom()(*Recipient) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
func (m *Message) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// GetImportance gets the importance property value. 
func (m *Message) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// GetInferenceClassification gets the inferenceClassification property value. 
func (m *Message) GetInferenceClassification()(*InferenceClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.inferenceClassification
    }
}
// GetInternetMessageHeaders gets the internetMessageHeaders property value. 
func (m *Message) GetInternetMessageHeaders()([]InternetMessageHeader) {
    if m == nil {
        return nil
    } else {
        return m.internetMessageHeaders
    }
}
// GetInternetMessageId gets the internetMessageId property value. 
func (m *Message) GetInternetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internetMessageId
    }
}
// GetIsDeliveryReceiptRequested gets the isDeliveryReceiptRequested property value. 
func (m *Message) GetIsDeliveryReceiptRequested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeliveryReceiptRequested
    }
}
// GetIsDraft gets the isDraft property value. 
func (m *Message) GetIsDraft()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDraft
    }
}
// GetIsRead gets the isRead property value. 
func (m *Message) GetIsRead()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRead
    }
}
// GetIsReadReceiptRequested gets the isReadReceiptRequested property value. 
func (m *Message) GetIsReadReceiptRequested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReadReceiptRequested
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
func (m *Message) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetParentFolderId gets the parentFolderId property value. 
func (m *Message) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// GetReceivedDateTime gets the receivedDateTime property value. 
func (m *Message) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.receivedDateTime
    }
}
// GetReplyTo gets the replyTo property value. 
func (m *Message) GetReplyTo()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.replyTo
    }
}
// GetSender gets the sender property value. 
func (m *Message) GetSender()(*Recipient) {
    if m == nil {
        return nil
    } else {
        return m.sender
    }
}
// GetSentDateTime gets the sentDateTime property value. 
func (m *Message) GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.sentDateTime
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
func (m *Message) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetSubject gets the subject property value. 
func (m *Message) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetToRecipients gets the toRecipients property value. 
func (m *Message) GetToRecipients()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.toRecipients
    }
}
// GetUniqueBody gets the uniqueBody property value. 
func (m *Message) GetUniqueBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.uniqueBody
    }
}
// GetWebLink gets the webLink property value. 
func (m *Message) GetWebLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webLink
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Message) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttachment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Attachment, len(val))
            for i, v := range val {
                res[i] = *(v.(*Attachment))
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["bccRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetBccRecipients(res)
        }
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(*ItemBody))
        }
        return nil
    }
    res["bodyPreview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyPreview(val)
        }
        return nil
    }
    res["ccRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetCcRecipients(res)
        }
        return nil
    }
    res["conversationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationId(val)
        }
        return nil
    }
    res["conversationIndex"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationIndex(val)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["flag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFollowupFlag() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlag(val.(*FollowupFlag))
        }
        return nil
    }
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val.(*Recipient))
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
    res["inferenceClassification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseInferenceClassificationType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(InferenceClassificationType)
            m.SetInferenceClassification(&cast)
        }
        return nil
    }
    res["internetMessageHeaders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInternetMessageHeader() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InternetMessageHeader, len(val))
            for i, v := range val {
                res[i] = *(v.(*InternetMessageHeader))
            }
            m.SetInternetMessageHeaders(res)
        }
        return nil
    }
    res["internetMessageId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetMessageId(val)
        }
        return nil
    }
    res["isDeliveryReceiptRequested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeliveryReceiptRequested(val)
        }
        return nil
    }
    res["isDraft"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDraft(val)
        }
        return nil
    }
    res["isRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRead(val)
        }
        return nil
    }
    res["isReadReceiptRequested"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReadReceiptRequested(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMultiValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*MultiValueLegacyExtendedProperty))
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["receivedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["replyTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetReplyTo(res)
        }
        return nil
    }
    res["sender"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSender(val.(*Recipient))
        }
        return nil
    }
    res["sentDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentDateTime(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSingleValueLegacyExtendedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*SingleValueLegacyExtendedProperty))
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["toRecipients"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetToRecipients(res)
        }
        return nil
    }
    res["uniqueBody"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueBody(val.(*ItemBody))
        }
        return nil
    }
    res["webLink"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebLink(val)
        }
        return nil
    }
    return res
}
func (m *Message) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Message) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetBccRecipients()))
        for i, v := range m.GetBccRecipients() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("bccRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("bodyPreview", m.GetBodyPreview())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCcRecipients()))
        for i, v := range m.GetCcRecipients() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("ccRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("conversationId", m.GetConversationId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("conversationIndex", m.GetConversationIndex())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("flag", m.GetFlag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("from", m.GetFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := m.GetImportance().String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInferenceClassification() != nil {
        cast := m.GetInferenceClassification().String()
        err = writer.WriteStringValue("inferenceClassification", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetInternetMessageHeaders()))
        for i, v := range m.GetInternetMessageHeaders() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("internetMessageHeaders", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("internetMessageId", m.GetInternetMessageId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDeliveryReceiptRequested", m.GetIsDeliveryReceiptRequested())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isDraft", m.GetIsDraft())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isRead", m.GetIsRead())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReadReceiptRequested", m.GetIsReadReceiptRequested())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("parentFolderId", m.GetParentFolderId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("receivedDateTime", m.GetReceivedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReplyTo()))
        for i, v := range m.GetReplyTo() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("replyTo", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sender", m.GetSender())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("sentDateTime", m.GetSentDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetToRecipients()))
        for i, v := range m.GetToRecipients() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("toRecipients", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("uniqueBody", m.GetUniqueBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webLink", m.GetWebLink())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttachments sets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
func (m *Message) SetAttachments(value []Attachment)() {
    m.attachments = value
}
// SetBccRecipients sets the bccRecipients property value. The Bcc: recipients for the message.
func (m *Message) SetBccRecipients(value []Recipient)() {
    m.bccRecipients = value
}
// SetBody sets the body property value. The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
func (m *Message) SetBody(value *ItemBody)() {
    m.body = value
}
// SetBodyPreview sets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
func (m *Message) SetBodyPreview(value *string)() {
    m.bodyPreview = value
}
// SetCcRecipients sets the ccRecipients property value. The Cc: recipients for the message.
func (m *Message) SetCcRecipients(value []Recipient)() {
    m.ccRecipients = value
}
// SetConversationId sets the conversationId property value. The ID of the conversation the email belongs to.
func (m *Message) SetConversationId(value *string)() {
    m.conversationId = value
}
// SetConversationIndex sets the conversationIndex property value. Indicates the position of the message within the conversation.
func (m *Message) SetConversationIndex(value []byte)() {
    m.conversationIndex = value
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the message. Nullable.
func (m *Message) SetExtensions(value []Extension)() {
    m.extensions = value
}
// SetFlag sets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the message.
func (m *Message) SetFlag(value *FollowupFlag)() {
    m.flag = value
}
// SetFrom sets the from property value. The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
func (m *Message) SetFrom(value *Recipient)() {
    m.from = value
}
// SetHasAttachments sets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
func (m *Message) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// SetImportance sets the importance property value. 
func (m *Message) SetImportance(value *Importance)() {
    m.importance = value
}
// SetInferenceClassification sets the inferenceClassification property value. 
func (m *Message) SetInferenceClassification(value *InferenceClassificationType)() {
    m.inferenceClassification = value
}
// SetInternetMessageHeaders sets the internetMessageHeaders property value. 
func (m *Message) SetInternetMessageHeaders(value []InternetMessageHeader)() {
    m.internetMessageHeaders = value
}
// SetInternetMessageId sets the internetMessageId property value. 
func (m *Message) SetInternetMessageId(value *string)() {
    m.internetMessageId = value
}
// SetIsDeliveryReceiptRequested sets the isDeliveryReceiptRequested property value. 
func (m *Message) SetIsDeliveryReceiptRequested(value *bool)() {
    m.isDeliveryReceiptRequested = value
}
// SetIsDraft sets the isDraft property value. 
func (m *Message) SetIsDraft(value *bool)() {
    m.isDraft = value
}
// SetIsRead sets the isRead property value. 
func (m *Message) SetIsRead(value *bool)() {
    m.isRead = value
}
// SetIsReadReceiptRequested sets the isReadReceiptRequested property value. 
func (m *Message) SetIsReadReceiptRequested(value *bool)() {
    m.isReadReceiptRequested = value
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
func (m *Message) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    m.multiValueExtendedProperties = value
}
// SetParentFolderId sets the parentFolderId property value. 
func (m *Message) SetParentFolderId(value *string)() {
    m.parentFolderId = value
}
// SetReceivedDateTime sets the receivedDateTime property value. 
func (m *Message) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.receivedDateTime = value
}
// SetReplyTo sets the replyTo property value. 
func (m *Message) SetReplyTo(value []Recipient)() {
    m.replyTo = value
}
// SetSender sets the sender property value. 
func (m *Message) SetSender(value *Recipient)() {
    m.sender = value
}
// SetSentDateTime sets the sentDateTime property value. 
func (m *Message) SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.sentDateTime = value
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
func (m *Message) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    m.singleValueExtendedProperties = value
}
// SetSubject sets the subject property value. 
func (m *Message) SetSubject(value *string)() {
    m.subject = value
}
// SetToRecipients sets the toRecipients property value. 
func (m *Message) SetToRecipients(value []Recipient)() {
    m.toRecipients = value
}
// SetUniqueBody sets the uniqueBody property value. 
func (m *Message) SetUniqueBody(value *ItemBody)() {
    m.uniqueBody = value
}
// SetWebLink sets the webLink property value. 
func (m *Message) SetWebLink(value *string)() {
    m.webLink = value
}
