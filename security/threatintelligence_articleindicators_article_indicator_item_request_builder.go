package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder provides operations to manage the articleIndicators property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderGetQueryParameters read the properties and relationships of an articleIndicator object.
type ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderGetQueryParameters
}
// ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Artifact provides operations to manage the artifact property of the microsoft.graph.security.indicator entity.
// returns a *ThreatintelligenceArticleindicatorsItemArtifactRequestBuilder when successful
func (m *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) Artifact()(*ThreatintelligenceArticleindicatorsItemArtifactRequestBuilder) {
    return NewThreatintelligenceArticleindicatorsItemArtifactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderInternal instantiates a new ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder and sets the default values.
func NewThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) {
    m := &ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/articleIndicators/{articleIndicator%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder instantiates a new ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder and sets the default values.
func NewThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property articleIndicators for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an articleIndicator object.
// returns a ArticleIndicatorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-articleindicator-get?view=graph-rest-1.0
func (m *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ArticleIndicatorable, error) {
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
// Patch update the navigation property articleIndicators in security
// returns a ArticleIndicatorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ArticleIndicatorable, requestConfiguration *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ArticleIndicatorable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property articleIndicators for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an articleIndicator object.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property articleIndicators in security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.ArticleIndicatorable, requestConfiguration *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder when successful
func (m *ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder) {
    return NewThreatintelligenceArticleindicatorsArticleIndicatorItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
