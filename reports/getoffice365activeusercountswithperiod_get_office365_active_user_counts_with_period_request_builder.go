package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder provides operations to call the getOffice365ActiveUserCounts method.
type Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal instantiates a new Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    m := &Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365ActiveUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder instantiates a new Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the count of daily active users in the reporting period by product.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getoffice365activeusercounts?view=graph-rest-1.0
func (m *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the count of daily active users in the reporting period by product.
// returns a *RequestInformation when successful
func (m *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder when successful
func (m *Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder) {
    return NewGetoffice365activeusercountswithperiodGetOffice365ActiveUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
