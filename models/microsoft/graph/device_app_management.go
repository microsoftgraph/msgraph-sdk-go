package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new deviceAppManagement and sets the default values.
func NewDeviceAppManagement()(*DeviceAppManagement) {
    m := &DeviceAppManagement{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the androidManagedAppProtections property value. Android managed app policies.
func (m *DeviceAppManagement) GetAndroidManagedAppProtections()([]AndroidManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.androidManagedAppProtections
    }
}
// Gets the defaultManagedAppProtections property value. Default managed app policies.
func (m *DeviceAppManagement) GetDefaultManagedAppProtections()([]DefaultManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.defaultManagedAppProtections
    }
}
// Gets the iosManagedAppProtections property value. iOS managed app policies.
func (m *DeviceAppManagement) GetIosManagedAppProtections()([]IosManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.iosManagedAppProtections
    }
}
// Gets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
func (m *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabledForMicrosoftStoreForBusiness
    }
}
// Gets the managedAppPolicies property value. Managed app policies.
func (m *DeviceAppManagement) GetManagedAppPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.managedAppPolicies
    }
}
// Gets the managedAppRegistrations property value. The managed app registrations.
func (m *DeviceAppManagement) GetManagedAppRegistrations()([]ManagedAppRegistration) {
    if m == nil {
        return nil
    } else {
        return m.managedAppRegistrations
    }
}
// Gets the managedAppStatuses property value. The managed app statuses.
func (m *DeviceAppManagement) GetManagedAppStatuses()([]ManagedAppStatus) {
    if m == nil {
        return nil
    } else {
        return m.managedAppStatuses
    }
}
// Gets the managedEBooks property value. The Managed eBook.
func (m *DeviceAppManagement) GetManagedEBooks()([]ManagedEBook) {
    if m == nil {
        return nil
    } else {
        return m.managedEBooks
    }
}
// Gets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
func (m *DeviceAppManagement) GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.mdmWindowsInformationProtectionPolicies
    }
}
// Gets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLanguage
    }
}
// Gets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastCompletedApplicationSyncTime
    }
}
// Gets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastSuccessfulSyncDateTime
    }
}
// Gets the mobileAppCategories property value. The mobile app categories.
func (m *DeviceAppManagement) GetMobileAppCategories()([]MobileAppCategory) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppCategories
    }
}
// Gets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
func (m *DeviceAppManagement) GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppConfigurations
    }
}
// Gets the mobileApps property value. The mobile apps.
func (m *DeviceAppManagement) GetMobileApps()([]MobileApp) {
    if m == nil {
        return nil
    } else {
        return m.mobileApps
    }
}
// Gets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
func (m *DeviceAppManagement) GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.targetedManagedAppConfigurations
    }
}
// Gets the vppTokens property value. List of Vpp tokens for this organization.
func (m *DeviceAppManagement) GetVppTokens()([]VppToken) {
    if m == nil {
        return nil
    } else {
        return m.vppTokens
    }
}
// Gets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
func (m *DeviceAppManagement) GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionPolicies
    }
}
// The deserialization information for the current model
func (m *DeviceAppManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["androidManagedAppProtections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAndroidManagedAppProtection() })
        if err != nil {
            return err
        }
        res := make([]AndroidManagedAppProtection, len(val))
        for i, v := range val {
            res[i] = *(v.(*AndroidManagedAppProtection))
        }
        m.SetAndroidManagedAppProtections(res)
        return nil
    }
    res["defaultManagedAppProtections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDefaultManagedAppProtection() })
        if err != nil {
            return err
        }
        res := make([]DefaultManagedAppProtection, len(val))
        for i, v := range val {
            res[i] = *(v.(*DefaultManagedAppProtection))
        }
        m.SetDefaultManagedAppProtections(res)
        return nil
    }
    res["iosManagedAppProtections"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIosManagedAppProtection() })
        if err != nil {
            return err
        }
        res := make([]IosManagedAppProtection, len(val))
        for i, v := range val {
            res[i] = *(v.(*IosManagedAppProtection))
        }
        m.SetIosManagedAppProtections(res)
        return nil
    }
    res["isEnabledForMicrosoftStoreForBusiness"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsEnabledForMicrosoftStoreForBusiness(val)
        return nil
    }
    res["managedAppPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicy() })
        if err != nil {
            return err
        }
        res := make([]ManagedAppPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppPolicy))
        }
        m.SetManagedAppPolicies(res)
        return nil
    }
    res["managedAppRegistrations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppRegistration() })
        if err != nil {
            return err
        }
        res := make([]ManagedAppRegistration, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppRegistration))
        }
        m.SetManagedAppRegistrations(res)
        return nil
    }
    res["managedAppStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppStatus() })
        if err != nil {
            return err
        }
        res := make([]ManagedAppStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppStatus))
        }
        m.SetManagedAppStatuses(res)
        return nil
    }
    res["managedEBooks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedEBook() })
        if err != nil {
            return err
        }
        res := make([]ManagedEBook, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedEBook))
        }
        m.SetManagedEBooks(res)
        return nil
    }
    res["mdmWindowsInformationProtectionPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMdmWindowsInformationProtectionPolicy() })
        if err != nil {
            return err
        }
        res := make([]MdmWindowsInformationProtectionPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*MdmWindowsInformationProtectionPolicy))
        }
        m.SetMdmWindowsInformationProtectionPolicies(res)
        return nil
    }
    res["microsoftStoreForBusinessLanguage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMicrosoftStoreForBusinessLanguage(val)
        return nil
    }
    res["microsoftStoreForBusinessLastCompletedApplicationSyncTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(val)
        return nil
    }
    res["microsoftStoreForBusinessLastSuccessfulSyncDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(val)
        return nil
    }
    res["mobileAppCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileAppCategory() })
        if err != nil {
            return err
        }
        res := make([]MobileAppCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileAppCategory))
        }
        m.SetMobileAppCategories(res)
        return nil
    }
    res["mobileAppConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceMobileAppConfiguration() })
        if err != nil {
            return err
        }
        res := make([]ManagedDeviceMobileAppConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDeviceMobileAppConfiguration))
        }
        m.SetMobileAppConfigurations(res)
        return nil
    }
    res["mobileApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileApp() })
        if err != nil {
            return err
        }
        res := make([]MobileApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileApp))
        }
        m.SetMobileApps(res)
        return nil
    }
    res["targetedManagedAppConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetedManagedAppConfiguration() })
        if err != nil {
            return err
        }
        res := make([]TargetedManagedAppConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*TargetedManagedAppConfiguration))
        }
        m.SetTargetedManagedAppConfigurations(res)
        return nil
    }
    res["vppTokens"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewVppToken() })
        if err != nil {
            return err
        }
        res := make([]VppToken, len(val))
        for i, v := range val {
            res[i] = *(v.(*VppToken))
        }
        m.SetVppTokens(res)
        return nil
    }
    res["windowsInformationProtectionPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionPolicy() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionPolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionPolicy))
        }
        m.SetWindowsInformationProtectionPolicies(res)
        return nil
    }
    return res
}
func (m *DeviceAppManagement) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the androidManagedAppProtections property value. Android managed app policies.
// Parameters:
//  - value : Value to set for the androidManagedAppProtections property.
func (m *DeviceAppManagement) SetAndroidManagedAppProtections(value []AndroidManagedAppProtection)() {
    m.androidManagedAppProtections = value
}
// Sets the defaultManagedAppProtections property value. Default managed app policies.
// Parameters:
//  - value : Value to set for the defaultManagedAppProtections property.
func (m *DeviceAppManagement) SetDefaultManagedAppProtections(value []DefaultManagedAppProtection)() {
    m.defaultManagedAppProtections = value
}
// Sets the iosManagedAppProtections property value. iOS managed app policies.
// Parameters:
//  - value : Value to set for the iosManagedAppProtections property.
func (m *DeviceAppManagement) SetIosManagedAppProtections(value []IosManagedAppProtection)() {
    m.iosManagedAppProtections = value
}
// Sets the isEnabledForMicrosoftStoreForBusiness property value. Whether the account is enabled for syncing applications from the Microsoft Store for Business.
// Parameters:
//  - value : Value to set for the isEnabledForMicrosoftStoreForBusiness property.
func (m *DeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(value *bool)() {
    m.isEnabledForMicrosoftStoreForBusiness = value
}
// Sets the managedAppPolicies property value. Managed app policies.
// Parameters:
//  - value : Value to set for the managedAppPolicies property.
func (m *DeviceAppManagement) SetManagedAppPolicies(value []ManagedAppPolicy)() {
    m.managedAppPolicies = value
}
// Sets the managedAppRegistrations property value. The managed app registrations.
// Parameters:
//  - value : Value to set for the managedAppRegistrations property.
func (m *DeviceAppManagement) SetManagedAppRegistrations(value []ManagedAppRegistration)() {
    m.managedAppRegistrations = value
}
// Sets the managedAppStatuses property value. The managed app statuses.
// Parameters:
//  - value : Value to set for the managedAppStatuses property.
func (m *DeviceAppManagement) SetManagedAppStatuses(value []ManagedAppStatus)() {
    m.managedAppStatuses = value
}
// Sets the managedEBooks property value. The Managed eBook.
// Parameters:
//  - value : Value to set for the managedEBooks property.
func (m *DeviceAppManagement) SetManagedEBooks(value []ManagedEBook)() {
    m.managedEBooks = value
}
// Sets the mdmWindowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are MDM enrolled.
// Parameters:
//  - value : Value to set for the mdmWindowsInformationProtectionPolicies property.
func (m *DeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicy)() {
    m.mdmWindowsInformationProtectionPolicies = value
}
// Sets the microsoftStoreForBusinessLanguage property value. The locale information used to sync applications from the Microsoft Store for Business. Cultures that are specific to a country/region. The names of these cultures follow RFC 4646 (Windows Vista and later). The format is -<country/regioncode2>, where  is a lowercase two-letter code derived from ISO 639-1 and <country/regioncode2> is an uppercase two-letter code derived from ISO 3166. For example, en-US for English (United States) is a specific culture.
// Parameters:
//  - value : Value to set for the microsoftStoreForBusinessLanguage property.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(value *string)() {
    m.microsoftStoreForBusinessLanguage = value
}
// Sets the microsoftStoreForBusinessLastCompletedApplicationSyncTime property value. The last time an application sync from the Microsoft Store for Business was completed.
// Parameters:
//  - value : Value to set for the microsoftStoreForBusinessLastCompletedApplicationSyncTime property.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastCompletedApplicationSyncTime = value
}
// Sets the microsoftStoreForBusinessLastSuccessfulSyncDateTime property value. The last time the apps from the Microsoft Store for Business were synced successfully for the account.
// Parameters:
//  - value : Value to set for the microsoftStoreForBusinessLastSuccessfulSyncDateTime property.
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastSuccessfulSyncDateTime = value
}
// Sets the mobileAppCategories property value. The mobile app categories.
// Parameters:
//  - value : Value to set for the mobileAppCategories property.
func (m *DeviceAppManagement) SetMobileAppCategories(value []MobileAppCategory)() {
    m.mobileAppCategories = value
}
// Sets the mobileAppConfigurations property value. The Managed Device Mobile Application Configurations.
// Parameters:
//  - value : Value to set for the mobileAppConfigurations property.
func (m *DeviceAppManagement) SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfiguration)() {
    m.mobileAppConfigurations = value
}
// Sets the mobileApps property value. The mobile apps.
// Parameters:
//  - value : Value to set for the mobileApps property.
func (m *DeviceAppManagement) SetMobileApps(value []MobileApp)() {
    m.mobileApps = value
}
// Sets the targetedManagedAppConfigurations property value. Targeted managed app configurations.
// Parameters:
//  - value : Value to set for the targetedManagedAppConfigurations property.
func (m *DeviceAppManagement) SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfiguration)() {
    m.targetedManagedAppConfigurations = value
}
// Sets the vppTokens property value. List of Vpp tokens for this organization.
// Parameters:
//  - value : Value to set for the vppTokens property.
func (m *DeviceAppManagement) SetVppTokens(value []VppToken)() {
    m.vppTokens = value
}
// Sets the windowsInformationProtectionPolicies property value. Windows information protection for apps running on devices which are not MDM enrolled.
// Parameters:
//  - value : Value to set for the windowsInformationProtectionPolicies property.
func (m *DeviceAppManagement) SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicy)() {
    m.windowsInformationProtectionPolicies = value
}
