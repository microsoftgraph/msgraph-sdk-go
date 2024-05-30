package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder provides operations to manage the tags property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderGetQueryParameters read the properties and relationships of an ediscoveryReviewTag object.
type CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChildTags provides operations to manage the childTags property of the microsoft.graph.security.ediscoveryReviewTag entity.
// returns a *CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder when successful
func (m *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) ChildTags()(*CasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilder) {
    return NewCasesEdiscoverycasesItemTagsItemChildtagsChildTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/tags/{ediscoveryReviewTag%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder instantiates a new CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete remove an ediscoveryReviewTag object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-delete-tags?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read the properties and relationships of an ediscoveryReviewTag object.
// returns a EdiscoveryReviewTagable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewtag-get?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryReviewTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable), nil
}
// Parent provides operations to manage the parent property of the microsoft.graph.security.ediscoveryReviewTag entity.
// returns a *CasesEdiscoverycasesItemTagsItemParentRequestBuilder when successful
func (m *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) Parent()(*CasesEdiscoverycasesItemTagsItemParentRequestBuilder) {
    return NewCasesEdiscoverycasesItemTagsItemParentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an ediscoveryReviewTag object.
// returns a EdiscoveryReviewTagable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoveryreviewtag-update?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable, requestConfiguration *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryReviewTagFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable), nil
}
// ToDeleteRequestInformation remove an ediscoveryReviewTag object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an ediscoveryReviewTag object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an ediscoveryReviewTag object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable, requestConfiguration *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemTagsEdiscoveryReviewTagItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
