package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
type UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetQueryParameters user experience analytics work from anywhere metrics.
type UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetQueryParameters struct {
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
// UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetQueryParameters
}
// UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserExperienceAnalyticsWorkFromAnywhereMetricId provides operations to manage the userExperienceAnalyticsWorkFromAnywhereMetrics property of the microsoft.graph.deviceManagement entity.
// returns a *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) ByUserExperienceAnalyticsWorkFromAnywhereMetricId(userExperienceAnalyticsWorkFromAnywhereMetricId string)(*UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userExperienceAnalyticsWorkFromAnywhereMetricId != "" {
        urlTplParams["userExperienceAnalyticsWorkFromAnywhereMetric%2Did"] = userExperienceAnalyticsWorkFromAnywhereMetricId
    }
    return NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal instantiates a new UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    m := &UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/userExperienceAnalyticsWorkFromAnywhereMetrics{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder instantiates a new UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder and sets the default values.
func NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *UserexperienceanalyticsworkfromanywheremetricsCountRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) Count()(*UserexperienceanalyticsworkfromanywheremetricsCountRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremetricsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get user experience analytics work from anywhere metrics.
// returns a UserExperienceAnalyticsWorkFromAnywhereMetricCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) Get(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserExperienceAnalyticsWorkFromAnywhereMetricCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricCollectionResponseable), nil
}
// Post create new navigation property to userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
// returns a UserExperienceAnalyticsWorkFromAnywhereMetricable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricable, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation user experience analytics work from anywhere metrics.
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userExperienceAnalyticsWorkFromAnywhereMetrics for deviceManagement
// returns a *RequestInformation when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserExperienceAnalyticsWorkFromAnywhereMetricable, requestConfiguration *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder when successful
func (m *UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) WithUrl(rawUrl string)(*UserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder) {
    return NewUserexperienceanalyticsworkfromanywheremetricsUserExperienceAnalyticsWorkFromAnywhereMetricsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
