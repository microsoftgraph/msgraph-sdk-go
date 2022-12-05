package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder provides operations to manage the ediscoveryCases property of the microsoft.graph.security.casesRoot entity.
type SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetQueryParameters get ediscoveryCases from security
type SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetQueryParameters
}
// SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Close provides operations to call the close method.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Close()(*SecurityCasesEdiscoveryCasesItemCloseRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemCloseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewSecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderInternal instantiates a new EdiscoveryCaseItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) {
    m := &SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder instantiates a new EdiscoveryCaseItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property ediscoveryCases for security
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get ediscoveryCases from security
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property ediscoveryCases in security
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, requestConfiguration *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Custodians provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Custodians()(*SecurityCasesEdiscoveryCasesItemCustodiansRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemCustodiansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CustodiansById provides operations to manage the custodians property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) CustodiansById(id string)(*SecurityCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryCustodian%2Did"] = id
    }
    return NewSecurityCasesEdiscoveryCasesItemCustodiansEdiscoveryCustodianItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property ediscoveryCases for security
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get ediscoveryCases from security
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable), nil
}
// NoncustodialDataSources provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) NoncustodialDataSources()(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NoncustodialDataSourcesById provides operations to manage the noncustodialDataSources property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) NoncustodialDataSourcesById(id string)(*SecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryNoncustodialDataSource%2Did"] = id
    }
    return NewSecurityCasesEdiscoveryCasesItemNoncustodialDataSourcesEdiscoveryNoncustodialDataSourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Operations()(*SecurityCasesEdiscoveryCasesItemOperationsRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) OperationsById(id string)(*SecurityCasesEdiscoveryCasesItemOperationsCaseOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["caseOperation%2Did"] = id
    }
    return NewSecurityCasesEdiscoveryCasesItemOperationsCaseOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property ediscoveryCases in security
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, requestConfiguration *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateEdiscoveryCaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.EdiscoveryCaseable), nil
}
// Reopen provides operations to call the reopen method.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Reopen()(*SecurityCasesEdiscoveryCasesItemReopenRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemReopenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewSets provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ReviewSets()(*SecurityCasesEdiscoveryCasesItemReviewSetsRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReviewSetsById provides operations to manage the reviewSets property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) ReviewSetsById(id string)(*SecurityCasesEdiscoveryCasesItemReviewSetsEdiscoveryReviewSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryReviewSet%2Did"] = id
    }
    return NewSecurityCasesEdiscoveryCasesItemReviewSetsEdiscoveryReviewSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Searches provides operations to manage the searches property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Searches()(*SecurityCasesEdiscoveryCasesItemSearchesRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemSearchesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchesById provides operations to manage the searches property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) SearchesById(id string)(*SecurityCasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoverySearch%2Did"] = id
    }
    return NewSecurityCasesEdiscoveryCasesItemSearchesEdiscoverySearchItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Settings provides operations to manage the settings property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Settings()(*SecurityCasesEdiscoveryCasesItemSettingsRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags provides operations to manage the tags property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) Tags()(*SecurityCasesEdiscoveryCasesItemTagsRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById provides operations to manage the tags property of the microsoft.graph.security.ediscoveryCase entity.
func (m *SecurityCasesEdiscoveryCasesEdiscoveryCaseItemRequestBuilder) TagsById(id string)(*SecurityCasesEdiscoveryCasesItemTagsEdiscoveryReviewTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["ediscoveryReviewTag%2Did"] = id
    }
    return NewSecurityCasesEdiscoveryCasesItemTagsEdiscoveryReviewTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
