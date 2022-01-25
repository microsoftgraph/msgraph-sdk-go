package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Conversation 
type Conversation struct {
    Entity
    // Indicates whether any of the posts within this Conversation has at least one attachment. Supports $filter (eq, ne) and $search.
    hasAttachments *bool;
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, le, ge).
    lastDeliveredDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A short summary from the body of the latest post in this conversation.
    preview *string;
    // A collection of all the conversation threads in the conversation. A navigation property. Read-only. Nullable.
    threads []ConversationThread;
    // The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated.
    topic *string;
    // All the users that sent a message to this Conversation.
    uniqueSenders []string;
}
// NewConversation instantiates a new conversation and sets the default values.
func NewConversation()(*Conversation) {
    m := &Conversation{
        Entity: *NewEntity(),
    }
    return m
}
// GetHasAttachments gets the hasAttachments property value. Indicates whether any of the posts within this Conversation has at least one attachment. Supports $filter (eq, ne) and $search.
func (m *Conversation) GetHasAttachments()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasAttachments
    }
}
// GetLastDeliveredDateTime gets the lastDeliveredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, le, ge).
func (m *Conversation) GetLastDeliveredDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastDeliveredDateTime
    }
}
// GetPreview gets the preview property value. A short summary from the body of the latest post in this conversation.
func (m *Conversation) GetPreview()(*string) {
    if m == nil {
        return nil
    } else {
        return m.preview
    }
}
// GetThreads gets the threads property value. A collection of all the conversation threads in the conversation. A navigation property. Read-only. Nullable.
func (m *Conversation) GetThreads()([]ConversationThread) {
    if m == nil {
        return nil
    } else {
        return m.threads
    }
}
// GetTopic gets the topic property value. The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated.
func (m *Conversation) GetTopic()(*string) {
    if m == nil {
        return nil
    } else {
        return m.topic
    }
}
// GetUniqueSenders gets the uniqueSenders property value. All the users that sent a message to this Conversation.
func (m *Conversation) GetUniqueSenders()([]string) {
    if m == nil {
        return nil
    } else {
        return m.uniqueSenders
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Conversation) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["threads"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversationThread() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationThread, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConversationThread))
            }
            m.SetThreads(res)
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
func (m *Conversation) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetHasAttachments sets the hasAttachments property value. Indicates whether any of the posts within this Conversation has at least one attachment. Supports $filter (eq, ne) and $search.
func (m *Conversation) SetHasAttachments(value *bool)() {
    if m != nil {
        m.hasAttachments = value
    }
}
// SetLastDeliveredDateTime sets the lastDeliveredDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Supports $filter (eq, ne, le, ge).
func (m *Conversation) SetLastDeliveredDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastDeliveredDateTime = value
    }
}
// SetPreview sets the preview property value. A short summary from the body of the latest post in this conversation.
func (m *Conversation) SetPreview(value *string)() {
    if m != nil {
        m.preview = value
    }
}
// SetThreads sets the threads property value. A collection of all the conversation threads in the conversation. A navigation property. Read-only. Nullable.
func (m *Conversation) SetThreads(value []ConversationThread)() {
    if m != nil {
        m.threads = value
    }
}
// SetTopic sets the topic property value. The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated.
func (m *Conversation) SetTopic(value *string)() {
    if m != nil {
        m.topic = value
    }
}
// SetUniqueSenders sets the uniqueSenders property value. All the users that sent a message to this Conversation.
func (m *Conversation) SetUniqueSenders(value []string)() {
    if m != nil {
        m.uniqueSenders = value
    }
}
