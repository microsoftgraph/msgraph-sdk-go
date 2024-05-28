package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder provides operations to manage the assignmentScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetQueryParameters read the properties and relationships of a privilegedAccessGroupAssignmentScheduleRequest object.
type PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest entity.
// returns a *PrivilegedaccessGroupAssignmentschedulerequestsItemActivatedusingActivatedUsingRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) ActivatedUsing()(*PrivilegedaccessGroupAssignmentschedulerequestsItemActivatedusingActivatedUsingRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulerequestsItemActivatedusingActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Cancel provides operations to call the cancel method.
// returns a *PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Cancel()(*PrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulerequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentScheduleRequests/{privilegedAccessGroupAssignmentScheduleRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentScheduleRequests for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a privilegedAccessGroupAssignmentScheduleRequest object.
// returns a PrivilegedAccessGroupAssignmentScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupassignmentschedulerequest-get?view=graph-rest-1.0
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest entity.
// returns a *PrivilegedaccessGroupAssignmentschedulerequestsItemGroupRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Group()(*PrivilegedaccessGroupAssignmentschedulerequestsItemGroupRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulerequestsItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property assignmentScheduleRequests in identityGovernance
// returns a PrivilegedAccessGroupAssignmentScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest entity.
// returns a *PrivilegedaccessGroupAssignmentschedulerequestsItemPrincipalRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) Principal()(*PrivilegedaccessGroupAssignmentschedulerequestsItemPrincipalRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulerequestsItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.privilegedAccessGroupAssignmentScheduleRequest entity.
// returns a *PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) TargetSchedule()(*PrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignmentScheduleRequests for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a privilegedAccessGroupAssignmentScheduleRequest object.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignmentScheduleRequests in identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleRequestable, requestConfiguration *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulerequestsPrivilegedAccessGroupAssignmentScheduleRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
