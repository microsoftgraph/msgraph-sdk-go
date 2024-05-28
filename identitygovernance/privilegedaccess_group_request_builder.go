package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupRequestBuilder provides operations to manage the group property of the microsoft.graph.privilegedAccessRoot entity.
type PrivilegedaccessGroupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedaccessGroupRequestBuilderGetQueryParameters a group that's governed through Privileged Identity Management (PIM).
type PrivilegedaccessGroupRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssignmentApprovals provides operations to manage the assignmentApprovals property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupAssignmentapprovalsAssignmentApprovalsRequestBuilder when successful
func (m *PrivilegedaccessGroupRequestBuilder) AssignmentApprovals()(*PrivilegedaccessGroupAssignmentapprovalsAssignmentApprovalsRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentapprovalsAssignmentApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentScheduleInstances provides operations to manage the assignmentScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder when successful
func (m *PrivilegedaccessGroupRequestBuilder) AssignmentScheduleInstances()(*PrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentscheduleinstancesAssignmentScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentScheduleRequests provides operations to manage the assignmentScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupAssignmentschedulerequestsAssignmentScheduleRequestsRequestBuilder when successful
func (m *PrivilegedaccessGroupRequestBuilder) AssignmentScheduleRequests()(*PrivilegedaccessGroupAssignmentschedulerequestsAssignmentScheduleRequestsRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulerequestsAssignmentScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentSchedules provides operations to manage the assignmentSchedules property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupAssignmentschedulesAssignmentSchedulesRequestBuilder when successful
func (m *PrivilegedaccessGroupRequestBuilder) AssignmentSchedules()(*PrivilegedaccessGroupAssignmentschedulesAssignmentSchedulesRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulesAssignmentSchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedaccessGroupRequestBuilderInternal instantiates a new PrivilegedaccessGroupRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupRequestBuilder) {
    m := &PrivilegedaccessGroupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupRequestBuilder instantiates a new PrivilegedaccessGroupRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property group for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedaccessGroupRequestBuilderDeleteRequestConfiguration)(error) {
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
// EligibilityScheduleInstances provides operations to manage the eligibilityScheduleInstances property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder when successful
func (m *PrivilegedaccessGroupRequestBuilder) EligibilityScheduleInstances()(*PrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityscheduleinstancesEligibilityScheduleInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EligibilityScheduleRequests provides operations to manage the eligibilityScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder when successful
func (m *PrivilegedaccessGroupRequestBuilder) EligibilityScheduleRequests()(*PrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsEligibilityScheduleRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// EligibilitySchedules provides operations to manage the eligibilitySchedules property of the microsoft.graph.privilegedAccessGroup entity.
// returns a *PrivilegedaccessGroupEligibilityschedulesEligibilitySchedulesRequestBuilder when successful
func (m *PrivilegedaccessGroupRequestBuilder) EligibilitySchedules()(*PrivilegedaccessGroupEligibilityschedulesEligibilitySchedulesRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulesEligibilitySchedulesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a group that's governed through Privileged Identity Management (PIM).
// returns a PrivilegedAccessGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupable), nil
}
// Patch update the navigation property group in identityGovernance
// returns a PrivilegedAccessGroupable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupable, requestConfiguration *PrivilegedaccessGroupRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupable), nil
}
// ToDeleteRequestInformation delete navigation property group for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a group that's governed through Privileged Identity Management (PIM).
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property group in identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupable, requestConfiguration *PrivilegedaccessGroupRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupRequestBuilder when successful
func (m *PrivilegedaccessGroupRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupRequestBuilder) {
    return NewPrivilegedaccessGroupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
