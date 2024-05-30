package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\tokenIssuancePolicies\{tokenIssuancePolicy-id}
type ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilderInternal instantiates a new ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder and sets the default values.
func NewItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder) {
    m := &ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/tokenIssuancePolicies/{tokenIssuancePolicy%2Did}", pathParameters),
    }
    return m
}
// NewItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder instantiates a new ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder and sets the default values.
func NewItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of application entities.
// returns a *ItemTokenissuancepoliciesItemRefRequestBuilder when successful
func (m *ItemTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder) Ref()(*ItemTokenissuancepoliciesItemRefRequestBuilder) {
    return NewItemTokenissuancepoliciesItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
