package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetSharePointSiteUsagePagesWithPeriodRequestBuilder provides operations to call the getSharePointSiteUsagePages method.
type GetSharePointSiteUsagePagesWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetSharePointSiteUsagePagesWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetSharePointSiteUsagePagesWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal instantiates a new GetSharePointSiteUsagePagesWithPeriodRequestBuilder and sets the default values.
func NewGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    m := &GetSharePointSiteUsagePagesWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSharePointSiteUsagePages(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetSharePointSiteUsagePagesWithPeriodRequestBuilder instantiates a new GetSharePointSiteUsagePagesWithPeriodRequestBuilder and sets the default values.
func NewGetSharePointSiteUsagePagesWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetSharePointSiteUsagePagesWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getSharePointSiteUsagePages
func (m *GetSharePointSiteUsagePagesWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetSharePointSiteUsagePagesWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// ToGetRequestInformation invoke function getSharePointSiteUsagePages
func (m *GetSharePointSiteUsagePagesWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetSharePointSiteUsagePagesWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *GetSharePointSiteUsagePagesWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetSharePointSiteUsagePagesWithPeriodRequestBuilder) {
    return NewGetSharePointSiteUsagePagesWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
