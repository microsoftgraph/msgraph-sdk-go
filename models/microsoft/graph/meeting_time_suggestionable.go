package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingTimeSuggestionable 
type MeetingTimeSuggestionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAttendeeAvailability()([]AttendeeAvailabilityable)
    GetConfidence()(*float64)
    GetLocations()([]Locationable)
    GetMeetingTimeSlot()(TimeSlotable)
    GetOrder()(*int32)
    GetOrganizerAvailability()(*FreeBusyStatus)
    GetSuggestionReason()(*string)
    SetAttendeeAvailability(value []AttendeeAvailabilityable)()
    SetConfidence(value *float64)()
    SetLocations(value []Locationable)()
    SetMeetingTimeSlot(value TimeSlotable)()
    SetOrder(value *int32)()
    SetOrganizerAvailability(value *FreeBusyStatus)()
    SetSuggestionReason(value *string)()
}
