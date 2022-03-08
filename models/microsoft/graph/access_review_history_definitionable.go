package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewHistoryDefinitionable 
type AccessReviewHistoryDefinitionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCreatedBy()(UserIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDecisions()([]AccessReviewHistoryDecisionFilter)
    GetDisplayName()(*string)
    GetInstances()([]AccessReviewHistoryInstanceable)
    GetReviewHistoryPeriodEndDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewHistoryPeriodStartDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetScheduleSettings()(AccessReviewHistoryScheduleSettingsable)
    GetScopes()([]AccessReviewScopeable)
    GetStatus()(*AccessReviewHistoryStatus)
    SetCreatedBy(value UserIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDecisions(value []AccessReviewHistoryDecisionFilter)()
    SetDisplayName(value *string)()
    SetInstances(value []AccessReviewHistoryInstanceable)()
    SetReviewHistoryPeriodEndDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewHistoryPeriodStartDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetScheduleSettings(value AccessReviewHistoryScheduleSettingsable)()
    SetScopes(value []AccessReviewScopeable)()
    SetStatus(value *AccessReviewHistoryStatus)()
}
