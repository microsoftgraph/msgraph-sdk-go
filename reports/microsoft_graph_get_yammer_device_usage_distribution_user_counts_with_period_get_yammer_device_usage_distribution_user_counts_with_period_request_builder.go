package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder provides operations to call the getYammerDeviceUsageDistributionUserCounts method.
type MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal instantiates a new GetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    m := &MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getYammerDeviceUsageDistributionUserCounts(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder instantiates a new GetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder and sets the default values.
func NewMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getYammerDeviceUsageDistributionUserCounts
func (m *MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation invoke function getYammerDeviceUsageDistributionUserCounts
func (m *MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphGetYammerDeviceUsageDistributionUserCountsWithPeriodGetYammerDeviceUsageDistributionUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
