package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AttendeeAvailability struct {
    additionalData map[string]interface{};
    attendee *AttendeeBase;
    availability *FreeBusyStatus;
}
func NewAttendeeAvailability()(*AttendeeAvailability) {
    m := &AttendeeAvailability{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AttendeeAvailability) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AttendeeAvailability) GetAttendee()(*AttendeeBase) {
    if m == nil {
        return nil
    } else {
        return m.attendee
    }
}
func (m *AttendeeAvailability) GetAvailability()(*FreeBusyStatus) {
    if m == nil {
        return nil
    } else {
        return m.availability
    }
}
func (m *AttendeeAvailability) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attendee"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttendeeBase() })
        if err != nil {
            return err
        }
        m.SetAttendee(val.(*AttendeeBase))
        return nil
    }
    res["availability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFreeBusyStatus)
        if err != nil {
            return err
        }
        cast := val.(FreeBusyStatus)
        m.SetAvailability(&cast)
        return nil
    }
    return res
}
func (m *AttendeeAvailability) IsNil()(bool) {
    return m == nil
}
func (m *AttendeeAvailability) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("attendee", m.GetAttendee())
        if err != nil {
            return err
        }
    }
    if m.GetAvailability() != nil {
        cast := m.GetAvailability().String()
        err := writer.WriteStringValue("availability", &cast)
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
func (m *AttendeeAvailability) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AttendeeAvailability) SetAttendee(value *AttendeeBase)() {
    m.attendee = value
}
func (m *AttendeeAvailability) SetAvailability(value *FreeBusyStatus)() {
    m.availability = value
}
