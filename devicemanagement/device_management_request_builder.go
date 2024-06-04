package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceManagementRequestBuilder provides operations to manage the deviceManagement singleton.
type DeviceManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceManagementRequestBuilderGetQueryParameters read properties and relationships of the deviceManagement object.
type DeviceManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementRequestBuilderGetQueryParameters
}
// DeviceManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ApplePushNotificationCertificate provides operations to manage the applePushNotificationCertificate property of the microsoft.graph.deviceManagement entity.
// returns a *ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ApplePushNotificationCertificate()(*ApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilder) {
    return NewApplepushnotificationcertificateApplePushNotificationCertificateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AuditEvents provides operations to manage the auditEvents property of the microsoft.graph.deviceManagement entity.
// returns a *AuditeventsAuditEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) AuditEvents()(*AuditeventsAuditEventsRequestBuilder) {
    return NewAuditeventsAuditEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ComplianceManagementPartners provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartners()(*CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) {
    return NewCompliancemanagementpartnersComplianceManagementPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConditionalAccessSettings provides operations to manage the conditionalAccessSettings property of the microsoft.graph.deviceManagement entity.
// returns a *ConditionalaccesssettingsConditionalAccessSettingsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ConditionalAccessSettings()(*ConditionalaccesssettingsConditionalAccessSettingsRequestBuilder) {
    return NewConditionalaccesssettingsConditionalAccessSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceManagementRequestBuilderInternal instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    m := &DeviceManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceManagementRequestBuilder instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// DetectedApps provides operations to manage the detectedApps property of the microsoft.graph.deviceManagement entity.
// returns a *DetectedappsDetectedAppsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DetectedApps()(*DetectedappsDetectedAppsRequestBuilder) {
    return NewDetectedappsDetectedAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCategories provides operations to manage the deviceCategories property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecategoriesDeviceCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCategories()(*DevicecategoriesDeviceCategoriesRequestBuilder) {
    return NewDevicecategoriesDeviceCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicies provides operations to manage the deviceCompliancePolicies property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicies()(*DevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilder) {
    return NewDevicecompliancepoliciesDeviceCompliancePoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicyDeviceStateSummary provides operations to manage the deviceCompliancePolicyDeviceStateSummary property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicyDeviceStateSummary()(*DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    return NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceCompliancePolicySettingStateSummaries provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurationDeviceStateSummaries provides operations to manage the deviceConfigurationDeviceStateSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurationDeviceStateSummaries()(*DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) {
    return NewDeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceConfigurations provides operations to manage the deviceConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceconfigurationsDeviceConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceConfigurations()(*DeviceconfigurationsDeviceConfigurationsRequestBuilder) {
    return NewDeviceconfigurationsDeviceConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceEnrollmentConfigurations provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurations()(*DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceManagementPartners provides operations to manage the deviceManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *DevicemanagementpartnersDeviceManagementPartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) DeviceManagementPartners()(*DevicemanagementpartnersDeviceManagementPartnersRequestBuilder) {
    return NewDevicemanagementpartnersDeviceManagementPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExchangeConnectors provides operations to manage the exchangeConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *ExchangeconnectorsExchangeConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ExchangeConnectors()(*ExchangeconnectorsExchangeConnectorsRequestBuilder) {
    return NewExchangeconnectorsExchangeConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the deviceManagement object.
// returns a DeviceManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicemanagement-get?view=graph-rest-1.0
func (m *DeviceManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementable), nil
}
// GetEffectivePermissionsWithScope provides operations to call the getEffectivePermissions method.
// returns a *GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) GetEffectivePermissionsWithScope(scope *string)(*GeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilder) {
    return NewGeteffectivepermissionswithscopeGetEffectivePermissionsWithScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, scope)
}
// ImportedWindowsAutopilotDeviceIdentities provides operations to manage the importedWindowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentities()(*ImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewImportedwindowsautopilotdeviceidentitiesImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IosUpdateStatuses provides operations to manage the iosUpdateStatuses property of the microsoft.graph.deviceManagement entity.
// returns a *IosupdatestatusesIosUpdateStatusesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) IosUpdateStatuses()(*IosupdatestatusesIosUpdateStatusesRequestBuilder) {
    return NewIosupdatestatusesIosUpdateStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDeviceOverview provides operations to manage the managedDeviceOverview property of the microsoft.graph.deviceManagement entity.
// returns a *ManageddeviceoverviewManagedDeviceOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDeviceOverview()(*ManageddeviceoverviewManagedDeviceOverviewRequestBuilder) {
    return NewManageddeviceoverviewManagedDeviceOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ManagedDevices provides operations to manage the managedDevices property of the microsoft.graph.deviceManagement entity.
// returns a *ManageddevicesManagedDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ManagedDevices()(*ManageddevicesManagedDevicesRequestBuilder) {
    return NewManageddevicesManagedDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileAppTroubleshootingEvents provides operations to manage the mobileAppTroubleshootingEvents property of the microsoft.graph.deviceManagement entity.
// returns a *MobileapptroubleshootingeventsMobileAppTroubleshootingEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MobileAppTroubleshootingEvents()(*MobileapptroubleshootingeventsMobileAppTroubleshootingEventsRequestBuilder) {
    return NewMobileapptroubleshootingeventsMobileAppTroubleshootingEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MobileThreatDefenseConnectors provides operations to manage the mobileThreatDefenseConnectors property of the microsoft.graph.deviceManagement entity.
// returns a *MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectors()(*MobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilder) {
    return NewMobilethreatdefenseconnectorsMobileThreatDefenseConnectorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NotificationMessageTemplates provides operations to manage the notificationMessageTemplates property of the microsoft.graph.deviceManagement entity.
// returns a *NotificationmessagetemplatesNotificationMessageTemplatesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplates()(*NotificationmessagetemplatesNotificationMessageTemplatesRequestBuilder) {
    return NewNotificationmessagetemplatesNotificationMessageTemplatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a deviceManagement object.
// returns a DeviceManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicemanagement-update?view=graph-rest-1.0
func (m *DeviceManagementRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementable), nil
}
// RemoteAssistancePartners provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
// returns a *RemoteassistancepartnersRemoteAssistancePartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartners()(*RemoteassistancepartnersRemoteAssistancePartnersRequestBuilder) {
    return NewRemoteassistancepartnersRemoteAssistancePartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reports provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
// returns a *ReportsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) Reports()(*ReportsRequestBuilder) {
    return NewReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourceOperations provides operations to manage the resourceOperations property of the microsoft.graph.deviceManagement entity.
// returns a *ResourceoperationsResourceOperationsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) ResourceOperations()(*ResourceoperationsResourceOperationsRequestBuilder) {
    return NewResourceoperationsResourceOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
// returns a *RoleassignmentsRoleAssignmentsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RoleAssignments()(*RoleassignmentsRoleAssignmentsRequestBuilder) {
    return NewRoleassignmentsRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.deviceManagement entity.
// returns a *RoledefinitionsRoleDefinitionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) RoleDefinitions()(*RoledefinitionsRoleDefinitionsRequestBuilder) {
    return NewRoledefinitionsRoleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SoftwareUpdateStatusSummary provides operations to manage the softwareUpdateStatusSummary property of the microsoft.graph.deviceManagement entity.
// returns a *SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) SoftwareUpdateStatusSummary()(*SoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilder) {
    return NewSoftwareupdatestatussummarySoftwareUpdateStatusSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TelecomExpenseManagementPartners provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartners()(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) {
    return NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermsAndConditions provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
// returns a *TermsandconditionsTermsAndConditionsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TermsAndConditions()(*TermsandconditionsTermsAndConditionsRequestBuilder) {
    return NewTermsandconditionsTermsAndConditionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read properties and relationships of the deviceManagement object.
// returns a *RequestInformation when successful
func (m *DeviceManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the properties of a deviceManagement object.
// returns a *RequestInformation when successful
func (m *DeviceManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementable, requestConfiguration *DeviceManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// TroubleshootingEvents provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
// returns a *TroubleshootingeventsTroubleshootingEventsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) TroubleshootingEvents()(*TroubleshootingeventsTroubleshootingEventsRequestBuilder) {
    return NewTroubleshootingeventsTroubleshootingEventsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformance provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformance()(*UserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformanceUserExperienceAnalyticsAppHealthApplicationPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetails()(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondetailsUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceId()(*UserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyappversiondeviceidUserExperienceAnalyticsAppHealthApplicationPerformanceByAppVersionDeviceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion provides operations to manage the userExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthapplicationperformancebyosversionUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersion()(*UserexperienceanalyticsapphealthapplicationperformancebyosversionUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilder) {
    return NewUserexperienceanalyticsapphealthapplicationperformancebyosversionUserExperienceAnalyticsAppHealthApplicationPerformanceByOSVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthDeviceModelPerformance provides operations to manage the userExperienceAnalyticsAppHealthDeviceModelPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDeviceModelPerformance()(*UserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsapphealthdevicemodelperformanceUserExperienceAnalyticsAppHealthDeviceModelPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthDevicePerformance provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformance()(*UserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilder) {
    return NewUserexperienceanalyticsapphealthdeviceperformanceUserExperienceAnalyticsAppHealthDevicePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthDevicePerformanceDetails provides operations to manage the userExperienceAnalyticsAppHealthDevicePerformanceDetails property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthdeviceperformancedetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthDevicePerformanceDetails()(*UserexperienceanalyticsapphealthdeviceperformancedetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilder) {
    return NewUserexperienceanalyticsapphealthdeviceperformancedetailsUserExperienceAnalyticsAppHealthDevicePerformanceDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthOSVersionPerformance provides operations to manage the userExperienceAnalyticsAppHealthOSVersionPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthosversionperformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOSVersionPerformance()(*UserexperienceanalyticsapphealthosversionperformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsapphealthosversionperformanceUserExperienceAnalyticsAppHealthOSVersionPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsAppHealthOverview provides operations to manage the userExperienceAnalyticsAppHealthOverview property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsapphealthoverviewUserExperienceAnalyticsAppHealthOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsAppHealthOverview()(*UserexperienceanalyticsapphealthoverviewUserExperienceAnalyticsAppHealthOverviewRequestBuilder) {
    return NewUserexperienceanalyticsapphealthoverviewUserExperienceAnalyticsAppHealthOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsBaselines provides operations to manage the userExperienceAnalyticsBaselines property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsBaselines()(*UserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilder) {
    return NewUserexperienceanalyticsbaselinesUserExperienceAnalyticsBaselinesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsCategories provides operations to manage the userExperienceAnalyticsCategories property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticscategoriesUserExperienceAnalyticsCategoriesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsCategories()(*UserexperienceanalyticscategoriesUserExperienceAnalyticsCategoriesRequestBuilder) {
    return NewUserexperienceanalyticscategoriesUserExperienceAnalyticsCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDevicePerformance provides operations to manage the userExperienceAnalyticsDevicePerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDevicePerformance()(*UserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilder) {
    return NewUserexperienceanalyticsdeviceperformanceUserExperienceAnalyticsDevicePerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceScores provides operations to manage the userExperienceAnalyticsDeviceScores property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicescoresUserExperienceAnalyticsDeviceScoresRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceScores()(*UserexperienceanalyticsdevicescoresUserExperienceAnalyticsDeviceScoresRequestBuilder) {
    return NewUserexperienceanalyticsdevicescoresUserExperienceAnalyticsDeviceScoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceStartupHistory provides operations to manage the userExperienceAnalyticsDeviceStartupHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicestartuphistoryUserExperienceAnalyticsDeviceStartupHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupHistory()(*UserexperienceanalyticsdevicestartuphistoryUserExperienceAnalyticsDeviceStartupHistoryRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartuphistoryUserExperienceAnalyticsDeviceStartupHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceStartupProcesses provides operations to manage the userExperienceAnalyticsDeviceStartupProcesses property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcesses()(*UserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessesUserExperienceAnalyticsDeviceStartupProcessesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsDeviceStartupProcessPerformance provides operations to manage the userExperienceAnalyticsDeviceStartupProcessPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsDeviceStartupProcessPerformance()(*UserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsdevicestartupprocessperformanceUserExperienceAnalyticsDeviceStartupProcessPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsMetricHistory provides operations to manage the userExperienceAnalyticsMetricHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsMetricHistory()(*UserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilder) {
    return NewUserexperienceanalyticsmetrichistoryUserExperienceAnalyticsMetricHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsModelScores provides operations to manage the userExperienceAnalyticsModelScores property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsModelScores()(*UserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilder) {
    return NewUserexperienceanalyticsmodelscoresUserExperienceAnalyticsModelScoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsOverview provides operations to manage the userExperienceAnalyticsOverview property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsOverview()(*UserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilder) {
    return NewUserexperienceanalyticsoverviewUserExperienceAnalyticsOverviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsScoreHistory provides operations to manage the userExperienceAnalyticsScoreHistory property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsScoreHistory()(*UserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilder) {
    return NewUserexperienceanalyticsscorehistoryUserExperienceAnalyticsScoreHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices provides operations to call the userExperienceAnalyticsSummarizeWorkFromAnywhereDevices method.
// returns a *UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsSummarizeWorkFromAnywhereDevices()(*UserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilder) {
    return NewUserexperienceanalyticssummarizeworkfromanywheredevicesUserExperienceAnalyticsSummarizeWorkFromAnywhereDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric provides operations to manage the userExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetric()(*UserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywherehardwarereadinessmetricUserExperienceAnalyticsWorkFromAnywhereHardwareReadinessMetricRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereMetrics provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereMetrics()(*UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserExperienceAnalyticsWorkFromAnywhereModelPerformance provides operations to manage the userExperienceAnalyticsWorkFromAnywhereModelPerformance property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) UserExperienceAnalyticsWorkFromAnywhereModelPerformance()(*UserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremodelperformanceUserExperienceAnalyticsWorkFromAnywhereModelPerformanceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// VerifyWindowsEnrollmentAutoDiscoveryWithDomainName provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
// returns a *VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) VerifyWindowsEnrollmentAutoDiscoveryWithDomainName(domainName *string)(*VerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    return NewVerifywindowsenrollmentautodiscoverywithdomainnameVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, domainName)
}
// VirtualEndpoint provides operations to manage the virtualEndpoint property of the microsoft.graph.deviceManagement entity.
// returns a *VirtualendpointVirtualEndpointRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) VirtualEndpoint()(*VirtualendpointVirtualEndpointRequestBuilder) {
    return NewVirtualendpointVirtualEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsAutopilotDeviceIdentities provides operations to manage the windowsAutopilotDeviceIdentities property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentities()(*WindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return NewWindowsautopilotdeviceidentitiesWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionAppLearningSummaries provides operations to manage the windowsInformationProtectionAppLearningSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummaries()(*WindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummariesRequestBuilder) {
    return NewWindowsinformationprotectionapplearningsummariesWindowsInformationProtectionAppLearningSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsInformationProtectionNetworkLearningSummaries provides operations to manage the windowsInformationProtectionNetworkLearningSummaries property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummariesRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummaries()(*WindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummariesRequestBuilder) {
    return NewWindowsinformationprotectionnetworklearningsummariesWindowsInformationProtectionNetworkLearningSummariesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WindowsMalwareInformation provides operations to manage the windowsMalwareInformation property of the microsoft.graph.deviceManagement entity.
// returns a *WindowsmalwareinformationWindowsMalwareInformationRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WindowsMalwareInformation()(*WindowsmalwareinformationWindowsMalwareInformationRequestBuilder) {
    return NewWindowsmalwareinformationWindowsMalwareInformationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceManagementRequestBuilder when successful
func (m *DeviceManagementRequestBuilder) WithUrl(rawUrl string)(*DeviceManagementRequestBuilder) {
    return NewDeviceManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
