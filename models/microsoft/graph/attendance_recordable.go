package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AttendanceRecordable 
type AttendanceRecordable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAttendanceIntervals()([]AttendanceIntervalable)
    GetEmailAddress()(*string)
    GetIdentity()(Identityable)
    GetRole()(*string)
    GetTotalAttendanceInSeconds()(*int32)
    SetAttendanceIntervals(value []AttendanceIntervalable)()
    SetEmailAddress(value *string)()
    SetIdentity(value Identityable)()
    SetRole(value *string)()
    SetTotalAttendanceInSeconds(value *int32)()
}
