package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// IdentityGovernanceable 
type IdentityGovernanceable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    GetAccessReviews()(AccessReviewSetable)
    GetAppConsent()(AppConsentApprovalRouteable)
    GetEntitlementManagement()(EntitlementManagementable)
    GetTermsOfUse()(TermsOfUseContainerable)
    SetAccessReviews(value AccessReviewSetable)()
    SetAppConsent(value AppConsentApprovalRouteable)()
    SetEntitlementManagement(value EntitlementManagementable)()
    SetTermsOfUse(value TermsOfUseContainerable)()
}
