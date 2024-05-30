package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder provides operations to manage the dataSource property of the microsoft.graph.security.ediscoveryNoncustodialDataSource entity.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderGetQueryParameters user source or SharePoint site data source as noncustodial data source.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) {
    m := &CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/{ediscoveryNoncustodialDataSource%2Did}/dataSource{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property dataSource for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get user source or SharePoint site data source as noncustodial data source.
// returns a DataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable, error) {
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
// Patch update the navigation property dataSource in security
// returns a DataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable, error) {
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
// ToDeleteRequestInformation delete navigation property dataSource for security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation user source or SharePoint site data source as noncustodial data source.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property dataSource in security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.DataSourceable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
