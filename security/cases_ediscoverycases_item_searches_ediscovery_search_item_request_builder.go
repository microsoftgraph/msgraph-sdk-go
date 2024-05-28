package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder provides operations to manage the searches property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderGetQueryParameters read the properties and relationships of an ediscoverySearch object.
type CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AdditionalSources provides operations to manage the additionalSources property of the microsoft.graph.security.ediscoverySearch entity.
// returns a *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesAdditionalSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) AdditionalSources()(*CasesEdiscoverycasesItemSearchesItemAdditionalsourcesAdditionalSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemAdditionalsourcesAdditionalSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AddToReviewSetOperation provides operations to manage the addToReviewSetOperation property of the microsoft.graph.security.ediscoverySearch entity.
// returns a *CasesEdiscoverycasesItemSearchesItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) AddToReviewSetOperation()(*CasesEdiscoverycasesItemSearchesItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemAddtoreviewsetoperationAddToReviewSetOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewCasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder instantiates a new CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustodianSources provides operations to manage the custodianSources property of the microsoft.graph.security.ediscoverySearch entity.
// returns a *CasesEdiscoverycasesItemSearchesItemCustodiansourcesCustodianSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) CustodianSources()(*CasesEdiscoverycasesItemSearchesItemCustodiansourcesCustodianSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemCustodiansourcesCustodianSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete an ediscoverySearch object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-delete-searches?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an ediscoverySearch object.
// returns a EdiscoverySearchable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverysearch-get?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoverySearchFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable), nil
}
// LastEstimateStatisticsOperation provides operations to manage the lastEstimateStatisticsOperation property of the microsoft.graph.security.ediscoverySearch entity.
// returns a *CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) LastEstimateStatisticsOperation()(*CasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemLastestimatestatisticsoperationLastEstimateStatisticsOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityEstimateStatistics provides operations to call the estimateStatistics method.
// returns a *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) MicrosoftGraphSecurityEstimateStatistics()(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecurityestimatestatisticsMicrosoftGraphSecurityEstimateStatisticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityPurgeData provides operations to call the purgeData method.
// returns a *CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataMicrosoftGraphSecurityPurgeDataRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) MicrosoftGraphSecurityPurgeData()(*CasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataMicrosoftGraphSecurityPurgeDataRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemMicrosoftgraphsecuritypurgedataMicrosoftGraphSecurityPurgeDataRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NoncustodialSources provides operations to manage the noncustodialSources property of the microsoft.graph.security.ediscoverySearch entity.
// returns a *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) NoncustodialSources()(*CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an ediscoverySearch object.
// returns a EdiscoverySearchable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverysearch-update?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable, requestConfiguration *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoverySearchFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable), nil
}
// ToDeleteRequestInformation delete an ediscoverySearch object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an ediscoverySearch object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an ediscoverySearch object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoverySearchable, requestConfiguration *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesEdiscoverySearchItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
