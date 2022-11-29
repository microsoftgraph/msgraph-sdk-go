package builders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\homeRealmDiscoveryPolicies\{homeRealmDiscoveryPolicy-id}
type ServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal instantiates a new HomeRealmDiscoveryPolicyItemRequestBuilder and sets the default values.
func NewServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    m := &ServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder instantiates a new HomeRealmDiscoveryPolicyItemRequestBuilder and sets the default values.
func NewServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of servicePrincipal entities.
func (m *ServicePrincipalsItemHomeRealmDiscoveryPoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) Ref()(*ServicePrincipalsItemHomeRealmDiscoveryPoliciesItemRefRequestBuilder) {
    return NewServicePrincipalsItemHomeRealmDiscoveryPoliciesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
