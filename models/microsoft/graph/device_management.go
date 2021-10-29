package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceManagement struct {
    Entity
    // Apple push notification certificate.
    applePushNotificationCertificate *ApplePushNotificationCertificate;
    // The list of Compliance Management Partners configured by the tenant.
    complianceManagementPartners []ComplianceManagementPartner;
    // The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
    conditionalAccessSettings *OnPremisesConditionalAccessSettings;
    // The list of detected apps associated with a device.
    detectedApps []DetectedApp;
    // The list of device categories with the tenant.
    deviceCategories []DeviceCategory;
    // The device compliance policies.
    deviceCompliancePolicies []DeviceCompliancePolicy;
    // The device compliance state summary for this account.
    deviceCompliancePolicyDeviceStateSummary *DeviceCompliancePolicyDeviceStateSummary;
    // The summary states of compliance policy settings for this account.
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummary;
    // The device configuration device state summary for this account.
    deviceConfigurationDeviceStateSummaries *DeviceConfigurationDeviceStateSummary;
    // The device configurations.
    deviceConfigurations []DeviceConfiguration;
    // The list of device enrollment configurations
    deviceEnrollmentConfigurations []DeviceEnrollmentConfiguration;
    // The list of Device Management Partners configured by the tenant.
    deviceManagementPartners []DeviceManagementPartner;
    // The list of Exchange Connectors configured by the tenant.
    exchangeConnectors []DeviceManagementExchangeConnector;
    // Collection of imported Windows autopilot devices.
    importedWindowsAutopilotDeviceIdentities []ImportedWindowsAutopilotDeviceIdentity;
    // Intune Account Id for given tenant
    intuneAccountId *string;
    // intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
    intuneBrand *IntuneBrand;
    // The IOS software update installation statuses for this account.
    iosUpdateStatuses []IosUpdateDeviceStatus;
    // Device overview
    managedDeviceOverview *ManagedDeviceOverview;
    // The list of managed devices.
    managedDevices []ManagedDevice;
    // The list of Mobile threat Defense connectors configured by the tenant.
    mobileThreatDefenseConnectors []MobileThreatDefenseConnector;
    // The Notification Message Templates.
    notificationMessageTemplates []NotificationMessageTemplate;
    // The remote assist partners.
    remoteAssistancePartners []RemoteAssistancePartner;
    // Reports singleton
    reports *DeviceManagementReports;
    // The Resource Operations.
    resourceOperations []ResourceOperation;
    // The Role Assignments.
    roleAssignments []DeviceAndAppManagementRoleAssignment;
    // The Role Definitions.
    roleDefinitions []RoleDefinition;
    // Account level settings.
    settings *DeviceManagementSettings;
    // The software update status summary.
    softwareUpdateStatusSummary *SoftwareUpdateStatusSummary;
    // Tenant mobile device management subscription state. Possible values are: pending, active, warning, disabled, deleted, blocked, lockedOut.
    subscriptionState *DeviceManagementSubscriptionState;
    // The telecom expense management partners.
    telecomExpenseManagementPartners []TelecomExpenseManagementPartner;
    // The terms and conditions associated with device management of the company.
    termsAndConditions []TermsAndConditions;
    // The list of troubleshooting events for the tenant.
    troubleshootingEvents []DeviceManagementTroubleshootingEvent;
    // The Windows autopilot device identities contained collection.
    windowsAutopilotDeviceIdentities []WindowsAutopilotDeviceIdentity;
    // The windows information protection app learning summaries.
    windowsInformationProtectionAppLearningSummaries []WindowsInformationProtectionAppLearningSummary;
    // The windows information protection network learning summaries.
    windowsInformationProtectionNetworkLearningSummaries []WindowsInformationProtectionNetworkLearningSummary;
}
// Instantiates a new deviceManagement and sets the default values.
func NewDeviceManagement()(*DeviceManagement) {
    m := &DeviceManagement{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the applePushNotificationCertificate property value. Apple push notification certificate.
func (m *DeviceManagement) GetApplePushNotificationCertificate()(*ApplePushNotificationCertificate) {
    if m == nil {
        return nil
    } else {
        return m.applePushNotificationCertificate
    }
}
// Gets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
func (m *DeviceManagement) GetComplianceManagementPartners()([]ComplianceManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.complianceManagementPartners
    }
}
// Gets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
func (m *DeviceManagement) GetConditionalAccessSettings()(*OnPremisesConditionalAccessSettings) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessSettings
    }
}
// Gets the detectedApps property value. The list of detected apps associated with a device.
func (m *DeviceManagement) GetDetectedApps()([]DetectedApp) {
    if m == nil {
        return nil
    } else {
        return m.detectedApps
    }
}
// Gets the deviceCategories property value. The list of device categories with the tenant.
func (m *DeviceManagement) GetDeviceCategories()([]DeviceCategory) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategories
    }
}
// Gets the deviceCompliancePolicies property value. The device compliance policies.
func (m *DeviceManagement) GetDeviceCompliancePolicies()([]DeviceCompliancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicies
    }
}
// Gets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicyDeviceStateSummary
    }
}
// Gets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
func (m *DeviceManagement) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicySettingStateSummaries
    }
}
// Gets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
func (m *DeviceManagement) GetDeviceConfigurationDeviceStateSummaries()(*DeviceConfigurationDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationDeviceStateSummaries
    }
}
// Gets the deviceConfigurations property value. The device configurations.
func (m *DeviceManagement) GetDeviceConfigurations()([]DeviceConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurations
    }
}
// Gets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
func (m *DeviceManagement) GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentConfigurations
    }
}
// Gets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
func (m *DeviceManagement) GetDeviceManagementPartners()([]DeviceManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementPartners
    }
}
// Gets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
func (m *DeviceManagement) GetExchangeConnectors()([]DeviceManagementExchangeConnector) {
    if m == nil {
        return nil
    } else {
        return m.exchangeConnectors
    }
}
// Gets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
func (m *DeviceManagement) GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedWindowsAutopilotDeviceIdentities
    }
}
// Gets the intuneAccountId property value. Intune Account Id for given tenant
func (m *DeviceManagement) GetIntuneAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.intuneAccountId
    }
}
// Gets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
func (m *DeviceManagement) GetIntuneBrand()(*IntuneBrand) {
    if m == nil {
        return nil
    } else {
        return m.intuneBrand
    }
}
// Gets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
func (m *DeviceManagement) GetIosUpdateStatuses()([]IosUpdateDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.iosUpdateStatuses
    }
}
// Gets the managedDeviceOverview property value. Device overview
func (m *DeviceManagement) GetManagedDeviceOverview()(*ManagedDeviceOverview) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceOverview
    }
}
// Gets the managedDevices property value. The list of managed devices.
func (m *DeviceManagement) GetManagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevices
    }
}
// Gets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
func (m *DeviceManagement) GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnector) {
    if m == nil {
        return nil
    } else {
        return m.mobileThreatDefenseConnectors
    }
}
// Gets the notificationMessageTemplates property value. The Notification Message Templates.
func (m *DeviceManagement) GetNotificationMessageTemplates()([]NotificationMessageTemplate) {
    if m == nil {
        return nil
    } else {
        return m.notificationMessageTemplates
    }
}
// Gets the remoteAssistancePartners property value. The remote assist partners.
func (m *DeviceManagement) GetRemoteAssistancePartners()([]RemoteAssistancePartner) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistancePartners
    }
}
// Gets the reports property value. Reports singleton
func (m *DeviceManagement) GetReports()(*DeviceManagementReports) {
    if m == nil {
        return nil
    } else {
        return m.reports
    }
}
// Gets the resourceOperations property value. The Resource Operations.
func (m *DeviceManagement) GetResourceOperations()([]ResourceOperation) {
    if m == nil {
        return nil
    } else {
        return m.resourceOperations
    }
}
// Gets the roleAssignments property value. The Role Assignments.
func (m *DeviceManagement) GetRoleAssignments()([]DeviceAndAppManagementRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
// Gets the roleDefinitions property value. The Role Definitions.
func (m *DeviceManagement) GetRoleDefinitions()([]RoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
// Gets the settings property value. Account level settings.
func (m *DeviceManagement) GetSettings()(*DeviceManagementSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// Gets the softwareUpdateStatusSummary property value. The software update status summary.
func (m *DeviceManagement) GetSoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummary) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateStatusSummary
    }
}
// Gets the subscriptionState property value. Tenant mobile device management subscription state. Possible values are: pending, active, warning, disabled, deleted, blocked, lockedOut.
func (m *DeviceManagement) GetSubscriptionState()(*DeviceManagementSubscriptionState) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionState
    }
}
// Gets the telecomExpenseManagementPartners property value. The telecom expense management partners.
func (m *DeviceManagement) GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.telecomExpenseManagementPartners
    }
}
// Gets the termsAndConditions property value. The terms and conditions associated with device management of the company.
func (m *DeviceManagement) GetTermsAndConditions()([]TermsAndConditions) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditions
    }
}
// Gets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
func (m *DeviceManagement) GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEvent) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingEvents
    }
}
// Gets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
func (m *DeviceManagement) GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotDeviceIdentities
    }
}
// Gets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummary) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionAppLearningSummaries
    }
}
// Gets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
func (m *DeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummary) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionNetworkLearningSummaries
    }
}
// The deserialization information for the current model
func (m *DeviceManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["applePushNotificationCertificate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApplePushNotificationCertificate() })
        if err != nil {
            return err
        }
        m.SetApplePushNotificationCertificate(val.(*ApplePushNotificationCertificate))
        return nil
    }
    res["complianceManagementPartners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComplianceManagementPartner() })
        if err != nil {
            return err
        }
        res := make([]ComplianceManagementPartner, len(val))
        for i, v := range val {
            res[i] = *(v.(*ComplianceManagementPartner))
        }
        m.SetComplianceManagementPartners(res)
        return nil
    }
    res["conditionalAccessSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOnPremisesConditionalAccessSettings() })
        if err != nil {
            return err
        }
        m.SetConditionalAccessSettings(val.(*OnPremisesConditionalAccessSettings))
        return nil
    }
    res["detectedApps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDetectedApp() })
        if err != nil {
            return err
        }
        res := make([]DetectedApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*DetectedApp))
        }
        m.SetDetectedApps(res)
        return nil
    }
    res["deviceCategories"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCategory() })
        if err != nil {
            return err
        }
        res := make([]DeviceCategory, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCategory))
        }
        m.SetDeviceCategories(res)
        return nil
    }
    res["deviceCompliancePolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicy() })
        if err != nil {
            return err
        }
        res := make([]DeviceCompliancePolicy, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCompliancePolicy))
        }
        m.SetDeviceCompliancePolicies(res)
        return nil
    }
    res["deviceCompliancePolicyDeviceStateSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicyDeviceStateSummary() })
        if err != nil {
            return err
        }
        m.SetDeviceCompliancePolicyDeviceStateSummary(val.(*DeviceCompliancePolicyDeviceStateSummary))
        return nil
    }
    res["deviceCompliancePolicySettingStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceCompliancePolicySettingStateSummary() })
        if err != nil {
            return err
        }
        res := make([]DeviceCompliancePolicySettingStateSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceCompliancePolicySettingStateSummary))
        }
        m.SetDeviceCompliancePolicySettingStateSummaries(res)
        return nil
    }
    res["deviceConfigurationDeviceStateSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfigurationDeviceStateSummary() })
        if err != nil {
            return err
        }
        m.SetDeviceConfigurationDeviceStateSummaries(val.(*DeviceConfigurationDeviceStateSummary))
        return nil
    }
    res["deviceConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceConfiguration() })
        if err != nil {
            return err
        }
        res := make([]DeviceConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceConfiguration))
        }
        m.SetDeviceConfigurations(res)
        return nil
    }
    res["deviceEnrollmentConfigurations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceEnrollmentConfiguration() })
        if err != nil {
            return err
        }
        res := make([]DeviceEnrollmentConfiguration, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceEnrollmentConfiguration))
        }
        m.SetDeviceEnrollmentConfigurations(res)
        return nil
    }
    res["deviceManagementPartners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementPartner() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementPartner, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementPartner))
        }
        m.SetDeviceManagementPartners(res)
        return nil
    }
    res["exchangeConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementExchangeConnector() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementExchangeConnector, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementExchangeConnector))
        }
        m.SetExchangeConnectors(res)
        return nil
    }
    res["importedWindowsAutopilotDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewImportedWindowsAutopilotDeviceIdentity() })
        if err != nil {
            return err
        }
        res := make([]ImportedWindowsAutopilotDeviceIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*ImportedWindowsAutopilotDeviceIdentity))
        }
        m.SetImportedWindowsAutopilotDeviceIdentities(res)
        return nil
    }
    res["intuneAccountId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetIntuneAccountId(val)
        return nil
    }
    res["intuneBrand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIntuneBrand() })
        if err != nil {
            return err
        }
        m.SetIntuneBrand(val.(*IntuneBrand))
        return nil
    }
    res["iosUpdateStatuses"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIosUpdateDeviceStatus() })
        if err != nil {
            return err
        }
        res := make([]IosUpdateDeviceStatus, len(val))
        for i, v := range val {
            res[i] = *(v.(*IosUpdateDeviceStatus))
        }
        m.SetIosUpdateStatuses(res)
        return nil
    }
    res["managedDeviceOverview"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceOverview() })
        if err != nil {
            return err
        }
        m.SetManagedDeviceOverview(val.(*ManagedDeviceOverview))
        return nil
    }
    res["managedDevices"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDevice() })
        if err != nil {
            return err
        }
        res := make([]ManagedDevice, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedDevice))
        }
        m.SetManagedDevices(res)
        return nil
    }
    res["mobileThreatDefenseConnectors"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewMobileThreatDefenseConnector() })
        if err != nil {
            return err
        }
        res := make([]MobileThreatDefenseConnector, len(val))
        for i, v := range val {
            res[i] = *(v.(*MobileThreatDefenseConnector))
        }
        m.SetMobileThreatDefenseConnectors(res)
        return nil
    }
    res["notificationMessageTemplates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewNotificationMessageTemplate() })
        if err != nil {
            return err
        }
        res := make([]NotificationMessageTemplate, len(val))
        for i, v := range val {
            res[i] = *(v.(*NotificationMessageTemplate))
        }
        m.SetNotificationMessageTemplates(res)
        return nil
    }
    res["remoteAssistancePartners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRemoteAssistancePartner() })
        if err != nil {
            return err
        }
        res := make([]RemoteAssistancePartner, len(val))
        for i, v := range val {
            res[i] = *(v.(*RemoteAssistancePartner))
        }
        m.SetRemoteAssistancePartners(res)
        return nil
    }
    res["reports"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementReports() })
        if err != nil {
            return err
        }
        m.SetReports(val.(*DeviceManagementReports))
        return nil
    }
    res["resourceOperations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewResourceOperation() })
        if err != nil {
            return err
        }
        res := make([]ResourceOperation, len(val))
        for i, v := range val {
            res[i] = *(v.(*ResourceOperation))
        }
        m.SetResourceOperations(res)
        return nil
    }
    res["roleAssignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceAndAppManagementRoleAssignment() })
        if err != nil {
            return err
        }
        res := make([]DeviceAndAppManagementRoleAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceAndAppManagementRoleAssignment))
        }
        m.SetRoleAssignments(res)
        return nil
    }
    res["roleDefinitions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRoleDefinition() })
        if err != nil {
            return err
        }
        res := make([]RoleDefinition, len(val))
        for i, v := range val {
            res[i] = *(v.(*RoleDefinition))
        }
        m.SetRoleDefinitions(res)
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementSettings() })
        if err != nil {
            return err
        }
        m.SetSettings(val.(*DeviceManagementSettings))
        return nil
    }
    res["softwareUpdateStatusSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSoftwareUpdateStatusSummary() })
        if err != nil {
            return err
        }
        m.SetSoftwareUpdateStatusSummary(val.(*SoftwareUpdateStatusSummary))
        return nil
    }
    res["subscriptionState"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceManagementSubscriptionState)
        if err != nil {
            return err
        }
        cast := val.(DeviceManagementSubscriptionState)
        m.SetSubscriptionState(&cast)
        return nil
    }
    res["telecomExpenseManagementPartners"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTelecomExpenseManagementPartner() })
        if err != nil {
            return err
        }
        res := make([]TelecomExpenseManagementPartner, len(val))
        for i, v := range val {
            res[i] = *(v.(*TelecomExpenseManagementPartner))
        }
        m.SetTelecomExpenseManagementPartners(res)
        return nil
    }
    res["termsAndConditions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTermsAndConditions() })
        if err != nil {
            return err
        }
        res := make([]TermsAndConditions, len(val))
        for i, v := range val {
            res[i] = *(v.(*TermsAndConditions))
        }
        m.SetTermsAndConditions(res)
        return nil
    }
    res["troubleshootingEvents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDeviceManagementTroubleshootingEvent() })
        if err != nil {
            return err
        }
        res := make([]DeviceManagementTroubleshootingEvent, len(val))
        for i, v := range val {
            res[i] = *(v.(*DeviceManagementTroubleshootingEvent))
        }
        m.SetTroubleshootingEvents(res)
        return nil
    }
    res["windowsAutopilotDeviceIdentities"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsAutopilotDeviceIdentity() })
        if err != nil {
            return err
        }
        res := make([]WindowsAutopilotDeviceIdentity, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsAutopilotDeviceIdentity))
        }
        m.SetWindowsAutopilotDeviceIdentities(res)
        return nil
    }
    res["windowsInformationProtectionAppLearningSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionAppLearningSummary() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionAppLearningSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionAppLearningSummary))
        }
        m.SetWindowsInformationProtectionAppLearningSummaries(res)
        return nil
    }
    res["windowsInformationProtectionNetworkLearningSummaries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewWindowsInformationProtectionNetworkLearningSummary() })
        if err != nil {
            return err
        }
        res := make([]WindowsInformationProtectionNetworkLearningSummary, len(val))
        for i, v := range val {
            res[i] = *(v.(*WindowsInformationProtectionNetworkLearningSummary))
        }
        m.SetWindowsInformationProtectionNetworkLearningSummaries(res)
        return nil
    }
    return res
}
func (m *DeviceManagement) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DeviceManagement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("applePushNotificationCertificate", m.GetApplePushNotificationCertificate())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComplianceManagementPartners()))
        for i, v := range m.GetComplianceManagementPartners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("complianceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("conditionalAccessSettings", m.GetConditionalAccessSettings())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDetectedApps()))
        for i, v := range m.GetDetectedApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("detectedApps", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCategories()))
        for i, v := range m.GetDeviceCategories() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCategories", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCompliancePolicies()))
        for i, v := range m.GetDeviceCompliancePolicies() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicies", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceCompliancePolicyDeviceStateSummary", m.GetDeviceCompliancePolicyDeviceStateSummary())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceCompliancePolicySettingStateSummaries()))
        for i, v := range m.GetDeviceCompliancePolicySettingStateSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceCompliancePolicySettingStateSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deviceConfigurationDeviceStateSummaries", m.GetDeviceConfigurationDeviceStateSummaries())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceConfigurations()))
        for i, v := range m.GetDeviceConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceEnrollmentConfigurations()))
        for i, v := range m.GetDeviceEnrollmentConfigurations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceEnrollmentConfigurations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDeviceManagementPartners()))
        for i, v := range m.GetDeviceManagementPartners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("deviceManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExchangeConnectors()))
        for i, v := range m.GetExchangeConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("exchangeConnectors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetImportedWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetImportedWindowsAutopilotDeviceIdentities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("importedWindowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("intuneAccountId", m.GetIntuneAccountId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("intuneBrand", m.GetIntuneBrand())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetIosUpdateStatuses()))
        for i, v := range m.GetIosUpdateStatuses() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("iosUpdateStatuses", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("managedDeviceOverview", m.GetManagedDeviceOverview())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetManagedDevices()))
        for i, v := range m.GetManagedDevices() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("managedDevices", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetMobileThreatDefenseConnectors()))
        for i, v := range m.GetMobileThreatDefenseConnectors() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("mobileThreatDefenseConnectors", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetNotificationMessageTemplates()))
        for i, v := range m.GetNotificationMessageTemplates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("notificationMessageTemplates", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRemoteAssistancePartners()))
        for i, v := range m.GetRemoteAssistancePartners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("remoteAssistancePartners", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reports", m.GetReports())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetResourceOperations()))
        for i, v := range m.GetResourceOperations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("resourceOperations", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleAssignments()))
        for i, v := range m.GetRoleAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleAssignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetRoleDefinitions()))
        for i, v := range m.GetRoleDefinitions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("roleDefinitions", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("softwareUpdateStatusSummary", m.GetSoftwareUpdateStatusSummary())
        if err != nil {
            return err
        }
    }
    if m.GetSubscriptionState() != nil {
        cast := m.GetSubscriptionState().String()
        err = writer.WriteStringValue("subscriptionState", &cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTelecomExpenseManagementPartners()))
        for i, v := range m.GetTelecomExpenseManagementPartners() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("telecomExpenseManagementPartners", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTermsAndConditions()))
        for i, v := range m.GetTermsAndConditions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("termsAndConditions", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTroubleshootingEvents()))
        for i, v := range m.GetTroubleshootingEvents() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("troubleshootingEvents", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsAutopilotDeviceIdentities()))
        for i, v := range m.GetWindowsAutopilotDeviceIdentities() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsAutopilotDeviceIdentities", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsInformationProtectionAppLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionAppLearningSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionAppLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetWindowsInformationProtectionNetworkLearningSummaries()))
        for i, v := range m.GetWindowsInformationProtectionNetworkLearningSummaries() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("windowsInformationProtectionNetworkLearningSummaries", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the applePushNotificationCertificate property value. Apple push notification certificate.
// Parameters:
//  - value : Value to set for the applePushNotificationCertificate property.
func (m *DeviceManagement) SetApplePushNotificationCertificate(value *ApplePushNotificationCertificate)() {
    m.applePushNotificationCertificate = value
}
// Sets the complianceManagementPartners property value. The list of Compliance Management Partners configured by the tenant.
// Parameters:
//  - value : Value to set for the complianceManagementPartners property.
func (m *DeviceManagement) SetComplianceManagementPartners(value []ComplianceManagementPartner)() {
    m.complianceManagementPartners = value
}
// Sets the conditionalAccessSettings property value. The Exchange on premises conditional access settings. On premises conditional access will require devices to be both enrolled and compliant for mail access
// Parameters:
//  - value : Value to set for the conditionalAccessSettings property.
func (m *DeviceManagement) SetConditionalAccessSettings(value *OnPremisesConditionalAccessSettings)() {
    m.conditionalAccessSettings = value
}
// Sets the detectedApps property value. The list of detected apps associated with a device.
// Parameters:
//  - value : Value to set for the detectedApps property.
func (m *DeviceManagement) SetDetectedApps(value []DetectedApp)() {
    m.detectedApps = value
}
// Sets the deviceCategories property value. The list of device categories with the tenant.
// Parameters:
//  - value : Value to set for the deviceCategories property.
func (m *DeviceManagement) SetDeviceCategories(value []DeviceCategory)() {
    m.deviceCategories = value
}
// Sets the deviceCompliancePolicies property value. The device compliance policies.
// Parameters:
//  - value : Value to set for the deviceCompliancePolicies property.
func (m *DeviceManagement) SetDeviceCompliancePolicies(value []DeviceCompliancePolicy)() {
    m.deviceCompliancePolicies = value
}
// Sets the deviceCompliancePolicyDeviceStateSummary property value. The device compliance state summary for this account.
// Parameters:
//  - value : Value to set for the deviceCompliancePolicyDeviceStateSummary property.
func (m *DeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(value *DeviceCompliancePolicyDeviceStateSummary)() {
    m.deviceCompliancePolicyDeviceStateSummary = value
}
// Sets the deviceCompliancePolicySettingStateSummaries property value. The summary states of compliance policy settings for this account.
// Parameters:
//  - value : Value to set for the deviceCompliancePolicySettingStateSummaries property.
func (m *DeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummary)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
// Sets the deviceConfigurationDeviceStateSummaries property value. The device configuration device state summary for this account.
// Parameters:
//  - value : Value to set for the deviceConfigurationDeviceStateSummaries property.
func (m *DeviceManagement) SetDeviceConfigurationDeviceStateSummaries(value *DeviceConfigurationDeviceStateSummary)() {
    m.deviceConfigurationDeviceStateSummaries = value
}
// Sets the deviceConfigurations property value. The device configurations.
// Parameters:
//  - value : Value to set for the deviceConfigurations property.
func (m *DeviceManagement) SetDeviceConfigurations(value []DeviceConfiguration)() {
    m.deviceConfigurations = value
}
// Sets the deviceEnrollmentConfigurations property value. The list of device enrollment configurations
// Parameters:
//  - value : Value to set for the deviceEnrollmentConfigurations property.
func (m *DeviceManagement) SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfiguration)() {
    m.deviceEnrollmentConfigurations = value
}
// Sets the deviceManagementPartners property value. The list of Device Management Partners configured by the tenant.
// Parameters:
//  - value : Value to set for the deviceManagementPartners property.
func (m *DeviceManagement) SetDeviceManagementPartners(value []DeviceManagementPartner)() {
    m.deviceManagementPartners = value
}
// Sets the exchangeConnectors property value. The list of Exchange Connectors configured by the tenant.
// Parameters:
//  - value : Value to set for the exchangeConnectors property.
func (m *DeviceManagement) SetExchangeConnectors(value []DeviceManagementExchangeConnector)() {
    m.exchangeConnectors = value
}
// Sets the importedWindowsAutopilotDeviceIdentities property value. Collection of imported Windows autopilot devices.
// Parameters:
//  - value : Value to set for the importedWindowsAutopilotDeviceIdentities property.
func (m *DeviceManagement) SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentity)() {
    m.importedWindowsAutopilotDeviceIdentities = value
}
// Sets the intuneAccountId property value. Intune Account Id for given tenant
// Parameters:
//  - value : Value to set for the intuneAccountId property.
func (m *DeviceManagement) SetIntuneAccountId(value *string)() {
    m.intuneAccountId = value
}
// Sets the intuneBrand property value. intuneBrand contains data which is used in customizing the appearance of the Company Portal applications as well as the end user web portal.
// Parameters:
//  - value : Value to set for the intuneBrand property.
func (m *DeviceManagement) SetIntuneBrand(value *IntuneBrand)() {
    m.intuneBrand = value
}
// Sets the iosUpdateStatuses property value. The IOS software update installation statuses for this account.
// Parameters:
//  - value : Value to set for the iosUpdateStatuses property.
func (m *DeviceManagement) SetIosUpdateStatuses(value []IosUpdateDeviceStatus)() {
    m.iosUpdateStatuses = value
}
// Sets the managedDeviceOverview property value. Device overview
// Parameters:
//  - value : Value to set for the managedDeviceOverview property.
func (m *DeviceManagement) SetManagedDeviceOverview(value *ManagedDeviceOverview)() {
    m.managedDeviceOverview = value
}
// Sets the managedDevices property value. The list of managed devices.
// Parameters:
//  - value : Value to set for the managedDevices property.
func (m *DeviceManagement) SetManagedDevices(value []ManagedDevice)() {
    m.managedDevices = value
}
// Sets the mobileThreatDefenseConnectors property value. The list of Mobile threat Defense connectors configured by the tenant.
// Parameters:
//  - value : Value to set for the mobileThreatDefenseConnectors property.
func (m *DeviceManagement) SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnector)() {
    m.mobileThreatDefenseConnectors = value
}
// Sets the notificationMessageTemplates property value. The Notification Message Templates.
// Parameters:
//  - value : Value to set for the notificationMessageTemplates property.
func (m *DeviceManagement) SetNotificationMessageTemplates(value []NotificationMessageTemplate)() {
    m.notificationMessageTemplates = value
}
// Sets the remoteAssistancePartners property value. The remote assist partners.
// Parameters:
//  - value : Value to set for the remoteAssistancePartners property.
func (m *DeviceManagement) SetRemoteAssistancePartners(value []RemoteAssistancePartner)() {
    m.remoteAssistancePartners = value
}
// Sets the reports property value. Reports singleton
// Parameters:
//  - value : Value to set for the reports property.
func (m *DeviceManagement) SetReports(value *DeviceManagementReports)() {
    m.reports = value
}
// Sets the resourceOperations property value. The Resource Operations.
// Parameters:
//  - value : Value to set for the resourceOperations property.
func (m *DeviceManagement) SetResourceOperations(value []ResourceOperation)() {
    m.resourceOperations = value
}
// Sets the roleAssignments property value. The Role Assignments.
// Parameters:
//  - value : Value to set for the roleAssignments property.
func (m *DeviceManagement) SetRoleAssignments(value []DeviceAndAppManagementRoleAssignment)() {
    m.roleAssignments = value
}
// Sets the roleDefinitions property value. The Role Definitions.
// Parameters:
//  - value : Value to set for the roleDefinitions property.
func (m *DeviceManagement) SetRoleDefinitions(value []RoleDefinition)() {
    m.roleDefinitions = value
}
// Sets the settings property value. Account level settings.
// Parameters:
//  - value : Value to set for the settings property.
func (m *DeviceManagement) SetSettings(value *DeviceManagementSettings)() {
    m.settings = value
}
// Sets the softwareUpdateStatusSummary property value. The software update status summary.
// Parameters:
//  - value : Value to set for the softwareUpdateStatusSummary property.
func (m *DeviceManagement) SetSoftwareUpdateStatusSummary(value *SoftwareUpdateStatusSummary)() {
    m.softwareUpdateStatusSummary = value
}
// Sets the subscriptionState property value. Tenant mobile device management subscription state. Possible values are: pending, active, warning, disabled, deleted, blocked, lockedOut.
// Parameters:
//  - value : Value to set for the subscriptionState property.
func (m *DeviceManagement) SetSubscriptionState(value *DeviceManagementSubscriptionState)() {
    m.subscriptionState = value
}
// Sets the telecomExpenseManagementPartners property value. The telecom expense management partners.
// Parameters:
//  - value : Value to set for the telecomExpenseManagementPartners property.
func (m *DeviceManagement) SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartner)() {
    m.telecomExpenseManagementPartners = value
}
// Sets the termsAndConditions property value. The terms and conditions associated with device management of the company.
// Parameters:
//  - value : Value to set for the termsAndConditions property.
func (m *DeviceManagement) SetTermsAndConditions(value []TermsAndConditions)() {
    m.termsAndConditions = value
}
// Sets the troubleshootingEvents property value. The list of troubleshooting events for the tenant.
// Parameters:
//  - value : Value to set for the troubleshootingEvents property.
func (m *DeviceManagement) SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEvent)() {
    m.troubleshootingEvents = value
}
// Sets the windowsAutopilotDeviceIdentities property value. The Windows autopilot device identities contained collection.
// Parameters:
//  - value : Value to set for the windowsAutopilotDeviceIdentities property.
func (m *DeviceManagement) SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentity)() {
    m.windowsAutopilotDeviceIdentities = value
}
// Sets the windowsInformationProtectionAppLearningSummaries property value. The windows information protection app learning summaries.
// Parameters:
//  - value : Value to set for the windowsInformationProtectionAppLearningSummaries property.
func (m *DeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummary)() {
    m.windowsInformationProtectionAppLearningSummaries = value
}
// Sets the windowsInformationProtectionNetworkLearningSummaries property value. The windows information protection network learning summaries.
// Parameters:
//  - value : Value to set for the windowsInformationProtectionNetworkLearningSummaries property.
func (m *DeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummary)() {
    m.windowsInformationProtectionNetworkLearningSummaries = value
}
