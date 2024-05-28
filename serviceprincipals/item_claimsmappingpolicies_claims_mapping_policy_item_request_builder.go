package serviceprincipals

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\claimsMappingPolicies\{claimsMappingPolicy-id}
type ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilderInternal instantiates a new ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder and sets the default values.
func NewItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder) {
    m := &ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/claimsMappingPolicies/{claimsMappingPolicy%2Did}", pathParameters),
    }
    return m
}
// NewItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder instantiates a new ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder and sets the default values.
func NewItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of servicePrincipal entities.
// returns a *ItemClaimsmappingpoliciesItemRefRequestBuilder when successful
func (m *ItemClaimsmappingpoliciesClaimsMappingPolicyItemRequestBuilder) Ref()(*ItemClaimsmappingpoliciesItemRefRequestBuilder) {
    return NewItemClaimsmappingpoliciesItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
