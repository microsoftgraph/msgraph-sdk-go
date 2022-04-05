package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Message 
type Message struct {
    OutlookItem
    // The fileAttachment and itemAttachment attachments for the message.
    attachments []Attachmentable;
    // The Bcc: recipients for the message.
    bccRecipients []Recipientable;
    // The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
    body ItemBodyable;
    // The first 255 characters of the message body. It is in text format.
    bodyPreview *string;
    // The Cc: recipients for the message.
    ccRecipients []Recipientable;
    // The ID of the conversation the email belongs to.
    conversationId *string;
    // Indicates the position of the message within the conversation.
    conversationIndex []byte;
    // The collection of open extensions defined for the message. Nullable.
    extensions []Extensionable;
    // The flag value that indicates the status, start date, due date, or completion date for the message.
    flag FollowupFlagable;
    // The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
    from Recipientable;
    // Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
    hasAttachments *bool;
    // The importance property
    importance *Importance;
    // The inferenceClassification property
    inferenceClassification *InferenceClassificationType;
    // The internetMessageHeaders property
    internetMessageHeaders []InternetMessageHeaderable;
    // The internetMessageId property
    internetMessageId *string;
    // The isDeliveryReceiptRequested property
    isDeliveryReceiptRequested *bool;
    // The isDraft property
    isDraft *bool;
    // The isRead property
    isRead *bool;
    // The isReadReceiptRequested property
    isReadReceiptRequested *bool;
    // The collection of multi-value extended properties defined for the message. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable;
    // The parentFolderId property
    parentFolderId *string;
    // The receivedDateTime property
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The replyTo property
    replyTo []Recipientable;
    // The sender property
    sender Recipientable;
    // The sentDateTime property
    sentDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The collection of single-value extended properties defined for the message. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable;
    // The subject property
    subject *string;
    // The toRecipients property
    toRecipients []Recipientable;
    // The uniqueBody property
    uniqueBody ItemBodyable;
    // The webLink property
    webLink *string;
}
// NewMessage instantiates a new message and sets the default values.
func NewMessage()(*Message) {
    m := &Message{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
// CreateMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMessageFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMessage(), nil
}
// GetAttachments gets the attachments property value. The fileAttachment and itemAttachment attachments for the message.
func (m *Message) GetAttachments()([]Attachmentable) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// GetBccRecipients gets the bccRecipients property value. The Bcc: recipients for the message.
func (m *Message) GetBccRecipients()([]Recipientable) {
    if m == nil {
        return nil
    } else {
        return m.bccRecipients
    }
}
// GetBody gets the body property value. The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
func (m *Message) GetBody()(ItemBodyable) {
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
func (m *Message) GetCcRecipients()([]Recipientable) {
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
func (m *Message) GetExtensions()([]Extensionable) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Message) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["attachments"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttachmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Attachmentable, len(val))
            for i, v := range val {
                res[i] = v.(Attachmentable)
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["bccRecipients"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetBccRecipients(res)
        }
        return nil
    }
    res["body"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(ItemBodyable))
        }
        return nil
    }
    res["bodyPreview"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyPreview(val)
        }
        return nil
    }
    res["ccRecipients"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetCcRecipients(res)
        }
        return nil
    }
    res["conversationId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationId(val)
        }
        return nil
    }
    res["conversationIndex"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationIndex(val)
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["flag"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateFollowupFlagFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlag(val.(FollowupFlagable))
        }
        return nil
    }
    res["from"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val.(Recipientable))
        }
        return nil
    }
    res["hasAttachments"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasAttachments(val)
        }
        return nil
    }
    res["importance"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*Importance))
        }
        return nil
    }
    res["inferenceClassification"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseInferenceClassificationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInferenceClassification(val.(*InferenceClassificationType))
        }
        return nil
    }
    res["internetMessageHeaders"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateInternetMessageHeaderFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]InternetMessageHeaderable, len(val))
            for i, v := range val {
                res[i] = v.(InternetMessageHeaderable)
            }
            m.SetInternetMessageHeaders(res)
        }
        return nil
    }
    res["internetMessageId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInternetMessageId(val)
        }
        return nil
    }
    res["isDeliveryReceiptRequested"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDeliveryReceiptRequested(val)
        }
        return nil
    }
    res["isDraft"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDraft(val)
        }
        return nil
    }
    res["isRead"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRead(val)
        }
        return nil
    }
    res["isReadReceiptRequested"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReadReceiptRequested(val)
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateMultiValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MultiValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(MultiValueLegacyExtendedPropertyable)
            }
            m.SetMultiValueExtendedProperties(res)
        }
        return nil
    }
    res["parentFolderId"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentFolderId(val)
        }
        return nil
    }
    res["receivedDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedDateTime(val)
        }
        return nil
    }
    res["replyTo"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetReplyTo(res)
        }
        return nil
    }
    res["sender"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSender(val.(Recipientable))
        }
        return nil
    }
    res["sentDateTime"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentDateTime(val)
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSingleValueLegacyExtendedPropertyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SingleValueLegacyExtendedPropertyable, len(val))
            for i, v := range val {
                res[i] = v.(SingleValueLegacyExtendedPropertyable)
            }
            m.SetSingleValueExtendedProperties(res)
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["toRecipients"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
            }
            m.SetToRecipients(res)
        }
        return nil
    }
    res["uniqueBody"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUniqueBody(val.(ItemBodyable))
        }
        return nil
    }
    res["webLink"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetFlag gets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the message.
func (m *Message) GetFlag()(FollowupFlagable) {
    if m == nil {
        return nil
    } else {
        return m.flag
    }
}
// GetFrom gets the from property value. The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
func (m *Message) GetFrom()(Recipientable) {
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
// GetImportance gets the importance property value. The importance property
func (m *Message) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// GetInferenceClassification gets the inferenceClassification property value. The inferenceClassification property
func (m *Message) GetInferenceClassification()(*InferenceClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.inferenceClassification
    }
}
// GetInternetMessageHeaders gets the internetMessageHeaders property value. The internetMessageHeaders property
func (m *Message) GetInternetMessageHeaders()([]InternetMessageHeaderable) {
    if m == nil {
        return nil
    } else {
        return m.internetMessageHeaders
    }
}
// GetInternetMessageId gets the internetMessageId property value. The internetMessageId property
func (m *Message) GetInternetMessageId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.internetMessageId
    }
}
// GetIsDeliveryReceiptRequested gets the isDeliveryReceiptRequested property value. The isDeliveryReceiptRequested property
func (m *Message) GetIsDeliveryReceiptRequested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDeliveryReceiptRequested
    }
}
// GetIsDraft gets the isDraft property value. The isDraft property
func (m *Message) GetIsDraft()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDraft
    }
}
// GetIsRead gets the isRead property value. The isRead property
func (m *Message) GetIsRead()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRead
    }
}
// GetIsReadReceiptRequested gets the isReadReceiptRequested property value. The isReadReceiptRequested property
func (m *Message) GetIsReadReceiptRequested()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReadReceiptRequested
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
func (m *Message) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetParentFolderId gets the parentFolderId property value. The parentFolderId property
func (m *Message) GetParentFolderId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentFolderId
    }
}
// GetReceivedDateTime gets the receivedDateTime property value. The receivedDateTime property
func (m *Message) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.receivedDateTime
    }
}
// GetReplyTo gets the replyTo property value. The replyTo property
func (m *Message) GetReplyTo()([]Recipientable) {
    if m == nil {
        return nil
    } else {
        return m.replyTo
    }
}
// GetSender gets the sender property value. The sender property
func (m *Message) GetSender()(Recipientable) {
    if m == nil {
        return nil
    } else {
        return m.sender
    }
}
// GetSentDateTime gets the sentDateTime property value. The sentDateTime property
func (m *Message) GetSentDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.sentDateTime
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
func (m *Message) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetSubject gets the subject property value. The subject property
func (m *Message) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetToRecipients gets the toRecipients property value. The toRecipients property
func (m *Message) GetToRecipients()([]Recipientable) {
    if m == nil {
        return nil
    } else {
        return m.toRecipients
    }
}
// GetUniqueBody gets the uniqueBody property value. The uniqueBody property
func (m *Message) GetUniqueBody()(ItemBodyable) {
    if m == nil {
        return nil
    } else {
        return m.uniqueBody
    }
}
// GetWebLink gets the webLink property value. The webLink property
func (m *Message) GetWebLink()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webLink
    }
}
// Serialize serializes information the current object
func (m *Message) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.OutlookItem.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttachments() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetBccRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetBccRecipients()))
        for i, v := range m.GetBccRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetCcRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCcRecipients()))
        for i, v := range m.GetCcRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := (*m.GetImportance()).String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInferenceClassification() != nil {
        cast := (*m.GetInferenceClassification()).String()
        err = writer.WriteStringValue("inferenceClassification", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetInternetMessageHeaders() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetInternetMessageHeaders()))
        for i, v := range m.GetInternetMessageHeaders() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetReplyTo() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetReplyTo()))
        for i, v := range m.GetReplyTo() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetToRecipients() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetToRecipients()))
        for i, v := range m.GetToRecipients() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
