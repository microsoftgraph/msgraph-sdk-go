package getcompliancesettingnoncompliancereport

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetComplianceSettingNonComplianceReportRequestBuilder provides operations to call the getComplianceSettingNonComplianceReport method.
type GetComplianceSettingNonComplianceReportRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetComplianceSettingNonComplianceReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetComplianceSettingNonComplianceReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetComplianceSettingNonComplianceReportRequestBuilderInternal instantiates a new GetComplianceSettingNonComplianceReportRequestBuilder and sets the default values.
func NewGetComplianceSettingNonComplianceReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetComplianceSettingNonComplianceReportRequestBuilder) {
    m := &GetComplianceSettingNonComplianceReportRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getComplianceSettingNonComplianceReport";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetComplianceSettingNonComplianceReportRequestBuilder instantiates a new GetComplianceSettingNonComplianceReportRequestBuilder and sets the default values.
func NewGetComplianceSettingNonComplianceReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetComplianceSettingNonComplianceReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetComplianceSettingNonComplianceReportRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getComplianceSettingNonComplianceReport
func (m *GetComplianceSettingNonComplianceReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetComplianceSettingNonComplianceReportRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getComplianceSettingNonComplianceReport
func (m *GetComplianceSettingNonComplianceReportRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body GetComplianceSettingNonComplianceReportRequestBodyable, requestConfiguration *GetComplianceSettingNonComplianceReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// PostWithResponseHandler invoke action getComplianceSettingNonComplianceReport
func (m *GetComplianceSettingNonComplianceReportRequestBuilder) PostWithResponseHandler(body GetComplianceSettingNonComplianceReportRequestBodyable, requestConfiguration *GetComplianceSettingNonComplianceReportRequestBuilderPostRequestConfiguration)(GetComplianceSettingNonComplianceReportResponseable, error) {
    return m.PostWithResponseHandler(body, requestConfiguration, nil);
}
// PostWithResponseHandler invoke action getComplianceSettingNonComplianceReport
func (m *GetComplianceSettingNonComplianceReportRequestBuilder) PostWithResponseHandler(body GetComplianceSettingNonComplianceReportRequestBodyable, requestConfiguration *GetComplianceSettingNonComplianceReportRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetComplianceSettingNonComplianceReportResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetComplianceSettingNonComplianceReportResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetComplianceSettingNonComplianceReportResponseable), nil
}
