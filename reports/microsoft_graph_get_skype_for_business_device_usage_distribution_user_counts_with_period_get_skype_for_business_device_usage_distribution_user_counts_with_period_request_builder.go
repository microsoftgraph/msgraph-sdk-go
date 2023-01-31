package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder provides operations to call the getSkypeForBusinessDeviceUsageDistributionUserCounts method.
type MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal instantiates a new GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    m := &MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getSkypeForBusinessDeviceUsageDistributionUserCounts(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder instantiates a new GetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getSkypeForBusinessDeviceUsageDistributionUserCounts
func (m *MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getSkypeForBusinessDeviceUsageDistributionUserCounts
func (m *MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodGetSkypeForBusinessDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
