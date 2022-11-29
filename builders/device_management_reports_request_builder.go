package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceManagementReportsRequestBuilder provides operations to manage the reports property of the microsoft.graph.deviceManagement entity.
type DeviceManagementReportsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementReportsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReportsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementReportsRequestBuilderGetQueryParameters reports singleton
type DeviceManagementReportsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementReportsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReportsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementReportsRequestBuilderGetQueryParameters
}
// DeviceManagementReportsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReportsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementReportsRequestBuilderInternal instantiates a new ReportsRequestBuilder and sets the default values.
func NewDeviceManagementReportsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReportsRequestBuilder) {
    m := &DeviceManagementReportsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementReportsRequestBuilder instantiates a new ReportsRequestBuilder and sets the default values.
func NewDeviceManagementReportsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReportsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementReportsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reports for deviceManagement
func (m *DeviceManagementReportsRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReportsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation reports singleton
func (m *DeviceManagementReportsRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReportsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property reports in deviceManagement
func (m *DeviceManagementReportsRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable, requestConfiguration *DeviceManagementReportsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property reports for deviceManagement
func (m *DeviceManagementReportsRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementReportsRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExportJobs provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *DeviceManagementReportsRequestBuilder) ExportJobs()(*DeviceManagementReportsExportJobsRequestBuilder) {
    return NewDeviceManagementReportsExportJobsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExportJobsById provides operations to manage the exportJobs property of the microsoft.graph.deviceManagementReports entity.
func (m *DeviceManagementReportsRequestBuilder) ExportJobsById(id string)(*DeviceManagementReportsExportJobsDeviceManagementExportJobItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExportJob%2Did"] = id
    }
    return NewDeviceManagementReportsExportJobsDeviceManagementExportJobItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get reports singleton
func (m *DeviceManagementReportsRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementReportsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable), nil
}
// GetCachedReport provides operations to call the getCachedReport method.
func (m *DeviceManagementReportsRequestBuilder) GetCachedReport()(*DeviceManagementReportsGetCachedReportRequestBuilder) {
    return NewDeviceManagementReportsGetCachedReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyNonComplianceReport provides operations to call the getCompliancePolicyNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetCompliancePolicyNonComplianceReport()(*DeviceManagementReportsGetCompliancePolicyNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetCompliancePolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCompliancePolicyNonComplianceSummaryReport provides operations to call the getCompliancePolicyNonComplianceSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetCompliancePolicyNonComplianceSummaryReport()(*DeviceManagementReportsGetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetComplianceSettingNonComplianceReport provides operations to call the getComplianceSettingNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetComplianceSettingNonComplianceReport()(*DeviceManagementReportsGetComplianceSettingNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetComplianceSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyNonComplianceReport provides operations to call the getConfigurationPolicyNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationPolicyNonComplianceReport()(*DeviceManagementReportsGetConfigurationPolicyNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationPolicyNonComplianceSummaryReport provides operations to call the getConfigurationPolicyNonComplianceSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationPolicyNonComplianceSummaryReport()(*DeviceManagementReportsGetConfigurationPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetConfigurationSettingNonComplianceReport provides operations to call the getConfigurationSettingNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetConfigurationSettingNonComplianceReport()(*DeviceManagementReportsGetConfigurationSettingNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetConfigurationSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceManagementIntentPerSettingContributingProfiles provides operations to call the getDeviceManagementIntentPerSettingContributingProfiles method.
func (m *DeviceManagementReportsRequestBuilder) GetDeviceManagementIntentPerSettingContributingProfiles()(*DeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilder) {
    return NewDeviceManagementReportsGetDeviceManagementIntentPerSettingContributingProfilesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceManagementIntentSettingsReport provides operations to call the getDeviceManagementIntentSettingsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetDeviceManagementIntentSettingsReport()(*DeviceManagementReportsGetDeviceManagementIntentSettingsReportRequestBuilder) {
    return NewDeviceManagementReportsGetDeviceManagementIntentSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetDeviceNonComplianceReport provides operations to call the getDeviceNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetDeviceNonComplianceReport()(*DeviceManagementReportsGetDeviceNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetDeviceNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetHistoricalReport provides operations to call the getHistoricalReport method.
func (m *DeviceManagementReportsRequestBuilder) GetHistoricalReport()(*DeviceManagementReportsGetHistoricalReportRequestBuilder) {
    return NewDeviceManagementReportsGetHistoricalReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetNoncompliantDevicesAndSettingsReport provides operations to call the getNoncompliantDevicesAndSettingsReport method.
func (m *DeviceManagementReportsRequestBuilder) GetNoncompliantDevicesAndSettingsReport()(*DeviceManagementReportsGetNoncompliantDevicesAndSettingsReportRequestBuilder) {
    return NewDeviceManagementReportsGetNoncompliantDevicesAndSettingsReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceMetadata provides operations to call the getPolicyNonComplianceMetadata method.
func (m *DeviceManagementReportsRequestBuilder) GetPolicyNonComplianceMetadata()(*DeviceManagementReportsGetPolicyNonComplianceMetadataRequestBuilder) {
    return NewDeviceManagementReportsGetPolicyNonComplianceMetadataRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceReport provides operations to call the getPolicyNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetPolicyNonComplianceReport()(*DeviceManagementReportsGetPolicyNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetPolicyNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetPolicyNonComplianceSummaryReport provides operations to call the getPolicyNonComplianceSummaryReport method.
func (m *DeviceManagementReportsRequestBuilder) GetPolicyNonComplianceSummaryReport()(*DeviceManagementReportsGetPolicyNonComplianceSummaryReportRequestBuilder) {
    return NewDeviceManagementReportsGetPolicyNonComplianceSummaryReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetReportFilters provides operations to call the getReportFilters method.
func (m *DeviceManagementReportsRequestBuilder) GetReportFilters()(*DeviceManagementReportsGetReportFiltersRequestBuilder) {
    return NewDeviceManagementReportsGetReportFiltersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetSettingNonComplianceReport provides operations to call the getSettingNonComplianceReport method.
func (m *DeviceManagementReportsRequestBuilder) GetSettingNonComplianceReport()(*DeviceManagementReportsGetSettingNonComplianceReportRequestBuilder) {
    return NewDeviceManagementReportsGetSettingNonComplianceReportRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property reports in deviceManagement
func (m *DeviceManagementReportsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable, requestConfiguration *DeviceManagementReportsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceManagementReportsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementReportsable), nil
}
