package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder provides operations to call the getEmailActivityCounts method.
type GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilderInternal instantiates a new GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder) {
    m := &GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getEmailActivityCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder instantiates a new GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get enables you to understand the trends of email activity (like how many were sent, read, and received) in your organization.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getemailactivitycounts?view=graph-rest-1.0
func (m *GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation enables you to understand the trends of email activity (like how many were sent, read, and received) in your organization.
// returns a *RequestInformation when successful
func (m *GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder when successful
func (m *GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder) {
    return NewGetemailactivitycountswithperiodGetEmailActivityCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
