package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OnlineMeetingable 
type OnlineMeetingable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowAttendeeToEnableCamera()(*bool)
    GetAllowAttendeeToEnableMic()(*bool)
    GetAllowedPresenters()(*OnlineMeetingPresenters)
    GetAllowMeetingChat()(*MeetingChatMode)
    GetAllowTeamworkReactions()(*bool)
    GetAttendanceReports()([]MeetingAttendanceReportable)
    GetAttendeeReport()([]byte)
    GetAudioConferencing()(AudioConferencingable)
    GetBroadcastSettings()(BroadcastMeetingSettingsable)
    GetChatInfo()(ChatInfoable)
    GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetExternalId()(*string)
    GetIsBroadcast()(*bool)
    GetIsEntryExitAnnounced()(*bool)
    GetJoinInformation()(ItemBodyable)
    GetJoinWebUrl()(*string)
    GetLobbyBypassSettings()(LobbyBypassSettingsable)
    GetParticipants()(MeetingParticipantsable)
    GetRecordAutomatically()(*bool)
    GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetSubject()(*string)
    GetVideoTeleconferenceId()(*string)
    SetAllowAttendeeToEnableCamera(value *bool)()
    SetAllowAttendeeToEnableMic(value *bool)()
    SetAllowedPresenters(value *OnlineMeetingPresenters)()
    SetAllowMeetingChat(value *MeetingChatMode)()
    SetAllowTeamworkReactions(value *bool)()
    SetAttendanceReports(value []MeetingAttendanceReportable)()
    SetAttendeeReport(value []byte)()
    SetAudioConferencing(value AudioConferencingable)()
    SetBroadcastSettings(value BroadcastMeetingSettingsable)()
    SetChatInfo(value ChatInfoable)()
    SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetExternalId(value *string)()
    SetIsBroadcast(value *bool)()
    SetIsEntryExitAnnounced(value *bool)()
    SetJoinInformation(value ItemBodyable)()
    SetJoinWebUrl(value *string)()
    SetLobbyBypassSettings(value LobbyBypassSettingsable)()
    SetParticipants(value MeetingParticipantsable)()
    SetRecordAutomatically(value *bool)()
    SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetSubject(value *string)()
    SetVideoTeleconferenceId(value *string)()
}
