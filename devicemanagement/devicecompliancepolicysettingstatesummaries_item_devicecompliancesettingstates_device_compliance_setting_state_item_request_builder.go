package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder provides operations to manage the deviceComplianceSettingStates property of the microsoft.graph.deviceCompliancePolicySettingStateSummary entity.
type DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderGetQueryParameters read properties and relationships of the deviceComplianceSettingState object.
type DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderGetQueryParameters
}
// DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderInternal instantiates a new DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder and sets the default values.
func NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) {
    m := &DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary%2Did}/deviceComplianceSettingStates/{deviceComplianceSettingState%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder instantiates a new DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder and sets the default values.
func NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a deviceComplianceSettingState.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancesettingstate-delete?view=graph-rest-1.0
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the deviceComplianceSettingState object.
// returns a DeviceComplianceSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancesettingstate-get?view=graph-rest-1.0
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the properties of a deviceComplianceSettingState object.
// returns a DeviceComplianceSettingStateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancesettingstate-update?view=graph-rest-1.0
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation deletes a deviceComplianceSettingState.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the deviceComplianceSettingState object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a deviceComplianceSettingState object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceSettingStateable, requestConfiguration *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder when successful
func (m *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder) {
    return NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
