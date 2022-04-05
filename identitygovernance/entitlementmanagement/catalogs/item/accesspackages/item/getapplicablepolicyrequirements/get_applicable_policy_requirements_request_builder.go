package getapplicablepolicyrequirements

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetApplicablePolicyRequirementsRequestBuilder provides operations to call the getApplicablePolicyRequirements method.
type GetApplicablePolicyRequirementsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetApplicablePolicyRequirementsRequestBuilderPostOptions options for Post
type GetApplicablePolicyRequirementsRequestBuilderPostOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewGetApplicablePolicyRequirementsRequestBuilderInternal instantiates a new GetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewGetApplicablePolicyRequirementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetApplicablePolicyRequirementsRequestBuilder) {
    m := &GetApplicablePolicyRequirementsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog_id}/accessPackages/{accessPackage_id}/microsoft.graph.getApplicablePolicyRequirements";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetApplicablePolicyRequirementsRequestBuilder instantiates a new GetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewGetApplicablePolicyRequirementsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetApplicablePolicyRequirementsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetApplicablePolicyRequirementsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action getApplicablePolicyRequirements
func (m *GetApplicablePolicyRequirementsRequestBuilder) CreatePostRequestInformation(options *GetApplicablePolicyRequirementsRequestBuilderPostOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
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
// Post invoke action getApplicablePolicyRequirements
func (m *GetApplicablePolicyRequirementsRequestBuilder) Post(options *GetApplicablePolicyRequirementsRequestBuilderPostOptions)(GetApplicablePolicyRequirementsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetApplicablePolicyRequirementsResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetApplicablePolicyRequirementsResponseable), nil
}
