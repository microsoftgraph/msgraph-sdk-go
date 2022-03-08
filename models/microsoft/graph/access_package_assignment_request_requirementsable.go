package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AccessPackageAssignmentRequestRequirementsable 
type AccessPackageAssignmentRequestRequirementsable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAllowCustomAssignmentSchedule()(*bool)
    GetIsApprovalRequiredForAdd()(*bool)
    GetIsApprovalRequiredForUpdate()(*bool)
    GetPolicyDescription()(*string)
    GetPolicyDisplayName()(*string)
    GetPolicyId()(*string)
    GetSchedule()(EntitlementManagementScheduleable)
    SetAllowCustomAssignmentSchedule(value *bool)()
    SetIsApprovalRequiredForAdd(value *bool)()
    SetIsApprovalRequiredForUpdate(value *bool)()
    SetPolicyDescription(value *string)()
    SetPolicyDisplayName(value *string)()
    SetPolicyId(value *string)()
    SetSchedule(value EntitlementManagementScheduleable)()
}
