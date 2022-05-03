package getcompliancepolicynoncompliancesummaryreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetCompliancePolicyNonComplianceSummaryReportRequestBuilder provides operations to call the getCompliancePolicyNonComplianceSummaryReport method.
type GetCompliancePolicyNonComplianceSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetCompliancePolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetCompliancePolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal instantiates a new GetCompliancePolicyNonComplianceSummaryReportRequestBuilder and sets the default values.
func NewGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    m := &GetCompliancePolicyNonComplianceSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getCompliancePolicyNonComplianceSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetCompliancePolicyNonComplianceSummaryReportRequestBuilder instantiates a new GetCompliancePolicyNonComplianceSummaryReportRequestBuilder and sets the default values.
func NewGetCompliancePolicyNonComplianceSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetCompliancePolicyNonComplianceSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getCompliancePolicyNonComplianceSummaryReport
func (m *GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) CreatePostRequestInformation(body GetCompliancePolicyNonComplianceSummaryReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getCompliancePolicyNonComplianceSummaryReport
func (m *GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetCompliancePolicyNonComplianceSummaryReportRequestBodyable, requestConfiguration *GetCompliancePolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getCompliancePolicyNonComplianceSummaryReport
func (m *GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) Post(body GetCompliancePolicyNonComplianceSummaryReportRequestBodyable)(GetCompliancePolicyNonComplianceSummaryReportResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getCompliancePolicyNonComplianceSummaryReport
func (m *GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetCompliancePolicyNonComplianceSummaryReportRequestBodyable, requestConfiguration *GetCompliancePolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetCompliancePolicyNonComplianceSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetCompliancePolicyNonComplianceSummaryReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetCompliancePolicyNonComplianceSummaryReportResponseable), nil
}
