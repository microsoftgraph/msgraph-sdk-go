package schedule

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i075694b2a8efa878fe00f87889862d739294455f6efff8199d2d989f3ebf4cba "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/schedulinggroups"
    i26a34e0f10f26b3bb561d271a4595092c5f4cd605c28eae42c233eb6f4658521 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/share"
    i2ad47c4606f7a2a188f343b2ac1db602b5e9e1dbfe7bfbea053a304d630990a6 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/timeoffrequests"
    i363cad643714e467edd342c9bf7adf7984461e48c9f613f331bdbda9805c7c29 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/openshiftchangerequests"
    i49bb3bef27839d6b6c772c52798a99b9ec445098226c0138d02c56b11d27fc35 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/timesoff"
    i56ad5fc9582390c984982fd18a9da92f7672dd0720777db0bb4e02d9a66e3478 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/swapshiftschangerequests"
    i8282dff0ec1ee0d108c1eb2b9bd96541dc02ca1291abc179f3125208e078cd76 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/shifts"
    ic60cd14fb26ecb0b8abe2f8d972cf6c98b37478730d7982231e50648bb1cf343 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/offershiftrequests"
    id52f0a18e9b113d2d9a076f88fc4daac41e43ad2dd5d9b1d9b75035c42631fdb "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/openshifts"
    ie5d5271becf3f8b7fe1400f8f22fed56b444f4b4d65ab9bddf1e300e584c7e39 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/timeoffreasons"
    i22d9083c01b60c357dfa3ef52a74e75974fece0d9cc374707c62582e6c670d78 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/offershiftrequests/item"
    i325e1dbc08b3c7de4ec1e0544ea652801f2a7e0e8f122804a01b02f15024054d "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/timeoffreasons/item"
    i53445d1cde7ea9f4c78f439e5c3177819401b1e542ff41fcfe9d701d56688231 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/timeoffrequests/item"
    i568663956637218756e049423f943bfd975e2fbcc942fda0b43fe2d336b24986 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/swapshiftschangerequests/item"
    i6eb8e2f139966e506839381b3b22a154b0cb6ca4a1e8db1c00104fe5676b5221 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/openshiftchangerequests/item"
    ia0480667d398c5699f3671622951a91abc2c43918021edfd9867308a881e5256 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/shifts/item"
    ib8462250f78a8053c2a681770d8f3127c44cb3b9f10de3c8eb69e7de6e4578f2 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/schedulinggroups/item"
    ibcd0f63ee32639148598b0d6925bc07eed0f7e18d3004b7f1cb46fc2e6e26562 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/openshifts/item"
    id54e49d27fe39eb2b19a73b7f41eb9c65318265f3e9be65ced863c7266f4755b "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/schedule/timesoff/item"
)

// ScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ScheduleRequestBuilderDeleteOptions options for Delete
type ScheduleRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ScheduleRequestBuilderGetOptions options for Get
type ScheduleRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ScheduleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ScheduleRequestBuilderGetQueryParameters the schedule of shifts for this team.
type ScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ScheduleRequestBuilderPatchOptions options for Patch
type ScheduleRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleRequestBuilder) {
    m := &ScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team_id}/schedule{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schedule for me
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation(options *ScheduleRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) CreateGetRequestInformation(options *ScheduleRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property schedule in me
func (m *ScheduleRequestBuilder) CreatePatchRequestInformation(options *ScheduleRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property schedule for me
func (m *ScheduleRequestBuilder) Delete(options *ScheduleRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) Get(options *ScheduleRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateScheduleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable), nil
}
// OfferShiftRequests the offerShiftRequests property
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*ic60cd14fb26ecb0b8abe2f8d972cf6c98b37478730d7982231e50648bb1cf343.OfferShiftRequestsRequestBuilder) {
    return ic60cd14fb26ecb0b8abe2f8d972cf6c98b37478730d7982231e50648bb1cf343.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.schedule.offerShiftRequests.item collection
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*i22d9083c01b60c357dfa3ef52a74e75974fece0d9cc374707c62582e6c670d78.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest_id"] = id
    }
    return i22d9083c01b60c357dfa3ef52a74e75974fece0d9cc374707c62582e6c670d78.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests the openShiftChangeRequests property
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*i363cad643714e467edd342c9bf7adf7984461e48c9f613f331bdbda9805c7c29.OpenShiftChangeRequestsRequestBuilder) {
    return i363cad643714e467edd342c9bf7adf7984461e48c9f613f331bdbda9805c7c29.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.schedule.openShiftChangeRequests.item collection
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i6eb8e2f139966e506839381b3b22a154b0cb6ca4a1e8db1c00104fe5676b5221.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest_id"] = id
    }
    return i6eb8e2f139966e506839381b3b22a154b0cb6ca4a1e8db1c00104fe5676b5221.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts the openShifts property
