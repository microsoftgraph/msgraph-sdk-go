package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DevicecompliancepoliciesItemAssignRequestBuilder provides operations to call the assign method.
type DevicecompliancepoliciesItemAssignRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepoliciesItemAssignRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemAssignRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecompliancepoliciesItemAssignRequestBuilderInternal instantiates a new DevicecompliancepoliciesItemAssignRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemAssignRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemAssignRequestBuilder) {
    m := &DevicecompliancepoliciesItemAssignRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/assign", pathParameters),
    }
    return m
}
// NewDevicecompliancepoliciesItemAssignRequestBuilder instantiates a new DevicecompliancepoliciesItemAssignRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemAssignRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemAssignRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepoliciesItemAssignRequestBuilderInternal(urlParams, requestAdapter)
}
// Post not yet documented
// Deprecated: This method is obsolete. Use PostAsAssignPostResponse instead.
// returns a DevicecompliancepoliciesItemAssignResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancepolicy-assign?view=graph-rest-1.0
func (m *DevicecompliancepoliciesItemAssignRequestBuilder) Post(ctx context.Context, body DevicecompliancepoliciesItemAssignPostRequestBodyable, requestConfiguration *DevicecompliancepoliciesItemAssignRequestBuilderPostRequestConfiguration)(DevicecompliancepoliciesItemAssignResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicecompliancepoliciesItemAssignResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicecompliancepoliciesItemAssignResponseable), nil
}
// PostAsAssignPostResponse not yet documented
// returns a DevicecompliancepoliciesItemAssignPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecompliancepolicy-assign?view=graph-rest-1.0
func (m *DevicecompliancepoliciesItemAssignRequestBuilder) PostAsAssignPostResponse(ctx context.Context, body DevicecompliancepoliciesItemAssignPostRequestBodyable, requestConfiguration *DevicecompliancepoliciesItemAssignRequestBuilderPostRequestConfiguration)(DevicecompliancepoliciesItemAssignPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDevicecompliancepoliciesItemAssignPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DevicecompliancepoliciesItemAssignPostResponseable), nil
}
// ToPostRequestInformation not yet documented
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemAssignRequestBuilder) ToPostRequestInformation(ctx context.Context, body DevicecompliancepoliciesItemAssignPostRequestBodyable, requestConfiguration *DevicecompliancepoliciesItemAssignRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepoliciesItemAssignRequestBuilder when successful
func (m *DevicecompliancepoliciesItemAssignRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepoliciesItemAssignRequestBuilder) {
    return NewDevicecompliancepoliciesItemAssignRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
