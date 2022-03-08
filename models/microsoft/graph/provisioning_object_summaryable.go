package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ProvisioningObjectSummaryable 
type ProvisioningObjectSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetChangeId()(*string)
    GetCycleId()(*string)
    GetDurationInMilliseconds()(*int32)
    GetInitiatedBy()(Initiatorable)
    GetJobId()(*string)
    GetModifiedProperties()([]ModifiedPropertyable)
    GetProvisioningAction()(*ProvisioningAction)
    GetProvisioningStatusInfo()(ProvisioningStatusInfoable)
    GetProvisioningSteps()([]ProvisioningStepable)
    GetServicePrincipal()(ProvisioningServicePrincipalable)
    GetSourceIdentity()(ProvisionedIdentityable)
    GetSourceSystem()(ProvisioningSystemable)
    GetTargetIdentity()(ProvisionedIdentityable)
    GetTargetSystem()(ProvisioningSystemable)
    GetTenantId()(*string)
    SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetChangeId(value *string)()
    SetCycleId(value *string)()
    SetDurationInMilliseconds(value *int32)()
    SetInitiatedBy(value Initiatorable)()
    SetJobId(value *string)()
    SetModifiedProperties(value []ModifiedPropertyable)()
    SetProvisioningAction(value *ProvisioningAction)()
    SetProvisioningStatusInfo(value ProvisioningStatusInfoable)()
    SetProvisioningSteps(value []ProvisioningStepable)()
    SetServicePrincipal(value ProvisioningServicePrincipalable)()
    SetSourceIdentity(value ProvisionedIdentityable)()
    SetSourceSystem(value ProvisioningSystemable)()
    SetTargetIdentity(value ProvisionedIdentityable)()
    SetTargetSystem(value ProvisioningSystemable)()
    SetTenantId(value *string)()
}
