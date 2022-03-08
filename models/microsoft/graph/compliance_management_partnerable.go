package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ComplianceManagementPartnerable 
type ComplianceManagementPartnerable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAndroidEnrollmentAssignments()([]ComplianceManagementPartnerAssignmentable)
    GetAndroidOnboarded()(*bool)
    GetDisplayName()(*string)
    GetIosEnrollmentAssignments()([]ComplianceManagementPartnerAssignmentable)
    GetIosOnboarded()(*bool)
    GetLastHeartbeatDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMacOsEnrollmentAssignments()([]ComplianceManagementPartnerAssignmentable)
    GetMacOsOnboarded()(*bool)
    GetPartnerState()(*DeviceManagementPartnerTenantState)
    SetAndroidEnrollmentAssignments(value []ComplianceManagementPartnerAssignmentable)()
    SetAndroidOnboarded(value *bool)()
    SetDisplayName(value *string)()
    SetIosEnrollmentAssignments(value []ComplianceManagementPartnerAssignmentable)()
    SetIosOnboarded(value *bool)()
    SetLastHeartbeatDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMacOsEnrollmentAssignments(value []ComplianceManagementPartnerAssignmentable)()
    SetMacOsOnboarded(value *bool)()
    SetPartnerState(value *DeviceManagementPartnerTenantState)()
}
