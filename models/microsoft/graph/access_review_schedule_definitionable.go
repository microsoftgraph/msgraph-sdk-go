package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessReviewScheduleDefinitionable 
type AccessReviewScheduleDefinitionable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAdditionalNotificationRecipients()([]AccessReviewNotificationRecipientItemable)
    GetCreatedBy()(UserIdentityable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescriptionForAdmins()(*string)
    GetDescriptionForReviewers()(*string)
    GetDisplayName()(*string)
    GetFallbackReviewers()([]AccessReviewReviewerScopeable)
    GetInstanceEnumerationScope()(AccessReviewScopeable)
    GetInstances()([]AccessReviewInstanceable)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetReviewers()([]AccessReviewReviewerScopeable)
    GetScope()(AccessReviewScopeable)
    GetSettings()(AccessReviewScheduleSettingsable)
    GetStatus()(*string)
    SetAdditionalNotificationRecipients(value []AccessReviewNotificationRecipientItemable)()
    SetCreatedBy(value UserIdentityable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescriptionForAdmins(value *string)()
    SetDescriptionForReviewers(value *string)()
    SetDisplayName(value *string)()
    SetFallbackReviewers(value []AccessReviewReviewerScopeable)()
    SetInstanceEnumerationScope(value AccessReviewScopeable)()
    SetInstances(value []AccessReviewInstanceable)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetReviewers(value []AccessReviewReviewerScopeable)()
    SetScope(value AccessReviewScopeable)()
    SetSettings(value AccessReviewScheduleSettingsable)()
    SetStatus(value *string)()
}
