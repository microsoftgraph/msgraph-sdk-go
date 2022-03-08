package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppPolicyDeploymentSummaryable 
type ManagedAppPolicyDeploymentSummaryable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetConfigurationDeployedUserCount()(*int32)
    GetConfigurationDeploymentSummaryPerApp()([]ManagedAppPolicyDeploymentSummaryPerAppable)
    GetDisplayName()(*string)
    GetLastRefreshTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetVersion()(*string)
    SetConfigurationDeployedUserCount(value *int32)()
    SetConfigurationDeploymentSummaryPerApp(value []ManagedAppPolicyDeploymentSummaryPerAppable)()
    SetDisplayName(value *string)()
    SetLastRefreshTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetVersion(value *string)()
}
