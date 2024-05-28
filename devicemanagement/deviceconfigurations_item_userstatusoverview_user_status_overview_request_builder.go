package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder provides operations to manage the userStatusOverview property of the microsoft.graph.deviceConfiguration entity.
type DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderGetQueryParameters read properties and relationships of the deviceConfigurationUserOverview object.
type DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderGetQueryParameters
}
// DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderInternal instantiates a new DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) {
    m := &DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/userStatusOverview{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder instantiates a new DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder and sets the default values.
func NewDeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userStatusOverview for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read properties and relationships of the deviceConfigurationUserOverview object.
// returns a DeviceConfigurationUserOverviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-deviceconfigurationuseroverview-get?view=graph-rest-1.0
func (m *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationUserOverviewable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceConfigurationUserOverviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationUserOverviewable), nil
}
// Patch update the properties of a deviceConfigurationUserOverview object.
// returns a DeviceConfigurationUserOverviewable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-deviceconfig-deviceconfigurationuseroverview-update?view=graph-rest-1.0
func (m *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationUserOverviewable, requestConfiguration *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationUserOverviewable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceConfigurationUserOverviewFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationUserOverviewable), nil
}
// ToDeleteRequestInformation delete navigation property userStatusOverview for deviceManagement
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the deviceConfigurationUserOverview object.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a deviceConfigurationUserOverview object.
// returns a *RequestInformation when successful
func (m *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceConfigurationUserOverviewable, requestConfiguration *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder when successful
func (m *DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) WithUrl(rawUrl string)(*DeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder) {
    return NewDeviceconfigurationsItemUserstatusoverviewUserStatusOverviewRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
