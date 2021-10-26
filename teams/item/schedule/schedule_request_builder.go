package schedule

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i01519c045ea151004683c20ea1540e9b2eb6fbd39331756c6ee08f264daf2b4d "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/schedulinggroups"
    i01750d724d2b41bff9529d9c52fb8daa5027ac6457daa092e948fe6d5c394b90 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/timeoffrequests"
    i0beb6e8e8ccc36855cae59765ef188e46887acee315756b324fea6463e0fde75 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/shifts"
    i3bddaf301c4fc5b8766f262b90a121596090b82c6065ad0672073263696b869d "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/share"
    i5701c54fcbf36fecf5984439549a540b0068134a34161e9a5bad6893c977e04c "github.com/microsoftgraph/msgraph-sdk-go/teams/item/schedule/swapshiftschangerequests"
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

// Builds and executes requests for operations under \teams\{team-id}\schedule
type ScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The schedule of shifts for this team.
type ScheduleRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new ScheduleRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleRequestBuilder) {
    m := &ScheduleRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/teams/{team_id}/schedule{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ScheduleRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewScheduleRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ScheduleRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewScheduleRequestBuilderInternal(urlParams, requestAdapter)
}
// The schedule of shifts for this team.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The schedule of shifts for this team.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ScheduleRequestBuilder) CreateGetRequestInformation(q func (value *ScheduleRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ScheduleRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The schedule of shifts for this team.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ScheduleRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Schedule, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// The schedule of shifts for this team.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ScheduleRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// The schedule of shifts for this team.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ScheduleRequestBuilder) Get(q func (value *ScheduleRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Schedule, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSchedule() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Schedule), nil
}
func (m *ScheduleRequestBuilder) OfferShiftRequests()(*if3cd36465d4e09946347174ffbc52694eb751bbd5204888f91a88e6beeefa786.OfferShiftRequestsRequestBuilder) {
    return if3cd36465d4e09946347174ffbc52694eb751bbd5204888f91a88e6beeefa786.NewOfferShiftRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.offerShiftRequests.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ScheduleRequestBuilder) OfferShiftRequestsById(id string)(*id6d2fdc6b308e5bcba0d92cb156fdfdf7e8749fb58174b0d3bfafaaba4972a32.OfferShiftRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["offerShiftRequest_id"] = id
    }
    return id6d2fdc6b308e5bcba0d92cb156fdfdf7e8749fb58174b0d3bfafaaba4972a32.NewOfferShiftRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) OpenShiftChangeRequests()(*iabb36575e0f3f7b01e31d31207c004138ea1a70eef52d2604f41b0f97b7dba02.OpenShiftChangeRequestsRequestBuilder) {
    return iabb36575e0f3f7b01e31d31207c004138ea1a70eef52d2604f41b0f97b7dba02.NewOpenShiftChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.openShiftChangeRequests.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ScheduleRequestBuilder) OpenShiftChangeRequestsById(id string)(*i3e97a7fafa20db003ef2a75a5e79b9165e6ae74b4a976948a89557917308b1ae.OpenShiftChangeRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShiftChangeRequest_id"] = id
    }
    return i3e97a7fafa20db003ef2a75a5e79b9165e6ae74b4a976948a89557917308b1ae.NewOpenShiftChangeRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) OpenShifts()(*i979f41b4670fe4856b006b9ac93a1df55909f5177b6c25f3c7e2dc8cb552296e.OpenShiftsRequestBuilder) {
    return i979f41b4670fe4856b006b9ac93a1df55909f5177b6c25f3c7e2dc8cb552296e.NewOpenShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.openShifts.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ScheduleRequestBuilder) OpenShiftsById(id string)(*i077cef4a84be1e18919076de6e3b973ebcf3c54ffb2a0db7aec5fdaf05f74873.OpenShiftRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["openShift_id"] = id
    }
    return i077cef4a84be1e18919076de6e3b973ebcf3c54ffb2a0db7aec5fdaf05f74873.NewOpenShiftRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The schedule of shifts for this team.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ScheduleRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Schedule, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *ScheduleRequestBuilder) SchedulingGroups()(*i01519c045ea151004683c20ea1540e9b2eb6fbd39331756c6ee08f264daf2b4d.SchedulingGroupsRequestBuilder) {
    return i01519c045ea151004683c20ea1540e9b2eb6fbd39331756c6ee08f264daf2b4d.NewSchedulingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.schedulingGroups.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ScheduleRequestBuilder) SchedulingGroupsById(id string)(*id08a6e1dd69521cd4e2e6c2c076205fa4095da47b2e1624b4768ab0739c44eb1.SchedulingGroupRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["schedulingGroup_id"] = id
    }
    return id08a6e1dd69521cd4e2e6c2c076205fa4095da47b2e1624b4768ab0739c44eb1.NewSchedulingGroupRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) Share()(*i3bddaf301c4fc5b8766f262b90a121596090b82c6065ad0672073263696b869d.ShareRequestBuilder) {
    return i3bddaf301c4fc5b8766f262b90a121596090b82c6065ad0672073263696b869d.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) Shifts()(*i0beb6e8e8ccc36855cae59765ef188e46887acee315756b324fea6463e0fde75.ShiftsRequestBuilder) {
    return i0beb6e8e8ccc36855cae59765ef188e46887acee315756b324fea6463e0fde75.NewShiftsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.shifts.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ScheduleRequestBuilder) ShiftsById(id string)(*i4717029473a05fb649af93bbf657da0c15821fab39c94b18fa872f9884772aa1.ShiftRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["shift_id"] = id
    }
    return i4717029473a05fb649af93bbf657da0c15821fab39c94b18fa872f9884772aa1.NewShiftRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequests()(*i5701c54fcbf36fecf5984439549a540b0068134a34161e9a5bad6893c977e04c.SwapShiftsChangeRequestsRequestBuilder) {
    return i5701c54fcbf36fecf5984439549a540b0068134a34161e9a5bad6893c977e04c.NewSwapShiftsChangeRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.swapShiftsChangeRequests.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ScheduleRequestBuilder) SwapShiftsChangeRequestsById(id string)(*ia798c3c70ac49063707043a913480d584e918c77145a9c2670279dbef879caa6.SwapShiftsChangeRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["swapShiftsChangeRequest_id"] = id
    }
    return ia798c3c70ac49063707043a913480d584e918c77145a9c2670279dbef879caa6.NewSwapShiftsChangeRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimeOffReasons()(*ic303fc988baaa576fdc32b83d9832baf43b33bdbe000c51e721bed4cbac7de4e.TimeOffReasonsRequestBuilder) {
    return ic303fc988baaa576fdc32b83d9832baf43b33bdbe000c51e721bed4cbac7de4e.NewTimeOffReasonsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.timeOffReasons.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ScheduleRequestBuilder) TimeOffReasonsById(id string)(*icd9016d4d1909d7883252b08c395dcf7ac68e8ff8385953a5652dd1e348add0f.TimeOffReasonRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffReason_id"] = id
    }
    return icd9016d4d1909d7883252b08c395dcf7ac68e8ff8385953a5652dd1e348add0f.NewTimeOffReasonRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimeOffRequests()(*i01750d724d2b41bff9529d9c52fb8daa5027ac6457daa092e948fe6d5c394b90.TimeOffRequestsRequestBuilder) {
    return i01750d724d2b41bff9529d9c52fb8daa5027ac6457daa092e948fe6d5c394b90.NewTimeOffRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.timeOffRequests.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ScheduleRequestBuilder) TimeOffRequestsById(id string)(*iabeccc97702e6f89f9a15e45a346d12c4f471cc17b4134ada692d86a7a7426f0.TimeOffRequestRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOffRequest_id"] = id
    }
    return iabeccc97702e6f89f9a15e45a346d12c4f471cc17b4134ada692d86a7a7426f0.NewTimeOffRequestRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ScheduleRequestBuilder) TimesOff()(*if0e6faf931107506fdb8d8969c5fb99869a339da13954a14cc90255ff691fae5.TimesOffRequestBuilder) {
    return if0e6faf931107506fdb8d8969c5fb99869a339da13954a14cc90255ff691fae5.NewTimesOffRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.schedule.timesOff.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ScheduleRequestBuilder) TimesOffById(id string)(*ia3f0ee46cad4161730260968933c300c8d5bd2479a2bb471cb27ede57b5aab3e.TimeOffRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["timeOff_id"] = id
    }
    return ia3f0ee46cad4161730260968933c300c8d5bd2479a2bb471cb27ede57b5aab3e.NewTimeOffRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
