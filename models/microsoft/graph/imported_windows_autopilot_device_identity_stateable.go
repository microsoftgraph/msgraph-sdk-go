package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportedWindowsAutopilotDeviceIdentityStateable 
type ImportedWindowsAutopilotDeviceIdentityStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetDeviceErrorCode()(*int32)
    GetDeviceErrorName()(*string)
    GetDeviceImportStatus()(*ImportedWindowsAutopilotDeviceIdentityImportStatus)
    GetDeviceRegistrationId()(*string)
    SetDeviceErrorCode(value *int32)()
    SetDeviceErrorName(value *string)()
    SetDeviceImportStatus(value *ImportedWindowsAutopilotDeviceIdentityImportStatus)()
    SetDeviceRegistrationId(value *string)()
}
