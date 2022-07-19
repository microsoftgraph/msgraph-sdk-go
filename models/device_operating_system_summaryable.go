package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceOperatingSystemSummaryable 
type DeviceOperatingSystemSummaryable interface {
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.AdditionalDataHolder
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAndroidCount()(*int32)
    GetIosCount()(*int32)
    GetMacOSCount()(*int32)
    GetOdataType()(*string)
    GetUnknownCount()(*int32)
    GetWindowsCount()(*int32)
    GetWindowsMobileCount()(*int32)
    SetAndroidCount(value *int32)()
    SetIosCount(value *int32)()
    SetMacOSCount(value *int32)()
    SetOdataType(value *string)()
    SetUnknownCount(value *int32)()
    SetWindowsCount(value *int32)()
    SetWindowsMobileCount(value *int32)()
}
