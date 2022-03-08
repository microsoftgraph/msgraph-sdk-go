package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Processable 
type Processable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAccountName()(*string)
    GetCommandLine()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetFileHash()(FileHashable)
    GetIntegrityLevel()(*ProcessIntegrityLevel)
    GetIsElevated()(*bool)
    GetName()(*string)
    GetParentProcessCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetParentProcessId()(*int32)
    GetParentProcessName()(*string)
    GetPath()(*string)
    GetProcessId()(*int32)
    SetAccountName(value *string)()
    SetCommandLine(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetFileHash(value FileHashable)()
    SetIntegrityLevel(value *ProcessIntegrityLevel)()
    SetIsElevated(value *bool)()
    SetName(value *string)()
    SetParentProcessCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetParentProcessId(value *int32)()
    SetParentProcessName(value *string)()
    SetPath(value *string)()
    SetProcessId(value *int32)()
}
