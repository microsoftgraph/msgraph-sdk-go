package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder provides operations to manage the assignmentSchedules property of the microsoft.graph.privilegedAccessGroup entity.
type PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderGetQueryParameters read the properties and relationships of a privilegedAccessGroupAssignmentSchedule object.
type PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderGetQueryParameters
}
// PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ActivatedUsing provides operations to manage the activatedUsing property of the microsoft.graph.privilegedAccessGroupAssignmentSchedule entity.
// returns a *PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) ActivatedUsing()(*PrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulesItemActivatedusingActivatedUsingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewPrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderInternal instantiates a new PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) {
    m := &PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/assignmentSchedules/{privilegedAccessGroupAssignmentSchedule%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder instantiates a new PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentSchedules for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a privilegedAccessGroupAssignmentSchedule object.
// returns a PrivilegedAccessGroupAssignmentScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupassignmentschedule-get?view=graph-rest-1.0
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) Get(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.privilegedAccessGroupAssignmentSchedule entity.
// returns a *PrivilegedaccessGroupAssignmentschedulesItemGroupRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) Group()(*PrivilegedaccessGroupAssignmentschedulesItemGroupRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulesItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property assignmentSchedules in identityGovernance
// returns a PrivilegedAccessGroupAssignmentScheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleable, requestConfiguration *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrivilegedAccessGroupAssignmentScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleable), nil
}
// Principal provides operations to manage the principal property of the microsoft.graph.privilegedAccessGroupAssignmentSchedule entity.
// returns a *PrivilegedaccessGroupAssignmentschedulesItemPrincipalRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) Principal()(*PrivilegedaccessGroupAssignmentschedulesItemPrincipalRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulesItemPrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignmentSchedules for identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a privilegedAccessGroupAssignmentSchedule object.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignmentSchedules in identityGovernance
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrivilegedAccessGroupAssignmentScheduleable, requestConfiguration *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder when successful
func (m *PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder) {
    return NewPrivilegedaccessGroupAssignmentschedulesPrivilegedAccessGroupAssignmentScheduleItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
