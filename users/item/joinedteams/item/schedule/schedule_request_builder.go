package schedule

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i26e1b177111f8a00c755a9493873e39cc8e54d97a7ed846fbab5e2b98e6086aa "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/shifts"
    i303f2b59aaf94036e4c52e1f6eec495c231b2ea4269c6c620f1968cc7f37d29e "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/offershiftrequests"
    i38b1d0cd4787f4c709deefdebc9011d65ea498ddd72bbe8c6d6753403d1da3a7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/openshifts"
    i4d899631f98079a2a1e192a51bd5f13a4ad4dcf317b82c1ae84986c572375f9d "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/timeoffreasons"
    i6977a5d84cfb8f3a7ca758621fa1892ac67095ec26831e278bead93844d09d6e "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/schedulinggroups"
    i6a080e6c11a55185c1b16614943b747c47e5b588cd0dda9a243670d6689511f5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/timeoffrequests"
    i76a90b8fe290b0eb35da52697b88cc104e65700992ff3f7f1ebf7568fc69de62 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/share"
    i86ba6ba234181696e23fb2cde502d1d1e9e082bb456412adf8467529ce308948 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/swapshiftschangerequests"
    i9ec7ccde7dc719a242dbb335c137486b82a7146d5e1fed1039608fadafbf62f6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/openshiftchangerequests"
    ib294d815ef40c29e26e10f3c1284fa16b88f1c5cd5c8c14103142614dd99c291 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/timesoff"
    i0d1a68bf16e77d2487a937577799dd6167ffee937d965666d69f6ad7f3c09a50 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/schedulinggroups/item"
    i1081b52037899720d0759af03675121003d474a31d48434a51a13f00717d1de6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/timeoffrequests/item"
    i25d31ae09339a87264a818f06856b020d1f7f458ad5bcd376f77616d64336b31 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/openshiftchangerequests/item"
    i544fd4f38bdf20701505814dc209e583f746b4b10fee5d775fea33949a2c1d8a "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/timesoff/item"
    i623731e26f5962ffc40ed47d9738052818c87601af17f1ad3e5cf3bfa9ff0675 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/offershiftrequests/item"
    i75132cc5252e03a3d7b7df052ceea5cf0151431f6a6ebc0935d43698207d46b3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/timeoffreasons/item"
    i78242729c849fb24be50e8793396cac5a437ebdcd77f00949f3f62b2d7bab779 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/openshifts/item"
    if4d4fbf86081722c0a7ebc0e19f822a786b3897e217dc778edba5c1aa53a731c "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/shifts/item"
    if7aca6b7d8b5075b0972208f9d700383af1ec4e05db2a2e3d7672d44075b206f "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/schedule/swapshiftschangerequests/item"
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/joinedTeams/{team_id}/schedule{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property schedule for users
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
// CreatePatchRequestInformation update the navigation property schedule in users
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
// Delete delete navigation property schedule for users
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
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*i303f2b59aaf94036e4c52e1f6eec495c231b2ea4269c6c620f1968cc7f37d29e.OfferShiftRequestsRequestBuilder) {
    return i303f2b59aaf94036e4c52e1f6eec495c231b2ea4269c6c620f1968cc7f37d29e.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.schedule.offerShiftRequests.item collection
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*i623731e26f5962ffc40ed47d9738052818c87601af17f1ad3e5cf3bfa9ff0675.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest_id"] = id
    }
    return i623731e26f5962ffc40ed47d9738052818c87601af17f1ad3e5cf3bfa9ff0675.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests the openShiftChangeRequests property
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*i9ec7ccde7dc719a242dbb335c137486b82a7146d5e1fed1039608fadafbf62f6.OpenShiftChangeRequestsRequestBuilder) {
    return i9ec7ccde7dc719a242dbb335c137486b82a7146d5e1fed1039608fadafbf62f6.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.schedule.openShiftChangeRequests.item collection
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i25d31ae09339a87264a818f06856b020d1f7f458ad5bcd376f77616d64336b31.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest_id"] = id
    }
    return i25d31ae09339a87264a818f06856b020d1f7f458ad5bcd376f77616d64336b31.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts the openShifts property
