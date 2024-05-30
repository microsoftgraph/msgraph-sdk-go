package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetQueryParameters user experience analytics work from anywhere metrics.
type UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal instantiates a new UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    m := &UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics/{userExperienceAnalyticsWorkFromAnywhereMetric%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder instantiates a new UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user experience analytics work from anywhere metrics.
// returns a UserExperienceAnalyticsWorkFromAnywhereMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricable), nil
}
// MetricDevices provides operations to manage the metricDevices property of the microsoft.graph.userExperienceAnalyticsWorkFromAnywhereMetric entity.
// returns a *UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesMetricDevicesRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) MetricDevices()(*UserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesMetricDevicesRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremetricsItemMetricdevicesMetricDevicesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property userExperienceAnalyticsWorkFromAnywhereMetrics in deviceManagement
// returns a UserExperienceAnalyticsWorkFromAnywhereMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricable, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereMetricFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricable), nil
}
// ToDeleteRequestInformation delete navigation property userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user experience analytics work from anywhere metrics.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userExperienceAnalyticsWorkFromAnywhereMetrics in deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricable, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
