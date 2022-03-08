package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DirectoryAuditable 
type DirectoryAuditable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetActivityDisplayName()(*string)
    GetAdditionalDetails()([]KeyValueable)
    GetCategory()(*string)
    GetCorrelationId()(*string)
    GetInitiatedBy()(AuditActivityInitiatorable)
    GetLoggedByService()(*string)
    GetOperationType()(*string)
    GetResult()(*OperationResult)
    GetResultReason()(*string)
    GetTargetResources()([]TargetResourceable)
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetActivityDisplayName(value *string)()
    SetAdditionalDetails(value []KeyValueable)()
    SetCategory(value *string)()
    SetCorrelationId(value *string)()
    SetInitiatedBy(value AuditActivityInitiatorable)()
    SetLoggedByService(value *string)()
    SetOperationType(value *string)()
    SetResult(value *OperationResult)()
    SetResultReason(value *string)()
    SetTargetResources(value []TargetResourceable)()
}
