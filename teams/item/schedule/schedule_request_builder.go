package schedule

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i01519c045ea151004683c20ea1540e9b2eb6fbd39331756c6ee08f264daf2b4d "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/schedulinggroups"
    i01750d724d2b41bff9529d9c52fb8daa5027ac6457daa092e948fe6d5c394b90 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/timeoffrequests"
    i0beb6e8e8ccc36855cae59765ef188e46887acee315756b324fea6463e0fde75 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/shifts"
    i3bddaf301c4fc5b8766f262b90a121596090b82c6065ad0672073263696b869d "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/share"
    i5701c54fcbf36fecf5984439549a540b0068134a34161e9a5bad6893c977e04c "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/swapshiftschangerequests"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i979f41b4670fe4856b006b9ac93a1df55909f5177b6c25f3c7e2dc8cb552296e "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/openshifts"
    iabb36575e0f3f7b01e31d31207c004138ea1a70eef52d2604f41b0f97b7dba02 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/openshiftchangerequests"
    ic303fc988baaa576fdc32b83d9832baf43b33bdbe000c51e721bed4cbac7de4e "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/timeoffreasons"
    if0e6faf931107506fdb8d8969c5fb99869a339da13954a14cc90255ff691fae5 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/timesoff"
    if3cd36465d4e09946347174ffbc52694eb751bbd5204888f91a88e6beeefa786 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/offershiftrequests"
    i077cef4a84be1e18919076de6e3b973ebcf3c54ffb2a0db7aec5fdaf05f74873 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/openshifts/item"
    i3e97a7fafa20db003ef2a75a5e79b9165e6ae74b4a976948a89557917308b1ae "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/openshiftchangerequests/item"
    i4717029473a05fb649af93bbf657da0c15821fab39c94b18fa872f9884772aa1 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/shifts/item"
    ia3f0ee46cad4161730260968933c300c8d5bd2479a2bb471cb27ede57b5aab3e "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/timesoff/item"
    ia798c3c70ac49063707043a913480d584e918c77145a9c2670279dbef879caa6 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/swapshiftschangerequests/item"
    iabeccc97702e6f89f9a15e45a346d12c4f471cc17b4134ada692d86a7a7426f0 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/timeoffrequests/item"
    icd9016d4d1909d7883252b08c395dcf7ac68e8ff8385953a5652dd1e348add0f "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/timeoffreasons/item"
    id08a6e1dd69521cd4e2e6c2c076205fa4095da47b2e1624b4768ab0739c44eb1 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/schedulinggroups/item"
    id6d2fdc6b308e5bcba0d92cb156fdfdf7e8749fb58174b0d3bfafaaba4972a32 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/offershiftrequests/item"
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
    m.urlTemplate = "{+baseurl}/teams/{team_id}/schedule{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property schedule for teams
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
// CreatePatchRequestInformation update the navigation property schedule in teams
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
// Delete delete navigation property schedule for teams
func (m *ScheduleRequestBuilder) Delete(options *ScheduleRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
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
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateScheduleFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Scheduleable), nil
}
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*if3cd36465d4e09946347174ffbc52694eb751bbd5204888f91a88e6beeefa786.OfferShiftRequestsRequestBuilder) {
    return if3cd36465d4e09946347174ffbc52694eb751bbd5204888f91a88e6beeefa786.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OfferShiftRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.offerShiftRequests.item collection
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*id6d2fdc6b308e5bcba0d92cb156fdfdf7e8749fb58174b0d3bfafaaba4972a32.OfferShiftRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest_id"] = id
    }
    return id6d2fdc6b308e5bcba0d92cb156fdfdf7e8749fb58174b0d3bfafaaba4972a32.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*iabb36575e0f3f7b01e31d31207c004138ea1a70eef52d2604f41b0f97b7dba02.OpenShiftChangeRequestsRequestBuilder) {
    return iabb36575e0f3f7b01e31d31207c004138ea1a70eef52d2604f41b0f97b7dba02.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.openShiftChangeRequests.item collection
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i3e97a7fafa20db003ef2a75a5e79b9165e6ae74b4a976948a89557917308b1ae.OpenShiftChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest_id"] = id
    }
    return i3e97a7fafa20db003ef2a75a5e79b9165e6ae74b4a976948a89557917308b1ae.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) OpenShifts()(*i979f41b4670fe4856b006b9ac93a1df55909f5177b6c25f3c7e2dc8cb552296e.OpenShiftsRequestBuilder) {
    return i979f41b4670fe4856b006b9ac93a1df55909f5177b6c25f3c7e2dc8cb552296e.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OpenShiftsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.openShifts.item collection
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*i077cef4a84be1e18919076de6e3b973ebcf3c54ffb2a0db7aec5fdaf05f74873.OpenShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift_id"] = id
    }
    return i077cef4a84be1e18919076de6e3b973ebcf3c54ffb2a0db7aec5fdaf05f74873.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in teams
