package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DeviceAppManagementable 
type DeviceAppManagementable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    Entityable
    GetAndroidManagedAppProtections()([]AndroidManagedAppProtectionable)
    GetDefaultManagedAppProtections()([]DefaultManagedAppProtectionable)
    GetIosManagedAppProtections()([]IosManagedAppProtectionable)
    GetIsEnabledForMicrosoftStoreForBusiness()(*bool)
    GetManagedAppPolicies()([]ManagedAppPolicyable)
    GetManagedAppRegistrations()([]ManagedAppRegistrationable)
    GetManagedAppStatuses()([]ManagedAppStatusable)
    GetManagedEBooks()([]ManagedEBookable)
    GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicyable)
    GetMicrosoftStoreForBusinessLanguage()(*string)
    GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMobileAppCategories()([]MobileAppCategoryable)
    GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfigurationable)
    GetMobileApps()([]MobileAppable)
    GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfigurationable)
    GetVppTokens()([]VppTokenable)
    GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicyable)
    SetAndroidManagedAppProtections(value []AndroidManagedAppProtectionable)()
    SetDefaultManagedAppProtections(value []DefaultManagedAppProtectionable)()
    SetIosManagedAppProtections(value []IosManagedAppProtectionable)()
    SetIsEnabledForMicrosoftStoreForBusiness(value *bool)()
    SetManagedAppPolicies(value []ManagedAppPolicyable)()
    SetManagedAppRegistrations(value []ManagedAppRegistrationable)()
    SetManagedAppStatuses(value []ManagedAppStatusable)()
    SetManagedEBooks(value []ManagedEBookable)()
    SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicyable)()
    SetMicrosoftStoreForBusinessLanguage(value *string)()
    SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMobileAppCategories(value []MobileAppCategoryable)()
    SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfigurationable)()
    SetMobileApps(value []MobileAppable)()
    SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfigurationable)()
    SetVppTokens(value []VppTokenable)()
    SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicyable)()
}
