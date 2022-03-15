package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DetectedAppable 
type DetectedAppable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceCount()(*int32)
    GetDisplayName()(*string)
    GetManagedDevices()([]ManagedDeviceable)
    GetSizeInByte()(*int32)
    GetVersion()(*string)
    SetDeviceCount(value *int32)()
    SetDisplayName(value *string)()
    SetManagedDevices(value []ManagedDeviceable)()
    SetSizeInByte(value *int32)()
    SetVersion(value *string)()
}
