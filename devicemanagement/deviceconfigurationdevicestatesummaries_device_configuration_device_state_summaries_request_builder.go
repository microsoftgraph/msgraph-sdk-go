package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder provides operations to manage the deviceConfigurationDeviceStateSummaries property of the microsoft.graph.deviceManagement entity.
type DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderGetQueryParameters read properties and relationships of the deviceConfigurationDeviceStateSummary object.
type DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderGetQueryParameters
}
// DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderInternal instantiates a new DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder and sets the default values.
func NewDeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) {
    m := &DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurationDeviceStateSummaries{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder instantiates a new DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder and sets the default values.
func NewDeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceConfigurationDeviceStateSummaries for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read properties and relationships of the deviceConfigurationDeviceStateSummary object.
// returns a DeviceConfigurationDeviceStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-deviceconfigurationdevicestatesummary-get?view=graph-rest-1.0
func (m *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationDeviceStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceConfigurationDeviceStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationDeviceStateSummaryable), nil
}
// Patch update the properties of a deviceConfigurationDeviceStateSummary object.
// returns a DeviceConfigurationDeviceStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-deviceconfigurationdevicestatesummary-update?view=graph-rest-1.0
func (m *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationDeviceStateSummaryable, requestConfiguration *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationDeviceStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceConfigurationDeviceStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationDeviceStateSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property deviceConfigurationDeviceStateSummaries for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the deviceConfigurationDeviceStateSummary object.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a deviceConfigurationDeviceStateSummary object.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationDeviceStateSummaryable, requestConfiguration *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder when successful
func (m *DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder) {
    return NewDeviceconfigurationdevicestatesummariesDeviceConfigurationDeviceStateSummariesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
