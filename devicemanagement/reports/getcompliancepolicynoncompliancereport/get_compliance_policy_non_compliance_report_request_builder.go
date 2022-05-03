package getcompliancepolicynoncompliancereport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetCompliancePolicyNonComplianceReportRequestBuilder provides operations to call the getCompliancePolicyNonComplianceReport method.
type GetCompliancePolicyNonComplianceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetCompliancePolicyNonComplianceReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetCompliancePolicyNonComplianceReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetCompliancePolicyNonComplianceReportRequestBuilderInternal instantiates a new GetCompliancePolicyNonComplianceReportRequestBuilder and sets the default values.
func NewGetCompliancePolicyNonComplianceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetCompliancePolicyNonComplianceReportRequestBuilder) {
    m := &GetCompliancePolicyNonComplianceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getCompliancePolicyNonComplianceReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetCompliancePolicyNonComplianceReportRequestBuilder instantiates a new GetCompliancePolicyNonComplianceReportRequestBuilder and sets the default values.
func NewGetCompliancePolicyNonComplianceReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetCompliancePolicyNonComplianceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetCompliancePolicyNonComplianceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getCompliancePolicyNonComplianceReport
func (m *GetCompliancePolicyNonComplianceReportRequestBuilder) CreatePostRequestInformation(body GetCompliancePolicyNonComplianceReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getCompliancePolicyNonComplianceReport
func (m *GetCompliancePolicyNonComplianceReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetCompliancePolicyNonComplianceReportRequestBodyable, requestConfiguration *GetCompliancePolicyNonComplianceReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getCompliancePolicyNonComplianceReport
func (m *GetCompliancePolicyNonComplianceReportRequestBuilder) Post(body GetCompliancePolicyNonComplianceReportRequestBodyable)(GetCompliancePolicyNonComplianceReportResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getCompliancePolicyNonComplianceReport
func (m *GetCompliancePolicyNonComplianceReportRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body GetCompliancePolicyNonComplianceReportRequestBodyable, requestConfiguration *GetCompliancePolicyNonComplianceReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetCompliancePolicyNonComplianceReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetCompliancePolicyNonComplianceReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetCompliancePolicyNonComplianceReportResponseable), nil
}
