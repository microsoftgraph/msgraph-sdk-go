package getpolicynoncompliancesummaryreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetPolicyNonComplianceSummaryReportRequestBuilder provides operations to call the getPolicyNonComplianceSummaryReport method.
type GetPolicyNonComplianceSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetPolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetPolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetPolicyNonComplianceSummaryReportRequestBuilderInternal instantiates a new GetPolicyNonComplianceSummaryReportRequestBuilder and sets the default values.
func NewGetPolicyNonComplianceSummaryReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetPolicyNonComplianceSummaryReportRequestBuilder) {
    m := &GetPolicyNonComplianceSummaryReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getPolicyNonComplianceSummaryReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetPolicyNonComplianceSummaryReportRequestBuilder instantiates a new GetPolicyNonComplianceSummaryReportRequestBuilder and sets the default values.
func NewGetPolicyNonComplianceSummaryReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetPolicyNonComplianceSummaryReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetPolicyNonComplianceSummaryReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getPolicyNonComplianceSummaryReport
func (m *GetPolicyNonComplianceSummaryReportRequestBuilder) CreatePostRequestInformation(body GetPolicyNonComplianceSummaryReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getPolicyNonComplianceSummaryReport
func (m *GetPolicyNonComplianceSummaryReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetPolicyNonComplianceSummaryReportRequestBodyable, requestConfiguration *GetPolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getPolicyNonComplianceSummaryReport
func (m *GetPolicyNonComplianceSummaryReportRequestBuilder) Post(body GetPolicyNonComplianceSummaryReportRequestBodyable)(GetPolicyNonComplianceSummaryReportResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getPolicyNonComplianceSummaryReport
func (m *GetPolicyNonComplianceSummaryReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetPolicyNonComplianceSummaryReportRequestBodyable, requestConfiguration *GetPolicyNonComplianceSummaryReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetPolicyNonComplianceSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetPolicyNonComplianceSummaryReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetPolicyNonComplianceSummaryReportResponseable), nil
}
