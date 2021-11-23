package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// deviceAppManagement 
type DeviceAppManagement struct {
    Entity
    // Android managed app policies.
    androidManagedAppProtections []AndroidManagedAppProtection;
    // Default managed app policies.
    defaultManagedAppProtections []DefaultManagedAppProtection;
    // iOS managed app policies.
    iosManagedAppProtections []IosManagedAppProtection;
    // Whether the account is enabled for syncing applications from the Microsoft Store for Business.
    isEnabledForMicrosoftStoreForBusiness *bool;
    // Managed app policies.
    managedAppPolicies []ManagedAppPolicy;
    // The managed app registrations.
    managedAppRegistrations []ManagedAppRegistration;
    // The managed app statuses.
    managedAppStatuses []ManagedAppStatus;
    // The Managed eBook.
    managedEBooks []ManagedEBook;
    // Windows information protection for apps running on devices which are MDM enrolled.
    mdmWindowsInformationProtectionPolicies []MdmWindowsInformationProtectionPolicy;
    // The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
    microsoftStoreForBusinessLanguage *string;
    // The last time an application sync from the Microsoft Store for Business was completed.
    microsoftStoreForBusinessLastCompletedApplicationSyncTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The last time the apps from the Microsoft Store for Business were synced successfully for the account.
    microsoftStoreForBusinessLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The mobile app categories.
    mobileAppCategories []MobileAppCategory;
    // The Managed Device Mobile Application Configurations.
    mobileAppConfigurations []ManagedDeviceMobileAppConfiguration;
    // The mobile apps.
    mobileApps []MobileApp;
    // Targeted managed app configurations.
    targetedManagedAppConfigurations []TargetedManagedAppConfiguration;
    // List of Vpp tokens for this organization.
    vppTokens []VppToken;
    // Windows information protection for apps running on devices which are not MDM enrolled.
    windowsInformationProtectionPolicies []WindowsInformationProtectionPolicy;
}
// NewDeviceAppManagement instantiates a new deviceAppManagement and sets the default values.
func NewDeviceAppManagement()(*DeviceAppManagement) {
    m := &DeviceAppManagement{
        Entity: *NewEntity(),
    }
    return m
}
// GetAndroidManagedAppProtections gets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) GetAndroidManagedAppProtections()([]AndroidManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.androidManagedAppProtections
    }
}
// GetDefaultManagedAppProtections gets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) GetDefaultManagedAppProtections()([]DefaultManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.defaultManagedAppProtections
    }
}
// GetIosManagedAppProtections gets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) GetIosManagedAppProtections()([]IosManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.iosManagedAppProtections
    }
}
// GetIsEnabledForMicrosoftStoreForBusiness gets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabledForMicrosoftStoreForBusiness
    }
}
// GetManagedAppPolicies gets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) GetManagedAppPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.managedAppPolicies
    }
}
// GetManagedAppRegistrations gets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) GetManagedAppRegistrations()([]ManagedAppRegistration) {
    if m == nil {
        return nil
    } else {
        return m.managedAppRegistrations
    }
}
// GetManagedAppStatuses gets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) GetManagedAppStatuses()([]ManagedAppStatus) {
    if m == nil {
        return nil
    } else {
        return m.managedAppStatuses
    }
}
// GetManagedEBooks gets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) GetManagedEBooks()([]ManagedEBook) {
    if m == nil {
        return nil
    } else {
        return m.managedEBooks
    }
}
// GetMdmWindowsInformationProtectionPolicies gets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.mdmWindowsInformationProtectionPolicies
    }
}
// GetMicrosoftStoreForBusinessLanguage gets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLanguage
    }
}
// GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime gets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastCompletedApplicationSyncTime
    }
}
// GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime gets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastSuccessfulSyncDateTime
    }
}
// GetMobileAppCategories gets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) GetMobileAppCategories()([]MobileAppCategory) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppCategories
    }
}
// GetMobileAppConfigurations gets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppConfigurations
    }
}
// GetMobileApps gets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) GetMobileApps()([]MobileApp) {
    if m == nil {
        return nil
    } else {
        return m.mobileApps
    }
}
// GetTargetedManagedAppConfigurations gets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.targetedManagedAppConfigurations
    }
}
// GetVppTokens gets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) GetVppTokens()([]VppToken) {
    if m == nil {
        return nil
    } else {
        return m.vppTokens
    }
}
// GetWindowsInformationProtectionPolicies gets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionPolicies
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceAppManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidManagedAppProtections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidManagedAppProtection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AndroidManagedAppProtection, len(val))
            for i, v := range val {
                res[i] = *(v.(*AndroidManagedAppProtection))
            }
            m.SetAndroidManagedAppProtections(res)
        }
        return nil
    }
    res["defaultManagedAppProtections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDefaultManagedAppProtection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]DefaultManagedAppProtection, len(val))
            for i, v := range val {
                res[i] = *(v.(*DefaultManagedAppProtection))
            }
            m.SetDefaultManagedAppProtections(res)
        }
        return nil
    }
    res["iosManagedAppProtections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIosManagedAppProtection() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]IosManagedAppProtection, len(val))
            for i, v := range val {
                res[i] = *(v.(*IosManagedAppProtection))
            }
            m.SetIosManagedAppProtections(res)
        }
        return nil
    }
    res["isEnabledForMicrosoftStoreForBusiness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsEnabledForMicrosoftStoreForBusiness(val)
        }
        return nil
    }
    res["managedAppPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppPolicy))
            }
            m.SetManagedAppPolicies(res)
        }
        return nil
    }
    res["managedAppRegistrations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppRegistration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppRegistration, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppRegistration))
            }
            m.SetManagedAppRegistrations(res)
        }
        return nil
    }
    res["managedAppStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppStatus, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppStatus))
            }
            m.SetManagedAppStatuses(res)
        }
        return nil
    }
    res["managedEBooks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedEBook() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedEBook, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedEBook))
            }
            m.SetManagedEBooks(res)
        }
        return nil
    }
    res["mdmWindowsInformationProtectionPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMdmWindowsInformationProtectionPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MdmWindowsInformationProtectionPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*MdmWindowsInformationProtectionPolicy))
            }
            m.SetMdmWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    res["microsoftStoreForBusinessLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLanguage(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastCompletedApplicationSyncTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(val)
        }
        return nil
    }
    res["microsoftStoreForBusinessLastSuccessfulSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(val)
        }
        return nil
    }
    res["mobileAppCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppCategory() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileAppCategory, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileAppCategory))
            }
            m.SetMobileAppCategories(res)
        }
        return nil
    }
    res["mobileAppConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedDeviceMobileAppConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedDeviceMobileAppConfiguration))
            }
            m.SetMobileAppConfigurations(res)
        }
        return nil
    }
    res["mobileApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileApp() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]MobileApp, len(val))
            for i, v := range val {
                res[i] = *(v.(*MobileApp))
            }
            m.SetMobileApps(res)
        }
        return nil
    }
    res["targetedManagedAppConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetedManagedAppConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TargetedManagedAppConfiguration, len(val))
            for i, v := range val {
                res[i] = *(v.(*TargetedManagedAppConfiguration))
            }
            m.SetTargetedManagedAppConfigurations(res)
        }
        return nil
    }
    res["vppTokens"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVppToken() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]VppToken, len(val))
            for i, v := range val {
                res[i] = *(v.(*VppToken))
            }
            m.SetVppTokens(res)
        }
        return nil
    }
    res["windowsInformationProtectionPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionPolicy() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]WindowsInformationProtectionPolicy, len(val))
            for i, v := range val {
                res[i] = *(v.(*WindowsInformationProtectionPolicy))
            }
            m.SetWindowsInformationProtectionPolicies(res)
        }
        return nil
    }
    return res
}
func (m *DeviceAppManagement) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DeviceAppManagement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAndroidManagedAppProtections()))
        for i, v := range m.GetAndroidManagedAppProtections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("androidManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDefaultManagedAppProtections()))
        for i, v := range m.GetDefaultManagedAppProtections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("defaultManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIosManagedAppProtections()))
        for i, v := range m.GetIosManagedAppProtections() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("iosManagedAppProtections", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isEnabledForMicrosoftStoreForBusiness", m.GetIsEnabledForMicrosoftStoreForBusiness())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedAppPolicies()))
        for i, v := range m.GetManagedAppPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedAppRegistrations()))
        for i, v := range m.GetManagedAppRegistrations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppRegistrations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedAppStatuses()))
        for i, v := range m.GetManagedAppStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedAppStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedEBooks()))
        for i, v := range m.GetManagedEBooks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedEBooks", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMdmWindowsInformationProtectionPolicies()))
        for i, v := range m.GetMdmWindowsInformationProtectionPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mdmWindowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("microsoftStoreForBusinessLanguage", m.GetMicrosoftStoreForBusinessLanguage())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("microsoftStoreForBusinessLastCompletedApplicationSyncTime", m.GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("microsoftStoreForBusinessLastSuccessfulSyncDateTime", m.GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileAppCategories()))
        for i, v := range m.GetMobileAppCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileAppConfigurations()))
        for i, v := range m.GetMobileAppConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileApps()))
        for i, v := range m.GetMobileApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileApps", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTargetedManagedAppConfigurations()))
        for i, v := range m.GetTargetedManagedAppConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("targetedManagedAppConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetVppTokens()))
        for i, v := range m.GetVppTokens() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("vppTokens", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsInformationProtectionPolicies()))
        for i, v := range m.GetWindowsInformationProtectionPolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionPolicies", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAndroidManagedAppProtections sets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) SetAndroidManagedAppProtections(value []AndroidManagedAppProtection)() {
    m.androidManagedAppProtections = value
}
// SetDefaultManagedAppProtections sets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) SetDefaultManagedAppProtections(value []DefaultManagedAppProtection)() {
    m.defaultManagedAppProtections = value
}
// SetIosManagedAppProtections sets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) SetIosManagedAppProtections(value []IosManagedAppProtection)() {
    m.iosManagedAppProtections = value
}
// SetIsEnabledForMicrosoftStoreForBusiness sets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(value *bool)() {
    m.isEnabledForMicrosoftStoreForBusiness = value
}
// SetManagedAppPolicies sets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) SetManagedAppPolicies(value []ManagedAppPolicy)() {
    m.managedAppPolicies = value
}
// SetManagedAppRegistrations sets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) SetManagedAppRegistrations(value []ManagedAppRegistration)() {
    m.managedAppRegistrations = value
}
// SetManagedAppStatuses sets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) SetManagedAppStatuses(value []ManagedAppStatus)() {
    m.managedAppStatuses = value
}
// SetManagedEBooks sets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) SetManagedEBooks(value []ManagedEBook)() {
    m.managedEBooks = value
}
// SetMdmWindowsInformationProtectionPolicies sets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicy)() {
    m.mdmWindowsInformationProtectionPolicies = value
}
// SetMicrosoftStoreForBusinessLanguage sets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(value *string)() {
    m.microsoftStoreForBusinessLanguage = value
}
// SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime sets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastCompletedApplicationSyncTime = value
}
// SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime sets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastSuccessfulSyncDateTime = value
}
// SetMobileAppCategories sets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) SetMobileAppCategories(value []MobileAppCategory)() {
    m.mobileAppCategories = value
}
// SetMobileAppConfigurations sets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfiguration)() {
    m.mobileAppConfigurations = value
}
// SetMobileApps sets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) SetMobileApps(value []MobileApp)() {
    m.mobileApps = value
}
// SetTargetedManagedAppConfigurations sets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfiguration)() {
    m.targetedManagedAppConfigurations = value
}
// SetVppTokens sets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) SetVppTokens(value []VppToken)() {
    m.vppTokens = value
}
// SetWindowsInformationProtectionPolicies sets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicy)() {
    m.windowsInformationProtectionPolicies = value
}
