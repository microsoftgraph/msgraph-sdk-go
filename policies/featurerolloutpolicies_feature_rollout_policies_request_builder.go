package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder provides operations to manage the featureRolloutPolicies property of the microsoft.graph.policyRoot entity.
type FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderGetQueryParameters retrieve a list of featureRolloutPolicy objects.
type FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderGetQueryParameters
}
// FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByFeatureRolloutPolicyId provides operations to manage the featureRolloutPolicies property of the microsoft.graph.policyRoot entity.
// returns a *FeaturerolloutpoliciesFeatureRolloutPolicyItemRequestBuilder when successful
func (m *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) ByFeatureRolloutPolicyId(featureRolloutPolicyId string)(*FeaturerolloutpoliciesFeatureRolloutPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if featureRolloutPolicyId != "" {
        urlTplParams["featureRolloutPolicy%2Did"] = featureRolloutPolicyId
    }
    return NewFeaturerolloutpoliciesFeatureRolloutPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewFeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderInternal instantiates a new FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder and sets the default values.
func NewFeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) {
    m := &FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/featureRolloutPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder instantiates a new FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder and sets the default values.
func NewFeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *FeaturerolloutpoliciesCountRequestBuilder when successful
func (m *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) Count()(*FeaturerolloutpoliciesCountRequestBuilder) {
    return NewFeaturerolloutpoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of featureRolloutPolicy objects.
// returns a FeatureRolloutPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/featurerolloutpolicies-list?view=graph-rest-1.0
func (m *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FeatureRolloutPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFeatureRolloutPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FeatureRolloutPolicyCollectionResponseable), nil
}
// Post create a new featureRolloutPolicy object.
// returns a FeatureRolloutPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/featurerolloutpolicies-post?view=graph-rest-1.0
func (m *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FeatureRolloutPolicyable, requestConfiguration *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FeatureRolloutPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFeatureRolloutPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FeatureRolloutPolicyable), nil
}
// ToGetRequestInformation retrieve a list of featureRolloutPolicy objects.
// returns a *RequestInformation when successful
func (m *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPostRequestInformation create a new featureRolloutPolicy object.
// returns a *RequestInformation when successful
func (m *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.FeatureRolloutPolicyable, requestConfiguration *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder when successful
func (m *FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) WithUrl(rawUrl string)(*FeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder) {
    return NewFeaturerolloutpoliciesFeatureRolloutPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
