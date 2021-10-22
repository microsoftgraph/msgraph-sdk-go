package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceManagement struct {
    Entity
    applePushNotificationCertificate *ApplePushNotificationCertificate;
    complianceManagementPartners []ComplianceManagementPartner;
    conditionalAccessSettings *OnPremisesConditionalAccessSettings;
    detectedApps []DetectedApp;
    deviceCategories []DeviceCategory;
    deviceCompliancePolicies []DeviceCompliancePolicy;
    deviceCompliancePolicyDeviceStateSummary *DeviceCompliancePolicyDeviceStateSummary;
    deviceCompliancePolicySettingStateSummaries []DeviceCompliancePolicySettingStateSummary;
    deviceConfigurationDeviceStateSummaries *DeviceConfigurationDeviceStateSummary;
    deviceConfigurations []DeviceConfiguration;
    deviceEnrollmentConfigurations []DeviceEnrollmentConfiguration;
    deviceManagementPartners []DeviceManagementPartner;
    exchangeConnectors []DeviceManagementExchangeConnector;
    importedWindowsAutopilotDeviceIdentities []ImportedWindowsAutopilotDeviceIdentity;
    intuneAccountId *string;
    intuneBrand *IntuneBrand;
    iosUpdateStatuses []IosUpdateDeviceStatus;
    managedDeviceOverview *ManagedDeviceOverview;
    managedDevices []ManagedDevice;
    mobileThreatDefenseConnectors []MobileThreatDefenseConnector;
    notificationMessageTemplates []NotificationMessageTemplate;
    remoteAssistancePartners []RemoteAssistancePartner;
    resourceOperations []ResourceOperation;
    roleAssignments []DeviceAndAppManagementRoleAssignment;
    roleDefinitions []RoleDefinition;
    settings *DeviceManagementSettings;
    softwareUpdateStatusSummary *SoftwareUpdateStatusSummary;
    subscriptionState *DeviceManagementSubscriptionState;
    telecomExpenseManagementPartners []TelecomExpenseManagementPartner;
    termsAndConditions []TermsAndConditions;
    troubleshootingEvents []DeviceManagementTroubleshootingEvent;
    windowsAutopilotDeviceIdentities []WindowsAutopilotDeviceIdentity;
    windowsInformationProtectionAppLearningSummaries []WindowsInformationProtectionAppLearningSummary;
    windowsInformationProtectionNetworkLearningSummaries []WindowsInformationProtectionNetworkLearningSummary;
}
func NewDeviceManagement()(*DeviceManagement) {
    m := &DeviceManagement{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DeviceManagement) GetApplePushNotificationCertificate()(*ApplePushNotificationCertificate) {
    if m == nil {
        return nil
    } else {
        return m.applePushNotificationCertificate
    }
}
func (m *DeviceManagement) GetComplianceManagementPartners()([]ComplianceManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.complianceManagementPartners
    }
}
func (m *DeviceManagement) GetConditionalAccessSettings()(*OnPremisesConditionalAccessSettings) {
    if m == nil {
        return nil
    } else {
        return m.conditionalAccessSettings
    }
}
func (m *DeviceManagement) GetDetectedApps()([]DetectedApp) {
    if m == nil {
        return nil
    } else {
        return m.detectedApps
    }
}
func (m *DeviceManagement) GetDeviceCategories()([]DeviceCategory) {
    if m == nil {
        return nil
    } else {
        return m.deviceCategories
    }
}
func (m *DeviceManagement) GetDeviceCompliancePolicies()([]DeviceCompliancePolicy) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicies
    }
}
func (m *DeviceManagement) GetDeviceCompliancePolicyDeviceStateSummary()(*DeviceCompliancePolicyDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicyDeviceStateSummary
    }
}
func (m *DeviceManagement) GetDeviceCompliancePolicySettingStateSummaries()([]DeviceCompliancePolicySettingStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceCompliancePolicySettingStateSummaries
    }
}
func (m *DeviceManagement) GetDeviceConfigurationDeviceStateSummaries()(*DeviceConfigurationDeviceStateSummary) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurationDeviceStateSummaries
    }
}
func (m *DeviceManagement) GetDeviceConfigurations()([]DeviceConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceConfigurations
    }
}
func (m *DeviceManagement) GetDeviceEnrollmentConfigurations()([]DeviceEnrollmentConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.deviceEnrollmentConfigurations
    }
}
func (m *DeviceManagement) GetDeviceManagementPartners()([]DeviceManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.deviceManagementPartners
    }
}
func (m *DeviceManagement) GetExchangeConnectors()([]DeviceManagementExchangeConnector) {
    if m == nil {
        return nil
    } else {
        return m.exchangeConnectors
    }
}
func (m *DeviceManagement) GetImportedWindowsAutopilotDeviceIdentities()([]ImportedWindowsAutopilotDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.importedWindowsAutopilotDeviceIdentities
    }
}
func (m *DeviceManagement) GetIntuneAccountId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.intuneAccountId
    }
}
func (m *DeviceManagement) GetIntuneBrand()(*IntuneBrand) {
    if m == nil {
        return nil
    } else {
        return m.intuneBrand
    }
}
func (m *DeviceManagement) GetIosUpdateStatuses()([]IosUpdateDeviceStatus) {
    if m == nil {
        return nil
    } else {
        return m.iosUpdateStatuses
    }
}
func (m *DeviceManagement) GetManagedDeviceOverview()(*ManagedDeviceOverview) {
    if m == nil {
        return nil
    } else {
        return m.managedDeviceOverview
    }
}
func (m *DeviceManagement) GetManagedDevices()([]ManagedDevice) {
    if m == nil {
        return nil
    } else {
        return m.managedDevices
    }
}
func (m *DeviceManagement) GetMobileThreatDefenseConnectors()([]MobileThreatDefenseConnector) {
    if m == nil {
        return nil
    } else {
        return m.mobileThreatDefenseConnectors
    }
}
func (m *DeviceManagement) GetNotificationMessageTemplates()([]NotificationMessageTemplate) {
    if m == nil {
        return nil
    } else {
        return m.notificationMessageTemplates
    }
}
func (m *DeviceManagement) GetRemoteAssistancePartners()([]RemoteAssistancePartner) {
    if m == nil {
        return nil
    } else {
        return m.remoteAssistancePartners
    }
}
func (m *DeviceManagement) GetResourceOperations()([]ResourceOperation) {
    if m == nil {
        return nil
    } else {
        return m.resourceOperations
    }
}
func (m *DeviceManagement) GetRoleAssignments()([]DeviceAndAppManagementRoleAssignment) {
    if m == nil {
        return nil
    } else {
        return m.roleAssignments
    }
}
func (m *DeviceManagement) GetRoleDefinitions()([]RoleDefinition) {
    if m == nil {
        return nil
    } else {
        return m.roleDefinitions
    }
}
func (m *DeviceManagement) GetSettings()(*DeviceManagementSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *DeviceManagement) GetSoftwareUpdateStatusSummary()(*SoftwareUpdateStatusSummary) {
    if m == nil {
        return nil
    } else {
        return m.softwareUpdateStatusSummary
    }
}
func (m *DeviceManagement) GetSubscriptionState()(*DeviceManagementSubscriptionState) {
    if m == nil {
        return nil
    } else {
        return m.subscriptionState
    }
}
func (m *DeviceManagement) GetTelecomExpenseManagementPartners()([]TelecomExpenseManagementPartner) {
    if m == nil {
        return nil
    } else {
        return m.telecomExpenseManagementPartners
    }
}
func (m *DeviceManagement) GetTermsAndConditions()([]TermsAndConditions) {
    if m == nil {
        return nil
    } else {
        return m.termsAndConditions
    }
}
func (m *DeviceManagement) GetTroubleshootingEvents()([]DeviceManagementTroubleshootingEvent) {
    if m == nil {
        return nil
    } else {
        return m.troubleshootingEvents
    }
}
func (m *DeviceManagement) GetWindowsAutopilotDeviceIdentities()([]WindowsAutopilotDeviceIdentity) {
    if m == nil {
        return nil
    } else {
        return m.windowsAutopilotDeviceIdentities
    }
}
func (m *DeviceManagement) GetWindowsInformationProtectionAppLearningSummaries()([]WindowsInformationProtectionAppLearningSummary) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionAppLearningSummaries
    }
}
func (m *DeviceManagement) GetWindowsInformationProtectionNetworkLearningSummaries()([]WindowsInformationProtectionNetworkLearningSummary) {
    if m == nil {
        return nil
    } else {
        return m.windowsInformationProtectionNetworkLearningSummaries
    }
}
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
func (m *DeviceManagement) SetApplePushNotificationCertificate(value *ApplePushNotificationCertificate)() {
    m.applePushNotificationCertificate = value
}
func (m *DeviceManagement) SetComplianceManagementPartners(value []ComplianceManagementPartner)() {
    m.complianceManagementPartners = value
}
func (m *DeviceManagement) SetConditionalAccessSettings(value *OnPremisesConditionalAccessSettings)() {
    m.conditionalAccessSettings = value
}
func (m *DeviceManagement) SetDetectedApps(value []DetectedApp)() {
    m.detectedApps = value
}
func (m *DeviceManagement) SetDeviceCategories(value []DeviceCategory)() {
    m.deviceCategories = value
}
func (m *DeviceManagement) SetDeviceCompliancePolicies(value []DeviceCompliancePolicy)() {
    m.deviceCompliancePolicies = value
}
func (m *DeviceManagement) SetDeviceCompliancePolicyDeviceStateSummary(value *DeviceCompliancePolicyDeviceStateSummary)() {
    m.deviceCompliancePolicyDeviceStateSummary = value
}
func (m *DeviceManagement) SetDeviceCompliancePolicySettingStateSummaries(value []DeviceCompliancePolicySettingStateSummary)() {
    m.deviceCompliancePolicySettingStateSummaries = value
}
func (m *DeviceManagement) SetDeviceConfigurationDeviceStateSummaries(value *DeviceConfigurationDeviceStateSummary)() {
    m.deviceConfigurationDeviceStateSummaries = value
}
func (m *DeviceManagement) SetDeviceConfigurations(value []DeviceConfiguration)() {
    m.deviceConfigurations = value
}
func (m *DeviceManagement) SetDeviceEnrollmentConfigurations(value []DeviceEnrollmentConfiguration)() {
    m.deviceEnrollmentConfigurations = value
}
func (m *DeviceManagement) SetDeviceManagementPartners(value []DeviceManagementPartner)() {
    m.deviceManagementPartners = value
}
func (m *DeviceManagement) SetExchangeConnectors(value []DeviceManagementExchangeConnector)() {
    m.exchangeConnectors = value
}
func (m *DeviceManagement) SetImportedWindowsAutopilotDeviceIdentities(value []ImportedWindowsAutopilotDeviceIdentity)() {
    m.importedWindowsAutopilotDeviceIdentities = value
}
func (m *DeviceManagement) SetIntuneAccountId(value *string)() {
    m.intuneAccountId = value
}
func (m *DeviceManagement) SetIntuneBrand(value *IntuneBrand)() {
    m.intuneBrand = value
}
func (m *DeviceManagement) SetIosUpdateStatuses(value []IosUpdateDeviceStatus)() {
    m.iosUpdateStatuses = value
}
func (m *DeviceManagement) SetManagedDeviceOverview(value *ManagedDeviceOverview)() {
    m.managedDeviceOverview = value
}
func (m *DeviceManagement) SetManagedDevices(value []ManagedDevice)() {
    m.managedDevices = value
}
func (m *DeviceManagement) SetMobileThreatDefenseConnectors(value []MobileThreatDefenseConnector)() {
    m.mobileThreatDefenseConnectors = value
}
func (m *DeviceManagement) SetNotificationMessageTemplates(value []NotificationMessageTemplate)() {
    m.notificationMessageTemplates = value
}
func (m *DeviceManagement) SetRemoteAssistancePartners(value []RemoteAssistancePartner)() {
    m.remoteAssistancePartners = value
}
func (m *DeviceManagement) SetResourceOperations(value []ResourceOperation)() {
    m.resourceOperations = value
}
func (m *DeviceManagement) SetRoleAssignments(value []DeviceAndAppManagementRoleAssignment)() {
    m.roleAssignments = value
}
func (m *DeviceManagement) SetRoleDefinitions(value []RoleDefinition)() {
    m.roleDefinitions = value
}
func (m *DeviceManagement) SetSettings(value *DeviceManagementSettings)() {
    m.settings = value
}
func (m *DeviceManagement) SetSoftwareUpdateStatusSummary(value *SoftwareUpdateStatusSummary)() {
    m.softwareUpdateStatusSummary = value
}
func (m *DeviceManagement) SetSubscriptionState(value *DeviceManagementSubscriptionState)() {
    m.subscriptionState = value
}
func (m *DeviceManagement) SetTelecomExpenseManagementPartners(value []TelecomExpenseManagementPartner)() {
    m.telecomExpenseManagementPartners = value
}
func (m *DeviceManagement) SetTermsAndConditions(value []TermsAndConditions)() {
    m.termsAndConditions = value
}
func (m *DeviceManagement) SetTroubleshootingEvents(value []DeviceManagementTroubleshootingEvent)() {
    m.troubleshootingEvents = value
}
func (m *DeviceManagement) SetWindowsAutopilotDeviceIdentities(value []WindowsAutopilotDeviceIdentity)() {
    m.windowsAutopilotDeviceIdentities = value
}
func (m *DeviceManagement) SetWindowsInformationProtectionAppLearningSummaries(value []WindowsInformationProtectionAppLearningSummary)() {
    m.windowsInformationProtectionAppLearningSummaries = value
}
func (m *DeviceManagement) SetWindowsInformationProtectionNetworkLearningSummaries(value []WindowsInformationProtectionNetworkLearningSummary)() {
    m.windowsInformationProtectionNetworkLearningSummaries = value
}
