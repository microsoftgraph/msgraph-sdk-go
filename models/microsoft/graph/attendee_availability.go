package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttendeeAvailability 
type AttendeeAvailability struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The email address and type of attendee - whether it's a person or a resource, and whether required or optional if it's a person.
    attendee *AttendeeBase;
    // The availability status of the attendee. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
    availability *FreeBusyStatus;
}
// NewAttendeeAvailability instantiates a new attendeeAvailability and sets the default values.
func NewAttendeeAvailability()(*AttendeeAvailability) {
    m := &AttendeeAvailability{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttendeeAvailability) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAttendee gets the attendee property value. The email address and type of attendee - whether it's a person or a resource, and whether required or optional if it's a person.
func (m *AttendeeAvailability) GetAttendee()(*AttendeeBase) {
    if m == nil {
        return nil
    } else {
        return m.attendee
    }
}
// GetAvailability gets the availability property value. The availability status of the attendee. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
func (m *AttendeeAvailability) GetAvailability()(*FreeBusyStatus) {
    if m == nil {
        return nil
    } else {
        return m.availability
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttendeeAvailability) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attendee"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttendeeBase() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAttendee(val.(*AttendeeBase))
        }
        return nil
    }
    res["availability"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFreeBusyStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(FreeBusyStatus)
            m.SetAvailability(&cast)
        }
        return nil
    }
    return res
}
func (m *AttendeeAvailability) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AttendeeAvailability) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAttendee sets the attendee property value. The email address and type of attendee - whether it's a person or a resource, and whether required or optional if it's a person.
func (m *AttendeeAvailability) SetAttendee(value *AttendeeBase)() {
    if m != nil {
        m.attendee = value
    }
}
// SetAvailability sets the availability property value. The availability status of the attendee. The possible values are: free, tentative, busy, oof, workingElsewhere, unknown.
func (m *AttendeeAvailability) SetAvailability(value *FreeBusyStatus)() {
    if m != nil {
        m.availability = value
    }
}
