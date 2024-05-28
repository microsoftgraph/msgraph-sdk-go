package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder provides operations to manage the catalogs property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters retrieve the properties and relationships of an accessPackageCatalog object.
type EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackages provides operations to manage the accessPackages property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementmanagementCatalogsItemAccesspackagesAccessPackagesRequestBuilder when successful
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) AccessPackages()(*EntitlementmanagementCatalogsItemAccesspackagesAccessPackagesRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemAccesspackagesAccessPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderInternal instantiates a new EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder and sets the default values.
func NewEntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) {
    m := &EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder instantiates a new EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder and sets the default values.
func NewEntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CustomWorkflowExtensions provides operations to manage the customWorkflowExtensions property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomWorkflowExtensionsRequestBuilder when successful
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) CustomWorkflowExtensions()(*EntitlementmanagementCatalogsItemCustomworkflowextensionsCustomWorkflowExtensionsRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemCustomworkflowextensionsCustomWorkflowExtensionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete an accessPackageCatalog.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackagecatalog-delete?view=graph-rest-1.0
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of an accessPackageCatalog object.
// returns a AccessPackageCatalogable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackagecatalog-get?view=graph-rest-1.0
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCatalogable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCatalogable), nil
}
// Patch update an existing accessPackageCatalog object to change one or more of its properties, such as the display name or description.
// returns a AccessPackageCatalogable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackagecatalog-update?view=graph-rest-1.0
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCatalogable, requestConfiguration *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCatalogable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageCatalogFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCatalogable), nil
}
// ResourceRoles provides operations to manage the resourceRoles property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementmanagementCatalogsItemResourcerolesResourceRolesRequestBuilder when successful
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) ResourceRoles()(*EntitlementmanagementCatalogsItemResourcerolesResourceRolesRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemResourcerolesResourceRolesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resources provides operations to manage the resources property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementmanagementCatalogsItemResourcesRequestBuilder when successful
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) Resources()(*EntitlementmanagementCatalogsItemResourcesRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourceScopes provides operations to manage the resourceScopes property of the microsoft.graph.accessPackageCatalog entity.
// returns a *EntitlementmanagementCatalogsItemResourcescopesResourceScopesRequestBuilder when successful
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) ResourceScopes()(*EntitlementmanagementCatalogsItemResourcescopesResourceScopesRequestBuilder) {
    return NewEntitlementmanagementCatalogsItemResourcescopesResourceScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete an accessPackageCatalog.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of an accessPackageCatalog object.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update an existing accessPackageCatalog object to change one or more of its properties, such as the display name or description.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCatalogable, requestConfiguration *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder when successful
func (m *EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder) {
    return NewEntitlementmanagementCatalogsAccessPackageCatalogItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
