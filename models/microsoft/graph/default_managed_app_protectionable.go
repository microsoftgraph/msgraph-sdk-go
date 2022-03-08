package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DefaultManagedAppProtectionable 
type DefaultManagedAppProtectionable interface {
    ManagedAppProtectionable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetAppDataEncryptionType()(*ManagedAppDataEncryptionType)
    GetApps()([]ManagedMobileAppable)
    GetCustomSettings()([]KeyValuePairable)
    GetDeployedAppCount()(*int32)
    GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable)
    GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool)
    GetEncryptAppData()(*bool)
    GetFaceIdBlocked()(*bool)
    GetMinimumRequiredPatchVersion()(*string)
    GetMinimumRequiredSdkVersion()(*string)
    GetMinimumWarningPatchVersion()(*string)
    GetScreenCaptureBlocked()(*bool)
    SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)()
    SetApps(value []ManagedMobileAppable)()
    SetCustomSettings(value []KeyValuePairable)()
    SetDeployedAppCount(value *int32)()
    SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)()
    SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)()
    SetEncryptAppData(value *bool)()
    SetFaceIdBlocked(value *bool)()
    SetMinimumRequiredPatchVersion(value *string)()
    SetMinimumRequiredSdkVersion(value *string)()
    SetMinimumWarningPatchVersion(value *string)()
    SetScreenCaptureBlocked(value *bool)()
}
