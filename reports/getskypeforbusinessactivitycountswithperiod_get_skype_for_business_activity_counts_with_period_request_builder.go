package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder provides operations to call the getSkypeForBusinessActivityCounts method.
type GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal instantiates a new GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    m := &GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSkypeForBusinessActivityCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder instantiates a new GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the trends on how many users organized and participated in conference sessions held in your organization through Skype for Business. The report also includes the number of peer-to-peer sessions.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getskypeforbusinessactivitycounts?view=graph-rest-1.0
func (m *GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the trends on how many users organized and participated in conference sessions held in your organization through Skype for Business. The report also includes the number of peer-to-peer sessions.
// returns a *RequestInformation when successful
func (m *GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder when successful
func (m *GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder) {
    return NewGetskypeforbusinessactivitycountswithperiodGetSkypeForBusinessActivityCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
