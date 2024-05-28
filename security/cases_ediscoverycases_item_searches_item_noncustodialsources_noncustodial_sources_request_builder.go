package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder provides operations to manage the noncustodialSources property of the microsoft.graph.security.ediscoverySearch entity.
type CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetQueryParameters noncustodialDataSource sources that are included in the eDiscovery search
type CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetQueryParameters
}
// ByEdiscoveryNoncustodialDataSourceId provides operations to manage the noncustodialSources property of the microsoft.graph.security.ediscoverySearch entity.
// returns a *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) ByEdiscoveryNoncustodialDataSourceId(ediscoveryNoncustodialDataSourceId string)(*CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ediscoveryNoncustodialDataSourceId != "" {
        urlTplParams["ediscoveryNoncustodialDataSource%2Did"] = ediscoveryNoncustodialDataSourceId
    }
    return NewCasesEdiscoverycasesItemSearchesItemNoncustodialsourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) {
    m := &CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/noncustodialSources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder instantiates a new CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesCountRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) Count()(*CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesCountRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemNoncustodialsourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get noncustodialDataSource sources that are included in the eDiscovery search
// returns a EdiscoveryNoncustodialDataSourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryNoncustodialDataSourceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceCollectionResponseable), nil
}
// ToGetRequestInformation noncustodialDataSource sources that are included in the eDiscovery search
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemNoncustodialsourcesNoncustodialSourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
