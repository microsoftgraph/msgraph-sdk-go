package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Chatable 
type Chatable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetChatType()(*ChatType)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetInstalledApps()([]TeamsAppInstallationable)
    GetLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMembers()([]ConversationMemberable)
    GetMessages()([]ChatMessageable)
    GetOnlineMeetingInfo()(TeamworkOnlineMeetingInfoable)
    GetTabs()([]TeamsTabable)
    GetTenantId()(*string)
    GetTopic()(*string)
    GetWebUrl()(*string)
    SetChatType(value *ChatType)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetInstalledApps(value []TeamsAppInstallationable)()
    SetLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMembers(value []ConversationMemberable)()
    SetMessages(value []ChatMessageable)()
    SetOnlineMeetingInfo(value TeamworkOnlineMeetingInfoable)()
    SetTabs(value []TeamsTabable)()
    SetTenantId(value *string)()
    SetTopic(value *string)()
    SetWebUrl(value *string)()
}
