package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder provides operations to manage the ediscoveryCases property of the microsoft.graph.security.casesRoot entity.
type CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderGetQueryParameters read the properties and relationships of an ediscoveryCase object.
type CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) {
    m := &CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder instantiates a new CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Custodians provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoverycasesItemCustodiansRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) Custodians()(*CasesEdiscoverycasesItemCustodiansRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete an ediscoveryCase object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-casesroot-delete-ediscoverycases?view=graph-rest-1.0
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an ediscoveryCase object.
// returns a EdiscoveryCaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-get?view=graph-rest-1.0
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable), nil
}
// MicrosoftGraphSecurityClose provides operations to call the close method.
// returns a *CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) MicrosoftGraphSecurityClose()(*CasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilder) {
    return NewCasesEdiscoverycasesItemMicrosoftgraphsecuritycloseMicrosoftGraphSecurityCloseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityReopen provides operations to call the reopen method.
// returns a *CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) MicrosoftGraphSecurityReopen()(*CasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilder) {
    return NewCasesEdiscoverycasesItemMicrosoftgraphsecurityreopenMicrosoftGraphSecurityReopenRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NoncustodialDataSources provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) NoncustodialDataSources()(*CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoverycasesItemOperationsRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) Operations()(*CasesEdiscoverycasesItemOperationsRequestBuilder) {
    return NewCasesEdiscoverycasesItemOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an ediscoveryCase object.
// returns a EdiscoveryCaseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-update?view=graph-rest-1.0
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, requestConfiguration *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable), nil
}
// ReviewSets provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) ReviewSets()(*CasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilder) {
    return NewCasesEdiscoverycasesItemReviewsetsReviewSetsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Searches provides operations to manage the searches property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoverycasesItemSearchesRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) Searches()(*CasesEdiscoverycasesItemSearchesRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoverycasesItemSettingsRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) Settings()(*CasesEdiscoverycasesItemSettingsRequestBuilder) {
    return NewCasesEdiscoverycasesItemSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoverycasesItemTagsRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) Tags()(*CasesEdiscoverycasesItemTagsRequestBuilder) {
    return NewCasesEdiscoverycasesItemTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete an ediscoveryCase object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an ediscoveryCase object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an ediscoveryCase object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, requestConfiguration *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder when successful
func (m *CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder) {
    return NewCasesEdiscoverycasesEdiscoveryCaseItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
