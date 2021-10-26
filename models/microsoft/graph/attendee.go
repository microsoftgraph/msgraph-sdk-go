package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Attendee struct {
    AttendeeBase
    // An alternate date/time proposed by the attendee for a meeting request to start and end. If the attendee hasn't proposed another time, then this property is not included in a response of a GET event.
    proposedNewTime *TimeSlot;
    // The attendee's response (none, accepted, declined, etc.) for the event and date-time that the response was sent.
    status *ResponseStatus;
}
// Instantiates a new attendee and sets the default values.
func NewAttendee()(*Attendee) {
    m := &Attendee{
        AttendeeBase: *NewAttendeeBase(),
    }
    return m
}
// Gets the proposedNewTime property value. An alternate date/time proposed by the attendee for a meeting request to start and end. If the attendee hasn't proposed another time, then this property is not included in a response of a GET event.
func (m *Attendee) GetProposedNewTime()(*TimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.proposedNewTime
    }
}
// Gets the status property value. The attendee's response (none, accepted, declined, etc.) for the event and date-time that the response was sent.
func (m *Attendee) GetStatus()(*ResponseStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *Attendee) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.AttendeeBase.GetFieldDeserializers()
    res["proposedNewTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeSlot() })
        if err != nil {
            return err
        }
        m.SetProposedNewTime(val.(*TimeSlot))
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResponseStatus() })
        if err != nil {
            return err
        }
        m.SetStatus(val.(*ResponseStatus))
        return nil
    }
    return res
}
func (m *Attendee) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *Attendee) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.AttendeeBase.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("proposedNewTime", m.GetProposedNewTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the proposedNewTime property value. An alternate date/time proposed by the attendee for a meeting request to start and end. If the attendee hasn't proposed another time, then this property is not included in a response of a GET event.
// Parameters:
//  - value : Value to set for the proposedNewTime property.
func (m *Attendee) SetProposedNewTime(value *TimeSlot)() {
    m.proposedNewTime = value
}
// Sets the status property value. The attendee's response (none, accepted, declined, etc.) for the event and date-time that the response was sent.
// Parameters:
//  - value : Value to set for the status property.
func (m *Attendee) SetStatus(value *ResponseStatus)() {
    m.status = value
}
