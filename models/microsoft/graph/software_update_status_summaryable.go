package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SoftwareUpdateStatusSummaryable 
type SoftwareUpdateStatusSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetCompliantDeviceCount()(*int32)
    GetCompliantUserCount()(*int32)
    GetConflictDeviceCount()(*int32)
    GetConflictUserCount()(*int32)
    GetDisplayName()(*string)
    GetErrorDeviceCount()(*int32)
    GetErrorUserCount()(*int32)
    GetNonCompliantDeviceCount()(*int32)
    GetNonCompliantUserCount()(*int32)
    GetNotApplicableDeviceCount()(*int32)
    GetNotApplicableUserCount()(*int32)
    GetRemediatedDeviceCount()(*int32)
    GetRemediatedUserCount()(*int32)
    GetUnknownDeviceCount()(*int32)
    GetUnknownUserCount()(*int32)
    SetCompliantDeviceCount(value *int32)()
    SetCompliantUserCount(value *int32)()
    SetConflictDeviceCount(value *int32)()
    SetConflictUserCount(value *int32)()
    SetDisplayName(value *string)()
    SetErrorDeviceCount(value *int32)()
    SetErrorUserCount(value *int32)()
    SetNonCompliantDeviceCount(value *int32)()
    SetNonCompliantUserCount(value *int32)()
    SetNotApplicableDeviceCount(value *int32)()
    SetNotApplicableUserCount(value *int32)()
    SetRemediatedDeviceCount(value *int32)()
    SetRemediatedUserCount(value *int32)()
    SetUnknownDeviceCount(value *int32)()
    SetUnknownUserCount(value *int32)()
}
