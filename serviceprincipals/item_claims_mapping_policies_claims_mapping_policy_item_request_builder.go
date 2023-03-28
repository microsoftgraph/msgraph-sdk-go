package serviceprincipals

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\claimsMappingPolicies\{claimsMappingPolicy-id}
type ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal instantiates a new ClaimsMappingPolicyItemRequestBuilder and sets the default values.
func NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) {
    m := &ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/claimsMappingPolicies/{claimsMappingPolicy%2Did}", pathParameters),
    }
    return m
}
// NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder instantiates a new ClaimsMappingPolicyItemRequestBuilder and sets the default values.
func NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of servicePrincipal entities.
func (m *ItemClaimsMappingPoliciesClaimsMappingPolicyItemRequestBuilder) Ref()(*ItemClaimsMappingPoliciesItemRefRequestBuilder) {
    return NewItemClaimsMappingPoliciesItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
