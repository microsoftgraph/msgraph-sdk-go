package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Attendee struct {
    AttendeeBase
    proposedNewTime *TimeSlot;
    status *ResponseStatus;
}
func NewAttendee()(*Attendee) {
    m := &Attendee{
        AttendeeBase: *NewAttendeeBase(),
    }
    return m
}
func (m *Attendee) GetProposedNewTime()(*TimeSlot) {
    if m == nil {
        return nil
    } else {
        return m.proposedNewTime
    }
}
func (m *Attendee) GetStatus()(*ResponseStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
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
func (m *Attendee) SetProposedNewTime(value *TimeSlot)() {
    m.proposedNewTime = value
}
func (m *Attendee) SetStatus(value *ResponseStatus)() {
    m.status = value
}
