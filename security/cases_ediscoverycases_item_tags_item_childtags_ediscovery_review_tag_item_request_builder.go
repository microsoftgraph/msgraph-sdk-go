package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder provides operations to manage the childTags property of the microsoft.graph.security.ediscoveryReviewTag entity.
type CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderGetQueryParameters returns the tags that are a child of a tag.
type CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderGetQueryParameters
}
// NewCasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/tags/{ediscoveryReviewTag%2Did}/childTags/{ediscoveryReviewTag%2Did1}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder instantiates a new CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get returns the tags that are a child of a tag.
// returns a EdiscoveryReviewTagable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryReviewTagable, error) {
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
// ToGetRequestInformation returns the tags that are a child of a tag.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemTagsItemChildtagsEdiscoveryReviewTagItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
