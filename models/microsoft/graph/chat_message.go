package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ChatMessage provides operations to manage the collection of chat entities.
type ChatMessage struct {
    Entity
    // References to attached objects like files, tabs, meetings etc.
    attachments []ChatMessageAttachmentable;
    // 
    body ItemBodyable;
    // If the message was sent in a channel, represents identity of the channel.
    channelIdentity ChannelIdentityable;
    // If the message was sent in a chat, represents the identity of the chat.
    chatId *string;
    // Timestamp of when the chat message was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read only. Timestamp at which the chat message was deleted, or null if not deleted.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. Version number of the chat message.
    etag *string;
    // Read-only. If present, represents details of an event that happened in a chat, a channel, or a team, for example, adding new members. For event messages, the messageType property will be set to systemEventMessage.
    eventDetail EventMessageDetailable;
    // Details of the sender of the chat message. Can only be set during migration.
    from ChatMessageFromIdentitySetable;
    // Content in a message hosted by Microsoft Teams - for example, images or code snippets.
    hostedContents []ChatMessageHostedContentable;
    // The importance of the chat message. The possible values are: normal, high, urgent.
    importance *ChatMessageImportance;
    // Read only. Timestamp when edits to the chat message were made. Triggers an 'Edited' flag in the Teams UI. If no edits are made the value is null.
    lastEditedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read only. Timestamp when the chat message is created (initial setting) or modified, including when a reaction is added or removed.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Locale of the chat message set by the client. Always set to en-us.
    locale *string;
    // List of entities mentioned in the chat message. Supported entities are: user, bot, team, and channel.
    mentions []ChatMessageMentionable;
    // The type of chat message. The possible values are: message, chatEvent, typing, unknownFutureValue, systemEventMessage. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: systemEventMessage.
    messageType *ChatMessageType;
    // Defines the properties of a policy violation set by a data loss prevention (DLP) application.
    policyViolation ChatMessagePolicyViolationable;
    // Reactions for this chat message (for example, Like).
    reactions []ChatMessageReactionable;
    // Replies for a specified message.
    replies []ChatMessageable;
    // Read-only. ID of the parent chat message or root chat message of the thread. (Only applies to chat messages in channels, not chats.)
    replyToId *string;
    // The subject of the chat message, in plaintext.
    subject *string;
    // Summary text of the chat message that could be used for push notifications and summary views or fall back views. Only applies to channel chat messages, not chat messages in a chat.
    summary *string;
    // Read-only. Link to the message in Microsoft Teams.
    webUrl *string;
}
// NewChatMessage instantiates a new chatMessage and sets the default values.
func NewChatMessage()(*ChatMessage) {
    m := &ChatMessage{
        Entity: *NewEntity(),
    }
    return m
}
// CreateChatMessageFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateChatMessageFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewChatMessage(), nil
}
// GetAttachments gets the attachments property value. References to attached objects like files, tabs, meetings etc.
func (m *ChatMessage) GetAttachments()([]ChatMessageAttachmentable) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// GetBody gets the body property value. 
func (m *ChatMessage) GetBody()(ItemBodyable) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// GetChannelIdentity gets the channelIdentity property value. If the message was sent in a channel, represents identity of the channel.
func (m *ChatMessage) GetChannelIdentity()(ChannelIdentityable) {
    if m == nil {
        return nil
    } else {
        return m.channelIdentity
    }
}
// GetChatId gets the chatId property value. If the message was sent in a chat, represents the identity of the chat.
func (m *ChatMessage) GetChatId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.chatId
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Timestamp of when the chat message was created.
func (m *ChatMessage) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDeletedDateTime gets the deletedDateTime property value. Read only. Timestamp at which the chat message was deleted, or null if not deleted.
func (m *ChatMessage) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deletedDateTime
    }
}
// GetEtag gets the etag property value. Read-only. Version number of the chat message.
func (m *ChatMessage) GetEtag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.etag
    }
}
// GetEventDetail gets the eventDetail property value. Read-only. If present, represents details of an event that happened in a chat, a channel, or a team, for example, adding new members. For event messages, the messageType property will be set to systemEventMessage.
func (m *ChatMessage) GetEventDetail()(EventMessageDetailable) {
    if m == nil {
        return nil
    } else {
        return m.eventDetail
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ChatMessage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatMessageAttachmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageAttachmentable, len(val))
            for i, v := range val {
                res[i] = v.(ChatMessageAttachmentable)
            }
            m.SetAttachments(res)
        }
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(ItemBodyable))
        }
        return nil
    }
    res["channelIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateChannelIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelIdentity(val.(ChannelIdentityable))
        }
        return nil
    }
    res["chatId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatId(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["deletedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeletedDateTime(val)
        }
        return nil
    }
    res["etag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEtag(val)
        }
        return nil
    }
    res["eventDetail"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEventMessageDetailFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDetail(val.(EventMessageDetailable))
        }
        return nil
    }
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessageFromIdentitySetFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val.(ChatMessageFromIdentitySetable))
        }
        return nil
    }
    res["hostedContents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatMessageHostedContentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageHostedContentable, len(val))
            for i, v := range val {
                res[i] = v.(ChatMessageHostedContentable)
            }
            m.SetHostedContents(res)
        }
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessageImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*ChatMessageImportance))
        }
        return nil
    }
    res["lastEditedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastEditedDateTime(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["locale"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLocale(val)
        }
        return nil
    }
    res["mentions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatMessageMentionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageMentionable, len(val))
            for i, v := range val {
                res[i] = v.(ChatMessageMentionable)
            }
            m.SetMentions(res)
        }
        return nil
    }
    res["messageType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessageType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMessageType(val.(*ChatMessageType))
        }
        return nil
    }
    res["policyViolation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateChatMessagePolicyViolationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyViolation(val.(ChatMessagePolicyViolationable))
        }
        return nil
    }
    res["reactions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatMessageReactionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageReactionable, len(val))
            for i, v := range val {
                res[i] = v.(ChatMessageReactionable)
            }
            m.SetReactions(res)
        }
        return nil
    }
    res["replies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChatMessageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageable, len(val))
            for i, v := range val {
                res[i] = v.(ChatMessageable)
            }
            m.SetReplies(res)
        }
        return nil
    }
    res["replyToId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReplyToId(val)
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
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSummary(val)
        }
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWebUrl(val)
        }
        return nil
    }
    return res
}
// GetFrom gets the from property value. Details of the sender of the chat message. Can only be set during migration.
func (m *ChatMessage) GetFrom()(ChatMessageFromIdentitySetable) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
// GetHostedContents gets the hostedContents property value. Content in a message hosted by Microsoft Teams - for example, images or code snippets.
func (m *ChatMessage) GetHostedContents()([]ChatMessageHostedContentable) {
    if m == nil {
        return nil
    } else {
        return m.hostedContents
    }
}
// GetImportance gets the importance property value. The importance of the chat message. The possible values are: normal, high, urgent.
func (m *ChatMessage) GetImportance()(*ChatMessageImportance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// GetLastEditedDateTime gets the lastEditedDateTime property value. Read only. Timestamp when edits to the chat message were made. Triggers an 'Edited' flag in the Teams UI. If no edits are made the value is null.
func (m *ChatMessage) GetLastEditedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastEditedDateTime
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Read only. Timestamp when the chat message is created (initial setting) or modified, including when a reaction is added or removed.
func (m *ChatMessage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLocale gets the locale property value. Locale of the chat message set by the client. Always set to en-us.
func (m *ChatMessage) GetLocale()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locale
    }
}
// GetMentions gets the mentions property value. List of entities mentioned in the chat message. Supported entities are: user, bot, team, and channel.
func (m *ChatMessage) GetMentions()([]ChatMessageMentionable) {
    if m == nil {
        return nil
    } else {
        return m.mentions
    }
}
// GetMessageType gets the messageType property value. The type of chat message. The possible values are: message, chatEvent, typing, unknownFutureValue, systemEventMessage. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: systemEventMessage.
func (m *ChatMessage) GetMessageType()(*ChatMessageType) {
    if m == nil {
        return nil
    } else {
        return m.messageType
    }
}
// GetPolicyViolation gets the policyViolation property value. Defines the properties of a policy violation set by a data loss prevention (DLP) application.
func (m *ChatMessage) GetPolicyViolation()(ChatMessagePolicyViolationable) {
    if m == nil {
        return nil
    } else {
        return m.policyViolation
    }
}
// GetReactions gets the reactions property value. Reactions for this chat message (for example, Like).
func (m *ChatMessage) GetReactions()([]ChatMessageReactionable) {
    if m == nil {
        return nil
    } else {
        return m.reactions
    }
}
// GetReplies gets the replies property value. Replies for a specified message.
func (m *ChatMessage) GetReplies()([]ChatMessageable) {
    if m == nil {
        return nil
    } else {
        return m.replies
    }
}
// GetReplyToId gets the replyToId property value. Read-only. ID of the parent chat message or root chat message of the thread. (Only applies to chat messages in channels, not chats.)
func (m *ChatMessage) GetReplyToId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.replyToId
    }
}
// GetSubject gets the subject property value. The subject of the chat message, in plaintext.
func (m *ChatMessage) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// GetSummary gets the summary property value. Summary text of the chat message that could be used for push notifications and summary views or fall back views. Only applies to channel chat messages, not chat messages in a chat.
func (m *ChatMessage) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
// GetWebUrl gets the webUrl property value. Read-only. Link to the message in Microsoft Teams.
func (m *ChatMessage) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *ChatMessage) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ChatMessage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttachments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttachments()))
        for i, v := range m.GetAttachments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attachments", cast)
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
        err = writer.WriteObjectValue("channelIdentity", m.GetChannelIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("chatId", m.GetChatId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("deletedDateTime", m.GetDeletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("etag", m.GetEtag())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("eventDetail", m.GetEventDetail())
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
    if m.GetHostedContents() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHostedContents()))
        for i, v := range m.GetHostedContents() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("hostedContents", cast)
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
    {
        err = writer.WriteTimeValue("lastEditedDateTime", m.GetLastEditedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("locale", m.GetLocale())
        if err != nil {
            return err
        }
    }
    if m.GetMentions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMentions()))
        for i, v := range m.GetMentions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("mentions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMessageType() != nil {
        cast := (*m.GetMessageType()).String()
        err = writer.WriteStringValue("messageType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("policyViolation", m.GetPolicyViolation())
        if err != nil {
            return err
        }
    }
    if m.GetReactions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReactions()))
        for i, v := range m.GetReactions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("reactions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetReplies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReplies()))
        for i, v := range m.GetReplies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("replies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("replyToId", m.GetReplyToId())
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
        err = writer.WriteStringValue("summary", m.GetSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("webUrl", m.GetWebUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttachments sets the attachments property value. References to attached objects like files, tabs, meetings etc.
func (m *ChatMessage) SetAttachments(value []ChatMessageAttachmentable)() {
    if m != nil {
        m.attachments = value
    }
}
// SetBody sets the body property value. 
func (m *ChatMessage) SetBody(value ItemBodyable)() {
    if m != nil {
        m.body = value
    }
}
// SetChannelIdentity sets the channelIdentity property value. If the message was sent in a channel, represents identity of the channel.
func (m *ChatMessage) SetChannelIdentity(value ChannelIdentityable)() {
    if m != nil {
        m.channelIdentity = value
    }
}
// SetChatId sets the chatId property value. If the message was sent in a chat, represents the identity of the chat.
func (m *ChatMessage) SetChatId(value *string)() {
    if m != nil {
        m.chatId = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Timestamp of when the chat message was created.
func (m *ChatMessage) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDeletedDateTime sets the deletedDateTime property value. Read only. Timestamp at which the chat message was deleted, or null if not deleted.
func (m *ChatMessage) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.deletedDateTime = value
    }
}
// SetEtag sets the etag property value. Read-only. Version number of the chat message.
func (m *ChatMessage) SetEtag(value *string)() {
    if m != nil {
        m.etag = value
    }
}
// SetEventDetail sets the eventDetail property value. Read-only. If present, represents details of an event that happened in a chat, a channel, or a team, for example, adding new members. For event messages, the messageType property will be set to systemEventMessage.
func (m *ChatMessage) SetEventDetail(value EventMessageDetailable)() {
    if m != nil {
        m.eventDetail = value
    }
}
// SetFrom sets the from property value. Details of the sender of the chat message. Can only be set during migration.
func (m *ChatMessage) SetFrom(value ChatMessageFromIdentitySetable)() {
    if m != nil {
        m.from = value
    }
}
// SetHostedContents sets the hostedContents property value. Content in a message hosted by Microsoft Teams - for example, images or code snippets.
func (m *ChatMessage) SetHostedContents(value []ChatMessageHostedContentable)() {
    if m != nil {
        m.hostedContents = value
    }
}
// SetImportance sets the importance property value. The importance of the chat message. The possible values are: normal, high, urgent.
func (m *ChatMessage) SetImportance(value *ChatMessageImportance)() {
    if m != nil {
        m.importance = value
    }
}
// SetLastEditedDateTime sets the lastEditedDateTime property value. Read only. Timestamp when edits to the chat message were made. Triggers an 'Edited' flag in the Teams UI. If no edits are made the value is null.
func (m *ChatMessage) SetLastEditedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastEditedDateTime = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Read only. Timestamp when the chat message is created (initial setting) or modified, including when a reaction is added or removed.
func (m *ChatMessage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLocale sets the locale property value. Locale of the chat message set by the client. Always set to en-us.
func (m *ChatMessage) SetLocale(value *string)() {
    if m != nil {
        m.locale = value
    }
}
// SetMentions sets the mentions property value. List of entities mentioned in the chat message. Supported entities are: user, bot, team, and channel.
func (m *ChatMessage) SetMentions(value []ChatMessageMentionable)() {
    if m != nil {
        m.mentions = value
    }
}
// SetMessageType sets the messageType property value. The type of chat message. The possible values are: message, chatEvent, typing, unknownFutureValue, systemEventMessage. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: systemEventMessage.
func (m *ChatMessage) SetMessageType(value *ChatMessageType)() {
    if m != nil {
        m.messageType = value
    }
}
// SetPolicyViolation sets the policyViolation property value. Defines the properties of a policy violation set by a data loss prevention (DLP) application.
func (m *ChatMessage) SetPolicyViolation(value ChatMessagePolicyViolationable)() {
    if m != nil {
        m.policyViolation = value
    }
}
// SetReactions sets the reactions property value. Reactions for this chat message (for example, Like).
func (m *ChatMessage) SetReactions(value []ChatMessageReactionable)() {
    if m != nil {
        m.reactions = value
    }
}
// SetReplies sets the replies property value. Replies for a specified message.
func (m *ChatMessage) SetReplies(value []ChatMessageable)() {
    if m != nil {
        m.replies = value
    }
}
// SetReplyToId sets the replyToId property value. Read-only. ID of the parent chat message or root chat message of the thread. (Only applies to chat messages in channels, not chats.)
func (m *ChatMessage) SetReplyToId(value *string)() {
    if m != nil {
        m.replyToId = value
    }
}
// SetSubject sets the subject property value. The subject of the chat message, in plaintext.
func (m *ChatMessage) SetSubject(value *string)() {
    if m != nil {
        m.subject = value
    }
}
// SetSummary sets the summary property value. Summary text of the chat message that could be used for push notifications and summary views or fall back views. Only applies to channel chat messages, not chat messages in a chat.
func (m *ChatMessage) SetSummary(value *string)() {
    if m != nil {
        m.summary = value
    }
}
// SetWebUrl sets the webUrl property value. Read-only. Link to the message in Microsoft Teams.
func (m *ChatMessage) SetWebUrl(value *string)() {
    if m != nil {
        m.webUrl = value
    }
}
