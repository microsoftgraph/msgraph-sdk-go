package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder provides operations to call the getOneDriveUsageAccountDetail method.
type GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal instantiates a new GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder and sets the default values.
func NewGetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, period *string)(*GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    m := &GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOneDriveUsageAccountDetail(period='{period}')", pathParameters),
    }
    if period != nil {
        m.BaseRequestBuilder.PathParameters["period"] = *period
    }
    return m
}
// NewGetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder instantiates a new GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder and sets the default values.
func NewGetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getOneDriveUsageAccountDetail
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) Get(ctx context.Context, requestConfiguration *GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// ToGetRequestInformation invoke function getOneDriveUsageAccountDetail
// returns a *RequestInformation when successful
func (m *GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder when successful
func (m *GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) WithUrl(rawUrl string)(*GetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder) {
    return NewGetonedriveusageaccountdetailwithperiodGetOneDriveUsageAccountDetailWithPeriodRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
