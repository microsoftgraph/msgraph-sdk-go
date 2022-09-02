package schedule

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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

// ScheduleRequestBuilder provides operations to manage the schedule property of the microsoft.graph.team entity.
type ScheduleRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ScheduleRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ScheduleRequestBuilderGetQueryParameters the schedule of shifts for this team.
type ScheduleRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ScheduleRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ScheduleRequestBuilderGetQueryParameters
}
// ScheduleRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ScheduleRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewScheduleRequestBuilderInternal instantiates a new ScheduleRequestBuilder and sets the default values.
func NewScheduleRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ScheduleRequestBuilder) {
    m := &ScheduleRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}/schedule{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property schedule for teams
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property schedule for teams
func (m *ScheduleRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property schedule in teams
func (m *ScheduleRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property schedule in teams
func (m *ScheduleRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *ScheduleRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property schedule for teams
func (m *ScheduleRequestBuilder) Delete(ctx context.Context, requestConfiguration *ScheduleRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// Get the schedule of shifts for this team.
func (m *ScheduleRequestBuilder) Get(ctx context.Context, requestConfiguration *ScheduleRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
// OfferShiftRequests the offerShiftRequests property
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
        urlTplParams["offerShiftRequest%2Did"] = id
    }
    return id6d2fdc6b308e5bcba0d92cb156fdfdf7e8749fb58174b0d3bfafaaba4972a32.NewOfferShiftRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShiftChangeRequests the openShiftChangeRequests property
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
        urlTplParams["openShiftChangeRequest%2Did"] = id
    }
    return i3e97a7fafa20db003ef2a75a5e79b9165e6ae74b4a976948a89557917308b1ae.NewOpenShiftChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// OpenShifts the openShifts property
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
        urlTplParams["openShift%2Did"] = id
    }
    return i077cef4a84be1e18919076de6e3b973ebcf3c54ffb2a0db7aec5fdaf05f74873.NewOpenShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property schedule in teams
func (m *ScheduleRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Scheduleable, requestConfiguration *ScheduleRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// SchedulingGroups the schedulingGroups property
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
        urlTplParams["schedulingGroup%2Did"] = id
    }
    return id08a6e1dd69521cd4e2e6c2c076205fa4095da47b2e1624b4768ab0739c44eb1.NewSchedulingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Share the share property
func (m *ScheduleRequestBuilder) Share()(*i3bddaf301c4fc5b8766f262b90a121596090b82c6065ad0672073263696b869d.ShareRequestBuilder) {
    return i3bddaf301c4fc5b8766f262b90a121596090b82c6065ad0672073263696b869d.NewShareRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Shifts the shifts property
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
        urlTplParams["shift%2Did"] = id
    }
    return i4717029473a05fb649af93bbf657da0c15821fab39c94b18fa872f9884772aa1.NewShiftItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SwapShiftsChangeRequests the swapShiftsChangeRequests property
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
        urlTplParams["swapShiftsChangeRequest%2Did"] = id
    }
    return ia798c3c70ac49063707043a913480d584e918c77145a9c2670279dbef879caa6.NewSwapShiftsChangeRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffReasons the timeOffReasons property
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
        urlTplParams["timeOffReason%2Did"] = id
    }
    return icd9016d4d1909d7883252b08c395dcf7ac68e8ff8385953a5652dd1e348add0f.NewTimeOffReasonItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimeOffRequests the timeOffRequests property
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
        urlTplParams["timeOffRequest%2Did"] = id
    }
    return iabeccc97702e6f89f9a15e45a346d12c4f471cc17b4134ada692d86a7a7426f0.NewTimeOffRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// TimesOff the timesOff property
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
        urlTplParams["timeOff%2Did"] = id
    }
    return ia3f0ee46cad4161730260968933c300c8d5bd2479a2bb471cb27ede57b5aab3e.NewTimeOffItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
