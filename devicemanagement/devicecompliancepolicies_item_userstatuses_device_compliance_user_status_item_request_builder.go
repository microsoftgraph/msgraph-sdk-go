package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder provides operations to manage the userStatuses property of the microsoft.graph.deviceCompliancePolicy entity.
type DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderGetQueryParameters read properties and relationships of the deviceComplianceUserStatus object.
type DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderGetQueryParameters
}
// DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderInternal instantiates a new DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) {
    m := &DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCompliancePolicies/{deviceCompliancePolicy%2Did}/userStatuses/{deviceComplianceUserStatus%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder instantiates a new DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder and sets the default values.
func NewDevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a deviceComplianceUserStatus.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecomplianceuserstatus-delete?view=graph-rest-1.0
func (m *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the deviceComplianceUserStatus object.
// returns a DeviceComplianceUserStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecomplianceuserstatus-get?view=graph-rest-1.0
func (m *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceUserStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceUserStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceUserStatusable), nil
}
// Patch update the properties of a deviceComplianceUserStatus object.
// returns a DeviceComplianceUserStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-devicecomplianceuserstatus-update?view=graph-rest-1.0
func (m *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceUserStatusable, requestConfiguration *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceUserStatusable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceComplianceUserStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceUserStatusable), nil
}
// ToDeleteRequestInformation deletes a deviceComplianceUserStatus.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the deviceComplianceUserStatus object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a deviceComplianceUserStatus object.
// returns a *RequestInformation when successful
func (m *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceComplianceUserStatusable, requestConfiguration *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder when successful
func (m *DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) WithUrl(rawUrl string)(*DevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder) {
    return NewDevicecompliancepoliciesItemUserstatusesDeviceComplianceUserStatusItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
