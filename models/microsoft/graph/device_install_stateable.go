package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceInstallStateable 
type DeviceInstallStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDeviceId()(*string)
    GetDeviceName()(*string)
    GetErrorCode()(*string)
    GetInstallState()(*InstallState)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetOsDescription()(*string)
    GetOsVersion()(*string)
    GetUserName()(*string)
    SetDeviceId(value *string)()
    SetDeviceName(value *string)()
    SetErrorCode(value *string)()
    SetInstallState(value *InstallState)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetOsDescription(value *string)()
    SetOsVersion(value *string)()
    SetUserName(value *string)()
}
