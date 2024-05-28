package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\tokenLifetimePolicies\{tokenLifetimePolicy-id}
type ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderInternal instantiates a new ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) {
    m := &ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/tokenLifetimePolicies/{tokenLifetimePolicy%2Did}", pathParameters),
    }
    return m
}
// NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder instantiates a new ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of application entities.
// returns a *ItemTokenlifetimepoliciesItemRefRequestBuilder when successful
func (m *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) Ref()(*ItemTokenlifetimepoliciesItemRefRequestBuilder) {
    return NewItemTokenlifetimepoliciesItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
