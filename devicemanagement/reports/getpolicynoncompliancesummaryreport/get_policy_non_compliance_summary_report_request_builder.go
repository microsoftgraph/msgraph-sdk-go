package getpolicynoncompliancesummaryreport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetPolicyNonComplianceSummaryReportRequestBuilder provides operations to call the getPolicyNonComplianceSummaryReport method.
type GetPolicyNonComplianceSummaryReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetPolicyNonComplianceSummaryReportRequestBuilderPostOptions options for Post
type GetPolicyNonComplianceSummaryReportRequestBuilderPostOptions struct {
    // 
    Body GetPolicyNonComplianceSummaryReportRequestBodyable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
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
func (m *GetPolicyNonComplianceSummaryReportRequestBuilder) CreatePostRequestInformation(options *GetPolicyNonComplianceSummaryReportRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action getPolicyNonComplianceSummaryReport
func (m *GetPolicyNonComplianceSummaryReportRequestBuilder) Post(options *GetPolicyNonComplianceSummaryReportRequestBuilderPostOptions)(GetPolicyNonComplianceSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetPolicyNonComplianceSummaryReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetPolicyNonComplianceSummaryReportResponseable), nil
}
