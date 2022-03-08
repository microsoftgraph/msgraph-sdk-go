package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ImportedWindowsAutopilotDeviceIdentityable 
type ImportedWindowsAutopilotDeviceIdentityable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAssignedUserPrincipalName()(*string)
    GetGroupTag()(*string)
    GetHardwareIdentifier()([]byte)
    GetImportId()(*string)
    GetProductKey()(*string)
    GetSerialNumber()(*string)
    GetState()(ImportedWindowsAutopilotDeviceIdentityStateable)
    SetAssignedUserPrincipalName(value *string)()
    SetGroupTag(value *string)()
    SetHardwareIdentifier(value []byte)()
    SetImportId(value *string)()
    SetProductKey(value *string)()
    SetSerialNumber(value *string)()
    SetState(value ImportedWindowsAutopilotDeviceIdentityStateable)()
}
