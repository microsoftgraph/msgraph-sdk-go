package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Post provides operations to manage the collection of group entities.
type Post struct {
    OutlookItem
    // Read-only. Nullable. Supports $expand.
    attachments []Attachmentable;
    // The contents of the post. This is a default property. This property can be null.
    body ItemBodyable;
    // Unique ID of the conversation. Read-only.
    conversationId *string;
    // Unique ID of the conversation thread. Read-only.
    conversationThreadId *string;
    // The collection of open extensions defined for the post. Read-only. Nullable. Supports $expand.
    extensions []Extensionable;
    // 
    from Recipientable;
    // Indicates whether the post has at least one attachment. This is a default property.
    hasAttachments *bool;
    // Read-only. Supports $expand.
    inReplyTo Postable;
    // The collection of multi-value extended properties defined for the post. Read-only. Nullable.
    multiValueExtendedProperties []MultiValueLegacyExtendedPropertyable;
    // Conversation participants that were added to the thread as part of this post.
    newParticipants []Recipientable;
    // Specifies when the post was received. The DateTimeOffset type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    receivedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Contains the address of the sender. The value of Sender is assumed to be the address of the authenticated user in the case when Sender is not specified. This is a default property.
    sender Recipientable;
    // The collection of single-value extended properties defined for the post. Read-only. Nullable.
    singleValueExtendedProperties []SingleValueLegacyExtendedPropertyable;
}
// NewPost instantiates a new post and sets the default values.
func NewPost()(*Post) {
    m := &Post{
        OutlookItem: *NewOutlookItem(),
    }
    return m
}
// CreatePostFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePostFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPost(), nil
}
// GetAttachments gets the attachments property value. Read-only. Nullable. Supports $expand.
func (m *Post) GetAttachments()([]Attachmentable) {
    if m == nil {
        return nil
    } else {
        return m.attachments
    }
}
// GetBody gets the body property value. The contents of the post. This is a default property. This property can be null.
func (m *Post) GetBody()(ItemBodyable) {
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
func (m *Post) GetExtensions()([]Extensionable) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Post) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OutlookItem.GetFieldDeserializers()
    res["attachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["from"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFrom(val.(Recipientable))
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
        val, err := n.GetObjectValue(CreatePostFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInReplyTo(val.(Postable))
        }
        return nil
    }
    res["multiValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    res["newParticipants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Recipientable, len(val))
            for i, v := range val {
                res[i] = v.(Recipientable)
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
        val, err := n.GetObjectValue(CreateRecipientFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSender(val.(Recipientable))
        }
        return nil
    }
    res["singleValueExtendedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
    return res
}
// GetFrom gets the from property value. 
func (m *Post) GetFrom()(Recipientable) {
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
func (m *Post) GetInReplyTo()(Postable) {
    if m == nil {
        return nil
    } else {
        return m.inReplyTo
    }
}
// GetMultiValueExtendedProperties gets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the post. Read-only. Nullable.
func (m *Post) GetMultiValueExtendedProperties()([]MultiValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.multiValueExtendedProperties
    }
}
// GetNewParticipants gets the newParticipants property value. Conversation participants that were added to the thread as part of this post.
func (m *Post) GetNewParticipants()([]Recipientable) {
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
func (m *Post) GetSender()(Recipientable) {
    if m == nil {
        return nil
    } else {
        return m.sender
    }
}
// GetSingleValueExtendedProperties gets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the post. Read-only. Nullable.
func (m *Post) GetSingleValueExtendedProperties()([]SingleValueLegacyExtendedPropertyable) {
    if m == nil {
        return nil
    } else {
        return m.singleValueExtendedProperties
    }
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
    if m.GetExtensions() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetMultiValueExtendedProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMultiValueExtendedProperties()))
        for i, v := range m.GetMultiValueExtendedProperties() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("multiValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetNewParticipants() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNewParticipants()))
        for i, v := range m.GetNewParticipants() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m.GetSingleValueExtendedProperties() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSingleValueExtendedProperties()))
        for i, v := range m.GetSingleValueExtendedProperties() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("singleValueExtendedProperties", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttachments sets the attachments property value. Read-only. Nullable. Supports $expand.
func (m *Post) SetAttachments(value []Attachmentable)() {
    if m != nil {
        m.attachments = value
    }
}
// SetBody sets the body property value. The contents of the post. This is a default property. This property can be null.
func (m *Post) SetBody(value ItemBodyable)() {
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
func (m *Post) SetExtensions(value []Extensionable)() {
    if m != nil {
        m.extensions = value
    }
}
// SetFrom sets the from property value. 
func (m *Post) SetFrom(value Recipientable)() {
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
func (m *Post) SetInReplyTo(value Postable)() {
    if m != nil {
        m.inReplyTo = value
    }
}
// SetMultiValueExtendedProperties sets the multiValueExtendedProperties property value. The collection of multi-value extended properties defined for the post. Read-only. Nullable.
func (m *Post) SetMultiValueExtendedProperties(value []MultiValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.multiValueExtendedProperties = value
    }
}
// SetNewParticipants sets the newParticipants property value. Conversation participants that were added to the thread as part of this post.
func (m *Post) SetNewParticipants(value []Recipientable)() {
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
func (m *Post) SetSender(value Recipientable)() {
    if m != nil {
        m.sender = value
    }
}
// SetSingleValueExtendedProperties sets the singleValueExtendedProperties property value. The collection of single-value extended properties defined for the post. Read-only. Nullable.
func (m *Post) SetSingleValueExtendedProperties(value []SingleValueLegacyExtendedPropertyable)() {
    if m != nil {
        m.singleValueExtendedProperties = value
    }
}
