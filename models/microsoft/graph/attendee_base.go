package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttendeeBase 
type AttendeeBase struct {
    Recipient
    // The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type.
    type_escaped *AttendeeType;
}
// NewAttendeeBase instantiates a new attendeeBase and sets the default values.
func NewAttendeeBase()(*AttendeeBase) {
    m := &AttendeeBase{
        Recipient: *NewRecipient(),
    }
    return m
}
// GetType_escaped gets the type property value. The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type.
func (m *AttendeeBase) GetType_escaped()(*AttendeeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttendeeBase) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Recipient.GetFieldDeserializers()
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAttendeeType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AttendeeType)
            m.SetType(&cast)
        }
        return nil
    }
    return res
}
func (m *AttendeeBase) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AttendeeBase) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Recipient.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetType() != nil {
        cast := m.GetType().String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetType_escaped sets the type property value. The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type.
func (m *AttendeeBase) SetType_escaped(value *AttendeeType)() {
    if m != nil {
        m.type_escaped = value
    }
}
