package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceExchangeAccessStateSummaryable 
type DeviceExchangeAccessStateSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAllowedDeviceCount()(*int32)
    GetBlockedDeviceCount()(*int32)
    GetQuarantinedDeviceCount()(*int32)
    GetUnavailableDeviceCount()(*int32)
    GetUnknownDeviceCount()(*int32)
    SetAllowedDeviceCount(value *int32)()
    SetBlockedDeviceCount(value *int32)()
    SetQuarantinedDeviceCount(value *int32)()
    SetUnavailableDeviceCount(value *int32)()
    SetUnknownDeviceCount(value *int32)()
}
