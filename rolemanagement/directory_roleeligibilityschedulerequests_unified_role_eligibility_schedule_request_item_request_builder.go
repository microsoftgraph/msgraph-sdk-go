package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters in PIM, read the details of a request for for a role eligibility request made through the unifiedRoleEligibilityScheduleRequest object.
type DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetQueryParameters
}
// DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) AppScope()(*DirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsItemAppscopeAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *DirectoryRoleeligibilityschedulerequestsItemCancelRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Cancel()(*DirectoryRoleeligibilityschedulerequestsItemCancelRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal instantiates a new DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    m := &DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder instantiates a new DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleEligibilityScheduleRequests for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DirectoryScope provides operations to manage the directoryScope property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *DirectoryRoleeligibilityschedulerequestsItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) DirectoryScope()(*DirectoryRoleeligibilityschedulerequestsItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsItemDirectoryscopeDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get in PIM, read the details of a request for for a role eligibility request made through the unifiedRoleEligibilityScheduleRequest object.
// returns a UnifiedRoleEligibilityScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleeligibilityschedulerequest-get?view=graph-rest-1.0
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable), nil
}
// Patch update the navigation property roleEligibilityScheduleRequests in roleManagement
// returns a UnifiedRoleEligibilityScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *DirectoryRoleeligibilityschedulerequestsItemPrincipalRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) Principal()(*DirectoryRoleeligibilityschedulerequestsItemPrincipalRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *DirectoryRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) RoleDefinition()(*DirectoryRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.unifiedRoleEligibilityScheduleRequest entity.
// returns a *DirectoryRoleeligibilityschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) TargetSchedule()(*DirectoryRoleeligibilityschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleEligibilityScheduleRequests for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation in PIM, read the details of a request for for a role eligibility request made through the unifiedRoleEligibilityScheduleRequest object.
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleEligibilityScheduleRequests in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleRequestable, requestConfiguration *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulerequestsUnifiedRoleEligibilityScheduleRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