func (m *ScheduleRequestBuilder) OpenShifts()(*id52f0a18e9b113d2d9a076f88fc4daac41e43ad2dd5d9b1d9b75035c42631fdb.OpenShiftsRequestBuilder) {
    return id52f0a18e9b113d2d9a076f88fc4daac41e43ad2dd5d9b1d9b75035c42631fdb.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.schedule.openShifts.item collection
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*ibcd0f63ee32639148598b0d6925bc07eed0f7e18d3004b7f1cb46fc2e6e26562.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift_id"] = id
    }
    return ibcd0f63ee32639148598b0d6925bc07eed0f7e18d3004b7f1cb46fc2e6e26562.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in me
func (m *ScheduleRequestBuilder) Patch(options *ScheduleRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SchedulingGroups the schedulingGroups property
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i075694b2a8efa878fe00f87889862d739294455f6efff8199d2d989f3ebf4cba.SchedulingGroupsRequestBuilder) {
    return i075694b2a8efa878fe00f87889862d739294455f6efff8199d2d989f3ebf4cba.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.schedule.schedulingGroups.item collection
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*ib8462250f78a8053c2a681770d8f3127c44cb3b9f10de3c8eb69e7de6e4578f2.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup_id"] = id
    }
    return ib8462250f78a8053c2a681770d8f3127c44cb3b9f10de3c8eb69e7de6e4578f2.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share the share property
func (m *ScheduleRequestBuilder) Share()(*i26a34e0f10f26b3bb561d271a4595092c5f4cd605c28eae42c233eb6f4658521.ShareRequestBuilder) {
    return i26a34e0f10f26b3bb561d271a4595092c5f4cd605c28eae42c233eb6f4658521.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts the shifts property
func (m *ScheduleRequestBuilder) Shifts()(*i8282dff0ec1ee0d108c1eb2b9bd96541dc02ca1291abc179f3125208e078cd76.ShiftsRequestBuilder) {
    return i8282dff0ec1ee0d108c1eb2b9bd96541dc02ca1291abc179f3125208e078cd76.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.schedule.shifts.item collection
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*ia0480667d398c5699f3671622951a91abc2c43918021edfd9867308a881e5256.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift_id"] = id
    }
    return ia0480667d398c5699f3671622951a91abc2c43918021edfd9867308a881e5256.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests the swapShiftsChangeRequests property
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i56ad5fc9582390c984982fd18a9da92f7672dd0720777db0bb4e02d9a66e3478.SwapShiftsChangeRequestsRequestBuilder) {
    return i56ad5fc9582390c984982fd18a9da92f7672dd0720777db0bb4e02d9a66e3478.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.schedule.swapShiftsChangeRequests.item collection
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*i568663956637218756e049423f943bfd975e2fbcc942fda0b43fe2d336b24986.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest_id"] = id
    }
    return i568663956637218756e049423f943bfd975e2fbcc942fda0b43fe2d336b24986.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons the timeOffReasons property
func (m *ScheduleRequestBuilder) TimeOffReasons()(*ie5d5271becf3f8b7fe1400f8f22fed56b444f4b4d65ab9bddf1e300e584c7e39.TimeOffReasonsRequestBuilder) {
    return ie5d5271becf3f8b7fe1400f8f22fed56b444f4b4d65ab9bddf1e300e584c7e39.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.schedule.timeOffReasons.item collection
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*i325e1dbc08b3c7de4ec1e0544ea652801f2a7e0e8f122804a01b02f15024054d.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason_id"] = id
    }
    return i325e1dbc08b3c7de4ec1e0544ea652801f2a7e0e8f122804a01b02f15024054d.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests the timeOffRequests property
func (m *ScheduleRequestBuilder) TimeOffRequests()(*i2ad47c4606f7a2a188f343b2ac1db602b5e9e1dbfe7bfbea053a304d630990a6.TimeOffRequestsRequestBuilder) {
    return i2ad47c4606f7a2a188f343b2ac1db602b5e9e1dbfe7bfbea053a304d630990a6.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.schedule.timeOffRequests.item collection
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*i53445d1cde7ea9f4c78f439e5c3177819401b1e542ff41fcfe9d701d56688231.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest_id"] = id
    }
    return i53445d1cde7ea9f4c78f439e5c3177819401b1e542ff41fcfe9d701d56688231.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff the timesOff property
func (m *ScheduleRequestBuilder) TimesOff()(*i49bb3bef27839d6b6c772c52798a99b9ec445098226c0138d02c56b11d27fc35.TimesOffRequestBuilder) {
    return i49bb3bef27839d6b6c772c52798a99b9ec445098226c0138d02c56b11d27fc35.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.schedule.timesOff.item collection
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*id54e49d27fe39eb2b19a73b7f41eb9c65318265f3e9be65ced863c7266f4755b.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff_id"] = id
    }
    return id54e49d27fe39eb2b19a73b7f41eb9c65318265f3e9be65ced863c7266f4755b.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
