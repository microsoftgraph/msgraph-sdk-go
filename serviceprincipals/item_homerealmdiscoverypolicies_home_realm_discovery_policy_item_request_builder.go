package serviceprincipals

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\homeRealmDiscoveryPolicies\{homeRealmDiscoveryPolicy-id}
type ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal instantiates a new ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder and sets the default values.
func NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    m := &ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/homeRealmDiscoveryPolicies/{homeRealmDiscoveryPolicy%2Did}", pathParameters),
    }
    return m
}
// NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder instantiates a new ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder and sets the default values.
func NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of servicePrincipal entities.
// returns a *ItemHomerealmdiscoverypoliciesItemRefRequestBuilder when successful
func (m *ItemHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) Ref()(*ItemHomerealmdiscoverypoliciesItemRefRequestBuilder) {
    return NewItemHomerealmdiscoverypoliciesItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
