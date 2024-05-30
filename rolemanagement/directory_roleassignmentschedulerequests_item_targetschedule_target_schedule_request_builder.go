package rolemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder provides operations to manage the targetSchedule property of the microsoft.graph.unifiedRoleAssignmentScheduleRequest entity.
type DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetQueryParameters the schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand and $select nested in $expand.
type DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetQueryParameters
}
// NewDirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal instantiates a new DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    m := &DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/roleManagement/directory/roleAssignmentScheduleRequests/{unifiedRoleAssignmentScheduleRequest%2Did}/targetSchedule{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder instantiates a new DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder and sets the default values.
func NewDirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand and $select nested in $expand.
// returns a UnifiedRoleAssignmentScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable, error) {
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
// ToGetRequestInformation the schedule for an eligible role assignment that is referenced through the targetScheduleId property. Supports $expand and $select nested in $expand.
// returns a *RequestInformation when successful
func (m *DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder when successful
func (m *DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) WithUrl(rawUrl string)(*DirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    return NewDirectoryRoleassignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
