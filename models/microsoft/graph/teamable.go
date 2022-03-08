package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Teamable 
type Teamable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetChannels()([]Channelable)
    GetClassification()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetFunSettings()(TeamFunSettingsable)
    GetGroup()(Groupable)
    GetGuestSettings()(TeamGuestSettingsable)
    GetInstalledApps()([]TeamsAppInstallationable)
    GetInternalId()(*string)
    GetIsArchived()(*bool)
    GetMembers()([]ConversationMemberable)
    GetMemberSettings()(TeamMemberSettingsable)
    GetMessagingSettings()(TeamMessagingSettingsable)
    GetOperations()([]TeamsAsyncOperationable)
    GetPrimaryChannel()(Channelable)
    GetSchedule()(Scheduleable)
    GetSpecialization()(*TeamSpecialization)
    GetTemplate()(TeamsTemplateable)
    GetVisibility()(*TeamVisibilityType)
    GetWebUrl()(*string)
    SetChannels(value []Channelable)()
    SetClassification(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetFunSettings(value TeamFunSettingsable)()
    SetGroup(value Groupable)()
    SetGuestSettings(value TeamGuestSettingsable)()
    SetInstalledApps(value []TeamsAppInstallationable)()
    SetInternalId(value *string)()
    SetIsArchived(value *bool)()
    SetMembers(value []ConversationMemberable)()
    SetMemberSettings(value TeamMemberSettingsable)()
    SetMessagingSettings(value TeamMessagingSettingsable)()
    SetOperations(value []TeamsAsyncOperationable)()
    SetPrimaryChannel(value Channelable)()
    SetSchedule(value Scheduleable)()
    SetSpecialization(value *TeamSpecialization)()
    SetTemplate(value TeamsTemplateable)()
    SetVisibility(value *TeamVisibilityType)()
    SetWebUrl(value *string)()
}
