package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
type DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetQueryParameters get the instance of a role eligibility.
type DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetQueryParameters
}
// DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleeligibilityscheduleinstancesItemAppscopeAppScopeRequestBuilder when successful
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) AppScope()(*DirectoryRoleeligibilityscheduleinstancesItemAppscopeAppScopeRequestBuilder) {
    return NewDirectoryRoleeligibilityscheduleinstancesItemAppscopeAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal instantiates a new DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    m := &DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleEligibilityScheduleInstances/{unifiedRoleEligibilityScheduleInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder instantiates a new DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder and sets the default values.
func NewDirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleEligibilityScheduleInstances for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *DirectoryRoleeligibilityscheduleinstancesItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) DirectoryScope()(*DirectoryRoleeligibilityscheduleinstancesItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewDirectoryRoleeligibilityscheduleinstancesItemDirectoryscopeDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the instance of a role eligibility.
// returns a UnifiedRoleEligibilityScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/unifiedroleeligibilityscheduleinstance-get?view=graph-rest-1.0
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable), nil
}
// Patch update the navigation property roleEligibilityScheduleInstances in roleManagement
// returns a UnifiedRoleEligibilityScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, requestConfiguration *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleEligibilityScheduleInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleeligibilityscheduleinstancesItemPrincipalRequestBuilder when successful
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Principal()(*DirectoryRoleeligibilityscheduleinstancesItemPrincipalRequestBuilder) {
    return NewDirectoryRoleeligibilityscheduleinstancesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *DirectoryRoleeligibilityscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) RoleDefinition()(*DirectoryRoleeligibilityscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewDirectoryRoleeligibilityscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleEligibilityScheduleInstances for roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the instance of a role eligibility.
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property roleEligibilityScheduleInstances in roleManagement
// returns a *RequestInformation when successful
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, requestConfiguration *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder when successful
func (m *DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    return NewDirectoryRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
