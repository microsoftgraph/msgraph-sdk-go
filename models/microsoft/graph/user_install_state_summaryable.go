package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UserInstallStateSummaryable 
type UserInstallStateSummaryable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceStates()([]DeviceInstallStateable)
    GetFailedDeviceCount()(*int32)
    GetInstalledDeviceCount()(*int32)
    GetNotInstalledDeviceCount()(*int32)
    GetUserName()(*string)
    SetDeviceStates(value []DeviceInstallStateable)()
    SetFailedDeviceCount(value *int32)()
    SetInstalledDeviceCount(value *int32)()
    SetNotInstalledDeviceCount(value *int32)()
    SetUserName(value *string)()
}
