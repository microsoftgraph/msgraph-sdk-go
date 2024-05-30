package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetQueryParameters retrieve the schedule for a role eligibility operation.
type DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetQueryParameters
}
// DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *DirectoryRoleeligibilityschedulesItemAppscopeAppScopeRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) AppScope()(*DirectoryRoleeligibilityschedulesItemAppscopeAppScopeRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulesItemAppscopeAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal instantiates a new DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    m := &DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleEligibilitySchedules/{unifiedRoleEligibilitySchedule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder instantiates a new DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleEligibilitySchedules for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *DirectoryRoleeligibilityschedulesItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) DirectoryScope()(*DirectoryRoleeligibilityschedulesItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulesItemDirectoryscopeDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve the schedule for a role eligibility operation.
// returns a UnifiedRoleEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleeligibilityschedule-get?view=graph-rest-1.0
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable), nil
}
// Patch update the navigation property roleEligibilitySchedules in roleManagement
// returns a UnifiedRoleEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable, requestConfiguration *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *DirectoryRoleeligibilityschedulesItemPrincipalRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) Principal()(*DirectoryRoleeligibilityschedulesItemPrincipalRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleBase entity.
// returns a *DirectoryRoleeligibilityschedulesItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) RoleDefinition()(*DirectoryRoleeligibilityschedulesItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulesItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleEligibilitySchedules for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the schedule for a role eligibility operation.
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleEligibilitySchedules in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleable, requestConfiguration *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder when successful
func (m *DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder) {
    return NewDirectoryRoleeligibilityschedulesUnifiedRoleEligibilityScheduleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
