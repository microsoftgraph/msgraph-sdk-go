package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Conversation struct {
    Entity
    // Indicates whether any of the posts within this Conversation has at least one attachment. Supports $filter (eq, ne) and $search.
    hasAttachments *bool;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    lastDeliveredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A short summary from the body of the latest post in this conversation. Supports $filter (eq, ne, le, ge).
    preview *string;
    // A collection of all the conversation threads in the conversation. A navigation property. Read-only. Nullable.
    threads []ConversationThread;
    // The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated.
    topic *string;
    // All the users that sent a message to this Conversation.
    uniqueSenders []string;
}
// Instantiates a new conversation and sets the default values.
func NewConversation()(*Conversation) {
    m := &Conversation{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the hasAttachments property value. Indicates whether any of the posts within this Conversation has at least one attachment. Supports $filter (eq, ne) and $search.
func (m *Conversation) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// Gets the lastDeliveredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *Conversation) GetLastDeliveredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastDeliveredDateTime
    }
}
// Gets the preview property value. A short summary from the body of the latest post in this conversation. Supports $filter (eq, ne, le, ge).
func (m *Conversation) GetPreview()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preview
    }
}
// Gets the threads property value. A collection of all the conversation threads in the conversation. A navigation property. Read-only. Nullable.
func (m *Conversation) GetThreads()([]ConversationThread) {
    if m == nil {
        return nil
    } else {
        return m.threads
    }
}
// Gets the topic property value. The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated.
func (m *Conversation) GetTopic()(*string) {
    if m == nil {
        return nil
    } else {
        return m.topic
    }
}
// Gets the uniqueSenders property value. All the users that sent a message to this Conversation.
func (m *Conversation) GetUniqueSenders()([]string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueSenders
    }
}
// The deserialization information for the current model
func (m *Conversation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["hasAttachments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetHasAttachments(val)
        return nil
    }
    res["lastDeliveredDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastDeliveredDateTime(val)
        return nil
    }
    res["preview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPreview(val)
        return nil
    }
    res["threads"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversationThread() })
        if err != nil {
            return err
        }
        res := make([]ConversationThread, len(val))
        for i, v := range val {
            res[i] = *(v.(*ConversationThread))
        }
        m.SetThreads(res)
        return nil
    }
    res["topic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTopic(val)
        return nil
    }
    res["uniqueSenders"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetUniqueSenders(res)
        return nil
    }
    return res
}
func (m *Conversation) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Conversation) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("hasAttachments", m.GetHasAttachments())
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
        err = writer.WriteStringValue("preview", m.GetPreview())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetThreads()))
        for i, v := range m.GetThreads() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("threads", cast)
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
        err = writer.WriteCollectionOfStringValues("uniqueSenders", m.GetUniqueSenders())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the hasAttachments property value. Indicates whether any of the posts within this Conversation has at least one attachment. Supports $filter (eq, ne) and $search.
// Parameters:
//  - value : Value to set for the hasAttachments property.
func (m *Conversation) SetHasAttachments(value *bool)() {
    m.hasAttachments = value
}
// Sets the lastDeliveredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
// Parameters:
//  - value : Value to set for the lastDeliveredDateTime property.
func (m *Conversation) SetLastDeliveredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastDeliveredDateTime = value
}
// Sets the preview property value. A short summary from the body of the latest post in this conversation. Supports $filter (eq, ne, le, ge).
// Parameters:
//  - value : Value to set for the preview property.
func (m *Conversation) SetPreview(value *string)() {
    m.preview = value
}
// Sets the threads property value. A collection of all the conversation threads in the conversation. A navigation property. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the threads property.
func (m *Conversation) SetThreads(value []ConversationThread)() {
    m.threads = value
}
// Sets the topic property value. The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated.
// Parameters:
//  - value : Value to set for the topic property.
func (m *Conversation) SetTopic(value *string)() {
    m.topic = value
}
// Sets the uniqueSenders property value. All the users that sent a message to this Conversation.
// Parameters:
//  - value : Value to set for the uniqueSenders property.
func (m *Conversation) SetUniqueSenders(value []string)() {
    m.uniqueSenders = value
}
