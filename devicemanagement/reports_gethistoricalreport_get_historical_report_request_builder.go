package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ReportsGethistoricalreportGetHistoricalReportRequestBuilder provides operations to call the getHistoricalReport method.
type ReportsGethistoricalreportGetHistoricalReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGethistoricalreportGetHistoricalReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGethistoricalreportGetHistoricalReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGethistoricalreportGetHistoricalReportRequestBuilderInternal instantiates a new ReportsGethistoricalreportGetHistoricalReportRequestBuilder and sets the default values.
func NewReportsGethistoricalreportGetHistoricalReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGethistoricalreportGetHistoricalReportRequestBuilder) {
    m := &ReportsGethistoricalreportGetHistoricalReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getHistoricalReport", pathParameters),
    }
    return m
}
// NewReportsGethistoricalreportGetHistoricalReportRequestBuilder instantiates a new ReportsGethistoricalreportGetHistoricalReportRequestBuilder and sets the default values.
func NewReportsGethistoricalreportGetHistoricalReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGethistoricalreportGetHistoricalReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGethistoricalreportGetHistoricalReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post not yet documented
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-reporting-devicemanagementreports-gethistoricalreport?view=graph-rest-1.0
func (m *ReportsGethistoricalreportGetHistoricalReportRequestBuilder) Post(ctx context.Context, body ReportsGethistoricalreportGetHistoricalReportPostRequestBodyable, requestConfiguration *ReportsGethistoricalreportGetHistoricalReportRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation not yet documented
// returns a *RequestInformation when successful
func (m *ReportsGethistoricalreportGetHistoricalReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGethistoricalreportGetHistoricalReportPostRequestBodyable, requestConfiguration *ReportsGethistoricalreportGetHistoricalReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ReportsGethistoricalreportGetHistoricalReportRequestBuilder when successful
func (m *ReportsGethistoricalreportGetHistoricalReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGethistoricalreportGetHistoricalReportRequestBuilder) {
    return NewReportsGethistoricalreportGetHistoricalReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
