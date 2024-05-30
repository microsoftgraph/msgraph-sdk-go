package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetQueryParameters schedules for active role assignment operations.
type EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.unifiedRoleAssignmentSchedule entity.
// returns a *EntitlementmanagementRoleassignmentschedulesItemActivatedusingActivatedUsingRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) ActivatedUsing()(*EntitlementmanagementRoleassignmentschedulesItemActivatedusingActivatedUsingRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulesItemActivatedusingActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) AppScope()(*EntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulesItemAppscopeAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal instantiates a new EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    m := &EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleAssignmentSchedules/{unifiedRoleAssignmentSchedule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder instantiates a new EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentSchedules for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *EntitlementmanagementRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) DirectoryScope()(*EntitlementmanagementRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulesItemDirectoryscopeDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get schedules for active role assignment operations.
// returns a UnifiedRoleAssignmentScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable), nil
}
// Patch update the navigation property roleAssignmentSchedules in roleManagement
// returns a UnifiedRoleAssignmentScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable, requestConfiguration *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *EntitlementmanagementRoleassignmentschedulesItemPrincipalRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) Principal()(*EntitlementmanagementRoleassignmentschedulesItemPrincipalRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *EntitlementmanagementRoleassignmentschedulesItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) RoleDefinition()(*EntitlementmanagementRoleassignmentschedulesItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulesItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentSchedules for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation schedules for active role assignment operations.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignmentSchedules in roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable, requestConfiguration *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder when successful
func (m *EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulesUnifiedRoleAssignmentScheduleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
