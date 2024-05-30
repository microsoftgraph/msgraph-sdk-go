package identity

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder builds and executes requests for operations under \identity\b2xUserFlows\{b2xIdentityUserFlow-id}\userFlowIdentityProviders\{identityProviderBase-id}
type B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewB2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilderInternal instantiates a new B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder) {
    m := &B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/b2xUserFlows/{b2xIdentityUserFlow%2Did}/userFlowIdentityProviders/{identityProviderBase%2Did}", pathParameters),
    }
    return m
}
// NewB2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder instantiates a new B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder and sets the default values.
func NewB2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewB2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityContainer entities.
// returns a *B2xuserflowsItemUserflowidentityprovidersItemRefRequestBuilder when successful
func (m *B2xuserflowsItemUserflowidentityprovidersIdentityProviderBaseItemRequestBuilder) Ref()(*B2xuserflowsItemUserflowidentityprovidersItemRefRequestBuilder) {
    return NewB2xuserflowsItemUserflowidentityprovidersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
