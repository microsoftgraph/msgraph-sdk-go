package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderGetQueryParameters resource to grant access to users or groups.
type EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleAssignment entity.
// returns a *EntitlementmanagementRoleassignmentsItemAppscopeAppScopeRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) AppScope()(*EntitlementmanagementRoleassignmentsItemAppscopeAppScopeRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentsItemAppscopeAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderInternal instantiates a new EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    m := &EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleAssignments/{unifiedRoleAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder instantiates a new EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignments for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleAssignment entity.
// returns a *EntitlementmanagementRoleassignmentsItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) DirectoryScope()(*EntitlementmanagementRoleassignmentsItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentsItemDirectoryscopeDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get resource to grant access to users or groups.
// returns a UnifiedRoleAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentable), nil
}
// Patch update the navigation property roleAssignments in roleManagement
// returns a UnifiedRoleAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentable, requestConfiguration *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleAssignment entity.
// returns a *EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) Principal()(*EntitlementmanagementRoleassignmentsItemPrincipalRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentsItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignment entity.
// returns a *EntitlementmanagementRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) RoleDefinition()(*EntitlementmanagementRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentsItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignments for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation resource to grant access to users or groups.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignments in roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentable, requestConfiguration *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentsUnifiedRoleAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
