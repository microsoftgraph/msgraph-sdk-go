package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder provides operations to manage the custodianSources property of the microsoft.graph.security.ediscoverySearch entity.
type CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderGetQueryParameters custodian sources that are included in the eDiscovery search.
type CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderGetQueryParameters
}
// NewCasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/custodianSources/{dataSource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder instantiates a new CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get custodian sources that are included in the eDiscovery search.
// returns a DataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable), nil
}
// ToGetRequestInformation custodian sources that are included in the eDiscovery search.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemCustodiansourcesDataSourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
