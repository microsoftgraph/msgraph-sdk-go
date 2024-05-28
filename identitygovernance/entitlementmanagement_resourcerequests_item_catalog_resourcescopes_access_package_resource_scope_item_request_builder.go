package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder provides operations to manage the resourceScopes property of the microsoft.graph.accessPackageCatalog entity.
type EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters get resourceScopes from identityGovernance
type EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderInternal instantiates a new EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder and sets the default values.
func NewEntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) {
    m := &EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/resourceRequests/{accessPackageResourceRequest%2Did}/catalog/resourceScopes/{accessPackageResourceScope%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder instantiates a new EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder and sets the default values.
func NewEntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property resourceScopes for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, error) {
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
func (m *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, error) {
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
// returns a *EntitlementmanagementResourcerequestsItemCatalogResourcescopesItemResourceRequestBuilder when successful
func (m *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) Resource()(*EntitlementmanagementResourcerequestsItemCatalogResourcescopesItemResourceRequestBuilder) {
    return NewEntitlementmanagementResourcerequestsItemCatalogResourcescopesItemResourceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property resourceScopes for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, requestConfiguration *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder when successful
func (m *EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder) {
    return NewEntitlementmanagementResourcerequestsItemCatalogResourcescopesAccessPackageResourceScopeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
