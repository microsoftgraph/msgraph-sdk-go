package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder provides operations to manage the deviceComplianceSettingStates property of the microsoft.graph.deviceCompliancePolicySettingStateSummary entity.
type DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderGetQueryParameters not yet documented
type DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderGetQueryParameters
}
// DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderInternal instantiates a new DeviceComplianceSettingStateItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder) {
    m := &DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary%2Did}/deviceComplianceSettingStates/{deviceComplianceSettingState%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder instantiates a new DeviceComplianceSettingStateItemRequestBuilder and sets the default values.
func NewDeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceComplianceSettingStates for deviceManagement
func (m *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation not yet documented
func (m *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceComplianceSettingStates in deviceManagement
func (m *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, requestConfiguration *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceComplianceSettingStates for deviceManagement
func (m *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get not yet documented
func (m *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable), nil
}
// Patch update the navigation property deviceComplianceSettingStates in deviceManagement
func (m *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, requestConfiguration *DeviceManagementDeviceCompliancePolicySettingStateSummariesItemDeviceComplianceSettingStatesDeviceComplianceSettingStateItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable), nil
}
