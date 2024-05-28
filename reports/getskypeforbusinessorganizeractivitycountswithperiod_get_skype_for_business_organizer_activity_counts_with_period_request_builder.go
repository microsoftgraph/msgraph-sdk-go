package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder provides operations to call the getSkypeForBusinessOrganizerActivityCounts method.
type GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal instantiates a new GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    m := &GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getSkypeForBusinessOrganizerActivityCounts(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder instantiates a new GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder and sets the default values.
func NewGetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get usage trends on the number and type of conference sessions held and organized by users in your organization. Types of conference sessions include IM, audio/video, application sharing, web, dial-in/out - third party, and Dial-in/out Microsoft.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-getskypeforbusinessorganizeractivitycounts?view=graph-rest-1.0
func (m *GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get usage trends on the number and type of conference sessions held and organized by users in your organization. Types of conference sessions include IM, audio/video, application sharing, web, dial-in/out - third party, and Dial-in/out Microsoft.
// returns a *RequestInformation when successful
func (m *GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder when successful
func (m *GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder) {
    return NewGetskypeforbusinessorganizeractivitycountswithperiodGetSkypeForBusinessOrganizerActivityCountsWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
