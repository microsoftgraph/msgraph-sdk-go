package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttendanceRecord 
type AttendanceRecord struct {
    Entity
    // List of time periods between joining and leaving a meeting.
    attendanceIntervals []AttendanceInterval;
    // Email address of the user associated with this atttendance record.
    emailAddress *string;
    // Identity of the user associated with this atttendance record.
    identity *Identity;
    // Role of the attendee. Possible values are: None, Attendee, Presenter, and Organizer.
    role *string;
    // Total duration of the attendances in seconds.
    totalAttendanceInSeconds *int32;
}
// NewAttendanceRecord instantiates a new attendanceRecord and sets the default values.
func NewAttendanceRecord()(*AttendanceRecord) {
    m := &AttendanceRecord{
        Entity: *NewEntity(),
    }
    return m
}
// GetAttendanceIntervals gets the attendanceIntervals property value. List of time periods between joining and leaving a meeting.
func (m *AttendanceRecord) GetAttendanceIntervals()([]AttendanceInterval) {
    if m == nil {
        return nil
    } else {
        return m.attendanceIntervals
    }
}
// GetEmailAddress gets the emailAddress property value. Email address of the user associated with this atttendance record.
func (m *AttendanceRecord) GetEmailAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.emailAddress
    }
}
// GetIdentity gets the identity property value. Identity of the user associated with this atttendance record.
func (m *AttendanceRecord) GetIdentity()(*Identity) {
    if m == nil {
        return nil
    } else {
        return m.identity
    }
}
// GetRole gets the role property value. Role of the attendee. Possible values are: None, Attendee, Presenter, and Organizer.
func (m *AttendanceRecord) GetRole()(*string) {
    if m == nil {
        return nil
    } else {
        return m.role
    }
}
// GetTotalAttendanceInSeconds gets the totalAttendanceInSeconds property value. Total duration of the attendances in seconds.
func (m *AttendanceRecord) GetTotalAttendanceInSeconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalAttendanceInSeconds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AttendanceRecord) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attendanceIntervals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAttendanceInterval() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttendanceInterval, len(val))
            for i, v := range val {
                res[i] = *(v.(*AttendanceInterval))
            }
            m.SetAttendanceIntervals(res)
        }
        return nil
    }
    res["emailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEmailAddress(val)
        }
        return nil
    }
    res["identity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(*Identity))
        }
        return nil
    }
    res["role"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRole(val)
        }
        return nil
    }
    res["totalAttendanceInSeconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalAttendanceInSeconds(val)
        }
        return nil
    }
    return res
}
func (m *AttendanceRecord) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AttendanceRecord) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendanceIntervals()))
        for i, v := range m.GetAttendanceIntervals() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("attendanceIntervals", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("emailAddress", m.GetEmailAddress())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("identity", m.GetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("role", m.GetRole())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalAttendanceInSeconds", m.GetTotalAttendanceInSeconds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttendanceIntervals sets the attendanceIntervals property value. List of time periods between joining and leaving a meeting.
func (m *AttendanceRecord) SetAttendanceIntervals(value []AttendanceInterval)() {
    if m != nil {
        m.attendanceIntervals = value
    }
}
// SetEmailAddress sets the emailAddress property value. Email address of the user associated with this atttendance record.
func (m *AttendanceRecord) SetEmailAddress(value *string)() {
    if m != nil {
        m.emailAddress = value
    }
}
// SetIdentity sets the identity property value. Identity of the user associated with this atttendance record.
func (m *AttendanceRecord) SetIdentity(value *Identity)() {
    if m != nil {
        m.identity = value
    }
}
// SetRole sets the role property value. Role of the attendee. Possible values are: None, Attendee, Presenter, and Organizer.
func (m *AttendanceRecord) SetRole(value *string)() {
    if m != nil {
        m.role = value
    }
}
// SetTotalAttendanceInSeconds sets the totalAttendanceInSeconds property value. Total duration of the attendances in seconds.
func (m *AttendanceRecord) SetTotalAttendanceInSeconds(value *int32)() {
    if m != nil {
        m.totalAttendanceInSeconds = value
    }
}
