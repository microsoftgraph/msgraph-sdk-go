package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder provides operations to manage the eligibilityScheduleRequests property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetQueryParameters read the properties and relationships of a privilegedAccessGroupEligibilityScheduleRequest object.
type PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Cancel provides operations to call the cancel method.
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Cancel()(*PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderInternal instantiates a new PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) {
    m := &PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/{privilegedAccessGroupEligibilityScheduleRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder instantiates a new PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property eligibilityScheduleRequests for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a privilegedAccessGroupEligibilityScheduleRequest object.
// returns a PrivilegedAccessGroupEligibilityScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupeligibilityschedulerequest-get?view=graph-rest-1.0
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleRequestable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest entity.
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsItemGroupRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Group()(*PrivilegedaccessGroupEligibilityschedulerequestsItemGroupRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property eligibilityScheduleRequests in identityGovernance
// returns a PrivilegedAccessGroupEligibilityScheduleRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleRequestable, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupEligibilityScheduleRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleRequestable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest entity.
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsItemPrincipalRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) Principal()(*PrivilegedaccessGroupEligibilityschedulerequestsItemPrincipalRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TargetSchedule provides operations to manage the targetSchedule property of the microsoft.graph.privilegedAccessGroupEligibilityScheduleRequest entity.
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) TargetSchedule()(*PrivilegedaccessGroupEligibilityschedulerequestsItemTargetscheduleTargetScheduleRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsItemTargetscheduleTargetScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property eligibilityScheduleRequests for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a privilegedAccessGroupEligibilityScheduleRequest object.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property eligibilityScheduleRequests in identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupEligibilityScheduleRequestable, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsPrivilegedAccessGroupEligibilityScheduleRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
