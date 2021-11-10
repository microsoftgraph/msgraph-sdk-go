package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type OnlineMeeting struct {
    Entity
    // Indicates whether attendees can turn on their camera.
    allowAttendeeToEnableCamera *bool;
    // Indicates whether attendees can turn on their microphone.
    allowAttendeeToEnableMic *bool;
    // Specifies who can be a presenter in a meeting. Possible values are listed in the following table.
    allowedPresenters *OnlineMeetingPresenters;
    // Specifies the mode of meeting chat.
    allowMeetingChat *MeetingChatMode;
    // Indicates whether Teams reactions are enabled for the meeting.
    allowTeamworkReactions *bool;
    // The content stream of the attendee report of a Microsoft Teams live event. Read-only.
    attendeeReport []byte;
    // The phone access (dial-in) information for an online meeting. Read-only.
    audioConferencing *AudioConferencing;
    // Settings related to a live event.
    broadcastSettings *BroadcastMeetingSettings;
    // The chat information associated with this online meeting.
    chatInfo *ChatInfo;
    // The meeting creation time in UTC. Read-only.
    creationDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The meeting end time in UTC.
    endDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The external ID. A custom ID. Optional.
    externalId *string;
    // Indicates if this is a Teams live event.
    isBroadcast *bool;
    // Indicates whether to announce when callers join or leave.
    isEntryExitAnnounced *bool;
    // The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
    joinInformation *ItemBody;
    // The join URL of the online meeting. Read-only.
    joinWebUrl *string;
    // Specifies which participants can bypass the meeting   lobby.
    lobbyBypassSettings *LobbyBypassSettings;
    // The participants associated with the online meeting.  This includes the organizer and the attendees.
    participants *MeetingParticipants;
    // Indicates whether to record the meeting automatically.
    recordAutomatically *bool;
    // The meeting start time in UTC.
    startDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The subject of the online meeting.
    subject *string;
    // The video teleconferencing ID. Read-only.
    videoTeleconferenceId *string;
}
// Instantiates a new onlineMeeting and sets the default values.
func NewOnlineMeeting()(*OnlineMeeting) {
    m := &OnlineMeeting{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
func (m *OnlineMeeting) GetAllowAttendeeToEnableCamera()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAttendeeToEnableCamera
    }
}
// Gets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
func (m *OnlineMeeting) GetAllowAttendeeToEnableMic()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowAttendeeToEnableMic
    }
}
// Gets the allowedPresenters property value. Specifies who can be a presenter in a meeting. Possible values are listed in the following table.
func (m *OnlineMeeting) GetAllowedPresenters()(*OnlineMeetingPresenters) {
    if m == nil {
        return nil
    } else {
        return m.allowedPresenters
    }
}
// Gets the allowMeetingChat property value. Specifies the mode of meeting chat.
func (m *OnlineMeeting) GetAllowMeetingChat()(*MeetingChatMode) {
    if m == nil {
        return nil
    } else {
        return m.allowMeetingChat
    }
}
// Gets the allowTeamworkReactions property value. Indicates whether Teams reactions are enabled for the meeting.
func (m *OnlineMeeting) GetAllowTeamworkReactions()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowTeamworkReactions
    }
}
// Gets the attendeeReport property value. The content stream of the attendee report of a Microsoft Teams live event. Read-only.
func (m *OnlineMeeting) GetAttendeeReport()([]byte) {
    if m == nil {
        return nil
    } else {
        return m.attendeeReport
    }
}
// Gets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
func (m *OnlineMeeting) GetAudioConferencing()(*AudioConferencing) {
    if m == nil {
        return nil
    } else {
        return m.audioConferencing
    }
}
// Gets the broadcastSettings property value. Settings related to a live event.
func (m *OnlineMeeting) GetBroadcastSettings()(*BroadcastMeetingSettings) {
    if m == nil {
        return nil
    } else {
        return m.broadcastSettings
    }
}
// Gets the chatInfo property value. The chat information associated with this online meeting.
func (m *OnlineMeeting) GetChatInfo()(*ChatInfo) {
    if m == nil {
        return nil
    } else {
        return m.chatInfo
    }
}
// Gets the creationDateTime property value. The meeting creation time in UTC. Read-only.
func (m *OnlineMeeting) GetCreationDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.creationDateTime
    }
}
// Gets the endDateTime property value. The meeting end time in UTC.
func (m *OnlineMeeting) GetEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.endDateTime
    }
}
// Gets the externalId property value. The external ID. A custom ID. Optional.
func (m *OnlineMeeting) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the isBroadcast property value. Indicates if this is a Teams live event.
func (m *OnlineMeeting) GetIsBroadcast()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isBroadcast
    }
}
// Gets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
func (m *OnlineMeeting) GetIsEntryExitAnnounced()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEntryExitAnnounced
    }
}
// Gets the joinInformation property value. The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
func (m *OnlineMeeting) GetJoinInformation()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.joinInformation
    }
}
// Gets the joinWebUrl property value. The join URL of the online meeting. Read-only.
func (m *OnlineMeeting) GetJoinWebUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.joinWebUrl
    }
}
// Gets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting   lobby.
func (m *OnlineMeeting) GetLobbyBypassSettings()(*LobbyBypassSettings) {
    if m == nil {
        return nil
    } else {
        return m.lobbyBypassSettings
    }
}
// Gets the participants property value. The participants associated with the online meeting.  This includes the organizer and the attendees.
func (m *OnlineMeeting) GetParticipants()(*MeetingParticipants) {
    if m == nil {
        return nil
    } else {
        return m.participants
    }
}
// Gets the recordAutomatically property value. Indicates whether to record the meeting automatically.
func (m *OnlineMeeting) GetRecordAutomatically()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.recordAutomatically
    }
}
// Gets the startDateTime property value. The meeting start time in UTC.
func (m *OnlineMeeting) GetStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// Gets the subject property value. The subject of the online meeting.
func (m *OnlineMeeting) GetSubject()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subject
    }
}
// Gets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
func (m *OnlineMeeting) GetVideoTeleconferenceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.videoTeleconferenceId
    }
}
// The deserialization information for the current model
func (m *OnlineMeeting) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["allowAttendeeToEnableCamera"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableCamera(val)
        }
        return nil
    }
    res["allowAttendeeToEnableMic"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowAttendeeToEnableMic(val)
        }
        return nil
    }
    res["allowedPresenters"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnlineMeetingPresenters)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(OnlineMeetingPresenters)
            m.SetAllowedPresenters(&cast)
        }
        return nil
    }
    res["allowMeetingChat"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseMeetingChatMode)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(MeetingChatMode)
            m.SetAllowMeetingChat(&cast)
        }
        return nil
    }
    res["allowTeamworkReactions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowTeamworkReactions(val)
        }
        return nil
    }
    res["attendeeReport"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetByteArrayValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttendeeReport(val)
        }
        return nil
    }
    res["audioConferencing"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAudioConferencing() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAudioConferencing(val.(*AudioConferencing))
        }
        return nil
    }
    res["broadcastSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewBroadcastMeetingSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBroadcastSettings(val.(*BroadcastMeetingSettings))
        }
        return nil
    }
    res["chatInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewChatInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChatInfo(val.(*ChatInfo))
        }
        return nil
    }
    res["creationDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreationDateTime(val)
        }
        return nil
    }
    res["endDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDateTime(val)
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["isBroadcast"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsBroadcast(val)
        }
        return nil
    }
    res["isEntryExitAnnounced"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEntryExitAnnounced(val)
        }
        return nil
    }
    res["joinInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinInformation(val.(*ItemBody))
        }
        return nil
    }
    res["joinWebUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJoinWebUrl(val)
        }
        return nil
    }
    res["lobbyBypassSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLobbyBypassSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLobbyBypassSettings(val.(*LobbyBypassSettings))
        }
        return nil
    }
    res["participants"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMeetingParticipants() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParticipants(val.(*MeetingParticipants))
        }
        return nil
    }
    res["recordAutomatically"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecordAutomatically(val)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val)
        }
        return nil
    }
    res["subject"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubject(val)
        }
        return nil
    }
    res["videoTeleconferenceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVideoTeleconferenceId(val)
        }
        return nil
    }
    return res
}
func (m *OnlineMeeting) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *OnlineMeeting) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("allowAttendeeToEnableCamera", m.GetAllowAttendeeToEnableCamera())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowAttendeeToEnableMic", m.GetAllowAttendeeToEnableMic())
        if err != nil {
            return err
        }
    }
    if m.GetAllowedPresenters() != nil {
        cast := m.GetAllowedPresenters().String()
        err = writer.WriteStringValue("allowedPresenters", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowMeetingChat() != nil {
        cast := m.GetAllowMeetingChat().String()
        err = writer.WriteStringValue("allowMeetingChat", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("allowTeamworkReactions", m.GetAllowTeamworkReactions())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteByteArrayValue("attendeeReport", m.GetAttendeeReport())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("audioConferencing", m.GetAudioConferencing())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("broadcastSettings", m.GetBroadcastSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("chatInfo", m.GetChatInfo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("creationDateTime", m.GetCreationDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("endDateTime", m.GetEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isBroadcast", m.GetIsBroadcast())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEntryExitAnnounced", m.GetIsEntryExitAnnounced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("joinInformation", m.GetJoinInformation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("joinWebUrl", m.GetJoinWebUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("lobbyBypassSettings", m.GetLobbyBypassSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("participants", m.GetParticipants())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("recordAutomatically", m.GetRecordAutomatically())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("subject", m.GetSubject())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("videoTeleconferenceId", m.GetVideoTeleconferenceId())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allowAttendeeToEnableCamera property value. Indicates whether attendees can turn on their camera.
// Parameters:
//  - value : Value to set for the allowAttendeeToEnableCamera property.
func (m *OnlineMeeting) SetAllowAttendeeToEnableCamera(value *bool)() {
    m.allowAttendeeToEnableCamera = value
}
// Sets the allowAttendeeToEnableMic property value. Indicates whether attendees can turn on their microphone.
// Parameters:
//  - value : Value to set for the allowAttendeeToEnableMic property.
func (m *OnlineMeeting) SetAllowAttendeeToEnableMic(value *bool)() {
    m.allowAttendeeToEnableMic = value
}
// Sets the allowedPresenters property value. Specifies who can be a presenter in a meeting. Possible values are listed in the following table.
// Parameters:
//  - value : Value to set for the allowedPresenters property.
func (m *OnlineMeeting) SetAllowedPresenters(value *OnlineMeetingPresenters)() {
    m.allowedPresenters = value
}
// Sets the allowMeetingChat property value. Specifies the mode of meeting chat.
// Parameters:
//  - value : Value to set for the allowMeetingChat property.
func (m *OnlineMeeting) SetAllowMeetingChat(value *MeetingChatMode)() {
    m.allowMeetingChat = value
}
// Sets the allowTeamworkReactions property value. Indicates whether Teams reactions are enabled for the meeting.
// Parameters:
//  - value : Value to set for the allowTeamworkReactions property.
func (m *OnlineMeeting) SetAllowTeamworkReactions(value *bool)() {
    m.allowTeamworkReactions = value
}
// Sets the attendeeReport property value. The content stream of the attendee report of a Microsoft Teams live event. Read-only.
// Parameters:
//  - value : Value to set for the attendeeReport property.
func (m *OnlineMeeting) SetAttendeeReport(value []byte)() {
    m.attendeeReport = value
}
// Sets the audioConferencing property value. The phone access (dial-in) information for an online meeting. Read-only.
// Parameters:
//  - value : Value to set for the audioConferencing property.
func (m *OnlineMeeting) SetAudioConferencing(value *AudioConferencing)() {
    m.audioConferencing = value
}
// Sets the broadcastSettings property value. Settings related to a live event.
// Parameters:
//  - value : Value to set for the broadcastSettings property.
func (m *OnlineMeeting) SetBroadcastSettings(value *BroadcastMeetingSettings)() {
    m.broadcastSettings = value
}
// Sets the chatInfo property value. The chat information associated with this online meeting.
// Parameters:
//  - value : Value to set for the chatInfo property.
func (m *OnlineMeeting) SetChatInfo(value *ChatInfo)() {
    m.chatInfo = value
}
// Sets the creationDateTime property value. The meeting creation time in UTC. Read-only.
// Parameters:
//  - value : Value to set for the creationDateTime property.
func (m *OnlineMeeting) SetCreationDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.creationDateTime = value
}
// Sets the endDateTime property value. The meeting end time in UTC.
// Parameters:
//  - value : Value to set for the endDateTime property.
func (m *OnlineMeeting) SetEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.endDateTime = value
}
// Sets the externalId property value. The external ID. A custom ID. Optional.
// Parameters:
//  - value : Value to set for the externalId property.
func (m *OnlineMeeting) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the isBroadcast property value. Indicates if this is a Teams live event.
// Parameters:
//  - value : Value to set for the isBroadcast property.
func (m *OnlineMeeting) SetIsBroadcast(value *bool)() {
    m.isBroadcast = value
}
// Sets the isEntryExitAnnounced property value. Indicates whether to announce when callers join or leave.
// Parameters:
//  - value : Value to set for the isEntryExitAnnounced property.
func (m *OnlineMeeting) SetIsEntryExitAnnounced(value *bool)() {
    m.isEntryExitAnnounced = value
}
// Sets the joinInformation property value. The join information in the language and locale variant specified in the Accept-Language request HTTP header. Read-only.
// Parameters:
//  - value : Value to set for the joinInformation property.
func (m *OnlineMeeting) SetJoinInformation(value *ItemBody)() {
    m.joinInformation = value
}
// Sets the joinWebUrl property value. The join URL of the online meeting. Read-only.
// Parameters:
//  - value : Value to set for the joinWebUrl property.
func (m *OnlineMeeting) SetJoinWebUrl(value *string)() {
    m.joinWebUrl = value
}
// Sets the lobbyBypassSettings property value. Specifies which participants can bypass the meeting   lobby.
// Parameters:
//  - value : Value to set for the lobbyBypassSettings property.
func (m *OnlineMeeting) SetLobbyBypassSettings(value *LobbyBypassSettings)() {
    m.lobbyBypassSettings = value
}
// Sets the participants property value. The participants associated with the online meeting.  This includes the organizer and the attendees.
// Parameters:
//  - value : Value to set for the participants property.
func (m *OnlineMeeting) SetParticipants(value *MeetingParticipants)() {
    m.participants = value
}
// Sets the recordAutomatically property value. Indicates whether to record the meeting automatically.
// Parameters:
//  - value : Value to set for the recordAutomatically property.
func (m *OnlineMeeting) SetRecordAutomatically(value *bool)() {
    m.recordAutomatically = value
}
// Sets the startDateTime property value. The meeting start time in UTC.
// Parameters:
//  - value : Value to set for the startDateTime property.
func (m *OnlineMeeting) SetStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.startDateTime = value
}
// Sets the subject property value. The subject of the online meeting.
// Parameters:
//  - value : Value to set for the subject property.
func (m *OnlineMeeting) SetSubject(value *string)() {
    m.subject = value
}
// Sets the videoTeleconferenceId property value. The video teleconferencing ID. Read-only.
// Parameters:
//  - value : Value to set for the videoTeleconferenceId property.
func (m *OnlineMeeting) SetVideoTeleconferenceId(value *string)() {
    m.videoTeleconferenceId = value
}
