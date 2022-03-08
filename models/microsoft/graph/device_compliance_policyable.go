package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceCompliancePolicyable 
type DeviceCompliancePolicyable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAssignments()([]DeviceCompliancePolicyAssignmentable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDescription()(*string)
    GetDeviceSettingStateSummaries()([]SettingStateDeviceSummaryable)
    GetDeviceStatuses()([]DeviceComplianceDeviceStatusable)
    GetDeviceStatusOverview()(DeviceComplianceDeviceOverviewable)
    GetDisplayName()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetScheduledActionsForRule()([]DeviceComplianceScheduledActionForRuleable)
    GetUserStatuses()([]DeviceComplianceUserStatusable)
    GetUserStatusOverview()(DeviceComplianceUserOverviewable)
    GetVersion()(*int32)
    SetAssignments(value []DeviceCompliancePolicyAssignmentable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDescription(value *string)()
    SetDeviceSettingStateSummaries(value []SettingStateDeviceSummaryable)()
    SetDeviceStatuses(value []DeviceComplianceDeviceStatusable)()
    SetDeviceStatusOverview(value DeviceComplianceDeviceOverviewable)()
    SetDisplayName(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetScheduledActionsForRule(value []DeviceComplianceScheduledActionForRuleable)()
    SetUserStatuses(value []DeviceComplianceUserStatusable)()
    SetUserStatusOverview(value DeviceComplianceUserOverviewable)()
    SetVersion(value *int32)()
}
