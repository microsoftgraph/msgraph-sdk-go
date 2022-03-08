package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RiskyUserable 
type RiskyUserable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetHistory()([]RiskyUserHistoryItemable)
    GetIsDeleted()(*bool)
    GetIsProcessing()(*bool)
    GetRiskDetail()(*RiskDetail)
    GetRiskLastUpdatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRiskLevel()(*RiskLevel)
    GetRiskState()(*RiskState)
    GetUserDisplayName()(*string)
    GetUserPrincipalName()(*string)
    SetHistory(value []RiskyUserHistoryItemable)()
    SetIsDeleted(value *bool)()
    SetIsProcessing(value *bool)()
    SetRiskDetail(value *RiskDetail)()
    SetRiskLastUpdatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRiskLevel(value *RiskLevel)()
    SetRiskState(value *RiskState)()
    SetUserDisplayName(value *string)()
    SetUserPrincipalName(value *string)()
}
