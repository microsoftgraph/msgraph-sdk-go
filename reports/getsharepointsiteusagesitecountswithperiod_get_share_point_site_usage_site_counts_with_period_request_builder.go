package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder provides operations to call the getSharePointSiteUsageSiteCounts method.
type GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal instantiates a new GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder and sets the default values.
func NewGetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    m := &GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSharePointSiteUsageSiteCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder instantiates a new GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder and sets the default values.
func NewGetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the trend of total and active site count during the reporting period.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getsharepointsiteusagesitecounts?view=graph-rest-1.0
func (m *GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the trend of total and active site count during the reporting period.
// returns a *RequestInformation when successful
func (m *GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder when successful
func (m *GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder) {
    return NewGetsharepointsiteusagesitecountswithperiodGetSharePointSiteUsageSiteCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
