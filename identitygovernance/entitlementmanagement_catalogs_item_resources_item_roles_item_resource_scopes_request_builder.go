package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder provides operations to manage the scopes property of the microsoft.graph.accessPackageResource entity.
type EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderGetQueryParameters read-only. Nullable. Supports $expand.
type EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderGetQueryParameters struct {
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
// EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderGetQueryParameters
}
// EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAccessPackageResourceScopeId provides operations to manage the scopes property of the microsoft.graph.accessPackageResource entity.
// returns a *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesAccessPackageResourceScopeItemRequestBuilder when successful
func (m *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) ByAccessPackageResourceScopeId(accessPackageResourceScopeId string)(*EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesAccessPackageResourceScopeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if accessPackageResourceScopeId != "" {
        urlTplParams["accessPackageResourceScope%2Did"] = accessPackageResourceScopeId
    }
    return NewEntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesAccessPackageResourceScopeItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderInternal instantiates a new EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder and sets the default values.
func NewEntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) {
    m := &EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/resources/{accessPackageResource%2Did}/roles/{accessPackageResourceRole%2Did}/resource/scopes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder instantiates a new EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder and sets the default values.
func NewEntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesCountRequestBuilder when successful
func (m *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) Count()(*EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesCountRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read-only. Nullable. Supports $expand.
// returns a AccessPackageResourceScopeCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageResourceScopeCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeCollectionResponseable), nil
}
// Post create new navigation property to scopes for identityGovernance
// returns a AccessPackageResourceScopeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, requestConfiguration *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation read-only. Nullable. Supports $expand.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to scopes for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageResourceScopeable, requestConfiguration *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder when successful
func (m *EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemResourcesItemRolesItemResourceScopesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
