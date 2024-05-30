package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
type GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal instantiates a new GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    m := &GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getYammerDeviceUsageDistributionUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder instantiates a new GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the number of users by device type.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getyammerdeviceusagedistributionusercounts?view=graph-rest-1.0
func (m *GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the number of users by device type.
// returns a *RequestInformation when successful
func (m *GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder when successful
func (m *GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    return NewGetyammerdeviceusagedistributionusercountswithperiodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
