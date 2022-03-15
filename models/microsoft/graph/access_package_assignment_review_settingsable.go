package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentReviewSettingsable 
type AccessPackageAssignmentReviewSettingsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetExpirationBehavior()(*AccessReviewExpirationBehavior)
    GetFallbackReviewers()([]SubjectSetable)
    GetIsEnabled()(*bool)
    GetIsRecommendationEnabled()(*bool)
    GetIsReviewerJustificationRequired()(*bool)
    GetIsSelfReview()(*bool)
    GetPrimaryReviewers()([]SubjectSetable)
    GetSchedule()(EntitlementManagementScheduleable)
    SetExpirationBehavior(value *AccessReviewExpirationBehavior)()
    SetFallbackReviewers(value []SubjectSetable)()
    SetIsEnabled(value *bool)()
    SetIsRecommendationEnabled(value *bool)()
    SetIsReviewerJustificationRequired(value *bool)()
    SetIsSelfReview(value *bool)()
    SetPrimaryReviewers(value []SubjectSetable)()
    SetSchedule(value EntitlementManagementScheduleable)()
}
