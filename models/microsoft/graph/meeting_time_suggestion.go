package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type MeetingTimeSuggestion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // An array that shows the availability status of each attendee for this meeting suggestion.
    attendeeAvailability []AttendeeAvailability;
    // A percentage that represents the likelhood of all the attendees attending.
    confidence *float64;
    // An array that specifies the name and geographic location of each meeting location for this meeting suggestion.
    locations []Location;
    // A time period suggested for the meeting.
    meetingTimeSlot *TimeSlot;
    // Order of meeting time suggestions sorted by their computed confidence value from high to low, then by chronology if there are suggestions with the same confidence.
    order *int32;
    // Availability of the meeting organizer for this meeting suggestion. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
    organizerAvailability *FreeBusyStatus;
    // Reason for suggesting the meeting time.
    suggestionReason *string;
}
// Instantiates a new meetingTimeSuggestion and sets the default values.
func NewMeetingTimeSuggestion()(*MeetingTimeSuggestion) {
    m := &MeetingTimeSuggestion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingTimeSuggestion) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the attendeeAvailability property value. An array that shows the availability status of each attendee for this meeting suggestion.
func (m *MeetingTimeSuggestion) GetAttendeeAvailability()([]AttendeeAvailability) {
    if m == nil {
        return nil
    } else {
        return m.attendeeAvailability
    }
}
// Gets the confidence property value. A percentage that represents the likelhood of all the attendees attending.
func (m *MeetingTimeSuggestion) GetConfidence()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// Gets the locations property value. An array that specifies the name and geographic location of each meeting location for this meeting suggestion.
func (m *MeetingTimeSuggestion) GetLocations()([]Location) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
// Gets the meetingTimeSlot property value. A time period suggested for the meeting.
func (m *MeetingTimeSuggestion) GetMeetingTimeSlot()(*TimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.meetingTimeSlot
    }
}
// Gets the order property value. Order of meeting time suggestions sorted by their computed confidence value from high to low, then by chronology if there are suggestions with the same confidence.
func (m *MeetingTimeSuggestion) GetOrder()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
// Gets the organizerAvailability property value. Availability of the meeting organizer for this meeting suggestion. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
func (m *MeetingTimeSuggestion) GetOrganizerAvailability()(*FreeBusyStatus) {
    if m == nil {
        return nil
    } else {
        return m.organizerAvailability
    }
}
// Gets the suggestionReason property value. Reason for suggesting the meeting time.
func (m *MeetingTimeSuggestion) GetSuggestionReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.suggestionReason
    }
}
// The deserialization information for the current model
func (m *MeetingTimeSuggestion) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attendeeAvailability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttendeeAvailability() })
        if err != nil {
            return err
        }
        res := make([]AttendeeAvailability, len(val))
        for i, v := range val {
            res[i] = *(v.(*AttendeeAvailability))
        }
        m.SetAttendeeAvailability(res)
        return nil
    }
    res["confidence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetConfidence(val)
        return nil
    }
    res["locations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLocation() })
        if err != nil {
            return err
        }
        res := make([]Location, len(val))
        for i, v := range val {
            res[i] = *(v.(*Location))
        }
        m.SetLocations(res)
        return nil
    }
    res["meetingTimeSlot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeSlot() })
        if err != nil {
            return err
        }
        m.SetMeetingTimeSlot(val.(*TimeSlot))
        return nil
    }
    res["order"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetOrder(val)
        return nil
    }
    res["organizerAvailability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFreeBusyStatus)
        if err != nil {
            return err
        }
        cast := val.(FreeBusyStatus)
        m.SetOrganizerAvailability(&cast)
        return nil
    }
    res["suggestionReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSuggestionReason(val)
        return nil
    }
    return res
}
func (m *MeetingTimeSuggestion) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *MeetingTimeSuggestion) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendeeAvailability()))
        for i, v := range m.GetAttendeeAvailability() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("attendeeAvailability", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat64Value("confidence", m.GetConfidence())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLocations()))
        for i, v := range m.GetLocations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err := writer.WriteCollectionOfObjectValues("locations", cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("meetingTimeSlot", m.GetMeetingTimeSlot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("order", m.GetOrder())
        if err != nil {
            return err
        }
    }
    if m.GetOrganizerAvailability() != nil {
        cast := m.GetOrganizerAvailability().String()
        err := writer.WriteStringValue("organizerAvailability", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("suggestionReason", m.GetSuggestionReason())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *MeetingTimeSuggestion) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the attendeeAvailability property value. An array that shows the availability status of each attendee for this meeting suggestion.
// Parameters:
//  - value : Value to set for the attendeeAvailability property.
func (m *MeetingTimeSuggestion) SetAttendeeAvailability(value []AttendeeAvailability)() {
    m.attendeeAvailability = value
}
// Sets the confidence property value. A percentage that represents the likelhood of all the attendees attending.
// Parameters:
//  - value : Value to set for the confidence property.
func (m *MeetingTimeSuggestion) SetConfidence(value *float64)() {
    m.confidence = value
}
// Sets the locations property value. An array that specifies the name and geographic location of each meeting location for this meeting suggestion.
// Parameters:
//  - value : Value to set for the locations property.
func (m *MeetingTimeSuggestion) SetLocations(value []Location)() {
    m.locations = value
}
// Sets the meetingTimeSlot property value. A time period suggested for the meeting.
// Parameters:
//  - value : Value to set for the meetingTimeSlot property.
func (m *MeetingTimeSuggestion) SetMeetingTimeSlot(value *TimeSlot)() {
    m.meetingTimeSlot = value
}
// Sets the order property value. Order of meeting time suggestions sorted by their computed confidence value from high to low, then by chronology if there are suggestions with the same confidence.
// Parameters:
//  - value : Value to set for the order property.
func (m *MeetingTimeSuggestion) SetOrder(value *int32)() {
    m.order = value
}
// Sets the organizerAvailability property value. Availability of the meeting organizer for this meeting suggestion. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
// Parameters:
//  - value : Value to set for the organizerAvailability property.
func (m *MeetingTimeSuggestion) SetOrganizerAvailability(value *FreeBusyStatus)() {
    m.organizerAvailability = value
}
// Sets the suggestionReason property value. Reason for suggesting the meeting time.
// Parameters:
//  - value : Value to set for the suggestionReason property.
func (m *MeetingTimeSuggestion) SetSuggestionReason(value *string)() {
    m.suggestionReason = value
}
