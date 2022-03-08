package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceCompliancePolicySettingStateSummaryable 
type DeviceCompliancePolicySettingStateSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCompliantDeviceCount()(*int32)
    GetConflictDeviceCount()(*int32)
    GetDeviceComplianceSettingStates()([]DeviceComplianceSettingStateable)
    GetErrorDeviceCount()(*int32)
    GetNonCompliantDeviceCount()(*int32)
    GetNotApplicableDeviceCount()(*int32)
    GetPlatformType()(*PolicyPlatformType)
    GetRemediatedDeviceCount()(*int32)
    GetSetting()(*string)
    GetSettingName()(*string)
    GetUnknownDeviceCount()(*int32)
    SetCompliantDeviceCount(value *int32)()
    SetConflictDeviceCount(value *int32)()
    SetDeviceComplianceSettingStates(value []DeviceComplianceSettingStateable)()
    SetErrorDeviceCount(value *int32)()
    SetNonCompliantDeviceCount(value *int32)()
    SetNotApplicableDeviceCount(value *int32)()
    SetPlatformType(value *PolicyPlatformType)()
    SetRemediatedDeviceCount(value *int32)()
    SetSetting(value *string)()
    SetSettingName(value *string)()
    SetUnknownDeviceCount(value *int32)()
}
