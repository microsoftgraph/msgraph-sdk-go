package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder provides operations to manage the deviceSettingStateSummaries property of the microsoft.graph.deviceConfiguration entity.
type DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderGetQueryParameters read properties and relationships of the settingStateDeviceSummary object.
type DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderGetQueryParameters
}
// DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderInternal instantiates a new DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) {
    m := &DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/deviceSettingStateSummaries/{settingStateDeviceSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder instantiates a new DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a settingStateDeviceSummary.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-settingstatedevicesummary-delete?view=graph-rest-1.0
func (m *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the settingStateDeviceSummary object.
// returns a SettingStateDeviceSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-settingstatedevicesummary-get?view=graph-rest-1.0
func (m *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SettingStateDeviceSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSettingStateDeviceSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SettingStateDeviceSummaryable), nil
}
// Patch update the properties of a settingStateDeviceSummary object.
// returns a SettingStateDeviceSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-settingstatedevicesummary-update?view=graph-rest-1.0
func (m *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SettingStateDeviceSummaryable, requestConfiguration *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SettingStateDeviceSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSettingStateDeviceSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SettingStateDeviceSummaryable), nil
}
// ToDeleteRequestInformation deletes a settingStateDeviceSummary.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the settingStateDeviceSummary object.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a settingStateDeviceSummary object.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SettingStateDeviceSummaryable, requestConfiguration *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder when successful
func (m *DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder) {
    return NewDeviceconfigurationsItemDevicesettingstatesummariesSettingStateDeviceSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
