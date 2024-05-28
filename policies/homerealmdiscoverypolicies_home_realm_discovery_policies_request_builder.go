package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.policyRoot entity.
type HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters get a list of homeRealmDiscoveryPolicy objects.
type HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters struct {
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
// HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetQueryParameters
}
// HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByHomeRealmDiscoveryPolicyId provides operations to manage the homeRealmDiscoveryPolicies property of the microsoft.graph.policyRoot entity.
// returns a *HomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder when successful
func (m *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) ByHomeRealmDiscoveryPolicyId(homeRealmDiscoveryPolicyId string)(*HomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if homeRealmDiscoveryPolicyId != "" {
        urlTplParams["homeRealmDiscoveryPolicy%2Did"] = homeRealmDiscoveryPolicyId
    }
    return NewHomerealmdiscoverypoliciesHomeRealmDiscoveryPolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderInternal instantiates a new HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) {
    m := &HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/homeRealmDiscoveryPolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder instantiates a new HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder and sets the default values.
func NewHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *HomerealmdiscoverypoliciesCountRequestBuilder when successful
func (m *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) Count()(*HomerealmdiscoverypoliciesCountRequestBuilder) {
    return NewHomerealmdiscoverypoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of homeRealmDiscoveryPolicy objects.
// returns a HomeRealmDiscoveryPolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/homerealmdiscoverypolicy-list?view=graph-rest-1.0
func (m *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHomeRealmDiscoveryPolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyCollectionResponseable), nil
}
// Post create a new homeRealmDiscoveryPolicy object.
// returns a HomeRealmDiscoveryPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/homerealmdiscoverypolicy-post-homerealmdiscoverypolicies?view=graph-rest-1.0
func (m *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyable, requestConfiguration *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHomeRealmDiscoveryPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyable), nil
}
// ToGetRequestInformation get a list of homeRealmDiscoveryPolicy objects.
// returns a *RequestInformation when successful
func (m *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new homeRealmDiscoveryPolicy object.
// returns a *RequestInformation when successful
func (m *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HomeRealmDiscoveryPolicyable, requestConfiguration *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder when successful
func (m *HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) WithUrl(rawUrl string)(*HomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder) {
    return NewHomerealmdiscoverypoliciesHomeRealmDiscoveryPoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
