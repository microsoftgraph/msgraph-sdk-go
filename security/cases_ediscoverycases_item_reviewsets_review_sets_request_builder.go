package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderGetQueryParameters get a list of ediscoveryReviewSet objects associated with an eDiscovery case.
type CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEdiscoveryReviewSetId provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoverycasesItemReviewsetsEdiscoveryReviewSetItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) ByEdiscoveryReviewSetId(ediscoveryReviewSetId string)(*CasesEdiscoverycasesItemReviewsetsEdiscoveryReviewSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ediscoveryReviewSetId != "" {
        urlTplParams["ediscoveryReviewSet%2Did"] = ediscoveryReviewSetId
    }
    return NewCasesEdiscoverycasesItemReviewsetsEdiscoveryReviewSetItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) {
    m := &CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/reviewSets{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder instantiates a new CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CasesEdiscoverycasesItemReviewsetsCountRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) Count()(*CasesEdiscoverycasesItemReviewsetsCountRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of ediscoveryReviewSet objects associated with an eDiscovery case.
// returns a EdiscoveryReviewSetCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-list-reviewsets?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryReviewSetCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetCollectionResponseable), nil
}
// Post create a new ediscoveryReviewSet object.
// returns a EdiscoveryReviewSetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-post-reviewsets?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) Post(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderPostRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryReviewSetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable), nil
}
// ToGetRequestInformation get a list of ediscoveryReviewSet objects associated with an eDiscovery case.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new ediscoveryReviewSet object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) ToPostRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewSetable, requestConfiguration *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder when successful
func (m *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
