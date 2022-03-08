package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Channelable 
type Channelable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetEmail()(*string)
    GetFilesFolder()(DriveItemable)
    GetIsFavoriteByDefault()(*bool)
    GetMembers()([]ConversationMemberable)
    GetMembershipType()(*ChannelMembershipType)
    GetMessages()([]ChatMessageable)
    GetTabs()([]TeamsTabable)
    GetWebUrl()(*string)
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetEmail(value *string)()
    SetFilesFolder(value DriveItemable)()
    SetIsFavoriteByDefault(value *bool)()
    SetMembers(value []ConversationMemberable)()
    SetMembershipType(value *ChannelMembershipType)()
    SetMessages(value []ChatMessageable)()
    SetTabs(value []TeamsTabable)()
    SetWebUrl(value *string)()
}
