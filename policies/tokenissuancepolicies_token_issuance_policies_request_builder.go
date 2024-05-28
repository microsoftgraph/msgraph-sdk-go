package policies

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.policyRoot entity.
type TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetQueryParameters get a list of tokenIssuancePolicy objects.
type TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetQueryParameters struct {
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
// TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetQueryParameters
}
// TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTokenIssuancePolicyId provides operations to manage the tokenIssuancePolicies property of the microsoft.graph.policyRoot entity.
// returns a *TokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder when successful
func (m *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) ByTokenIssuancePolicyId(tokenIssuancePolicyId string)(*TokenissuancepoliciesTokenIssuancePolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if tokenIssuancePolicyId != "" {
        urlTplParams["tokenIssuancePolicy%2Did"] = tokenIssuancePolicyId
    }
    return NewTokenissuancepoliciesTokenIssuancePolicyItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderInternal instantiates a new TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder and sets the default values.
func NewTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) {
    m := &TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/policies/tokenIssuancePolicies{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder instantiates a new TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder and sets the default values.
func NewTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTokenissuancepoliciesTokenIssuancePoliciesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TokenissuancepoliciesCountRequestBuilder when successful
func (m *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) Count()(*TokenissuancepoliciesCountRequestBuilder) {
    return NewTokenissuancepoliciesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of tokenIssuancePolicy objects.
// returns a TokenIssuancePolicyCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tokenissuancepolicy-list?view=graph-rest-1.0
func (m *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) Get(ctx context.Context, requestConfiguration *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenIssuancePolicyCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTokenIssuancePolicyCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenIssuancePolicyCollectionResponseable), nil
}
// Post create a new tokenIssuancePolicy object.
// returns a TokenIssuancePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tokenissuancepolicy-post-tokenissuancepolicy?view=graph-rest-1.0
func (m *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenIssuancePolicyable, requestConfiguration *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenIssuancePolicyable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTokenIssuancePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenIssuancePolicyable), nil
}
// ToGetRequestInformation get a list of tokenIssuancePolicy objects.
// returns a *RequestInformation when successful
func (m *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new tokenIssuancePolicy object.
// returns a *RequestInformation when successful
func (m *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenIssuancePolicyable, requestConfiguration *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder when successful
func (m *TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) WithUrl(rawUrl string)(*TokenissuancepoliciesTokenIssuancePoliciesRequestBuilder) {
    return NewTokenissuancepoliciesTokenIssuancePoliciesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
