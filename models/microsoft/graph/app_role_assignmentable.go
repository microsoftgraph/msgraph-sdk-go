package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AppRoleAssignmentable 
type AppRoleAssignmentable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    DirectoryObjectable
    GetAppRoleId()(*string)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPrincipalDisplayName()(*string)
    GetPrincipalId()(*string)
    GetPrincipalType()(*string)
    GetResourceDisplayName()(*string)
    GetResourceId()(*string)
    SetAppRoleId(value *string)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPrincipalDisplayName(value *string)()
    SetPrincipalId(value *string)()
    SetPrincipalType(value *string)()
    SetResourceDisplayName(value *string)()
    SetResourceId(value *string)()
}
