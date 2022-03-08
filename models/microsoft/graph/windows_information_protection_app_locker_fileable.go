package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// WindowsInformationProtectionAppLockerFileable 
type WindowsInformationProtectionAppLockerFileable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetDisplayName()(*string)
    GetFile()([]byte)
    GetFileHash()(*string)
    GetVersion()(*string)
    SetDisplayName(value *string)()
    SetFile(value []byte)()
    SetFileHash(value *string)()
    SetVersion(value *string)()
}
