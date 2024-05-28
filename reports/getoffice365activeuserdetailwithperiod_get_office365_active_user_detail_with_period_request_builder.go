package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder provides operations to call the getOffice365ActiveUserDetail method.
type Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal instantiates a new Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    m := &Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365ActiveUserDetail(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder instantiates a new Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOffice365ActiveUserDetail
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation invoke function getOffice365ActiveUserDetail
// returns a *RequestInformation when successful
func (m *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder when successful
func (m *Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) WithUrl(rawUrl string)(*Getoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder) {
    return NewGetoffice365activeuserdetailwithperiodGetOffice365ActiveUserDetailWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
