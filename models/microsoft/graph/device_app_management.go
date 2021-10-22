package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceAppManagement struct {
    Entity
    androidManagedAppProtections []AndroidManagedAppProtection;
    defaultManagedAppProtections []DefaultManagedAppProtection;
    iosManagedAppProtections []IosManagedAppProtection;
    isEnabledForMicrosoftStoreForBusiness *bool;
    managedAppPolicies []ManagedAppPolicy;
    managedAppRegistrations []ManagedAppRegistration;
    managedAppStatuses []ManagedAppStatus;
    managedEBooks []ManagedEBook;
    mdmWindowsInformationProtectionPolicies []MdmWindowsInformationProtectionPolicy;
    microsoftStoreForBusinessLanguage *string;
    microsoftStoreForBusinessLastCompletedApplicationSyncTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    microsoftStoreForBusinessLastSuccessfulSyncDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    mobileAppCategories []MobileAppCategory;
    mobileAppConfigurations []ManagedDeviceMobileAppConfiguration;
    mobileApps []MobileApp;
    targetedManagedAppConfigurations []TargetedManagedAppConfiguration;
    vppTokens []VppToken;
    windowsInformationProtectionPolicies []WindowsInformationProtectionPolicy;
}
func NewDeviceAppManagement()(*DeviceAppManagement) {
    m := &DeviceAppManagement{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceAppManagement) GetAndroidManagedAppProtections()([]AndroidManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.androidManagedAppProtections
    }
}
func (m *DeviceAppManagement) GetDefaultManagedAppProtections()([]DefaultManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.defaultManagedAppProtections
    }
}
func (m *DeviceAppManagement) GetIosManagedAppProtections()([]IosManagedAppProtection) {
    if m == nil {
        return nil
    } else {
        return m.iosManagedAppProtections
    }
}
func (m *DeviceAppManagement) GetIsEnabledForMicrosoftStoreForBusiness()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isEnabledForMicrosoftStoreForBusiness
    }
}
func (m *DeviceAppManagement) GetManagedAppPolicies()([]ManagedAppPolicy) {
    if m == nil {
        return nil
    } else {
        return m.managedAppPolicies
    }
}
func (m *DeviceAppManagement) GetManagedAppRegistrations()([]ManagedAppRegistration) {
    if m == nil {
        return nil
    } else {
        return m.managedAppRegistrations
    }
}
func (m *DeviceAppManagement) GetManagedAppStatuses()([]ManagedAppStatus) {
    if m == nil {
        return nil
    } else {
        return m.managedAppStatuses
    }
}
func (m *DeviceAppManagement) GetManagedEBooks()([]ManagedEBook) {
    if m == nil {
        return nil
    } else {
        return m.managedEBooks
    }
}
func (m *DeviceAppManagement) GetMdmWindowsInformationProtectionPolicies()([]MdmWindowsInformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.mdmWindowsInformationProtectionPolicies
    }
}
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLanguage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLanguage
    }
}
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastCompletedApplicationSyncTime
    }
}
func (m *DeviceAppManagement) GetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.microsoftStoreForBusinessLastSuccessfulSyncDateTime
    }
}
func (m *DeviceAppManagement) GetMobileAppCategories()([]MobileAppCategory) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppCategories
    }
}
func (m *DeviceAppManagement) GetMobileAppConfigurations()([]ManagedDeviceMobileAppConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.mobileAppConfigurations
    }
}
func (m *DeviceAppManagement) GetMobileApps()([]MobileApp) {
    if m == nil {
        return nil
    } else {
        return m.mobileApps
    }
}
func (m *DeviceAppManagement) GetTargetedManagedAppConfigurations()([]TargetedManagedAppConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.targetedManagedAppConfigurations
    }
}
func (m *DeviceAppManagement) GetVppTokens()([]VppToken) {
    if m == nil {
        return nil
    } else {
        return m.vppTokens
    }
}
func (m *DeviceAppManagement) GetWindowsInformationProtectionPolicies()([]WindowsInformationProtectionPolicy) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionPolicies
    }
}
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
func (m *DeviceAppManagement) SetAndroidManagedAppProtections(value []AndroidManagedAppProtection)() {
    m.androidManagedAppProtections = value
}
func (m *DeviceAppManagement) SetDefaultManagedAppProtections(value []DefaultManagedAppProtection)() {
    m.defaultManagedAppProtections = value
}
func (m *DeviceAppManagement) SetIosManagedAppProtections(value []IosManagedAppProtection)() {
    m.iosManagedAppProtections = value
}
func (m *DeviceAppManagement) SetIsEnabledForMicrosoftStoreForBusiness(value *bool)() {
    m.isEnabledForMicrosoftStoreForBusiness = value
}
func (m *DeviceAppManagement) SetManagedAppPolicies(value []ManagedAppPolicy)() {
    m.managedAppPolicies = value
}
func (m *DeviceAppManagement) SetManagedAppRegistrations(value []ManagedAppRegistration)() {
    m.managedAppRegistrations = value
}
func (m *DeviceAppManagement) SetManagedAppStatuses(value []ManagedAppStatus)() {
    m.managedAppStatuses = value
}
func (m *DeviceAppManagement) SetManagedEBooks(value []ManagedEBook)() {
    m.managedEBooks = value
}
func (m *DeviceAppManagement) SetMdmWindowsInformationProtectionPolicies(value []MdmWindowsInformationProtectionPolicy)() {
    m.mdmWindowsInformationProtectionPolicies = value
}
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLanguage(value *string)() {
    m.microsoftStoreForBusinessLanguage = value
}
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastCompletedApplicationSyncTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastCompletedApplicationSyncTime = value
}
func (m *DeviceAppManagement) SetMicrosoftStoreForBusinessLastSuccessfulSyncDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.microsoftStoreForBusinessLastSuccessfulSyncDateTime = value
}
func (m *DeviceAppManagement) SetMobileAppCategories(value []MobileAppCategory)() {
    m.mobileAppCategories = value
}
func (m *DeviceAppManagement) SetMobileAppConfigurations(value []ManagedDeviceMobileAppConfiguration)() {
    m.mobileAppConfigurations = value
}
func (m *DeviceAppManagement) SetMobileApps(value []MobileApp)() {
    m.mobileApps = value
}
func (m *DeviceAppManagement) SetTargetedManagedAppConfigurations(value []TargetedManagedAppConfiguration)() {
    m.targetedManagedAppConfigurations = value
}
func (m *DeviceAppManagement) SetVppTokens(value []VppToken)() {
    m.vppTokens = value
}
func (m *DeviceAppManagement) SetWindowsInformationProtectionPolicies(value []WindowsInformationProtectionPolicy)() {
    m.windowsInformationProtectionPolicies = value
}
