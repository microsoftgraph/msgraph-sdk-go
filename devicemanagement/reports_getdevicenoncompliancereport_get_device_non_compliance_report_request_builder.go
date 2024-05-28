package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder provides operations to call the getDeviceNonComplianceReport method.
type ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilderInternal instantiates a new ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder and sets the default values.
func NewReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder) {
    m := &ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reports/getDeviceNonComplianceReport", pathParameters),
    }
    return m
}
// NewReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder instantiates a new ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder and sets the default values.
func NewReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post not yet documented
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-reporting-devicemanagementreports-getdevicenoncompliancereport?view=graph-rest-1.0
func (m *ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder) Post(ctx context.Context, body ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportPostRequestBodyable, requestConfiguration *ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilderPostRequestConfiguration)([]byte, error) {
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
func (m *ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportPostRequestBodyable, requestConfiguration *ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder when successful
func (m *ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder) WithUrl(rawUrl string)(*ReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder) {
    return NewReportsGetdevicenoncompliancereportGetDeviceNonComplianceReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
