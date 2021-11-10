package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Channel struct {
    Entity
    // Read only. Timestamp at which the channel was created.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Optional textual description for the channel.
    description *string;
    // Channel name as it will appear to the user in Microsoft Teams.
    displayName *string;
    // The email address for sending messages to the channel. Read-only.
    email *string;
    // Metadata for the location where the channel's files are stored.
    filesFolder *DriveItem;
    // Indicates whether the channel should automatically be marked 'favorite' for all members of the team. Can only be set programmatically with Create team. Default: false.
    isFavoriteByDefault *bool;
    // A collection of membership records associated with the channel.
    members []ConversationMember;
    // The type of the channel. Can be set during creation and can't be changed. Possible values are: standard - Channel inherits the list of members of the parent team; private - Channel can have members that are a subset of all the members on the parent team.
    membershipType *ChannelMembershipType;
    // A collection of all the messages in the channel. A navigation property. Nullable.
    messages []ChatMessage;
    // A collection of all the tabs in the channel. A navigation property.
    tabs []TeamsTab;
    // A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not parsed. Read-only.
    webUrl *string;
}
// Instantiates a new channel and sets the default values.
func NewChannel()(*Channel) {
    m := &Channel{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. Read only. Timestamp at which the channel was created.
func (m *Channel) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the description property value. Optional textual description for the channel.
func (m *Channel) GetDescription()(*string) {
    if m == nil {
        return nil
    } else {
        return m.description
    }
}
// Gets the displayName property value. Channel name as it will appear to the user in Microsoft Teams.
func (m *Channel) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the email property value. The email address for sending messages to the channel. Read-only.
func (m *Channel) GetEmail()(*string) {
    if m == nil {
        return nil
    } else {
        return m.email
    }
}
// Gets the filesFolder property value. Metadata for the location where the channel's files are stored.
func (m *Channel) GetFilesFolder()(*DriveItem) {
    if m == nil {
        return nil
    } else {
        return m.filesFolder
    }
}
// Gets the isFavoriteByDefault property value. Indicates whether the channel should automatically be marked 'favorite' for all members of the team. Can only be set programmatically with Create team. Default: false.
func (m *Channel) GetIsFavoriteByDefault()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFavoriteByDefault
    }
}
// Gets the members property value. A collection of membership records associated with the channel.
func (m *Channel) GetMembers()([]ConversationMember) {
    if m == nil {
        return nil
    } else {
        return m.members
    }
}
// Gets the membershipType property value. The type of the channel. Can be set during creation and can't be changed. Possible values are: standard - Channel inherits the list of members of the parent team; private - Channel can have members that are a subset of all the members on the parent team.
func (m *Channel) GetMembershipType()(*ChannelMembershipType) {
    if m == nil {
        return nil
    } else {
        return m.membershipType
    }
}
// Gets the messages property value. A collection of all the messages in the channel. A navigation property. Nullable.
func (m *Channel) GetMessages()([]ChatMessage) {
    if m == nil {
        return nil
    } else {
        return m.messages
    }
}
// Gets the tabs property value. A collection of all the tabs in the channel. A navigation property.
func (m *Channel) GetTabs()([]TeamsTab) {
    if m == nil {
        return nil
    } else {
        return m.tabs
    }
}
// Gets the webUrl property value. A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not parsed. Read-only.
func (m *Channel) GetWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.webUrl
    }
}
// The deserialization information for the current model
func (m *Channel) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
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
    res["description"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDescription(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["email"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmail(val)
        }
        return nil
    }
    res["filesFolder"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDriveItem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFilesFolder(val.(*DriveItem))
        }
        return nil
    }
    res["isFavoriteByDefault"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFavoriteByDefault(val)
        }
        return nil
    }
    res["members"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConversationMember() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConversationMember, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConversationMember))
            }
            m.SetMembers(res)
        }
        return nil
    }
    res["membershipType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseChannelMembershipType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ChannelMembershipType)
            m.SetMembershipType(&cast)
        }
        return nil
    }
    res["messages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatMessage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChatMessage, len(val))
            for i, v := range val {
                res[i] = *(v.(*ChatMessage))
            }
            m.SetMessages(res)
        }
        return nil
    }
    res["tabs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTeamsTab() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TeamsTab, len(val))
            for i, v := range val {
                res[i] = *(v.(*TeamsTab))
            }
            m.SetTabs(res)
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
func (m *Channel) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Channel) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("email", m.GetEmail())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("filesFolder", m.GetFilesFolder())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFavoriteByDefault", m.GetIsFavoriteByDefault())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMembers()))
        for i, v := range m.GetMembers() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("members", cast)
        if err != nil {
            return err
        }
    }
    if m.GetMembershipType() != nil {
        cast := m.GetMembershipType().String()
        err = writer.WriteStringValue("membershipType", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMessages()))
        for i, v := range m.GetMessages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("messages", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTabs()))
        for i, v := range m.GetTabs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tabs", cast)
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
// Sets the createdDateTime property value. Read only. Timestamp at which the channel was created.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Channel) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the description property value. Optional textual description for the channel.
// Parameters:
//  - value : Value to set for the description property.
func (m *Channel) SetDescription(value *string)() {
    m.description = value
}
// Sets the displayName property value. Channel name as it will appear to the user in Microsoft Teams.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *Channel) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the email property value. The email address for sending messages to the channel. Read-only.
// Parameters:
//  - value : Value to set for the email property.
func (m *Channel) SetEmail(value *string)() {
    m.email = value
}
// Sets the filesFolder property value. Metadata for the location where the channel's files are stored.
// Parameters:
//  - value : Value to set for the filesFolder property.
func (m *Channel) SetFilesFolder(value *DriveItem)() {
    m.filesFolder = value
}
// Sets the isFavoriteByDefault property value. Indicates whether the channel should automatically be marked 'favorite' for all members of the team. Can only be set programmatically with Create team. Default: false.
// Parameters:
//  - value : Value to set for the isFavoriteByDefault property.
func (m *Channel) SetIsFavoriteByDefault(value *bool)() {
    m.isFavoriteByDefault = value
}
// Sets the members property value. A collection of membership records associated with the channel.
// Parameters:
//  - value : Value to set for the members property.
func (m *Channel) SetMembers(value []ConversationMember)() {
    m.members = value
}
// Sets the membershipType property value. The type of the channel. Can be set during creation and can't be changed. Possible values are: standard - Channel inherits the list of members of the parent team; private - Channel can have members that are a subset of all the members on the parent team.
// Parameters:
//  - value : Value to set for the membershipType property.
func (m *Channel) SetMembershipType(value *ChannelMembershipType)() {
    m.membershipType = value
}
// Sets the messages property value. A collection of all the messages in the channel. A navigation property. Nullable.
// Parameters:
//  - value : Value to set for the messages property.
func (m *Channel) SetMessages(value []ChatMessage)() {
    m.messages = value
}
// Sets the tabs property value. A collection of all the tabs in the channel. A navigation property.
// Parameters:
//  - value : Value to set for the tabs property.
func (m *Channel) SetTabs(value []TeamsTab)() {
    m.tabs = value
}
// Sets the webUrl property value. A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not parsed. Read-only.
// Parameters:
//  - value : Value to set for the webUrl property.
func (m *Channel) SetWebUrl(value *string)() {
    m.webUrl = value
}
