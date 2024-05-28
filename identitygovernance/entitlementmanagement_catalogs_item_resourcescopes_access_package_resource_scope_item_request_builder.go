package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder provides operations to manage the resourceScopes property of the microsoft.graph.accessPackageCatalog entity.
type EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters get resourceScopes from identityGovernance
type EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderInternal instantiates a new EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder and sets the default values.
func NewEntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) {
    m := &EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/resourceScopes/{accessPackageResourceScope%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder instantiates a new EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder and sets the default values.
func NewEntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resourceScopes for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get resourceScopes from identityGovernance
// returns a AccessPackageResourceScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable), nil
}
// Patch update the navigation property resourceScopes in identityGovernance
// returns a AccessPackageResourceScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, requestConfiguration *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceScopeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable), nil
}
// Resource provides operations to manage the resource property of the microsoft.graph.accessPackageResourceScope entity.
// returns a *EntitlementmanagementCatalogsItemResourcescopesItemResourceRequestBuilder when successful
func (m *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) Resource()(*EntitlementmanagementCatalogsItemResourcescopesItemResourceRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemResourcescopesItemResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property resourceScopes for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get resourceScopes from identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property resourceScopes in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, requestConfiguration *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder when successful
func (m *EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemResourcescopesAccessPackageResourceScopeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
