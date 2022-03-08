package schedule

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i072c72d2684659fe7cbdb65aef47d6c1918f631d9ab0628e1f00b14b8b8c2188 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/shifts"
    i252a222eb77235b52b18b0c9f8303553c657839a4db9ead11772772fd8630fe4 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/openshiftchangerequests"
    i2547a7f1d67b3d63ac7cd6ad852788ce52638547a6ec23de203e52fa44a92944 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/share"
    i3b0fd31e4b254703a84b00af06b87d47018b7f09e267fbbdec09a339330613fc "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/timesoff"
    i4ee57924c99393c68749c89c3f52497a0a762f96e8c85319a78b103466ae6d15 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/swapshiftschangerequests"
    i6d9dcfcd3f7fb04246276019aa4e2779a6d73d2fdefb23aeb6b7727980ae72a8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/openshifts"
    i98ee724eb951b642ca62569abf1d3cda0e9ca342b1c35e6bb93d1ab0b5946191 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/offershiftrequests"
    ib957f55cc6bceca14111f9d522afca85a7a87ee4297cb1a52c18263455a1be9f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/timeoffreasons"
    ibe75ffd15e1edfe2fb53fc58a7e53cc582e467c94a026b855b863e17e863e4e2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/schedulinggroups"
    ie9c0fe875266534e995b493c7e216891e7e63ab8cb66e0a1bb94a0ed4bd476f0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/timeoffrequests"
    i25835bf2ea4b2783655a613062849327eae2c462a3fb0bb246ab4e81c0aeefcd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/schedulinggroups/item"
    i4371af37049a130938818ac064e9a7234fa37759b90ddd60c724a9f2e40312aa "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/openshifts/item"
    i49f769d63eb0f564f5e3fd6dcde8d8b66f068577823c2849fb223a24af8c4b6d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/shifts/item"
    i4d1975d833f0871642b081140a8181ebd5b9b14b52c8daf8c17e9fc462496d93 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/timeoffrequests/item"
    i7c1a8cd3fa594c1fc7433210be87d46a669f79ffa033b2ed121559b1a7836b36 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/openshiftchangerequests/item"
    i8d73b6115390f96f3f9effc6c115d9a8c01e4bccf3a6ffb524fe8f9d6d189521 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/swapshiftschangerequests/item"
    i91916952245ecadf3a132c25d6712618462caf09e7999a937a6bdd80260050a1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/timesoff/item"
    iad1beb593b32b2b10705cc5c3c2e390baf898c1d2ec6abdecf140dd092a393c5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/timeoffreasons/item"
    iea0766a6107f3958d23f6d55165a212b7c18810c581726007daff6658ae36796 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/schedule/offershiftrequests/item"
)

// ScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ScheduleRequestBuilderDeleteOptions options for Delete
type ScheduleRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ScheduleRequestBuilderGetOptions options for Get
type ScheduleRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ScheduleRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Scheduleable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleRequestBuilder) {
    m := &ScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/team/schedule{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewScheduleRequestBuilder instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property schedule for groups
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation(options *ScheduleRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) CreateGetRequestInformation(options *ScheduleRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property schedule in groups
func (m *ScheduleRequestBuilder) CreatePatchRequestInformation(options *ScheduleRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property schedule for groups
func (m *ScheduleRequestBuilder) Delete(options *ScheduleRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) Get(options *ScheduleRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Scheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateScheduleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Scheduleable), nil
}
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*i98ee724eb951b642ca62569abf1d3cda0e9ca342b1c35e6bb93d1ab0b5946191.OfferShiftRequestsRequestBuilder) {
    return i98ee724eb951b642ca62569abf1d3cda0e9ca342b1c35e6bb93d1ab0b5946191.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.schedule.offerShiftRequests.item collection
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*iea0766a6107f3958d23f6d55165a212b7c18810c581726007daff6658ae36796.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest_id"] = id
    }
    return iea0766a6107f3958d23f6d55165a212b7c18810c581726007daff6658ae36796.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*i252a222eb77235b52b18b0c9f8303553c657839a4db9ead11772772fd8630fe4.OpenShiftChangeRequestsRequestBuilder) {
    return i252a222eb77235b52b18b0c9f8303553c657839a4db9ead11772772fd8630fe4.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.schedule.openShiftChangeRequests.item collection
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i7c1a8cd3fa594c1fc7433210be87d46a669f79ffa033b2ed121559b1a7836b36.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest_id"] = id
    }
    return i7c1a8cd3fa594c1fc7433210be87d46a669f79ffa033b2ed121559b1a7836b36.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) OpenShifts()(*i6d9dcfcd3f7fb04246276019aa4e2779a6d73d2fdefb23aeb6b7727980ae72a8.OpenShiftsRequestBuilder) {
    return i6d9dcfcd3f7fb04246276019aa4e2779a6d73d2fdefb23aeb6b7727980ae72a8.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.schedule.openShifts.item collection
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*i4371af37049a130938818ac064e9a7234fa37759b90ddd60c724a9f2e40312aa.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift_id"] = id
    }
    return i4371af37049a130938818ac064e9a7234fa37759b90ddd60c724a9f2e40312aa.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in groups