func (m *ScheduleRequestBuilder) Patch(options *ScheduleRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i01519c045ea151004683c20ea1540e9b2eb6fbd39331756c6ee08f264daf2b4d.SchedulingGroupsRequestBuilder) {
    return i01519c045ea151004683c20ea1540e9b2eb6fbd39331756c6ee08f264daf2b4d.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchedulingGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.schedulingGroups.item collection
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*id08a6e1dd69521cd4e2e6c2c076205fa4095da47b2e1624b4768ab0739c44eb1.SchedulingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup_id"] = id
    }
    return id08a6e1dd69521cd4e2e6c2c076205fa4095da47b2e1624b4768ab0739c44eb1.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) Share()(*i3bddaf301c4fc5b8766f262b90a121596090b82c6065ad0672073263696b869d.ShareRequestBuilder) {
    return i3bddaf301c4fc5b8766f262b90a121596090b82c6065ad0672073263696b869d.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) Shifts()(*i0beb6e8e8ccc36855cae59765ef188e46887acee315756b324fea6463e0fde75.ShiftsRequestBuilder) {
    return i0beb6e8e8ccc36855cae59765ef188e46887acee315756b324fea6463e0fde75.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ShiftsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.shifts.item collection
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*i4717029473a05fb649af93bbf657da0c15821fab39c94b18fa872f9884772aa1.ShiftItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift_id"] = id
    }
    return i4717029473a05fb649af93bbf657da0c15821fab39c94b18fa872f9884772aa1.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i5701c54fcbf36fecf5984439549a540b0068134a34161e9a5bad6893c977e04c.SwapShiftsChangeRequestsRequestBuilder) {
    return i5701c54fcbf36fecf5984439549a540b0068134a34161e9a5bad6893c977e04c.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SwapShiftsChangeRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.swapShiftsChangeRequests.item collection
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*ia798c3c70ac49063707043a913480d584e918c77145a9c2670279dbef879caa6.SwapShiftsChangeRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest_id"] = id
    }
    return ia798c3c70ac49063707043a913480d584e918c77145a9c2670279dbef879caa6.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimeOffReasons()(*ic303fc988baaa576fdc32b83d9832baf43b33bdbe000c51e721bed4cbac7de4e.TimeOffReasonsRequestBuilder) {
    return ic303fc988baaa576fdc32b83d9832baf43b33bdbe000c51e721bed4cbac7de4e.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffReasonsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.timeOffReasons.item collection
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*icd9016d4d1909d7883252b08c395dcf7ac68e8ff8385953a5652dd1e348add0f.TimeOffReasonItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason_id"] = id
    }
    return icd9016d4d1909d7883252b08c395dcf7ac68e8ff8385953a5652dd1e348add0f.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimeOffRequests()(*i01750d724d2b41bff9529d9c52fb8daa5027ac6457daa092e948fe6d5c394b90.TimeOffRequestsRequestBuilder) {
    return i01750d724d2b41bff9529d9c52fb8daa5027ac6457daa092e948fe6d5c394b90.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimeOffRequestsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.timeOffRequests.item collection
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*iabeccc97702e6f89f9a15e45a346d12c4f471cc17b4134ada692d86a7a7426f0.TimeOffRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest_id"] = id
    }
    return iabeccc97702e6f89f9a15e45a346d12c4f471cc17b4134ada692d86a7a7426f0.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimesOff()(*if0e6faf931107506fdb8d8969c5fb99869a339da13954a14cc90255ff691fae5.TimesOffRequestBuilder) {
    return if0e6faf931107506fdb8d8969c5fb99869a339da13954a14cc90255ff691fae5.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TimesOffById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.timesOff.item collection
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*ia3f0ee46cad4161730260968933c300c8d5bd2479a2bb471cb27ede57b5aab3e.TimeOffItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff_id"] = id
    }
    return ia3f0ee46cad4161730260968933c300c8d5bd2479a2bb471cb27ede57b5aab3e.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
