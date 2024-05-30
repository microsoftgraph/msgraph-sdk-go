package policies

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder builds and executes requests for operations under \policies\featureRolloutPolicies\{featureRolloutPolicy-id}\appliesTo\{directoryObject-id}
type FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewFeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilderInternal instantiates a new FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder and sets the default values.
func NewFeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) {
    m := &FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewFeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder instantiates a new FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder and sets the default values.
func NewFeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of policyRoot entities.
// returns a *FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder when successful
func (m *FeaturerolloutpoliciesItemAppliestoDirectoryObjectItemRequestBuilder) Ref()(*FeaturerolloutpoliciesItemAppliestoItemRefRequestBuilder) {
    return NewFeaturerolloutpoliciesItemAppliestoItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
