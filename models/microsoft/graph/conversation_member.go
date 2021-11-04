package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConversationMember struct {
    Entity
    // The display name of the user.
    displayName *string;
    // The roles for that user.
    roles []string;
    // The timestamp denoting how far back a conversation's history is shared with the conversation member. This property is settable only for members of a chat.
    visibleHistoryStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
}
// Instantiates a new conversationMember and sets the default values.
func NewConversationMember()(*ConversationMember) {
    m := &ConversationMember{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The display name of the user.
func (m *ConversationMember) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the roles property value. The roles for that user.
func (m *ConversationMember) GetRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.roles
    }
}
// Gets the visibleHistoryStartDateTime property value. The timestamp denoting how far back a conversation's history is shared with the conversation member. This property is settable only for members of a chat.
func (m *ConversationMember) GetVisibleHistoryStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.visibleHistoryStartDateTime
    }
}
// The deserialization information for the current model
func (m *ConversationMember) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["roles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = *(v.(*string))
        }
        m.SetRoles(res)
        return nil
    }
    res["visibleHistoryStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetVisibleHistoryStartDateTime(val)
        return nil
    }
    return res
}
func (m *ConversationMember) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConversationMember) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("roles", m.GetRoles())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("visibleHistoryStartDateTime", m.GetVisibleHistoryStartDateTime())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The display name of the user.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *ConversationMember) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the roles property value. The roles for that user.
// Parameters:
//  - value : Value to set for the roles property.
func (m *ConversationMember) SetRoles(value []string)() {
    m.roles = value
}
// Sets the visibleHistoryStartDateTime property value. The timestamp denoting how far back a conversation's history is shared with the conversation member. This property is settable only for members of a chat.
// Parameters:
//  - value : Value to set for the visibleHistoryStartDateTime property.
func (m *ConversationMember) SetVisibleHistoryStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.visibleHistoryStartDateTime = value
}
