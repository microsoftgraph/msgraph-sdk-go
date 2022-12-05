package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder provides operations to manage the unifiedGroupSources property of the microsoft.graph.security.ediscoveryCustodian entity.
type SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetQueryParameters data source entity for groups associated with the custodian.
type SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetQueryParameters
}
// SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal instantiates a new UnifiedGroupSourceItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) {
    m := &SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/cases/ediscoveryCases/{ediscoveryCase%2Did}/custodians/{ediscoveryCustodian%2Did}/unifiedGroupSources/{unifiedGroupSource%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder instantiates a new UnifiedGroupSourceItemRequestBuilder and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property unifiedGroupSources for security
func (m *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation data source entity for groups associated with the custodian.
func (m *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property unifiedGroupSources in security
func (m *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable, requestConfiguration *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property unifiedGroupSources for security
func (m *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get data source entity for groups associated with the custodian.
func (m *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateUnifiedGroupSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.security.unifiedGroupSource entity.
func (m *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) Group()(*SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilder) {
    return NewSecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property unifiedGroupSources in security
func (m *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable, requestConfiguration *SecurityCasesEdiscoveryCasesItemCustodiansItemUnifiedGroupSourcesUnifiedGroupSourceItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateUnifiedGroupSourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.UnifiedGroupSourceable), nil
}
