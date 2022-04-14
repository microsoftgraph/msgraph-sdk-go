package getpolicynoncompliancemetadata

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetPolicyNonComplianceMetadataRequestBuilder provides operations to call the getPolicyNonComplianceMetadata method.
type GetPolicyNonComplianceMetadataRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetPolicyNonComplianceMetadataRequestBuilderPostOptions options for Post
type GetPolicyNonComplianceMetadataRequestBuilderPostOptions struct {
    // 
    Body GetPolicyNonComplianceMetadataRequestBodyable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetPolicyNonComplianceMetadataRequestBuilderInternal instantiates a new GetPolicyNonComplianceMetadataRequestBuilder and sets the default values.
func NewGetPolicyNonComplianceMetadataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetPolicyNonComplianceMetadataRequestBuilder) {
    m := &GetPolicyNonComplianceMetadataRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reports/microsoft.graph.getPolicyNonComplianceMetadata";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetPolicyNonComplianceMetadataRequestBuilder instantiates a new GetPolicyNonComplianceMetadataRequestBuilder and sets the default values.
func NewGetPolicyNonComplianceMetadataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetPolicyNonComplianceMetadataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetPolicyNonComplianceMetadataRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getPolicyNonComplianceMetadata
func (m *GetPolicyNonComplianceMetadataRequestBuilder) CreatePostRequestInformation(options *GetPolicyNonComplianceMetadataRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Post invoke action getPolicyNonComplianceMetadata
func (m *GetPolicyNonComplianceMetadataRequestBuilder) Post(options *GetPolicyNonComplianceMetadataRequestBuilderPostOptions)(GetPolicyNonComplianceMetadataResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetPolicyNonComplianceMetadataResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetPolicyNonComplianceMetadataResponseable), nil
}
