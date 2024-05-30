package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder provides operations to manage the indicators property of the microsoft.graph.security.article entity.
type ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderGetQueryParameters indicators related to this article.
type ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderGetQueryParameters
}
// NewThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderInternal instantiates a new ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder and sets the default values.
func NewThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder) {
    m := &ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/articles/{article%2Did}/indicators/{articleIndicator%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder instantiates a new ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder and sets the default values.
func NewThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get indicators related to this article.
// returns a ArticleIndicatorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ArticleIndicatorable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateArticleIndicatorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ArticleIndicatorable), nil
}
// ToGetRequestInformation indicators related to this article.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder when successful
func (m *ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder) {
    return NewThreatintelligenceArticlesItemIndicatorsArticleIndicatorItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
