package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AndroidManagedAppProtectionable 
type AndroidManagedAppProtectionable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    TargetedManagedAppProtectionable
    GetApps()([]ManagedMobileAppable)
    GetCustomBrowserDisplayName()(*string)
    GetCustomBrowserPackageId()(*string)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool)
    GetEncryptAppData()(*bool)
    GetMinimumRequiredPatchVersion()(*string)
    GetMinimumWarningPatchVersion()(*string)
    GetScreenCaptureBlocked()(*bool)
    SetApps(value []ManagedMobileAppable)()
    SetCustomBrowserDisplayName(value *string)()
    SetCustomBrowserPackageId(value *string)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)()
    SetEncryptAppData(value *bool)()
    SetMinimumRequiredPatchVersion(value *string)()
    SetMinimumWarningPatchVersion(value *string)()
    SetScreenCaptureBlocked(value *bool)()
}
