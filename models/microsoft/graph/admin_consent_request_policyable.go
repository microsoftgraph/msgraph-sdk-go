package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AdminConsentRequestPolicyable 
type AdminConsentRequestPolicyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetIsEnabled()(*bool)
    GetNotifyReviewers()(*bool)
    GetRemindersEnabled()(*bool)
    GetRequestDurationInDays()(*int32)
    GetReviewers()([]AccessReviewReviewerScopeable)
    GetVersion()(*int32)
    SetIsEnabled(value *bool)()
    SetNotifyReviewers(value *bool)()
    SetRemindersEnabled(value *bool)()
    SetRequestDurationInDays(value *int32)()
    SetReviewers(value []AccessReviewReviewerScopeable)()
    SetVersion(value *int32)()
}
