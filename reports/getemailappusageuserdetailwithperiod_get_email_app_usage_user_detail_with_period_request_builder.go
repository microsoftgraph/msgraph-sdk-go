package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder provides operations to call the getEmailAppUsageUserDetail method.
type GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal instantiates a new GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    m := &GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getEmailAppUsageUserDetail(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder instantiates a new GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder and sets the default values.
func NewGetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get get details about which activities users performed on the various email apps.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation get details about which activities users performed on the various email apps.
// returns a *RequestInformation when successful
func (m *GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder when successful
func (m *GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder) {
    return NewGetemailappusageuserdetailwithperiodGetEmailAppUsageUserDetailWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
