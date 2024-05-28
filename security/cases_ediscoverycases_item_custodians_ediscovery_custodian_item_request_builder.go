package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetQueryParameters read the properties and relationships of an ediscoveryCustodian object.
type CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder instantiates a new CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property custodians for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an ediscoveryCustodian object.
// returns a EdiscoveryCustodianable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycustodian-get?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable), nil
}
// LastIndexOperation provides operations to manage the lastIndexOperation property of the microsoft.graph.security.ediscoveryCustodian entity.
// returns a *CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) LastIndexOperation()(*CasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemLastindexoperationLastIndexOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityActivate provides operations to call the activate method.
// returns a *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityActivate()(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityactivateMicrosoftGraphSecurityActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityApplyHold provides operations to call the applyHold method.
// returns a *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityApplyHold()(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityRelease provides operations to call the release method.
// returns a *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityRelease()(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityRemoveHold provides operations to call the removeHold method.
// returns a *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityRemoveHold()(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityUpdateIndex provides operations to call the updateIndex method.
// returns a *CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) MicrosoftGraphSecurityUpdateIndex()(*CasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property custodians in security
// returns a EdiscoveryCustodianable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, requestConfiguration *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCustodianFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable), nil
}
// SiteSources provides operations to manage the siteSources property of the microsoft.graph.security.ediscoveryCustodian entity.
// returns a *CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) SiteSources()(*CasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemSitesourcesSiteSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property custodians for security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an ediscoveryCustodian object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property custodians in security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCustodianable, requestConfiguration *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UnifiedGroupSources provides operations to manage the unifiedGroupSources property of the microsoft.graph.security.ediscoveryCustodian entity.
// returns a *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) UnifiedGroupSources()(*CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserSources provides operations to manage the userSources property of the microsoft.graph.security.ediscoveryCustodian entity.
// returns a *CasesEdiscoverycasesItemCustodiansItemUsersourcesUserSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) UserSources()(*CasesEdiscoverycasesItemCustodiansItemUsersourcesUserSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemUsersourcesUserSourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansEdiscoveryCustodianItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
