package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// GroupsItemTeamScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type GroupsItemTeamScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemTeamScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemTeamScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemTeamScheduleRequestBuilderGetQueryParameters retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
type GroupsItemTeamScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemTeamScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemTeamScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemTeamScheduleRequestBuilderGetQueryParameters
}
// GroupsItemTeamScheduleRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemTeamScheduleRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsItemTeamScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewGroupsItemTeamScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemTeamScheduleRequestBuilder) {
    m := &GroupsItemTeamScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/team/schedule{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemTeamScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewGroupsItemTeamScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemTeamScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemTeamScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schedule for groups
func (m *GroupsItemTeamScheduleRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemTeamScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
func (m *GroupsItemTeamScheduleRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemTeamScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePutRequestInformation update the navigation property schedule in groups
func (m *GroupsItemTeamScheduleRequestBuilder) CreatePutRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *GroupsItemTeamScheduleRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property schedule for groups
func (m *GroupsItemTeamScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemTeamScheduleRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get retrieve the properties and relationships of a schedule object. The schedule creation process conforms to the One API guideline for resource based long running operations (RELO).When clients use the PUT method, if the schedule is provisioned, the operation updates the schedule; otherwise, the operation starts the schedule provisioning process in the background. During schedule provisioning, clients can use the GET method to get the schedule and look at the `provisionStatus` property for the current state of the provisioning. If the provisioning failed, clients can get additional information from the `provisionStatusCode` property. Clients can also inspect the configuration of the schedule.
func (m *GroupsItemTeamScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemTeamScheduleRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable), nil
}
// OfferShiftRequests provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) OfferShiftRequests()(*GroupsItemTeamScheduleOfferShiftRequestsRequestBuilder) {
    return NewGroupsItemTeamScheduleOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById provides operations to manage the offerShiftRequests property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) OfferShiftRequestsById(id string)(*GroupsItemTeamScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return NewGroupsItemTeamScheduleOfferShiftRequestsOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) OpenShiftChangeRequests()(*GroupsItemTeamScheduleOpenShiftChangeRequestsRequestBuilder) {
    return NewGroupsItemTeamScheduleOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById provides operations to manage the openShiftChangeRequests property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*GroupsItemTeamScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return NewGroupsItemTeamScheduleOpenShiftChangeRequestsOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) OpenShifts()(*GroupsItemTeamScheduleOpenShiftsRequestBuilder) {
    return NewGroupsItemTeamScheduleOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById provides operations to manage the openShifts property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) OpenShiftsById(id string)(*GroupsItemTeamScheduleOpenShiftsOpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift%2Did"] = id
    }
    return NewGroupsItemTeamScheduleOpenShiftsOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Put update the navigation property schedule in groups
func (m *GroupsItemTeamScheduleRequestBuilder) Put(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *GroupsItemTeamScheduleRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.CreatePutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable), nil
}
// SchedulingGroups provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) SchedulingGroups()(*GroupsItemTeamScheduleSchedulingGroupsRequestBuilder) {
    return NewGroupsItemTeamScheduleSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById provides operations to manage the schedulingGroups property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) SchedulingGroupsById(id string)(*GroupsItemTeamScheduleSchedulingGroupsSchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return NewGroupsItemTeamScheduleSchedulingGroupsSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share provides operations to call the share method.
func (m *GroupsItemTeamScheduleRequestBuilder) Share()(*GroupsItemTeamScheduleShareRequestBuilder) {
    return NewGroupsItemTeamScheduleShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) Shifts()(*GroupsItemTeamScheduleShiftsRequestBuilder) {
    return NewGroupsItemTeamScheduleShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById provides operations to manage the shifts property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) ShiftsById(id string)(*GroupsItemTeamScheduleShiftsShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift%2Did"] = id
    }
    return NewGroupsItemTeamScheduleShiftsShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) SwapShiftsChangeRequests()(*GroupsItemTeamScheduleSwapShiftsChangeRequestsRequestBuilder) {
    return NewGroupsItemTeamScheduleSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById provides operations to manage the swapShiftsChangeRequests property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*GroupsItemTeamScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return NewGroupsItemTeamScheduleSwapShiftsChangeRequestsSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) TimeOffReasons()(*GroupsItemTeamScheduleTimeOffReasonsRequestBuilder) {
    return NewGroupsItemTeamScheduleTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById provides operations to manage the timeOffReasons property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) TimeOffReasonsById(id string)(*GroupsItemTeamScheduleTimeOffReasonsTimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason%2Did"] = id
    }
    return NewGroupsItemTeamScheduleTimeOffReasonsTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) TimeOffRequests()(*GroupsItemTeamScheduleTimeOffRequestsRequestBuilder) {
    return NewGroupsItemTeamScheduleTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById provides operations to manage the timeOffRequests property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) TimeOffRequestsById(id string)(*GroupsItemTeamScheduleTimeOffRequestsTimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return NewGroupsItemTeamScheduleTimeOffRequestsTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) TimesOff()(*GroupsItemTeamScheduleTimesOffRequestBuilder) {
    return NewGroupsItemTeamScheduleTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById provides operations to manage the timesOff property of the microsoft.graph.schedule entity.
func (m *GroupsItemTeamScheduleRequestBuilder) TimesOffById(id string)(*GroupsItemTeamScheduleTimesOffTimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff%2Did"] = id
    }
    return NewGroupsItemTeamScheduleTimesOffTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
