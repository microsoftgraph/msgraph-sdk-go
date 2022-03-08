package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// MeetingAttendanceReport provides operations to manage the cloudCommunications singleton.
type MeetingAttendanceReport struct {
    Entity
    // List of attendance records of an attendance report. Read-only.
    attendanceRecords []AttendanceRecordable;
    // UTC time when the meeting ended. Read-only.
    meetingEndDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // UTC time when the meeting started. Read-only.
    meetingStartDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Total number of participants. Read-only.
    totalParticipantCount *int32;
}
// NewMeetingAttendanceReport instantiates a new meetingAttendanceReport and sets the default values.
func NewMeetingAttendanceReport()(*MeetingAttendanceReport) {
    m := &MeetingAttendanceReport{
        Entity: *NewEntity(),
    }
    return m
}
// CreateMeetingAttendanceReportFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateMeetingAttendanceReportFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewMeetingAttendanceReport(), nil
}
// GetAttendanceRecords gets the attendanceRecords property value. List of attendance records of an attendance report. Read-only.
func (m *MeetingAttendanceReport) GetAttendanceRecords()([]AttendanceRecordable) {
    if m == nil {
        return nil
    } else {
        return m.attendanceRecords
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *MeetingAttendanceReport) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["attendanceRecords"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAttendanceRecordFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AttendanceRecordable, len(val))
            for i, v := range val {
                res[i] = v.(AttendanceRecordable)
            }
            m.SetAttendanceRecords(res)
        }
        return nil
    }
    res["meetingEndDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingEndDateTime(val)
        }
        return nil
    }
    res["meetingStartDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMeetingStartDateTime(val)
        }
        return nil
    }
    res["totalParticipantCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTotalParticipantCount(val)
        }
        return nil
    }
    return res
}
// GetMeetingEndDateTime gets the meetingEndDateTime property value. UTC time when the meeting ended. Read-only.
func (m *MeetingAttendanceReport) GetMeetingEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.meetingEndDateTime
    }
}
// GetMeetingStartDateTime gets the meetingStartDateTime property value. UTC time when the meeting started. Read-only.
func (m *MeetingAttendanceReport) GetMeetingStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.meetingStartDateTime
    }
}
// GetTotalParticipantCount gets the totalParticipantCount property value. Total number of participants. Read-only.
func (m *MeetingAttendanceReport) GetTotalParticipantCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.totalParticipantCount
    }
}
func (m *MeetingAttendanceReport) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *MeetingAttendanceReport) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAttendanceRecords() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAttendanceRecords()))
        for i, v := range m.GetAttendanceRecords() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("attendanceRecords", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("meetingEndDateTime", m.GetMeetingEndDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("meetingStartDateTime", m.GetMeetingStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("totalParticipantCount", m.GetTotalParticipantCount())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAttendanceRecords sets the attendanceRecords property value. List of attendance records of an attendance report. Read-only.
func (m *MeetingAttendanceReport) SetAttendanceRecords(value []AttendanceRecordable)() {
    if m != nil {
        m.attendanceRecords = value
    }
}
// SetMeetingEndDateTime sets the meetingEndDateTime property value. UTC time when the meeting ended. Read-only.
func (m *MeetingAttendanceReport) SetMeetingEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.meetingEndDateTime = value
    }
}
// SetMeetingStartDateTime sets the meetingStartDateTime property value. UTC time when the meeting started. Read-only.
func (m *MeetingAttendanceReport) SetMeetingStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.meetingStartDateTime = value
    }
}
// SetTotalParticipantCount sets the totalParticipantCount property value. Total number of participants. Read-only.
func (m *MeetingAttendanceReport) SetTotalParticipantCount(value *int32)() {
    if m != nil {
        m.totalParticipantCount = value
    }
}
