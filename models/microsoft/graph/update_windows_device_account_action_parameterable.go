package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdateWindowsDeviceAccountActionParameterable 
type UpdateWindowsDeviceAccountActionParameterable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetCalendarSyncEnabled()(*bool)
    GetDeviceAccount()(WindowsDeviceAccountable)
    GetDeviceAccountEmail()(*string)
    GetExchangeServer()(*string)
    GetPasswordRotationEnabled()(*bool)
    GetSessionInitiationProtocalAddress()(*string)
    SetCalendarSyncEnabled(value *bool)()
    SetDeviceAccount(value WindowsDeviceAccountable)()
    SetDeviceAccountEmail(value *string)()
    SetExchangeServer(value *string)()
    SetPasswordRotationEnabled(value *bool)()
    SetSessionInitiationProtocalAddress(value *string)()
}
