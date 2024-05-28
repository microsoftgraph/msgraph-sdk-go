package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderGetQueryParameters returns a list of case ediscoveryNoncustodialDataSource objects for this case.
type CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderGetQueryParameters struct {
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
// CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByEdiscoveryNoncustodialDataSourceId provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) ByEdiscoveryNoncustodialDataSourceId(ediscoveryNoncustodialDataSourceId string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if ediscoveryNoncustodialDataSourceId != "" {
        urlTplParams["ediscoveryNoncustodialDataSource%2Did"] = ediscoveryNoncustodialDataSourceId
    }
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) {
    m := &CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesCountRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) Count()(*CasesEdiscoverycasesItemNoncustodialdatasourcesCountRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get returns a list of case ediscoveryNoncustodialDataSource objects for this case.
// returns a EdiscoveryNoncustodialDataSourceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceCollectionResponseable, error) {
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
// MicrosoftGraphSecurityApplyHold provides operations to call the applyHold method.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) MicrosoftGraphSecurityApplyHold()(*CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityRemoveHold provides operations to call the removeHold method.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) MicrosoftGraphSecurityRemoveHold()(*CasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create a new ediscoveryNoncustodialDataSource object.
// returns a EdiscoveryNoncustodialDataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-ediscoverycase-post-noncustodialdatasources?view=graph-rest-1.0
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) Post(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderPostRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryNoncustodialDataSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable), nil
}
// ToGetRequestInformation returns a list of case ediscoveryNoncustodialDataSource objects for this case.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new ediscoveryNoncustodialDataSource object.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) ToPostRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesNoncustodialDataSourcesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