func (m *ScheduleRequestBuilder) OpenShifts()(*i38b1d0cd4787f4c709deefdebc9011d65ea498ddd72bbe8c6d6753403d1da3a7.OpenShiftsRequestBuilder) {
    return i38b1d0cd4787f4c709deefdebc9011d65ea498ddd72bbe8c6d6753403d1da3a7.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.schedule.openShifts.item collection
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*i78242729c849fb24be50e8793396cac5a437ebdcd77f00949f3f62b2d7bab779.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift_id"] = id
    }
    return i78242729c849fb24be50e8793396cac5a437ebdcd77f00949f3f62b2d7bab779.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in users
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
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i6977a5d84cfb8f3a7ca758621fa1892ac67095ec26831e278bead93844d09d6e.SchedulingGroupsRequestBuilder) {
    return i6977a5d84cfb8f3a7ca758621fa1892ac67095ec26831e278bead93844d09d6e.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.schedule.schedulingGroups.item collection
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*i0d1a68bf16e77d2487a937577799dd6167ffee937d965666d69f6ad7f3c09a50.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup_id"] = id
    }
    return i0d1a68bf16e77d2487a937577799dd6167ffee937d965666d69f6ad7f3c09a50.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share the share property
func (m *ScheduleRequestBuilder) Share()(*i76a90b8fe290b0eb35da52697b88cc104e65700992ff3f7f1ebf7568fc69de62.ShareRequestBuilder) {
    return i76a90b8fe290b0eb35da52697b88cc104e65700992ff3f7f1ebf7568fc69de62.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts the shifts property
func (m *ScheduleRequestBuilder) Shifts()(*i26e1b177111f8a00c755a9493873e39cc8e54d97a7ed846fbab5e2b98e6086aa.ShiftsRequestBuilder) {
    return i26e1b177111f8a00c755a9493873e39cc8e54d97a7ed846fbab5e2b98e6086aa.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.schedule.shifts.item collection
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*if4d4fbf86081722c0a7ebc0e19f822a786b3897e217dc778edba5c1aa53a731c.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift_id"] = id
    }
    return if4d4fbf86081722c0a7ebc0e19f822a786b3897e217dc778edba5c1aa53a731c.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests the swapShiftsChangeRequests property
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i86ba6ba234181696e23fb2cde502d1d1e9e082bb456412adf8467529ce308948.SwapShiftsChangeRequestsRequestBuilder) {
    return i86ba6ba234181696e23fb2cde502d1d1e9e082bb456412adf8467529ce308948.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.schedule.swapShiftsChangeRequests.item collection
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*if7aca6b7d8b5075b0972208f9d700383af1ec4e05db2a2e3d7672d44075b206f.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest_id"] = id
    }
    return if7aca6b7d8b5075b0972208f9d700383af1ec4e05db2a2e3d7672d44075b206f.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons the timeOffReasons property
func (m *ScheduleRequestBuilder) TimeOffReasons()(*i4d899631f98079a2a1e192a51bd5f13a4ad4dcf317b82c1ae84986c572375f9d.TimeOffReasonsRequestBuilder) {
    return i4d899631f98079a2a1e192a51bd5f13a4ad4dcf317b82c1ae84986c572375f9d.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.schedule.timeOffReasons.item collection
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*i75132cc5252e03a3d7b7df052ceea5cf0151431f6a6ebc0935d43698207d46b3.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason_id"] = id
    }
    return i75132cc5252e03a3d7b7df052ceea5cf0151431f6a6ebc0935d43698207d46b3.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests the timeOffRequests property
func (m *ScheduleRequestBuilder) TimeOffRequests()(*i6a080e6c11a55185c1b16614943b747c47e5b588cd0dda9a243670d6689511f5.TimeOffRequestsRequestBuilder) {
    return i6a080e6c11a55185c1b16614943b747c47e5b588cd0dda9a243670d6689511f5.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.schedule.timeOffRequests.item collection
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*i1081b52037899720d0759af03675121003d474a31d48434a51a13f00717d1de6.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest_id"] = id
    }
    return i1081b52037899720d0759af03675121003d474a31d48434a51a13f00717d1de6.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff the timesOff property
func (m *ScheduleRequestBuilder) TimesOff()(*ib294d815ef40c29e26e10f3c1284fa16b88f1c5cd5c8c14103142614dd99c291.TimesOffRequestBuilder) {
    return ib294d815ef40c29e26e10f3c1284fa16b88f1c5cd5c8c14103142614dd99c291.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.schedule.timesOff.item collection
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*i544fd4f38bdf20701505814dc209e583f746b4b10fee5d775fea33949a2c1d8a.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff_id"] = id
    }
    return i544fd4f38bdf20701505814dc209e583f746b4b10fee5d775fea33949a2c1d8a.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
