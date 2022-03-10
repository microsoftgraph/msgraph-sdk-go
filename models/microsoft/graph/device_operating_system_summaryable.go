package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceOperatingSystemSummaryable 
type DeviceOperatingSystemSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAndroidCount()(*int32)
    GetIosCount()(*int32)
    GetMacOSCount()(*int32)
    GetUnknownCount()(*int32)
    GetWindowsCount()(*int32)
    GetWindowsMobileCount()(*int32)
    SetAndroidCount(value *int32)()
    SetIosCount(value *int32)()
    SetMacOSCount(value *int32)()
    SetUnknownCount(value *int32)()
    SetWindowsCount(value *int32)()
    SetWindowsMobileCount(value *int32)()
}
