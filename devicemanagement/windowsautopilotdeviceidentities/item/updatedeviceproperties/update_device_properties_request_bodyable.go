package updatedeviceproperties

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// UpdateDevicePropertiesRequestBodyable 
type UpdateDevicePropertiesRequestBodyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAddressableUserName()(*string)
    GetDisplayName()(*string)
    GetGroupTag()(*string)
    GetUserPrincipalName()(*string)
    SetAddressableUserName(value *string)()
    SetDisplayName(value *string)()
    SetGroupTag(value *string)()
    SetUserPrincipalName(value *string)()
}
