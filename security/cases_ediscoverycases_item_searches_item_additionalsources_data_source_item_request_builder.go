package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder provides operations to manage the additionalSources property of the microsoft.graph.security.ediscoverySearch entity.
type CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderGetQueryParameters adds an additional source to the eDiscovery search.
type CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/searches/{ediscoverySearch%2Did}/additionalSources/{dataSource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder instantiates a new CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property additionalSources for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get adds an additional source to the eDiscovery search.
// returns a DataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable, error) {
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
// Patch update the navigation property additionalSources in security
// returns a DataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable, requestConfiguration *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property additionalSources for security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation adds an additional source to the eDiscovery search.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property additionalSources in security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable, requestConfiguration *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemSearchesItemAdditionalsourcesDataSourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
