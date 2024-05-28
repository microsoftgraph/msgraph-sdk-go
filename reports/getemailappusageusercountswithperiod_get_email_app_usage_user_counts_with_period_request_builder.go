package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder provides operations to call the getEmailAppUsageUserCounts method.
type GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal instantiates a new GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    m := &GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getEmailAppUsageUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder instantiates a new GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the count of unique users that connected to Exchange Online using any email app.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getemailappusageusercounts?view=graph-rest-1.0
func (m *GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the count of unique users that connected to Exchange Online using any email app.
// returns a *RequestInformation when successful
func (m *GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder when successful
func (m *GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder) {
    return NewGetemailappusageusercountswithperiodGetEmailAppUsageUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