func (m *ScheduleRequestBuilder) Patch(options *ScheduleRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ScheduleRequestBuilder) SchedulingGroups()(*ibe75ffd15e1edfe2fb53fc58a7e53cc582e467c94a026b855b863e17e863e4e2.SchedulingGroupsRequestBuilder) {
    return ibe75ffd15e1edfe2fb53fc58a7e53cc582e467c94a026b855b863e17e863e4e2.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.schedule.schedulingGroups.item collection
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*i25835bf2ea4b2783655a613062849327eae2c462a3fb0bb246ab4e81c0aeefcd.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup_id"] = id
    }
    return i25835bf2ea4b2783655a613062849327eae2c462a3fb0bb246ab4e81c0aeefcd.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) Share()(*i2547a7f1d67b3d63ac7cd6ad852788ce52638547a6ec23de203e52fa44a92944.ShareRequestBuilder) {
    return i2547a7f1d67b3d63ac7cd6ad852788ce52638547a6ec23de203e52fa44a92944.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) Shifts()(*i072c72d2684659fe7cbdb65aef47d6c1918f631d9ab0628e1f00b14b8b8c2188.ShiftsRequestBuilder) {
    return i072c72d2684659fe7cbdb65aef47d6c1918f631d9ab0628e1f00b14b8b8c2188.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.schedule.shifts.item collection
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*i49f769d63eb0f564f5e3fd6dcde8d8b66f068577823c2849fb223a24af8c4b6d.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift_id"] = id
    }
    return i49f769d63eb0f564f5e3fd6dcde8d8b66f068577823c2849fb223a24af8c4b6d.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i4ee57924c99393c68749c89c3f52497a0a762f96e8c85319a78b103466ae6d15.SwapShiftsChangeRequestsRequestBuilder) {
    return i4ee57924c99393c68749c89c3f52497a0a762f96e8c85319a78b103466ae6d15.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.schedule.swapShiftsChangeRequests.item collection
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*i8d73b6115390f96f3f9effc6c115d9a8c01e4bccf3a6ffb524fe8f9d6d189521.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest_id"] = id
    }
    return i8d73b6115390f96f3f9effc6c115d9a8c01e4bccf3a6ffb524fe8f9d6d189521.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimeOffReasons()(*ib957f55cc6bceca14111f9d522afca85a7a87ee4297cb1a52c18263455a1be9f.TimeOffReasonsRequestBuilder) {
    return ib957f55cc6bceca14111f9d522afca85a7a87ee4297cb1a52c18263455a1be9f.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.schedule.timeOffReasons.item collection
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*iad1beb593b32b2b10705cc5c3c2e390baf898c1d2ec6abdecf140dd092a393c5.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason_id"] = id
    }
    return iad1beb593b32b2b10705cc5c3c2e390baf898c1d2ec6abdecf140dd092a393c5.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimeOffRequests()(*ie9c0fe875266534e995b493c7e216891e7e63ab8cb66e0a1bb94a0ed4bd476f0.TimeOffRequestsRequestBuilder) {
    return ie9c0fe875266534e995b493c7e216891e7e63ab8cb66e0a1bb94a0ed4bd476f0.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.schedule.timeOffRequests.item collection
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*i4d1975d833f0871642b081140a8181ebd5b9b14b52c8daf8c17e9fc462496d93.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest_id"] = id
    }
    return i4d1975d833f0871642b081140a8181ebd5b9b14b52c8daf8c17e9fc462496d93.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimesOff()(*i3b0fd31e4b254703a84b00af06b87d47018b7f09e267fbbdec09a339330613fc.TimesOffRequestBuilder) {
    return i3b0fd31e4b254703a84b00af06b87d47018b7f09e267fbbdec09a339330613fc.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.schedule.timesOff.item collection
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*i91916952245ecadf3a132c25d6712618462caf09e7999a937a6bdd80260050a1.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff_id"] = id
    }
    return i91916952245ecadf3a132c25d6712618462caf09e7999a937a6bdd80260050a1.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
