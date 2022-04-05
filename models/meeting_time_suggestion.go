package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// MeetingTimeSuggestion 
type MeetingTimeSuggestion struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // An array that shows the availability status of each attendee for this meeting suggestion.
    attendeeAvailability []AttendeeAvailabilityable;
    // A percentage that represents the likelhood of all the attendees attending.
    confidence *float64;
    // An array that specifies the name and geographic location of each meeting location for this meeting suggestion.
    locations []Locationable;
    // A time period suggested for the meeting.
    meetingTimeSlot TimeSlotable;
    // Order of meeting time suggestions sorted by their computed confidence value from high to low, then by chronology if there are suggestions with the same confidence.
    order *int32;
    // Availability of the meeting organizer for this meeting suggestion. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
    organizerAvailability *FreeBusyStatus;
    // Reason for suggesting the meeting time.
    suggestionReason *string;
}
// NewMeetingTimeSuggestion instantiates a new meetingTimeSuggestion and sets the default values.
func NewMeetingTimeSuggestion()(*MeetingTimeSuggestion) {
    m := &MeetingTimeSuggestion{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateMeetingTimeSuggestionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingTimeSuggestionFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewMeetingTimeSuggestion(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingTimeSuggestion) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttendeeAvailability gets the attendeeAvailability property value. An array that shows the availability status of each attendee for this meeting suggestion.
func (m *MeetingTimeSuggestion) GetAttendeeAvailability()([]AttendeeAvailabilityable) {
    if m == nil {
        return nil
    } else {
        return m.attendeeAvailability
    }
}
// GetConfidence gets the confidence property value. A percentage that represents the likelhood of all the attendees attending.
func (m *MeetingTimeSuggestion) GetConfidence()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.confidence
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingTimeSuggestion) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["attendeeAvailability"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttendeeAvailabilityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttendeeAvailabilityable, len(val))
            for i, v := range val {
                res[i] = v.(AttendeeAvailabilityable)
            }
            m.SetAttendeeAvailability(res)
        }
        return nil
    }
    res["confidence"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfidence(val)
        }
        return nil
    }
    res["locations"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLocationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Locationable, len(val))
            for i, v := range val {
                res[i] = v.(Locationable)
            }
            m.SetLocations(res)
        }
        return nil
    }
    res["meetingTimeSlot"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateTimeSlotFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingTimeSlot(val.(TimeSlotable))
        }
        return nil
    }
    res["order"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrder(val)
        }
        return nil
    }
    res["organizerAvailability"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFreeBusyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizerAvailability(val.(*FreeBusyStatus))
        }
        return nil
    }
    res["suggestionReason"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSuggestionReason(val)
        }
        return nil
    }
    return res
}
// GetLocations gets the locations property value. An array that specifies the name and geographic location of each meeting location for this meeting suggestion.
func (m *MeetingTimeSuggestion) GetLocations()([]Locationable) {
    if m == nil {
        return nil
    } else {
        return m.locations
    }
}
// GetMeetingTimeSlot gets the meetingTimeSlot property value. A time period suggested for the meeting.
func (m *MeetingTimeSuggestion) GetMeetingTimeSlot()(TimeSlotable) {
    if m == nil {
        return nil
    } else {
        return m.meetingTimeSlot
    }
}
// GetOrder gets the order property value. Order of meeting time suggestions sorted by their computed confidence value from high to low, then by chronology if there are suggestions with the same confidence.
func (m *MeetingTimeSuggestion) GetOrder()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.order
    }
}
// GetOrganizerAvailability gets the organizerAvailability property value. Availability of the meeting organizer for this meeting suggestion. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
func (m *MeetingTimeSuggestion) GetOrganizerAvailability()(*FreeBusyStatus) {
    if m == nil {
        return nil
    } else {
        return m.organizerAvailability
    }
}
// GetSuggestionReason gets the suggestionReason property value. Reason for suggesting the meeting time.
func (m *MeetingTimeSuggestion) GetSuggestionReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.suggestionReason
    }
}
// Serialize serializes information the current object
func (m *MeetingTimeSuggestion) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAttendeeAvailability() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAttendeeAvailability()))
        for i, v := range m.GetAttendeeAvailability() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
    if m.GetLocations() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLocations()))
        for i, v := range m.GetLocations() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := (*m.GetOrganizerAvailability()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *MeetingTimeSuggestion) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttendeeAvailability sets the attendeeAvailability property value. An array that shows the availability status of each attendee for this meeting suggestion.
func (m *MeetingTimeSuggestion) SetAttendeeAvailability(value []AttendeeAvailabilityable)() {
    if m != nil {
        m.attendeeAvailability = value
    }
}
// SetConfidence sets the confidence property value. A percentage that represents the likelhood of all the attendees attending.
func (m *MeetingTimeSuggestion) SetConfidence(value *float64)() {
    if m != nil {
        m.confidence = value
    }
}
// SetLocations sets the locations property value. An array that specifies the name and geographic location of each meeting location for this meeting suggestion.
func (m *MeetingTimeSuggestion) SetLocations(value []Locationable)() {
    if m != nil {
        m.locations = value
    }
}
// SetMeetingTimeSlot sets the meetingTimeSlot property value. A time period suggested for the meeting.
func (m *MeetingTimeSuggestion) SetMeetingTimeSlot(value TimeSlotable)() {
    if m != nil {
        m.meetingTimeSlot = value
    }
}
// SetOrder sets the order property value. Order of meeting time suggestions sorted by their computed confidence value from high to low, then by chronology if there are suggestions with the same confidence.
func (m *MeetingTimeSuggestion) SetOrder(value *int32)() {
    if m != nil {
        m.order = value
    }
}
// SetOrganizerAvailability sets the organizerAvailability property value. Availability of the meeting organizer for this meeting suggestion. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
func (m *MeetingTimeSuggestion) SetOrganizerAvailability(value *FreeBusyStatus)() {
    if m != nil {
        m.organizerAvailability = value
    }
}
// SetSuggestionReason sets the suggestionReason property value. Reason for suggesting the meeting time.
func (m *MeetingTimeSuggestion) SetSuggestionReason(value *string)() {
    if m != nil {
        m.suggestionReason = value
    }
}
