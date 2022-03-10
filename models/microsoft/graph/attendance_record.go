package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttendanceRecord provides operations to manage the cloudCommunications singleton.
type AttendanceRecord struct {
    Entity
    // List of time periods between joining and leaving a meeting.
    attendanceIntervals []AttendanceIntervalable;
    // Email address of the user associated with this atttendance record.
    emailAddress *string;
    // Identity of the user associated with this atttendance record.
    identity Identityable;
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
// CreateAttendanceRecordFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAttendanceRecordFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewAttendanceRecord(), nil
}
// GetAttendanceIntervals gets the attendanceIntervals property value. List of time periods between joining and leaving a meeting.
func (m *AttendanceRecord) GetAttendanceIntervals()([]AttendanceIntervalable) {
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
// GetFieldDeserializers the deserialization information for the current model
func (m *AttendanceRecord) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attendanceIntervals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttendanceIntervalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttendanceIntervalable, len(val))
            for i, v := range val {
                res[i] = v.(AttendanceIntervalable)
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
        val, err := n.GetObjectValue(CreateIdentityFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIdentity(val.(Identityable))
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
// GetIdentity gets the identity property value. Identity of the user associated with this atttendance record.
func (m *AttendanceRecord) GetIdentity()(Identityable) {
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
func (m *AttendanceRecord) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AttendanceRecord) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttendanceIntervals() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendanceIntervals()))
        for i, v := range m.GetAttendanceIntervals() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
func (m *AttendanceRecord) SetAttendanceIntervals(value []AttendanceIntervalable)() {
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
func (m *AttendanceRecord) SetIdentity(value Identityable)() {
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
