package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
type EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetQueryParameters instances for role eligibility requests.
type EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppScope provides operations to manage the appScope property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesItemAppscopeAppScopeRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) AppScope()(*EntitlementmanagementRoleeligibilityscheduleinstancesItemAppscopeAppScopeRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesItemAppscopeAppScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal instantiates a new EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    m := &EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleInstances/{unifiedRoleEligibilityScheduleInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder instantiates a new EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder and sets the default values.
func NewEntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property roleEligibilityScheduleInstances for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesItemDirectoryscopeDirectoryScopeRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) DirectoryScope()(*EntitlementmanagementRoleeligibilityscheduleinstancesItemDirectoryscopeDirectoryScopeRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesItemDirectoryscopeDirectoryScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get instances for role eligibility requests.
// returns a UnifiedRoleEligibilityScheduleInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, error) {
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
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, error) {
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
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesItemPrincipalRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) Principal()(*EntitlementmanagementRoleeligibilityscheduleinstancesItemPrincipalRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinition provides operations to manage the roleDefinition property of the microsoft.graph.unifiedRoleScheduleInstanceBase entity.
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) RoleDefinition()(*EntitlementmanagementRoleeligibilityscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesItemRoledefinitionRoleDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property roleEligibilityScheduleInstances for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation instances for role eligibility requests.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleEligibilityScheduleInstanceable, requestConfiguration *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder when successful
func (m *EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesUnifiedRoleEligibilityScheduleInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
