package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConversationThread struct {
    Entity
    // The Cc: recipients for the thread. Returned only on $select.
    ccRecipients []Recipient;
    // Indicates whether any of the posts within this thread has at least one attachment. Returned by default.
    hasAttachments *bool;
    // Indicates if the thread is locked. Returned by default.
    isLocked *bool;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.Returned by default.
    lastDeliveredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only. Nullable.
    posts []Post;
    // A short summary from the body of the latest post in this conversation. Returned by default.
    preview *string;
    // The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated. Returned by default.
    topic *string;
    // The To: recipients for the thread. Returned only on $select.
    toRecipients []Recipient;
    // All the users that sent a message to this thread. Returned by default.
    uniqueSenders []string;
}
// Instantiates a new conversationThread and sets the default values.
func NewConversationThread()(*ConversationThread) {
    m := &ConversationThread{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the ccRecipients property value. The Cc: recipients for the thread. Returned only on $select.
func (m *ConversationThread) GetCcRecipients()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.ccRecipients
    }
}
// Gets the hasAttachments property value. Indicates whether any of the posts within this thread has at least one attachment. Returned by default.
func (m *ConversationThread) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// Gets the isLocked property value. Indicates if the thread is locked. Returned by default.
func (m *ConversationThread) GetIsLocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isLocked
    }
}
// Gets the lastDeliveredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.Returned by default.
func (m *ConversationThread) GetLastDeliveredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastDeliveredDateTime
    }
}
// Gets the posts property value. Read-only. Nullable.
func (m *ConversationThread) GetPosts()([]Post) {
    if m == nil {
        return nil
    } else {
        return m.posts
    }
}
// Gets the preview property value. A short summary from the body of the latest post in this conversation. Returned by default.
func (m *ConversationThread) GetPreview()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preview
    }
}
// Gets the topic property value. The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated. Returned by default.
func (m *ConversationThread) GetTopic()(*string) {
    if m == nil {
        return nil
    } else {
        return m.topic
    }
}
// Gets the toRecipients property value. The To: recipients for the thread. Returned only on $select.
func (m *ConversationThread) GetToRecipients()([]Recipient) {
    if m == nil {
        return nil
    } else {
        return m.toRecipients
    }
}
// Gets the uniqueSenders property value. All the users that sent a message to this thread. Returned by default.
func (m *ConversationThread) GetUniqueSenders()([]string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueSenders
    }
}
// The deserialization information for the current model
func (m *ConversationThread) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["isLocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsLocked(val)
        }
        return nil
    }
    res["lastDeliveredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastDeliveredDateTime(val)
        }
        return nil
    }
    res["posts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPost() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Post, len(val))
            for i, v := range val {
                res[i] = *(v.(*Post))
            }
            m.SetPosts(res)
        }
        return nil
    }
    res["preview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPreview(val)
        }
        return nil
    }
    res["topic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTopic(val)
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
    res["uniqueSenders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetUniqueSenders(res)
        }
        return nil
    }
    return res
}
func (m *ConversationThread) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConversationThread) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
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
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isLocked", m.GetIsLocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastDeliveredDateTime", m.GetLastDeliveredDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetPosts()))
        for i, v := range m.GetPosts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("posts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("preview", m.GetPreview())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("topic", m.GetTopic())
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
        err = writer.WriteCollectionOfStringValues("uniqueSenders", m.GetUniqueSenders())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the ccRecipients property value. The Cc: recipients for the thread. Returned only on $select.
// Parameters:
//  - value : Value to set for the ccRecipients property.
func (m *ConversationThread) SetCcRecipients(value []Recipient)() {
    m.ccRecipients = value
}
// Sets the hasAttachments property value. Indicates whether any of the posts within this thread has at least one attachment. Returned by default.
// Parameters:
//  - value : Value to set for the hasAttachments property.
func (m *ConversationThread) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// Sets the isLocked property value. Indicates if the thread is locked. Returned by default.
// Parameters:
//  - value : Value to set for the isLocked property.
func (m *ConversationThread) SetIsLocked(value *bool)() {
    m.isLocked = value
}
// Sets the lastDeliveredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.Returned by default.
// Parameters:
//  - value : Value to set for the lastDeliveredDateTime property.
func (m *ConversationThread) SetLastDeliveredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDeliveredDateTime = value
}
// Sets the posts property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the posts property.
func (m *ConversationThread) SetPosts(value []Post)() {
    m.posts = value
}
// Sets the preview property value. A short summary from the body of the latest post in this conversation. Returned by default.
// Parameters:
//  - value : Value to set for the preview property.
func (m *ConversationThread) SetPreview(value *string)() {
    m.preview = value
}
// Sets the topic property value. The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated. Returned by default.
// Parameters:
//  - value : Value to set for the topic property.
func (m *ConversationThread) SetTopic(value *string)() {
    m.topic = value
}
// Sets the toRecipients property value. The To: recipients for the thread. Returned only on $select.
// Parameters:
//  - value : Value to set for the toRecipients property.
func (m *ConversationThread) SetToRecipients(value []Recipient)() {
    m.toRecipients = value
}
// Sets the uniqueSenders property value. All the users that sent a message to this thread. Returned by default.
// Parameters:
//  - value : Value to set for the uniqueSenders property.
func (m *ConversationThread) SetUniqueSenders(value []string)() {
    m.uniqueSenders = value
}
