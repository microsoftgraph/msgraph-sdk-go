package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters in PIM, read the details of a request for an active and persistent role assignment made through the unifiedRoleAssignmentScheduleRequest object.
type DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetQueryParameters
}
// DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
// returns a *DirectoryRoleassignmentschedulerequestsItemActivatedusingActivatedUsingRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*DirectoryRoleassignmentschedulerequestsItemActivatedusingActivatedUsingRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsItemActivatedusingActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
// returns a *DirectoryRoleassignmentschedulerequestsItemAppscopeAppScopeRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) AppScope()(*DirectoryRoleassignmentschedulerequestsItemAppscopeAppScopeRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsItemAppscopeAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *DirectoryRoleassignmentschedulerequestsItemCancelRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Cancel()(*DirectoryRoleassignmentschedulerequestsItemCancelRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    m := &DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder instantiates a new DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentScheduleRequests for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
// returns a *DirectoryRoleassignmentschedulerequestsItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) DirectoryScope()(*DirectoryRoleassignmentschedulerequestsItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsItemDirectoryscopeDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get in PIM, read the details of a request for an active and persistent role assignment made through the unifiedRoleAssignmentScheduleRequest object.
// returns a UnifiedRoleAssignmentScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentschedulerequest-get?view=graph-rest-1.0
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable), nil
}
// Patch update the navigation property roleAssignmentScheduleRequests in roleManagement
// returns a UnifiedRoleAssignmentScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
// returns a *DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) Principal()(*DirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
// returns a *DirectoryRoleassignmentschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) RoleDefinition()(*DirectoryRoleassignmentschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
// returns a *DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentScheduleRequests for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation in PIM, read the details of a request for an active and persistent role assignment made through the unifiedRoleAssignmentScheduleRequest object.
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignmentScheduleRequests in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleRequestable, requestConfiguration *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsUnifiedRoleAssignmentScheduleRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