func (m *Message) SetAttachments(value []Attachmentable)() {
    if m != nil {
        m.attachments = value
    }
}
// SetBccRecipients sets the bccRecipients property value. The Bcc: recipients for the message.
func (m *Message) SetBccRecipients(value []Recipientable)() {
    if m != nil {
        m.bccRecipients = value
    }
}
// SetBody sets the body property value. The body of the message. It can be in HTML or text format. Find out about safe HTML in a message body.
func (m *Message) SetBody(value ItemBodyable)() {
    if m != nil {
        m.body = value
    }
}
// SetBodyPreview sets the bodyPreview property value. The first 255 characters of the message body. It is in text format.
func (m *Message) SetBodyPreview(value *string)() {
    if m != nil {
        m.bodyPreview = value
    }
}
// SetCcRecipients sets the ccRecipients property value. The Cc: recipients for the message.
func (m *Message) SetCcRecipients(value []Recipientable)() {
    if m != nil {
        m.ccRecipients = value
    }
}
// SetConversationId sets the conversationId property value. The ID of the conversation the email belongs to.
func (m *Message) SetConversationId(value *string)() {
    if m != nil {
        m.conversationId = value
    }
}
// SetConversationIndex sets the conversationIndex property value. Indicates the position of the message within the conversation.
func (m *Message) SetConversationIndex(value []byte)() {
    if m != nil {
        m.conversationIndex = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the message. Nullable.
func (m *Message) SetExtensions(value []Extensionable)() {
    if m != nil {
        m.extensions = value
    }
}
// SetFlag sets the flag property value. The flag value that indicates the status, start date, due date, or completion date for the message.
func (m *Message) SetFlag(value FollowupFlagable)() {
    if m != nil {
        m.flag = value
    }
}
// SetFrom sets the from property value. The owner of the mailbox from which the message is sent. In most cases, this value is the same as the sender property, except for sharing or delegation scenarios. The value must correspond to the actual mailbox used. Find out more about setting the from and sender properties of a message.
func (m *Message) SetFrom(value Recipientable)() {
    if m != nil {
        m.from = value
    }
}
// SetHasAttachments sets the hasAttachments property value. Indicates whether the message has attachments. This property doesn't include inline attachments, so if a message contains only inline attachments, this property is false. To verify the existence of inline attachments, parse the body property to look for a src attribute, such as <IMG src='cid:image001.jpg@01D26CD8.6C05F070'>.
func (m *Message) SetHasAttachments(value *bool)() {
    if m != nil {
        m.hasAttachments = value
    }
}
// SetImportance sets the importance property value. The importance property
func (m *Message) SetImportance(value *Importance)() {
    if m != nil {
        m.importance = value
    }
}
// SetInferenceClassification sets the inferenceClassification property value. The inferenceClassification property
func (m *Message) SetInferenceClassification(value *InferenceClassificationType)() {
    if m != nil {
        m.inferenceClassification = value
    }
}
// SetInternetMessageHeaders sets the internetMessageHeaders property value. The internetMessageHeaders property
func (m *Message) SetInternetMessageHeaders(value []InternetMessageHeaderable)() {
    if m != nil {
        m.internetMessageHeaders = value
    }
}
// SetInternetMessageId sets the internetMessageId property value. The internetMessageId property
func (m *Message) SetInternetMessageId(value *string)() {
    if m != nil {
        m.internetMessageId = value
    }
}
// SetIsDeliveryReceiptRequested sets the isDeliveryReceiptRequested property value. The isDeliveryReceiptRequested property
func (m *Message) SetIsDeliveryReceiptRequested(value *bool)() {
    if m != nil {
        m.isDeliveryReceiptRequested = value
    }
}
// SetIsDraft sets the isDraft property value. The isDraft property
func (m *Message) SetIsDraft(value *bool)() {
    if m != nil {
        m.isDraft = value
    }
}
// SetIsRead sets the isRead property value. The isRead property
func (m *Message) SetIsRead(value *bool)() {
    if m != nil {
        m.isRead = value
    }
}
// SetIsReadReceiptRequested sets the isReadReceiptRequested property value. The isReadReceiptRequested property
func (m *Message) SetIsReadReceiptRequested(value *bool)() {
    if m != nil {
        m.isReadReceiptRequested = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the message. Nullable.
func (m *Message) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetParentFolderId sets the parentFolderId property value. The parentFolderId property
func (m *Message) SetParentFolderId(value *string)() {
    if m != nil {
        m.parentFolderId = value
    }
}
// SetReceivedDateTime sets the receivedDateTime property value. The receivedDateTime property
func (m *Message) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.receivedDateTime = value
    }
}
// SetReplyTo sets the replyTo property value. The replyTo property
func (m *Message) SetReplyTo(value []Recipientable)() {
    if m != nil {
        m.replyTo = value
    }
}
// SetSender sets the sender property value. The sender property
func (m *Message) SetSender(value Recipientable)() {
    if m != nil {
        m.sender = value
    }
}
// SetSentDateTime sets the sentDateTime property value. The sentDateTime property
func (m *Message) SetSentDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.sentDateTime = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the message. Nullable.
func (m *Message) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
// SetSubject sets the subject property value. The subject property
func (m *Message) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
// SetToRecipients sets the toRecipients property value. The toRecipients property
func (m *Message) SetToRecipients(value []Recipientable)() {
    if m != nil {
        m.toRecipients = value
    }
}
// SetUniqueBody sets the uniqueBody property value. The uniqueBody property
func (m *Message) SetUniqueBody(value ItemBodyable)() {
    if m != nil {
        m.uniqueBody = value
    }
}
// SetWebLink sets the webLink property value. The webLink property
func (m *Message) SetWebLink(value *string)() {
    if m != nil {
        m.webLink = value
    }
}
