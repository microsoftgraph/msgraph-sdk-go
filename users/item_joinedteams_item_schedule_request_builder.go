package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ItemJoinedteamsItemScheduleRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemJoinedteamsItemScheduleRequestBuilderGetQueryParameters the schedule of shifts for this team.
type ItemJoinedteamsItemScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedteamsItemScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemScheduleRequestBuilderGetQueryParameters
}
// ItemJoinedteamsItemScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedteamsItemScheduleRequestBuilderInternal instantiates a new ItemJoinedteamsItemScheduleRequestBuilder and sets the default values.
func NewItemJoinedteamsItemScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemScheduleRequestBuilder) {
    m := &ItemJoinedteamsItemScheduleRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/schedule{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemScheduleRequestBuilder instantiates a new ItemJoinedteamsItemScheduleRequestBuilder and sets the default values.
func NewItemJoinedteamsItemScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property schedule for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemJoinedteamsItemScheduleRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the schedule of shifts for this team.
// returns a Scheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemScheduleRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable), nil
}
// OfferShiftRequests provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
// returns a *ItemJoinedteamsItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) OfferShiftRequests()(*ItemJoinedteamsItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilder) {
    return NewItemJoinedteamsItemScheduleOffershiftrequestsOfferShiftRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
// returns a *ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestsRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) OpenShiftChangeRequests()(*ItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestsRequestBuilder) {
    return NewItemJoinedteamsItemScheduleOpenshiftchangerequestsOpenShiftChangeRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
// returns a *ItemJoinedteamsItemScheduleOpenshiftsOpenShiftsRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) OpenShifts()(*ItemJoinedteamsItemScheduleOpenshiftsOpenShiftsRequestBuilder) {
    return NewItemJoinedteamsItemScheduleOpenshiftsOpenShiftsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Put update the navigation property schedule in users
// returns a Scheduleable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemScheduleRequestBuilder) Put(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *ItemJoinedteamsItemScheduleRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable), nil
}
// SchedulingGroups provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
// returns a *ItemJoinedteamsItemScheduleSchedulinggroupsSchedulingGroupsRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) SchedulingGroups()(*ItemJoinedteamsItemScheduleSchedulinggroupsSchedulingGroupsRequestBuilder) {
    return NewItemJoinedteamsItemScheduleSchedulinggroupsSchedulingGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Share provides operations to call the share method.
// returns a *ItemJoinedteamsItemScheduleShareRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) Share()(*ItemJoinedteamsItemScheduleShareRequestBuilder) {
    return NewItemJoinedteamsItemScheduleShareRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
// returns a *ItemJoinedteamsItemScheduleShiftsRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) Shifts()(*ItemJoinedteamsItemScheduleShiftsRequestBuilder) {
    return NewItemJoinedteamsItemScheduleShiftsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
// returns a *ItemJoinedteamsItemScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) SwapShiftsChangeRequests()(*ItemJoinedteamsItemScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilder) {
    return NewItemJoinedteamsItemScheduleSwapshiftschangerequestsSwapShiftsChangeRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
// returns a *ItemJoinedteamsItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) TimeOffReasons()(*ItemJoinedteamsItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilder) {
    return NewItemJoinedteamsItemScheduleTimeoffreasonsTimeOffReasonsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
// returns a *ItemJoinedteamsItemScheduleTimeoffrequestsTimeOffRequestsRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) TimeOffRequests()(*ItemJoinedteamsItemScheduleTimeoffrequestsTimeOffRequestsRequestBuilder) {
    return NewItemJoinedteamsItemScheduleTimeoffrequestsTimeOffRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
// returns a *ItemJoinedteamsItemScheduleTimesoffTimesOffRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) TimesOff()(*ItemJoinedteamsItemScheduleTimesoffTimesOffRequestBuilder) {
    return NewItemJoinedteamsItemScheduleTimesoffTimesOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property schedule for users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the schedule of shifts for this team.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPutRequestInformation update the navigation property schedule in users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) ToPutRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *ItemJoinedteamsItemScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemJoinedteamsItemScheduleRequestBuilder when successful
func (m *ItemJoinedteamsItemScheduleRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemScheduleRequestBuilder) {
    return NewItemJoinedteamsItemScheduleRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
