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
// GetCompliancePolicyNonComplianceSummaryReportRequestBuilderPostOptions options for Post
type GetCompliancePolicyNonComplianceSummaryReportRequestBuilderPostOptions struct {
    // 
    Body GetCompliancePolicyNonComplianceSummaryReportRequestBodyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
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
func (m *GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) CreatePostRequestInformation(options *GetCompliancePolicyNonComplianceSummaryReportRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getCompliancePolicyNonComplianceSummaryReport
func (m *GetCompliancePolicyNonComplianceSummaryReportRequestBuilder) Post(options *GetCompliancePolicyNonComplianceSummaryReportRequestBuilderPostOptions)(GetCompliancePolicyNonComplianceSummaryReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetCompliancePolicyNonComplianceSummaryReportResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetCompliancePolicyNonComplianceSummaryReportResponseable), nil
}
