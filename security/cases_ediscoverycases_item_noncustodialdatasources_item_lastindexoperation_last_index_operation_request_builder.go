package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder provides operations to manage the lastIndexOperation property of the microsoft.graph.security.ediscoveryNoncustodialDataSource entity.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderGetQueryParameters operation entity that represents the latest indexing for the noncustodial data source.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderGetQueryParameters
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder) {
    m := &CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/{ediscoveryNoncustodialDataSource%2Did}/lastIndexOperation{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get operation entity that represents the latest indexing for the noncustodial data source.
// returns a EdiscoveryIndexOperationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryIndexOperationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryIndexOperationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryIndexOperationable), nil
}
// ToGetRequestInformation operation entity that represents the latest indexing for the noncustodial data source.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
