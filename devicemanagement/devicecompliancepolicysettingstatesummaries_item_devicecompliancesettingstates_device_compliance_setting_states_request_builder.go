package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder provides operations to manage the deviceComplianceSettingStates property of the microsoft.graph.deviceCompliancePolicySettingStateSummary entity.
type DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderGetQueryParameters list properties and relationships of the deviceComplianceSettingState objects.
type DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderGetQueryParameters
}
// DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceComplianceSettingStateId provides operations to manage the deviceComplianceSettingStates property of the microsoft.graph.deviceCompliancePolicySettingStateSummary entity.
// returns a *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder when successful
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) ByDeviceComplianceSettingStateId(deviceComplianceSettingStateId string)(*DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceComplianceSettingStateId != "" {
        urlTplParams["deviceComplianceSettingState%2Did"] = deviceComplianceSettingStateId
    }
    return NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderInternal instantiates a new DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder and sets the default values.
func NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) {
    m := &DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary%2Did}/deviceComplianceSettingStates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder instantiates a new DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder and sets the default values.
func NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesCountRequestBuilder when successful
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) Count()(*DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesCountRequestBuilder) {
    return NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the deviceComplianceSettingState objects.
// returns a DeviceComplianceSettingStateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancesettingstate-list?view=graph-rest-1.0
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceSettingStateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateCollectionResponseable), nil
}
// Post create a new deviceComplianceSettingState object.
// returns a DeviceComplianceSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancesettingstate-create?view=graph-rest-1.0
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceSettingStateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable), nil
}
// ToGetRequestInformation list properties and relationships of the deviceComplianceSettingState objects.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new deviceComplianceSettingState object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder when successful
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) {
    return NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
