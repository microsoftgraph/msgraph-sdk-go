package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Post 
type Post struct {
    OutlookItem
    // Read-only. Nullable. Supports $expand.
    attachments []Attachment;
    // The contents of the post. This is a default property. This property can be null.
    body *ItemBody;
    // Unique ID of the conversation. Read-only.
    conversationId *string;
    // Unique ID of the conversation thread. Read-only.
    conversationThreadId *string;
    // The collection of open extensions defined for the post. Read-only. Nullable. Supports $expand.
    extensions []Extension;
    // 
    from *Recipient;
    // Indicates whether the post has at least one attachment. This is a default property.
    hasAttachments *bool;
    // Read-only. Supports $expand.
    inReplyTo *Post;
    // The collection of multi-value extended properties defined for the post. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedProperty;
    // Conversation participants that were added to the thread as part of this post.
    newParticipants []Recipient;
    // Specifies when the post was received. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Contains the address of the sender. The value of Sender is assumed to be the address of the authenticated user in the case when Sender is not specified. This is a default property.
    sender *Recipient;
    // The collection of single-value extended properties defined for the post. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedProperty;
}
// NewPost instantiates a new post and sets the default values.
func NewPost()(*Post) {
    m := &Post{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
// GetAttachments gets the attachments property value. Read-only. Nullable. Supports $expand.
func (m *Post) GetAttachments()([]Attachment) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// GetBody gets the body property value. The contents of the post. This is a default property. This property can be null.
func (m *Post) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// GetConversationId gets the conversationId property value. Unique ID of the conversation. Read-only.
func (m *Post) GetConversationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conversationId
    }
}
// GetConversationThreadId gets the conversationThreadId property value. Unique ID of the conversation thread. Read-only.
func (m *Post) GetConversationThreadId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.conversationThreadId
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the post. Read-only. Nullable. Supports $expand.
func (m *Post) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFrom gets the from property value. 
func (m *Post) GetFrom()(*Recipient) {
    if m == nil {
        return nil
    } else {
        return m.from
    }
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether the post has at least one attachment. This is a default property.
func (m *Post) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// GetInReplyTo gets the inReplyTo property value. Read-only. Supports $expand.
func (m *Post) GetInReplyTo()(*Post) {
    if m == nil {
        return nil
    } else {
        return m.inReplyTo
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the post. Read-only. Nullable.
func (m *Post) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetNewParticipants gets the newParticipants property value. Conversation participants that were added to the thread as part of this post.
func (m *Post) GetNewParticipants()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.newParticipants
    }
}
// GetReceivedDateTime gets the receivedDateTime property value. Specifies when the post was received. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Post) GetReceivedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.receivedDateTime
    }
}
// GetSender gets the sender property value. Contains the address of the sender. The value of Sender is assumed to be the address of the authenticated user in the case when Sender is not specified. This is a default property.
func (m *Post) GetSender()(*Recipient) {
    if m == nil {
        return nil
    } else {
        return m.sender
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the post. Read-only. Nullable.
func (m *Post) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedProperty) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Post) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
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
    res["conversationThreadId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConversationThreadId(val)
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
    res["inReplyTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPost() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInReplyTo(val.(*Post))
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
    res["newParticipants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecipient() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipient, len(val))
            for i, v := range val {
                res[i] = *(v.(*Recipient))
            }
            m.SetNewParticipants(res)
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
    return res
}
func (m *Post) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Post) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteObjectValue("body", m.GetBody())
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
        err = writer.WriteStringValue("conversationThreadId", m.GetConversationThreadId())
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
    {
        err = writer.WriteObjectValue("inReplyTo", m.GetInReplyTo())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNewParticipants()))
        for i, v := range m.GetNewParticipants() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("newParticipants", cast)
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
        err = writer.WriteObjectValue("sender", m.GetSender())
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
    return nil
}
// SetAttachments sets the attachments property value. Read-only. Nullable. Supports $expand.
func (m *Post) SetAttachments(value []Attachment)() {
    if m != nil {
        m.attachments = value
    }
}
// SetBody sets the body property value. The contents of the post. This is a default property. This property can be null.
func (m *Post) SetBody(value *ItemBody)() {
    if m != nil {
        m.body = value
    }
}
// SetConversationId sets the conversationId property value. Unique ID of the conversation. Read-only.
func (m *Post) SetConversationId(value *string)() {
    if m != nil {
        m.conversationId = value
    }
}
// SetConversationThreadId sets the conversationThreadId property value. Unique ID of the conversation thread. Read-only.
func (m *Post) SetConversationThreadId(value *string)() {
    if m != nil {
        m.conversationThreadId = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the post. Read-only. Nullable. Supports $expand.
func (m *Post) SetExtensions(value []Extension)() {
    if m != nil {
        m.extensions = value
    }
}
// SetFrom sets the from property value. 
func (m *Post) SetFrom(value *Recipient)() {
    if m != nil {
        m.from = value
    }
}
// SetHasAttachments sets the hasAttachments property value. Indicates whether the post has at least one attachment. This is a default property.
func (m *Post) SetHasAttachments(value *bool)() {
    if m != nil {
        m.hasAttachments = value
    }
}
// SetInReplyTo sets the inReplyTo property value. Read-only. Supports $expand.
func (m *Post) SetInReplyTo(value *Post)() {
    if m != nil {
        m.inReplyTo = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the post. Read-only. Nullable.
func (m *Post) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedProperty)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetNewParticipants sets the newParticipants property value. Conversation participants that were added to the thread as part of this post.
func (m *Post) SetNewParticipants(value []Recipient)() {
    if m != nil {
        m.newParticipants = value
    }
}
// SetReceivedDateTime sets the receivedDateTime property value. Specifies when the post was received. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Post) SetReceivedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.receivedDateTime = value
    }
}
// SetSender sets the sender property value. Contains the address of the sender. The value of Sender is assumed to be the address of the authenticated user in the case when Sender is not specified. This is a default property.
func (m *Post) SetSender(value *Recipient)() {
    if m != nil {
        m.sender = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the post. Read-only. Nullable.
func (m *Post) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedProperty)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
