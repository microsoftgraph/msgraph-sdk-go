package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BroadcastMeetingSettingsable 
type BroadcastMeetingSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAllowedAudience()(*BroadcastMeetingAudience)
    GetIsAttendeeReportEnabled()(*bool)
    GetIsQuestionAndAnswerEnabled()(*bool)
    GetIsRecordingEnabled()(*bool)
    GetIsVideoOnDemandEnabled()(*bool)
    SetAllowedAudience(value *BroadcastMeetingAudience)()
    SetIsAttendeeReportEnabled(value *bool)()
    SetIsQuestionAndAnswerEnabled(value *bool)()
    SetIsRecordingEnabled(value *bool)()
    SetIsVideoOnDemandEnabled(value *bool)()
}
