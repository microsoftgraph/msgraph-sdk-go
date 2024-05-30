package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters get the instance of an active role assignment.
type DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetQueryParameters
}
// DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.unifiedRoleAssignmentScheduleInstance entity.
// returns a *DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder when successful
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ActivatedUsing()(*DirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilder) {
    return NewDirectoryRoleassignmentscheduleinstancesItemActivatedusingActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleassignmentscheduleinstancesItemAppscopeAppScopeRequestBuilder when successful
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) AppScope()(*DirectoryRoleassignmentscheduleinstancesItemAppscopeAppScopeRequestBuilder) {
    return NewDirectoryRoleassignmentscheduleinstancesItemAppscopeAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal instantiates a new DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    m := &DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentScheduleInstances/{unifiedRoleAssignmentScheduleInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder instantiates a new DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleAssignmentScheduleInstances for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleassignmentscheduleinstancesItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) DirectoryScope()(*DirectoryRoleassignmentscheduleinstancesItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewDirectoryRoleassignmentscheduleinstancesItemDirectoryscopeDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the instance of an active role assignment.
// returns a UnifiedRoleAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleassignmentscheduleinstance-get?view=graph-rest-1.0
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable), nil
}
// Patch update the navigation property roleAssignmentScheduleInstances in roleManagement
// returns a UnifiedRoleAssignmentScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable, requestConfiguration *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleassignmentscheduleinstancesItemPrincipalRequestBuilder when successful
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) Principal()(*DirectoryRoleassignmentscheduleinstancesItemPrincipalRequestBuilder) {
    return NewDirectoryRoleassignmentscheduleinstancesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) RoleDefinition()(*DirectoryRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewDirectoryRoleassignmentscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleAssignmentScheduleInstances for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the instance of an active role assignment.
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleAssignmentScheduleInstances in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleInstanceable, requestConfiguration *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder when successful
func (m *DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder) {
    return NewDirectoryRoleassignmentscheduleinstancesUnifiedRoleAssignmentScheduleInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
