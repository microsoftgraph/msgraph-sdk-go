package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EntitlementManagementable 
type EntitlementManagementable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAccessPackageAssignmentApprovals()([]Approvalable)
    GetAccessPackages()([]AccessPackageable)
    GetAssignmentRequests()([]AccessPackageAssignmentRequestable)
    GetAssignments()([]AccessPackageAssignmentable)
    GetCatalogs()([]AccessPackageCatalogable)
    GetConnectedOrganizations()([]ConnectedOrganizationable)
    GetSettings()(EntitlementManagementSettingsable)
    SetAccessPackageAssignmentApprovals(value []Approvalable)()
    SetAccessPackages(value []AccessPackageable)()
    SetAssignmentRequests(value []AccessPackageAssignmentRequestable)()
    SetAssignments(value []AccessPackageAssignmentable)()
    SetCatalogs(value []AccessPackageCatalogable)()
    SetConnectedOrganizations(value []ConnectedOrganizationable)()
    SetSettings(value EntitlementManagementSettingsable)()
}
