package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttendeeBase provides operations to manage the collection of drive entities.
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
// CreateAttendeeBaseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttendeeBaseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAttendeeBase(), nil
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
            m.SetType(val.(*AttendeeType))
        }
        return nil
    }
    return res
}
// GetType gets the type property value. The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type.
func (m *AttendeeBase) GetType()(*AttendeeType) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
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
        cast := (*m.GetType()).String()
        err = writer.WriteStringValue("type", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetType sets the type property value. The type of attendee. The possible values are: required, optional, resource. Currently if the attendee is a person, findMeetingTimes always considers the person is of the Required type.
func (m *AttendeeBase) SetType(value *AttendeeType)() {
    if m != nil {
        m.type_escaped = value
    }
}
