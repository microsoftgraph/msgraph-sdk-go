package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SettingStateDeviceSummaryable 
type SettingStateDeviceSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetCompliantDeviceCount()(*int32)
    GetConflictDeviceCount()(*int32)
    GetErrorDeviceCount()(*int32)
    GetInstancePath()(*string)
    GetNonCompliantDeviceCount()(*int32)
    GetNotApplicableDeviceCount()(*int32)
    GetRemediatedDeviceCount()(*int32)
    GetSettingName()(*string)
    GetUnknownDeviceCount()(*int32)
    SetCompliantDeviceCount(value *int32)()
    SetConflictDeviceCount(value *int32)()
    SetErrorDeviceCount(value *int32)()
    SetInstancePath(value *string)()
    SetNonCompliantDeviceCount(value *int32)()
    SetNotApplicableDeviceCount(value *int32)()
    SetRemediatedDeviceCount(value *int32)()
    SetSettingName(value *string)()
    SetUnknownDeviceCount(value *int32)()
}
