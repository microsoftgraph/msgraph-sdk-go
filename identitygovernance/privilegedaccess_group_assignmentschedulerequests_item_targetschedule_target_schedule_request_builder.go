package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder provides operations to manage the targetSchedule property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest entity.
type PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetQueryParameters schedule created by this request. Supports $expand.
type PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetQueryParameters
}
// NewPrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/{privilegedAccessGroupAssignmentScheduleRequest%2Did}/targetSchedule{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Get schedule created by this request. Supports $expand.
// returns a PrivilegedAccessGroupEligibilityScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleable), nil
}
// ToGetRequestInformation schedule created by this request. Supports $expand.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
