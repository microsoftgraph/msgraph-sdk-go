package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder provides operations to call the getOffice365ServicesUserCounts method.
type Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal instantiates a new Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    m := &Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365ServicesUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder instantiates a new Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get the count of users by activity type and service.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getoffice365servicesusercounts?view=graph-rest-1.0
func (m *Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get the count of users by activity type and service.
// returns a *RequestInformation when successful
func (m *Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder when successful
func (m *Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder) {
    return NewGetoffice365servicesusercountswithperiodGetOffice365ServicesUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
