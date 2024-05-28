package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder provides operations to manage the deviceCompliancePolicyDeviceStateSummary property of the microsoft.graph.deviceManagement entity.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetQueryParameters read properties and relationships of the deviceCompliancePolicyDeviceStateSummary object.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetQueryParameters
}
// DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal instantiates a new DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder and sets the default values.
func NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    m := &DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicyDeviceStateSummary{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder instantiates a new DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder and sets the default values.
func NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceCompliancePolicyDeviceStateSummary for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the deviceCompliancePolicyDeviceStateSummary object.
// returns a DeviceCompliancePolicyDeviceStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancepolicydevicestatesummary-get?view=graph-rest-1.0
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyDeviceStateSummaryable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyDeviceStateSummaryable), nil
}
// Patch update the properties of a deviceCompliancePolicyDeviceStateSummary object.
// returns a DeviceCompliancePolicyDeviceStateSummaryable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancepolicydevicestatesummary-update?view=graph-rest-1.0
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyDeviceStateSummaryable, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyDeviceStateSummaryable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceCompliancePolicyDeviceStateSummaryFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyDeviceStateSummaryable), nil
}
// ToDeleteRequestInformation delete navigation property deviceCompliancePolicyDeviceStateSummary for deviceManagement
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the deviceCompliancePolicyDeviceStateSummary object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a deviceCompliancePolicyDeviceStateSummary object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceCompliancePolicyDeviceStateSummaryable, requestConfiguration *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder when successful
func (m *DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    return NewDevicecompliancepolicydevicestatesummaryDeviceCompliancePolicyDeviceStateSummaryRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
