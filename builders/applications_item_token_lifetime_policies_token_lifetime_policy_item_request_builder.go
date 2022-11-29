package builders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\tokenLifetimePolicies\{tokenLifetimePolicy-id}
type ApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal instantiates a new TokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    m := &ApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/tokenLifetimePolicies/{tokenLifetimePolicy%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder instantiates a new TokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of application entities.
func (m *ApplicationsItemTokenLifetimePoliciesTokenLifetimePolicyItemRequestBuilder) Ref()(*ApplicationsItemTokenLifetimePoliciesItemRefRequestBuilder) {
    return NewApplicationsItemTokenLifetimePoliciesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
