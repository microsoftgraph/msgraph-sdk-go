package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder provides operations to manage the collection of policyRoot entities.
type FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteQueryParameters remove an appliesTo on a featureRolloutPolicy object to remove the directoryObject from feature rollout. This API is available in the following national cloud deployments.
type FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteQueryParameters
}
// NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) {
    m := &FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/featureRolloutPolicies/{featureRolloutPolicy%2Did}/appliesTo/{directoryObject%2Did}/$ref{?%40id*}", pathParameters),
    }
    return m
}
// NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove an appliesTo on a featureRolloutPolicy object to remove the directoryObject from feature rollout. This API is available in the following national cloud deployments.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/featurerolloutpolicy-delete-appliesto?view=graph-rest-1.0
func (m *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation remove an appliesTo on a featureRolloutPolicy object to remove the directoryObject from feature rollout. This API is available in the following national cloud deployments.
func (m *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) WithUrl(rawUrl string)(*FeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder) {
    return NewFeatureRolloutPoliciesItemAppliesToItemRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
