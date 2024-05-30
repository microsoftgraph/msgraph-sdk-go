package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder provides operations to manage the deviceCompliancePolicySettingStateSummaries property of the microsoft.graph.deviceManagement entity.
type DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetQueryParameters read properties and relationships of the deviceCompliancePolicySettingStateSummary object.
type DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetQueryParameters
}
// DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal instantiates a new DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder and sets the default values.
func NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    m := &DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicySettingStateSummaries/{deviceCompliancePolicySettingStateSummary%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder instantiates a new DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder and sets the default values.
func NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a deviceCompliancePolicySettingStateSummary.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancepolicysettingstatesummary-delete?view=graph-rest-1.0
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceComplianceSettingStates provides operations to manage the deviceComplianceSettingStates property of the microsoft.graph.deviceCompliancePolicySettingStateSummary entity.
// returns a *DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) DeviceComplianceSettingStates()(*DevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilder) {
    return NewDevicecompliancepolicysettingstatesummariesItemDevicecompliancesettingstatesDeviceComplianceSettingStatesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the deviceCompliancePolicySettingStateSummary object.
// returns a DeviceCompliancePolicySettingStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancepolicysettingstatesummary-get?view=graph-rest-1.0
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicySettingStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicySettingStateSummaryable), nil
}
// Patch update the properties of a deviceCompliancePolicySettingStateSummary object.
// returns a DeviceCompliancePolicySettingStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancepolicysettingstatesummary-update?view=graph-rest-1.0
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicySettingStateSummaryable, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicySettingStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceCompliancePolicySettingStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicySettingStateSummaryable), nil
}
// ToDeleteRequestInformation deletes a deviceCompliancePolicySettingStateSummary.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the deviceCompliancePolicySettingStateSummary object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a deviceCompliancePolicySettingStateSummary object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicySettingStateSummaryable, requestConfiguration *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder when successful
func (m *DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder) {
    return NewDevicecompliancepolicysettingstatesummariesDeviceCompliancePolicySettingStateSummaryItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
