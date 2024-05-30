package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder provides operations to call the getM365AppUserDetail method.
type Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilderInternal instantiates a new Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder) {
    m := &Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getM365AppUserDetail(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder instantiates a new Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get a report that provides the details about which apps and platforms users have used.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get a report that provides the details about which apps and platforms users have used.
// returns a *RequestInformation when successful
func (m *Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder when successful
func (m *Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder) {
    return NewGetm365appuserdetailwithperiodGetM365AppUserDetailWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
