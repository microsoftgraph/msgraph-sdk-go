package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
type CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters returns a list of case ediscoveryNoncustodialDataSource objects for this case.
type CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetQueryParameters
}
// CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    m := &CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/noncustodialDataSources/{ediscoveryNoncustodialDataSource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewCasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder instantiates a new CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder and sets the default values.
func NewCasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// DataSource provides operations to manage the dataSource property of the microsoft.graph.security.ediscoveryNoncustodialDataSource entity.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) DataSource()(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemDatasourceDataSourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property noncustodialDataSources for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get returns a list of case ediscoveryNoncustodialDataSource objects for this case.
// returns a EdiscoveryNoncustodialDataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// LastIndexOperation provides operations to manage the lastIndexOperation property of the microsoft.graph.security.ediscoveryNoncustodialDataSource entity.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) LastIndexOperation()(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemLastindexoperationLastIndexOperationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityApplyHold provides operations to call the applyHold method.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphSecurityApplyHold()(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityapplyholdMicrosoftGraphSecurityApplyHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityRelease provides operations to call the release method.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphSecurityRelease()(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityreleaseMicrosoftGraphSecurityReleaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityRemoveHold provides operations to call the removeHold method.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphSecurityRemoveHold()(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityremoveholdMicrosoftGraphSecurityRemoveHoldRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphSecurityUpdateIndex provides operations to call the updateIndex method.
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) MicrosoftGraphSecurityUpdateIndex()(*CasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesItemMicrosoftgraphsecurityupdateindexMicrosoftGraphSecurityUpdateIndexRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property noncustodialDataSources in security
// returns a EdiscoveryNoncustodialDataSourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property noncustodialDataSources for security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns a list of case ediscoveryNoncustodialDataSource objects for this case.
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property noncustodialDataSources in security
// returns a *RequestInformation when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryNoncustodialDataSourceable, requestConfiguration *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder when successful
func (m *CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) WithUrl(rawUrl string)(*CasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    return NewCasesEdiscoverycasesItemNoncustodialdatasourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
