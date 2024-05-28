package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder provides operations to call the getSharePointActivityUserCounts method.
type GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal instantiates a new GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    m := &GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSharePointActivityUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder instantiates a new GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the trend in the number of active users. A user is considered active if they have executed a file activity (save, sync, modify, or share) or visited a page within the specified time period.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getsharepointactivityusercounts?view=graph-rest-1.0
func (m *GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the trend in the number of active users. A user is considered active if they have executed a file activity (save, sync, modify, or share) or visited a page within the specified time period.
// returns a *RequestInformation when successful
func (m *GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder when successful
func (m *GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetsharepointactivityusercountswithperiodGetSharePointActivityUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
