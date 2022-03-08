package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppRegistrationable 
type ManagedAppRegistrationable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppIdentifier()(MobileAppIdentifierable)
    GetApplicationVersion()(*string)
    GetAppliedPolicies()([]ManagedAppPolicyable)
    GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetDeviceName()(*string)
    GetDeviceTag()(*string)
    GetDeviceType()(*string)
    GetFlaggedReasons()([]ManagedAppFlaggedReason)
    GetIntendedPolicies()([]ManagedAppPolicyable)
    GetLastSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetManagementSdkVersion()(*string)
    GetOperations()([]ManagedAppOperationable)
    GetPlatformVersion()(*string)
    GetUserId()(*string)
    GetVersion()(*string)
    SetAppIdentifier(value MobileAppIdentifierable)()
    SetApplicationVersion(value *string)()
    SetAppliedPolicies(value []ManagedAppPolicyable)()
    SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetDeviceName(value *string)()
    SetDeviceTag(value *string)()
    SetDeviceType(value *string)()
    SetFlaggedReasons(value []ManagedAppFlaggedReason)()
    SetIntendedPolicies(value []ManagedAppPolicyable)()
    SetLastSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetManagementSdkVersion(value *string)()
    SetOperations(value []ManagedAppOperationable)()
    SetPlatformVersion(value *string)()
    SetUserId(value *string)()
    SetVersion(value *string)()
}
