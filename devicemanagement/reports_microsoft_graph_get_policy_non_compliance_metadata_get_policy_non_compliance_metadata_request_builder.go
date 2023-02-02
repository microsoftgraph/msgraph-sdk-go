package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder provides operations to call the getPolicyNonComplianceMetadata method.
type ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilderInternal instantiates a new GetPolicyNonComplianceMetadataRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder) {
    m := &ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getPolicyNonComplianceMetadata";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder instantiates a new GetPolicyNonComplianceMetadataRequestBuilder and sets the default values.
func NewReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPolicyNonComplianceMetadata
func (m *ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder) Post(ctx context.Context, body ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilderPostRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToPostRequestInformation invoke action getPolicyNonComplianceMetadata
func (m *ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilder) ToPostRequestInformation(ctx context.Context, body ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataPostRequestBodyable, requestConfiguration *ReportsMicrosoftGraphGetPolicyNonComplianceMetadataGetPolicyNonComplianceMetadataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
