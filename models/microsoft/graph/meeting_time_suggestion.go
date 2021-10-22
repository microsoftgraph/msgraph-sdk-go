package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type MeetingTimeSuggestion struct {
    additionalData map[string]interface{};
    attendeeAvailability []AttendeeAvailability;
    confidence *float64;
    locations []Location;
    meetingTimeSlot *TimeSlot;
    order *int32;
    organizerAvailability *FreeBusyStatus;
    suggestionReason *string;
}
func NewMeetingTimeSuggestion()(*MeetingTimeSuggestion) {
    m := &MeetingTimeSuggestion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *MeetingTimeSuggestion) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *MeetingTimeSuggestion) GetAttendeeAvailability()([]AttendeeAvailability) {
    if m == nil {
        return nil
    } else {
        return m.attendeeAvailability
    }
}
func (m *MeetingTimeSuggestion) GetConfidence()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
func (m *MeetingTimeSuggestion) GetLocations()([]Location) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
func (m *MeetingTimeSuggestion) GetMeetingTimeSlot()(*TimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.meetingTimeSlot
    }
}
func (m *MeetingTimeSuggestion) GetOrder()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
func (m *MeetingTimeSuggestion) GetOrganizerAvailability()(*FreeBusyStatus) {
    if m == nil {
        return nil
    } else {
        return m.organizerAvailability
    }
}
func (m *MeetingTimeSuggestion) GetSuggestionReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.suggestionReason
    }
}
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
func (m *MeetingTimeSuggestion) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *MeetingTimeSuggestion) SetAttendeeAvailability(value []AttendeeAvailability)() {
    m.attendeeAvailability = value
}
func (m *MeetingTimeSuggestion) SetConfidence(value *float64)() {
    m.confidence = value
}
func (m *MeetingTimeSuggestion) SetLocations(value []Location)() {
    m.locations = value
}
func (m *MeetingTimeSuggestion) SetMeetingTimeSlot(value *TimeSlot)() {
    m.meetingTimeSlot = value
}
func (m *MeetingTimeSuggestion) SetOrder(value *int32)() {
    m.order = value
}
func (m *MeetingTimeSuggestion) SetOrganizerAvailability(value *FreeBusyStatus)() {
    m.organizerAvailability = value
}
func (m *MeetingTimeSuggestion) SetSuggestionReason(value *string)() {
    m.suggestionReason = value
}
