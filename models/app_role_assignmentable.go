package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AppRoleAssignmentable 
type AppRoleAssignmentable interface {
    DirectoryObjectable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetAppRoleId()(*UUID)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetPrincipalDisplayName()(*string)
    GetPrincipalId()(*UUID)
    GetPrincipalType()(*string)
    GetResourceDisplayName()(*string)
    GetResourceId()(*UUID)
    SetAppRoleId(value *UUID)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetPrincipalDisplayName(value *string)()
    SetPrincipalId(value *UUID)()
    SetPrincipalType(value *string)()
    SetResourceDisplayName(value *string)()
    SetResourceId(value *UUID)()
}
