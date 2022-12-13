package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuditEventable 
type AuditEventable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActivity()(*string)
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityOperationType()(*string)
    GetActivityResult()(*string)
    GetActivityType()(*string)
    GetActor()(AuditActorable)
    GetCategory()(*string)
    GetComponentName()(*string)
    GetCorrelationId()(*UUID)
    GetDisplayName()(*string)
    GetResources()([]AuditResourceable)
    SetActivity(value *string)()
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityOperationType(value *string)()
    SetActivityResult(value *string)()
    SetActivityType(value *string)()
    SetActor(value AuditActorable)()
    SetCategory(value *string)()
    SetComponentName(value *string)()
    SetCorrelationId(value *UUID)()
    SetDisplayName(value *string)()
    SetResources(value []AuditResourceable)()
}
