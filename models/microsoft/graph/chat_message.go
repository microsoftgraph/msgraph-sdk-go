package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ChatMessage struct {
    Entity
    // References to attached objects like files, tabs, meetings etc.
    attachments []ChatMessageAttachment;
    // 
    body *ItemBody;
    // If the message was sent in a channel, represents identity of the channel.
    channelIdentity *ChannelIdentity;
    // If the message was sent in a chat, represents the identity of the chat.
    chatId *string;
    // Timestamp of when the chat message was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read only. Timestamp at which the chat message was deleted, or null if not deleted.
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. Version number of the chat message.
    etag *string;
    // Read-only.  If present, represents details of an event that happened in a chat, a channel, or a team, for example, members were added, and so on. For event messages, the messageType property will be set to systemEventMessage.
    eventDetail *EventMessageDetail;
    // Details of the sender of the chat message. Can only be set during migration.
    from *ChatMessageFromIdentitySet;
    // Content in a message hosted by Microsoft Teams - for example, images or code snippets.
    hostedContents []ChatMessageHostedContent;
    // The importance of the chat message. The possible values are: normal, high, urgent.
    importance *ChatMessageImportance;
    // Read only. Timestamp when edits to the chat message were made. Triggers an 'Edited' flag in the Teams UI. If no edits are made the value is null.
    lastEditedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read only. Timestamp when the chat message is created (initial setting) or modified, including when a reaction is added or removed.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Locale of the chat message set by the client. Always set to en-us.
    locale *string;
    // List of entities mentioned in the chat message. Supported entities are: user, bot, team, and channel.
    mentions []ChatMessageMention;
    // The type of chat message. The possible values are: message, chatEvent, typing, unknownFutureValue, systemEventMessage. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: systemEventMessage.
    messageType *ChatMessageType;
    // Defines the properties of a policy violation set by a data loss prevention (DLP) application.
    policyViolation *ChatMessagePolicyViolation;
    // Reactions for this chat message (for example, Like).
    reactions []ChatMessageReaction;
    // Replies for a specified message.
    replies []ChatMessage;
    // Read-only. ID of the parent chat message or root chat message of the thread. (Only applies to chat messages in channels, not chats.)
    replyToId *string;
    // The subject of the chat message, in plaintext.
    subject *string;
    // Summary text of the chat message that could be used for push notifications and summary views or fall back views. Only applies to channel chat messages, not chat messages in a chat.
    summary *string;
    // Read-only. Link to the message in Microsoft Teams.
    webUrl *string;
}
// Instantiates a new chatMessage and sets the default values.
func NewChatMessage()(*ChatMessage) {
    m := &ChatMessage{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the attachments property value. References to attached objects like files, tabs, meetings etc.
func (m *ChatMessage) GetAttachments()([]ChatMessageAttachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// Gets the body property value. 
func (m *ChatMessage) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// Gets the channelIdentity property value. If the message was sent in a channel, represents identity of the channel.
func (m *ChatMessage) GetChannelIdentity()(*ChannelIdentity) {
    if m == nil {
        return nil
    } else {
        return m.channelIdentity
    }
}
// Gets the chatId property value. If the message was sent in a chat, represents the identity of the chat.
func (m *ChatMessage) GetChatId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.chatId
    }
}
// Gets the createdDateTime property value. Timestamp of when the chat message was created.
func (m *ChatMessage) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the deletedDateTime property value. Read only. Timestamp at which the chat message was deleted, or null if not deleted.
func (m *ChatMessage) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deletedDateTime
    }
}
// Gets the etag property value. Read-only. Version number of the chat message.
func (m *ChatMessage) GetEtag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.etag
    }
}
// Gets the eventDetail property value. Read-only.  If present, represents details of an event that happened in a chat, a channel, or a team, for example, members were added, and so on. For event messages, the messageType property will be set to systemEventMessage.
func (m *ChatMessage) GetEventDetail()(*EventMessageDetail) {
    if m == nil {
        return nil
    } else {
        return m.eventDetail
    }
}
// Gets the from property value. Details of the sender of the chat message. Can only be set during migration.
func (m *ChatMessage) GetFrom()(*ChatMessageFromIdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
// Gets the hostedContents property value. Content in a message hosted by Microsoft Teams - for example, images or code snippets.
func (m *ChatMessage) GetHostedContents()([]ChatMessageHostedContent) {
    if m == nil {
        return nil
    } else {
        return m.hostedContents
    }
}
// Gets the importance property value. The importance of the chat message. The possible values are: normal, high, urgent.
func (m *ChatMessage) GetImportance()(*ChatMessageImportance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// Gets the lastEditedDateTime property value. Read only. Timestamp when edits to the chat message were made. Triggers an 'Edited' flag in the Teams UI. If no edits are made the value is null.
func (m *ChatMessage) GetLastEditedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastEditedDateTime
    }
}
// Gets the lastModifiedDateTime property value. Read only. Timestamp when the chat message is created (initial setting) or modified, including when a reaction is added or removed.
func (m *ChatMessage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the locale property value. Locale of the chat message set by the client. Always set to en-us.
func (m *ChatMessage) GetLocale()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locale
    }
}
// Gets the mentions property value. List of entities mentioned in the chat message. Supported entities are: user, bot, team, and channel.
func (m *ChatMessage) GetMentions()([]ChatMessageMention) {
    if m == nil {
        return nil
    } else {
        return m.mentions
    }
}
// Gets the messageType property value. The type of chat message. The possible values are: message, chatEvent, typing, unknownFutureValue, systemEventMessage. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: systemEventMessage.
func (m *ChatMessage) GetMessageType()(*ChatMessageType) {
    if m == nil {
        return nil
    } else {
        return m.messageType
    }
}
// Gets the policyViolation property value. Defines the properties of a policy violation set by a data loss prevention (DLP) application.
func (m *ChatMessage) GetPolicyViolation()(*ChatMessagePolicyViolation) {
    if m == nil {
        return nil
    } else {
        return m.policyViolation
    }
}
// Gets the reactions property value. Reactions for this chat message (for example, Like).
func (m *ChatMessage) GetReactions()([]ChatMessageReaction) {
    if m == nil {
        return nil
    } else {
        return m.reactions
    }
}
// Gets the replies property value. Replies for a specified message.
func (m *ChatMessage) GetReplies()([]ChatMessage) {
    if m == nil {
        return nil
    } else {
        return m.replies
    }
}
// Gets the replyToId property value. Read-only. ID of the parent chat message or root chat message of the thread. (Only applies to chat messages in channels, not chats.)
func (m *ChatMessage) GetReplyToId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.replyToId
    }
}
// Gets the subject property value. The subject of the chat message, in plaintext.
func (m *ChatMessage) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// Gets the summary property value. Summary text of the chat message that could be used for push notifications and summary views or fall back views. Only applies to channel chat messages, not chat messages in a chat.
func (m *ChatMessage) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
// Gets the webUrl property value. Read-only. Link to the message in Microsoft Teams.
func (m *ChatMessage) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *ChatMessage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageAttachment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageAttachment, len(val))
            for i, v := range val {
                res[i] = *(v.(*ChatMessageAttachment))
            }
            m.SetAttachments(res)
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
    res["channelIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChannelIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChannelIdentity(val.(*ChannelIdentity))
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
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEventMessageDetail() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDetail(val.(*EventMessageDetail))
        }
        return nil
    }
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageFromIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val.(*ChatMessageFromIdentitySet))
        }
        return nil
    }
    res["hostedContents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageHostedContent() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageHostedContent, len(val))
            for i, v := range val {
                res[i] = *(v.(*ChatMessageHostedContent))
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
            cast := val.(ChatMessageImportance)
            m.SetImportance(&cast)
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
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageMention() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageMention, len(val))
            for i, v := range val {
                res[i] = *(v.(*ChatMessageMention))
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
            cast := val.(ChatMessageType)
            m.SetMessageType(&cast)
        }
        return nil
    }
    res["policyViolation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessagePolicyViolation() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPolicyViolation(val.(*ChatMessagePolicyViolation))
        }
        return nil
    }
    res["reactions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageReaction() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessageReaction, len(val))
            for i, v := range val {
                res[i] = *(v.(*ChatMessageReaction))
            }
            m.SetReactions(res)
        }
        return nil
    }
    res["replies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessage, len(val))
            for i, v := range val {
                res[i] = *(v.(*ChatMessage))
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
func (m *ChatMessage) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ChatMessage) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetHostedContents()))
        for i, v := range m.GetHostedContents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("hostedContents", cast)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMentions()))
        for i, v := range m.GetMentions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mentions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMessageType() != nil {
        cast := m.GetMessageType().String()
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReactions()))
        for i, v := range m.GetReactions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("reactions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetReplies()))
        for i, v := range m.GetReplies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// Sets the attachments property value. References to attached objects like files, tabs, meetings etc.
// Parameters:
//  - value : Value to set for the attachments property.
func (m *ChatMessage) SetAttachments(value []ChatMessageAttachment)() {
    m.attachments = value
}
// Sets the body property value. 
// Parameters:
//  - value : Value to set for the body property.
func (m *ChatMessage) SetBody(value *ItemBody)() {
    m.body = value
}
// Sets the channelIdentity property value. If the message was sent in a channel, represents identity of the channel.
// Parameters:
//  - value : Value to set for the channelIdentity property.
func (m *ChatMessage) SetChannelIdentity(value *ChannelIdentity)() {
    m.channelIdentity = value
}
// Sets the chatId property value. If the message was sent in a chat, represents the identity of the chat.
// Parameters:
//  - value : Value to set for the chatId property.
func (m *ChatMessage) SetChatId(value *string)() {
    m.chatId = value
}
// Sets the createdDateTime property value. Timestamp of when the chat message was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ChatMessage) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the deletedDateTime property value. Read only. Timestamp at which the chat message was deleted, or null if not deleted.
// Parameters:
//  - value : Value to set for the deletedDateTime property.
func (m *ChatMessage) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
// Sets the etag property value. Read-only. Version number of the chat message.
// Parameters:
//  - value : Value to set for the etag property.
func (m *ChatMessage) SetEtag(value *string)() {
    m.etag = value
}
// Sets the eventDetail property value. Read-only.  If present, represents details of an event that happened in a chat, a channel, or a team, for example, members were added, and so on. For event messages, the messageType property will be set to systemEventMessage.
// Parameters:
//  - value : Value to set for the eventDetail property.
func (m *ChatMessage) SetEventDetail(value *EventMessageDetail)() {
    m.eventDetail = value
}
// Sets the from property value. Details of the sender of the chat message. Can only be set during migration.
// Parameters:
//  - value : Value to set for the from property.
func (m *ChatMessage) SetFrom(value *ChatMessageFromIdentitySet)() {
    m.from = value
}
// Sets the hostedContents property value. Content in a message hosted by Microsoft Teams - for example, images or code snippets.
// Parameters:
//  - value : Value to set for the hostedContents property.
func (m *ChatMessage) SetHostedContents(value []ChatMessageHostedContent)() {
    m.hostedContents = value
}
// Sets the importance property value. The importance of the chat message. The possible values are: normal, high, urgent.
// Parameters:
//  - value : Value to set for the importance property.
func (m *ChatMessage) SetImportance(value *ChatMessageImportance)() {
    m.importance = value
}
// Sets the lastEditedDateTime property value. Read only. Timestamp when edits to the chat message were made. Triggers an 'Edited' flag in the Teams UI. If no edits are made the value is null.
// Parameters:
//  - value : Value to set for the lastEditedDateTime property.
func (m *ChatMessage) SetLastEditedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastEditedDateTime = value
}
// Sets the lastModifiedDateTime property value. Read only. Timestamp when the chat message is created (initial setting) or modified, including when a reaction is added or removed.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *ChatMessage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the locale property value. Locale of the chat message set by the client. Always set to en-us.
// Parameters:
//  - value : Value to set for the locale property.
func (m *ChatMessage) SetLocale(value *string)() {
    m.locale = value
}
// Sets the mentions property value. List of entities mentioned in the chat message. Supported entities are: user, bot, team, and channel.
// Parameters:
//  - value : Value to set for the mentions property.
func (m *ChatMessage) SetMentions(value []ChatMessageMention)() {
    m.mentions = value
}
// Sets the messageType property value. The type of chat message. The possible values are: message, chatEvent, typing, unknownFutureValue, systemEventMessage. Note that you must use the Prefer: include-unknown-enum-members request header to get the following value in this evolvable enum: systemEventMessage.
// Parameters:
//  - value : Value to set for the messageType property.
func (m *ChatMessage) SetMessageType(value *ChatMessageType)() {
    m.messageType = value
}
// Sets the policyViolation property value. Defines the properties of a policy violation set by a data loss prevention (DLP) application.
// Parameters:
//  - value : Value to set for the policyViolation property.
func (m *ChatMessage) SetPolicyViolation(value *ChatMessagePolicyViolation)() {
    m.policyViolation = value
}
// Sets the reactions property value. Reactions for this chat message (for example, Like).
// Parameters:
//  - value : Value to set for the reactions property.
func (m *ChatMessage) SetReactions(value []ChatMessageReaction)() {
    m.reactions = value
}
// Sets the replies property value. Replies for a specified message.
// Parameters:
//  - value : Value to set for the replies property.
func (m *ChatMessage) SetReplies(value []ChatMessage)() {
    m.replies = value
}
// Sets the replyToId property value. Read-only. ID of the parent chat message or root chat message of the thread. (Only applies to chat messages in channels, not chats.)
// Parameters:
//  - value : Value to set for the replyToId property.
func (m *ChatMessage) SetReplyToId(value *string)() {
    m.replyToId = value
}
// Sets the subject property value. The subject of the chat message, in plaintext.
// Parameters:
//  - value : Value to set for the subject property.
func (m *ChatMessage) SetSubject(value *string)() {
    m.subject = value
}
// Sets the summary property value. Summary text of the chat message that could be used for push notifications and summary views or fall back views. Only applies to channel chat messages, not chat messages in a chat.
// Parameters:
//  - value : Value to set for the summary property.
func (m *ChatMessage) SetSummary(value *string)() {
    m.summary = value
}
// Sets the webUrl property value. Read-only. Link to the message in Microsoft Teams.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *ChatMessage) SetWebUrl(value *string)() {
    m.webUrl = value
}
