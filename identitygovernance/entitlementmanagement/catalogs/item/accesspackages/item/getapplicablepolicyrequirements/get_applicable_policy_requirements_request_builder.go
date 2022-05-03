package getapplicablepolicyrequirements

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GetApplicablePolicyRequirementsRequestBuilder provides operations to call the getApplicablePolicyRequirements method.
type GetApplicablePolicyRequirementsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetApplicablePolicyRequirementsRequestBuilderInternal instantiates a new GetApplicablePolicyRequirementsRequestBuilder and sets the default values.
func NewGetApplicablePolicyRequirementsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetApplicablePolicyRequirementsRequestBuilder) {
    m := &GetApplicablePolicyRequirementsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/microsoft.graph.getApplicablePolicyRequirements";
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
func (m *GetApplicablePolicyRequirementsRequestBuilder) CreatePostRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action getApplicablePolicyRequirements
func (m *GetApplicablePolicyRequirementsRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(requestConfiguration *GetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action getApplicablePolicyRequirements
func (m *GetApplicablePolicyRequirementsRequestBuilder) Post()(GetApplicablePolicyRequirementsResponseable, error) {
    return m.PostWithRequestConfigurationAndResponseHandler(nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action getApplicablePolicyRequirements
func (m *GetApplicablePolicyRequirementsRequestBuilder) PostWithRequestConfigurationAndResponseHandler(requestConfiguration *GetApplicablePolicyRequirementsRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(GetApplicablePolicyRequirementsResponseable, error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetApplicablePolicyRequirementsResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetApplicablePolicyRequirementsResponseable), nil
}
