package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementEntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.roleManagement entity.
type EntitlementmanagementEntitlementManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters container for roles and assignments for entitlement management resources.
type EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters
}
// EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementEntitlementManagementRequestBuilderInternal instantiates a new EntitlementmanagementEntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementmanagementEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    m := &EntitlementmanagementEntitlementManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/entitlementManagement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementEntitlementManagementRequestBuilder instantiates a new EntitlementmanagementEntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementmanagementEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property entitlementManagement for roleManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get container for roles and assignments for entitlement management resources.
// returns a RbacApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable), nil
}
// Patch update the navigation property entitlementManagement in roleManagement
// returns a RbacApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRbacApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable), nil
}
// ResourceNamespaces provides operations to manage the resourceNamespaces property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementResourcenamespacesResourceNamespacesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ResourceNamespaces()(*EntitlementmanagementResourcenamespacesResourceNamespacesRequestBuilder) {
    return NewEntitlementmanagementResourcenamespacesResourceNamespacesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignments provides operations to manage the roleAssignments property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentsRoleAssignmentsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleAssignments()(*EntitlementmanagementRoleassignmentsRoleAssignmentsRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentsRoleAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignmentScheduleInstances provides operations to manage the roleAssignmentScheduleInstances property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleAssignmentScheduleInstances()(*EntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentscheduleinstancesRoleAssignmentScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignmentScheduleRequests provides operations to manage the roleAssignmentScheduleRequests property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentschedulerequestsRoleAssignmentScheduleRequestsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleAssignmentScheduleRequests()(*EntitlementmanagementRoleassignmentschedulerequestsRoleAssignmentScheduleRequestsRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulerequestsRoleAssignmentScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleAssignmentSchedules provides operations to manage the roleAssignmentSchedules property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleassignmentschedulesRoleAssignmentSchedulesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleAssignmentSchedules()(*EntitlementmanagementRoleassignmentschedulesRoleAssignmentSchedulesRequestBuilder) {
    return NewEntitlementmanagementRoleassignmentschedulesRoleAssignmentSchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleDefinitions provides operations to manage the roleDefinitions property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleDefinitions()(*EntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilder) {
    return NewEntitlementmanagementRoledefinitionsRoleDefinitionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleEligibilityScheduleInstances provides operations to manage the roleEligibilityScheduleInstances property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleEligibilityScheduleInstances()(*EntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityscheduleinstancesRoleEligibilityScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleEligibilityScheduleRequests provides operations to manage the roleEligibilityScheduleRequests property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleEligibilityScheduleRequests()(*EntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulerequestsRoleEligibilityScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RoleEligibilitySchedules provides operations to manage the roleEligibilitySchedules property of the microsoft.graph.rbacApplication entity.
// returns a *EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) RoleEligibilitySchedules()(*EntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilder) {
    return NewEntitlementmanagementRoleeligibilityschedulesRoleEligibilitySchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property entitlementManagement for roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation container for roles and assignments for entitlement management resources.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property entitlementManagement in roleManagement
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RbacApplicationable, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementEntitlementManagementRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    return NewEntitlementmanagementEntitlementManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
