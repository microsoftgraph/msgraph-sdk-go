package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder provides operations to call the getSkypeForBusinessOrganizerActivityUserCounts method.
type GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal instantiates a new GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    m := &GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSkypeForBusinessOrganizerActivityUserCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder instantiates a new GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder and sets the default values.
func NewGetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get usage trends on the number of unique users and type of conference sessions held and organized by users in your organization. Types of conference sessions include IM, audio/video, application sharing, web, dial-in/out - third party, and dial-in/out Microsoft.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getskypeforbusinessorganizeractivityusercounts?view=graph-rest-1.0
func (m *GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get usage trends on the number of unique users and type of conference sessions held and organized by users in your organization. Types of conference sessions include IM, audio/video, application sharing, web, dial-in/out - third party, and dial-in/out Microsoft.
// returns a *RequestInformation when successful
func (m *GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder when successful
func (m *GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder) {
    return NewGetskypeforbusinessorganizeractivityusercountswithperiodGetSkypeForBusinessOrganizerActivityUserCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
