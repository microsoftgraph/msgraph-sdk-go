package builders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\userFlowIdentityProviders\{identityProviderBase-id}
type IdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewIdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal instantiates a new IdentityProviderBaseItemRequestBuilder and sets the default values.
func NewIdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    m := &IdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userFlowIdentityProviders/{identityProviderBase%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder instantiates a new IdentityProviderBaseItemRequestBuilder and sets the default values.
func NewIdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityContainer entities.
func (m *IdentityB2xUserFlowsItemUserFlowIdentityProvidersIdentityProviderBaseItemRequestBuilder) Ref()(*IdentityB2xUserFlowsItemUserFlowIdentityProvidersItemRefRequestBuilder) {
    return NewIdentityB2xUserFlowsItemUserFlowIdentityProvidersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
