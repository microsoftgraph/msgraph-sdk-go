package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentPolicyable 
type AccessPackageAssignmentPolicyable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackage()(AccessPackageable)
    GetAllowedTargetScope()(*AllowedTargetScope)
    GetCatalog()(AccessPackageCatalogable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDisplayName()(*string)
    GetExpiration()(ExpirationPatternable)
    GetModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetRequestApprovalSettings()(AccessPackageAssignmentApprovalSettingsable)
    GetRequestorSettings()(AccessPackageAssignmentRequestorSettingsable)
    GetReviewSettings()(AccessPackageAssignmentReviewSettingsable)
    GetSpecificAllowedTargets()([]SubjectSetable)
    SetAccessPackage(value AccessPackageable)()
    SetAllowedTargetScope(value *AllowedTargetScope)()
    SetCatalog(value AccessPackageCatalogable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDisplayName(value *string)()
    SetExpiration(value ExpirationPatternable)()
    SetModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetRequestApprovalSettings(value AccessPackageAssignmentApprovalSettingsable)()
    SetRequestorSettings(value AccessPackageAssignmentRequestorSettingsable)()
    SetReviewSettings(value AccessPackageAssignmentReviewSettingsable)()
    SetSpecificAllowedTargets(value []SubjectSetable)()
}
