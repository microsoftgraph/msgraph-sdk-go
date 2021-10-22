package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ChatMessage struct {
    Entity
    attachments []ChatMessageAttachment;
    body *ItemBody;
    channelIdentity *ChannelIdentity;
    chatId *string;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    deletedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    etag *string;
    from *ChatMessageFromIdentitySet;
    hostedContents []ChatMessageHostedContent;
    importance *ChatMessageImportance;
    lastEditedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    locale *string;
    mentions []ChatMessageMention;
    messageType *ChatMessageType;
    policyViolation *ChatMessagePolicyViolation;
    reactions []ChatMessageReaction;
    replies []ChatMessage;
    replyToId *string;
    subject *string;
    summary *string;
    webUrl *string;
}
func NewChatMessage()(*ChatMessage) {
    m := &ChatMessage{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ChatMessage) GetAttachments()([]ChatMessageAttachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
func (m *ChatMessage) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
func (m *ChatMessage) GetChannelIdentity()(*ChannelIdentity) {
    if m == nil {
        return nil
    } else {
        return m.channelIdentity
    }
}
func (m *ChatMessage) GetChatId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.chatId
    }
}
func (m *ChatMessage) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *ChatMessage) GetDeletedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.deletedDateTime
    }
}
func (m *ChatMessage) GetEtag()(*string) {
    if m == nil {
        return nil
    } else {
        return m.etag
    }
}
func (m *ChatMessage) GetFrom()(*ChatMessageFromIdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
func (m *ChatMessage) GetHostedContents()([]ChatMessageHostedContent) {
    if m == nil {
        return nil
    } else {
        return m.hostedContents
    }
}
func (m *ChatMessage) GetImportance()(*ChatMessageImportance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
func (m *ChatMessage) GetLastEditedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastEditedDateTime
    }
}
func (m *ChatMessage) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *ChatMessage) GetLocale()(*string) {
    if m == nil {
        return nil
    } else {
        return m.locale
    }
}
func (m *ChatMessage) GetMentions()([]ChatMessageMention) {
    if m == nil {
        return nil
    } else {
        return m.mentions
    }
}
func (m *ChatMessage) GetMessageType()(*ChatMessageType) {
    if m == nil {
        return nil
    } else {
        return m.messageType
    }
}
func (m *ChatMessage) GetPolicyViolation()(*ChatMessagePolicyViolation) {
    if m == nil {
        return nil
    } else {
        return m.policyViolation
    }
}
func (m *ChatMessage) GetReactions()([]ChatMessageReaction) {
    if m == nil {
        return nil
    } else {
        return m.reactions
    }
}
func (m *ChatMessage) GetReplies()([]ChatMessage) {
    if m == nil {
        return nil
    } else {
        return m.replies
    }
}
func (m *ChatMessage) GetReplyToId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.replyToId
    }
}
func (m *ChatMessage) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
func (m *ChatMessage) GetSummary()(*string) {
    if m == nil {
        return nil
    } else {
        return m.summary
    }
}
func (m *ChatMessage) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
func (m *ChatMessage) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageAttachment() })
        if err != nil {
            return err
        }
        res := make([]ChatMessageAttachment, len(val))
        for i, v := range val {
            res[i] = *(v.(*ChatMessageAttachment))
        }
        m.SetAttachments(res)
        return nil
    }
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        m.SetBody(val.(*ItemBody))
        return nil
    }
    res["channelIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChannelIdentity() })
        if err != nil {
            return err
        }
        m.SetChannelIdentity(val.(*ChannelIdentity))
        return nil
    }
    res["chatId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChatId(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["deletedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetDeletedDateTime(val)
        return nil
    }
    res["etag"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEtag(val)
        return nil
    }
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageFromIdentitySet() })
        if err != nil {
            return err
        }
        m.SetFrom(val.(*ChatMessageFromIdentitySet))
        return nil
    }
    res["hostedContents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageHostedContent() })
        if err != nil {
            return err
        }
        res := make([]ChatMessageHostedContent, len(val))
        for i, v := range val {
            res[i] = *(v.(*ChatMessageHostedContent))
        }
        m.SetHostedContents(res)
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessageImportance)
        if err != nil {
            return err
        }
        cast := val.(ChatMessageImportance)
        m.SetImportance(&cast)
        return nil
    }
    res["lastEditedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastEditedDateTime(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["locale"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLocale(val)
        return nil
    }
    res["mentions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageMention() })
        if err != nil {
            return err
        }
        res := make([]ChatMessageMention, len(val))
        for i, v := range val {
            res[i] = *(v.(*ChatMessageMention))
        }
        m.SetMentions(res)
        return nil
    }
    res["messageType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChatMessageType)
        if err != nil {
            return err
        }
        cast := val.(ChatMessageType)
        m.SetMessageType(&cast)
        return nil
    }
    res["policyViolation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessagePolicyViolation() })
        if err != nil {
            return err
        }
        m.SetPolicyViolation(val.(*ChatMessagePolicyViolation))
        return nil
    }
    res["reactions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessageReaction() })
        if err != nil {
            return err
        }
        res := make([]ChatMessageReaction, len(val))
        for i, v := range val {
            res[i] = *(v.(*ChatMessageReaction))
        }
        m.SetReactions(res)
        return nil
    }
    res["replies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessage() })
        if err != nil {
            return err
        }
        res := make([]ChatMessage, len(val))
        for i, v := range val {
            res[i] = *(v.(*ChatMessage))
        }
        m.SetReplies(res)
        return nil
    }
    res["replyToId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetReplyToId(val)
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubject(val)
        return nil
    }
    res["summary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSummary(val)
        return nil
    }
    res["webUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWebUrl(val)
        return nil
    }
    return res
}
func (m *ChatMessage) IsNil()(bool) {
    return m == nil
}
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
func (m *ChatMessage) SetAttachments(value []ChatMessageAttachment)() {
    m.attachments = value
}
func (m *ChatMessage) SetBody(value *ItemBody)() {
    m.body = value
}
func (m *ChatMessage) SetChannelIdentity(value *ChannelIdentity)() {
    m.channelIdentity = value
}
func (m *ChatMessage) SetChatId(value *string)() {
    m.chatId = value
}
func (m *ChatMessage) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *ChatMessage) SetDeletedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.deletedDateTime = value
}
func (m *ChatMessage) SetEtag(value *string)() {
    m.etag = value
}
func (m *ChatMessage) SetFrom(value *ChatMessageFromIdentitySet)() {
    m.from = value
}
func (m *ChatMessage) SetHostedContents(value []ChatMessageHostedContent)() {
    m.hostedContents = value
}
func (m *ChatMessage) SetImportance(value *ChatMessageImportance)() {
    m.importance = value
}
func (m *ChatMessage) SetLastEditedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastEditedDateTime = value
}
func (m *ChatMessage) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *ChatMessage) SetLocale(value *string)() {
    m.locale = value
}
func (m *ChatMessage) SetMentions(value []ChatMessageMention)() {
    m.mentions = value
}
func (m *ChatMessage) SetMessageType(value *ChatMessageType)() {
    m.messageType = value
}
func (m *ChatMessage) SetPolicyViolation(value *ChatMessagePolicyViolation)() {
    m.policyViolation = value
}
func (m *ChatMessage) SetReactions(value []ChatMessageReaction)() {
    m.reactions = value
}
func (m *ChatMessage) SetReplies(value []ChatMessage)() {
    m.replies = value
}
func (m *ChatMessage) SetReplyToId(value *string)() {
    m.replyToId = value
}
func (m *ChatMessage) SetSubject(value *string)() {
    m.subject = value
}
func (m *ChatMessage) SetSummary(value *string)() {
    m.summary = value
}
func (m *ChatMessage) SetWebUrl(value *string)() {
    m.webUrl = value
}
