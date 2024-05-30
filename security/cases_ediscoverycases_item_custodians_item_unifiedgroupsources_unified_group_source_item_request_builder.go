package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder provides operations to manage the unifiedGroupSources property of the microsoft.graph.security.ediscoveryCustodian entity.
type CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderGetQueryParameters data source entity for groups associated with the custodian.
type CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/unifiedGroupSources/{unifiedGroupSource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder instantiates a new CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property unifiedGroupSources for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get data source entity for groups associated with the custodian.
// returns a UnifiedGroupSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateUnifiedGroupSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.security.unifiedGroupSource entity.
// returns a *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesItemGroupRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) Group()(*CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesItemGroupRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property unifiedGroupSources in security
// returns a UnifiedGroupSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateUnifiedGroupSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable), nil
}
// ToDeleteRequestInformation delete navigation property unifiedGroupSources for security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation data source entity for groups associated with the custodian.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property unifiedGroupSources in security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable, requestConfiguration *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemCustodiansItemUnifiedgroupsourcesUnifiedGroupSourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
