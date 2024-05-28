package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder provides operations to call the getM365AppPlatformUserCounts method.
type Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal instantiates a new Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    m := &Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getM365AppPlatformUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder instantiates a new Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get a report that provides the trend of active users across all apps for each platform (Windows, Mac, web, and mobile) in your organization.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getm365appplatformusercounts?view=graph-rest-1.0
func (m *Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation get a report that provides the trend of active users across all apps for each platform (Windows, Mac, web, and mobile) in your organization.
// returns a *RequestInformation when successful
func (m *Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder when successful
func (m *Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder) {
    return NewGetm365appplatformusercountswithperiodGetM365AppPlatformUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
