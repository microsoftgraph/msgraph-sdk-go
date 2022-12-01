package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceAppManagementRequestBuilder provides operations to manage the deviceAppManagement singleton.
type DeviceAppManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementRequestBuilderGetQueryParameters get deviceAppManagement
type DeviceAppManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementRequestBuilderGetQueryParameters
}
// DeviceAppManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AndroidManagedAppProtections provides operations to manage the androidManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtections()(*DeviceAppManagementAndroidManagedAppProtectionsRequestBuilder) {
    return NewDeviceAppManagementAndroidManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidManagedAppProtectionsById provides operations to manage the androidManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtectionsById(id string)(*DeviceAppManagementAndroidManagedAppProtectionsAndroidManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidManagedAppProtection%2Did"] = id
    }
    return NewDeviceAppManagementAndroidManagedAppProtectionsAndroidManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementRequestBuilderInternal instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    m := &DeviceAppManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementRequestBuilder instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAppManagementable, requestConfiguration *DeviceAppManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// DefaultManagedAppProtections provides operations to manage the defaultManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtections()(*DeviceAppManagementDefaultManagedAppProtectionsRequestBuilder) {
    return NewDeviceAppManagementDefaultManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultManagedAppProtectionsById provides operations to manage the defaultManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtectionsById(id string)(*DeviceAppManagementDefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["defaultManagedAppProtection%2Did"] = id
    }
    return NewDeviceAppManagementDefaultManagedAppProtectionsDefaultManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAppManagementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceAppManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAppManagementable), nil
}
// IosManagedAppProtections provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtections()(*DeviceAppManagementIosManagedAppProtectionsRequestBuilder) {
    return NewDeviceAppManagementIosManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IosManagedAppProtectionsById provides operations to manage the iosManagedAppProtections property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtectionsById(id string)(*DeviceAppManagementIosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosManagedAppProtection%2Did"] = id
    }
    return NewDeviceAppManagementIosManagedAppProtectionsIosManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppPolicies provides operations to manage the managedAppPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppPolicies()(*DeviceAppManagementManagedAppPoliciesRequestBuilder) {
    return NewDeviceAppManagementManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppPoliciesById provides operations to manage the managedAppPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppPoliciesById(id string)(*DeviceAppManagementManagedAppPoliciesManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy%2Did"] = id
    }
    return NewDeviceAppManagementManagedAppPoliciesManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppRegistrations provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrations()(*DeviceAppManagementManagedAppRegistrationsRequestBuilder) {
    return NewDeviceAppManagementManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppRegistrationsById provides operations to manage the managedAppRegistrations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrationsById(id string)(*DeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration%2Did"] = id
    }
    return NewDeviceAppManagementManagedAppRegistrationsManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppStatuses provides operations to manage the managedAppStatuses property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatuses()(*DeviceAppManagementManagedAppStatusesRequestBuilder) {
    return NewDeviceAppManagementManagedAppStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppStatusesById provides operations to manage the managedAppStatuses property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatusesById(id string)(*DeviceAppManagementManagedAppStatusesManagedAppStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppStatus%2Did"] = id
    }
    return NewDeviceAppManagementManagedAppStatusesManagedAppStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedEBooks provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedEBooks()(*DeviceAppManagementManagedEBooksRequestBuilder) {
    return NewDeviceAppManagementManagedEBooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedEBooksById provides operations to manage the managedEBooks property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) ManagedEBooksById(id string)(*DeviceAppManagementManagedEBooksManagedEBookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBook%2Did"] = id
    }
    return NewDeviceAppManagementManagedEBooksManagedEBookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MdmWindowsInformationProtectionPolicies provides operations to manage the mdmWindowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPolicies()(*DeviceAppManagementMdmWindowsInformationProtectionPoliciesRequestBuilder) {
    return NewDeviceAppManagementMdmWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MdmWindowsInformationProtectionPoliciesById provides operations to manage the mdmWindowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPoliciesById(id string)(*DeviceAppManagementMdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mdmWindowsInformationProtectionPolicy%2Did"] = id
    }
    return NewDeviceAppManagementMdmWindowsInformationProtectionPoliciesMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppCategories provides operations to manage the mobileAppCategories property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppCategories()(*DeviceAppManagementMobileAppCategoriesRequestBuilder) {
    return NewDeviceAppManagementMobileAppCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppCategoriesById provides operations to manage the mobileAppCategories property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppCategoriesById(id string)(*DeviceAppManagementMobileAppCategoriesMobileAppCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppCategory%2Did"] = id
    }
    return NewDeviceAppManagementMobileAppCategoriesMobileAppCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppConfigurations provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurations()(*DeviceAppManagementMobileAppConfigurationsRequestBuilder) {
    return NewDeviceAppManagementMobileAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppConfigurationsById provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurationsById(id string)(*DeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfiguration%2Did"] = id
    }
    return NewDeviceAppManagementMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileApps provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileApps()(*DeviceAppManagementMobileAppsRequestBuilder) {
    return NewDeviceAppManagementMobileAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppsById provides operations to manage the mobileApps property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) MobileAppsById(id string)(*DeviceAppManagementMobileAppsMobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileApp%2Did"] = id
    }
    return NewDeviceAppManagementMobileAppsMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAppManagementable, requestConfiguration *DeviceAppManagementRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAppManagementable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceAppManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAppManagementable), nil
}
// SyncMicrosoftStoreForBusinessApps provides operations to call the syncMicrosoftStoreForBusinessApps method.
func (m *DeviceAppManagementRequestBuilder) SyncMicrosoftStoreForBusinessApps()(*DeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    return NewDeviceAppManagementSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetedManagedAppConfigurations provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurations()(*DeviceAppManagementTargetedManagedAppConfigurationsRequestBuilder) {
    return NewDeviceAppManagementTargetedManagedAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetedManagedAppConfigurationsById provides operations to manage the targetedManagedAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurationsById(id string)(*DeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppConfiguration%2Did"] = id
    }
    return NewDeviceAppManagementTargetedManagedAppConfigurationsTargetedManagedAppConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VppTokens provides operations to manage the vppTokens property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) VppTokens()(*DeviceAppManagementVppTokensRequestBuilder) {
    return NewDeviceAppManagementVppTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VppTokensById provides operations to manage the vppTokens property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) VppTokensById(id string)(*DeviceAppManagementVppTokensVppTokenItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["vppToken%2Did"] = id
    }
    return NewDeviceAppManagementVppTokensVppTokenItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionPolicies provides operations to manage the windowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPolicies()(*DeviceAppManagementWindowsInformationProtectionPoliciesRequestBuilder) {
    return NewDeviceAppManagementWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionPoliciesById provides operations to manage the windowsInformationProtectionPolicies property of the microsoft.graph.deviceAppManagement entity.
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPoliciesById(id string)(*DeviceAppManagementWindowsInformationProtectionPoliciesWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionPolicy%2Did"] = id
    }
    return NewDeviceAppManagementWindowsInformationProtectionPoliciesWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
