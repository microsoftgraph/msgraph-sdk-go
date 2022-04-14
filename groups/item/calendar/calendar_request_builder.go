package calendar

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i325e8c5a186a2ac5ba63357864499a2daeca2d6b1fe397a8e242931a81e90fd6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/singlevalueextendedproperties"
    i564e44237d1fd29264e29294fe61a277a2b68966ba3da977d4dc4188473efc6f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/allowedcalendarsharingroleswithuser"
    i59567e44c84457909a7acd2691855472309ae72d46905ad5b927e32ac1bc6775 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarpermissions"
    i8a7b3ddf32454c04e11c42fafddbe8a65a1b8c29699e948e42fb5fb489efff7a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/getschedule"
    i91f19a09d718f883e167c3c1fea6e1dc0bafbd34f8ca677f95c72d7745bcd847 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/multivalueextendedproperties"
    ia4056a24d617675f13dc550d9eb8d60439d535c22994554b767157caae834324 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview"
    iae525d555fa999780edca2ea54ea681baf7b616326144106b66e173f4930b985 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events"
    i066e4b225a90f837574cba0d06582cb5234a89a58269aad5ccca1e0d592385ed "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarview/item"
    i32f1c0e0a1af3233116ddc476f2459674db4e63222c4a5762ca7049d50e06a83 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/singlevalueextendedproperties/item"
    i84f844d52d5e497fca0d718f40dd524efb38a3a493f27e5b266d0f0b0d06bd58 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/calendarpermissions/item"
    ieb71583a78a26be4af7a9008915150796bd995e96567f7d812206ea5c47e34ee "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/events/item"
    iff61068eadc42d2652152ea70da0c96f72f2b3e78201eeb2e58a48230426fa5a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/calendar/multivalueextendedproperties/item"
)

// CalendarRequestBuilder provides operations to manage the calendar property of the microsoft.graph.group entity.
type CalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CalendarRequestBuilderDeleteOptions options for Delete
type CalendarRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CalendarRequestBuilderGetOptions options for Get
type CalendarRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CalendarRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CalendarRequestBuilderGetQueryParameters the group's calendar. Read-only.
type CalendarRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CalendarRequestBuilderPatchOptions options for Patch
type CalendarRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i564e44237d1fd29264e29294fe61a277a2b68966ba3da977d4dc4188473efc6f.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i564e44237d1fd29264e29294fe61a277a2b68966ba3da977d4dc4188473efc6f.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
// CalendarPermissions the calendarPermissions property
func (m *CalendarRequestBuilder) CalendarPermissions()(*i59567e44c84457909a7acd2691855472309ae72d46905ad5b927e32ac1bc6775.CalendarPermissionsRequestBuilder) {
    return i59567e44c84457909a7acd2691855472309ae72d46905ad5b927e32ac1bc6775.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i84f844d52d5e497fca0d718f40dd524efb38a3a493f27e5b266d0f0b0d06bd58.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission%2Did"] = id
    }
    return i84f844d52d5e497fca0d718f40dd524efb38a3a493f27e5b266d0f0b0d06bd58.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *CalendarRequestBuilder) CalendarView()(*ia4056a24d617675f13dc550d9eb8d60439d535c22994554b767157caae834324.CalendarViewRequestBuilder) {
    return ia4056a24d617675f13dc550d9eb8d60439d535c22994554b767157caae834324.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i066e4b225a90f837574cba0d06582cb5234a89a58269aad5ccca1e0d592385ed.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return i066e4b225a90f837574cba0d06582cb5234a89a58269aad5ccca1e0d592385ed.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/calendar{?%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarRequestBuilder instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calendar for groups
func (m *CalendarRequestBuilder) CreateDeleteRequestInformation(options *CalendarRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the group's calendar. Read-only.
func (m *CalendarRequestBuilder) CreateGetRequestInformation(options *CalendarRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calendar in groups
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(options *CalendarRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property calendar for groups
func (m *CalendarRequestBuilder) Delete(options *CalendarRequestBuilderDeleteOptions)(error) {
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
// Events the events property
func (m *CalendarRequestBuilder) Events()(*iae525d555fa999780edca2ea54ea681baf7b616326144106b66e173f4930b985.EventsRequestBuilder) {
    return iae525d555fa999780edca2ea54ea681baf7b616326144106b66e173f4930b985.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*ieb71583a78a26be4af7a9008915150796bd995e96567f7d812206ea5c47e34ee.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event%2Did"] = id
    }
    return ieb71583a78a26be4af7a9008915150796bd995e96567f7d812206ea5c47e34ee.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the group's calendar. Read-only.
func (m *CalendarRequestBuilder) Get(options *CalendarRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCalendarFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable), nil
}
// GetSchedule the getSchedule property
func (m *CalendarRequestBuilder) GetSchedule()(*i8a7b3ddf32454c04e11c42fafddbe8a65a1b8c29699e948e42fb5fb489efff7a.GetScheduleRequestBuilder) {
    return i8a7b3ddf32454c04e11c42fafddbe8a65a1b8c29699e948e42fb5fb489efff7a.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i91f19a09d718f883e167c3c1fea6e1dc0bafbd34f8ca677f95c72d7745bcd847.MultiValueExtendedPropertiesRequestBuilder) {
    return i91f19a09d718f883e167c3c1fea6e1dc0bafbd34f8ca677f95c72d7745bcd847.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*iff61068eadc42d2652152ea70da0c96f72f2b3e78201eeb2e58a48230426fa5a.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty%2Did"] = id
    }
    return iff61068eadc42d2652152ea70da0c96f72f2b3e78201eeb2e58a48230426fa5a.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendar in groups
func (m *CalendarRequestBuilder) Patch(options *CalendarRequestBuilderPatchOptions)(error) {
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
// SingleValueExtendedProperties the singleValueExtendedProperties property
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i325e8c5a186a2ac5ba63357864499a2daeca2d6b1fe397a8e242931a81e90fd6.SingleValueExtendedPropertiesRequestBuilder) {
    return i325e8c5a186a2ac5ba63357864499a2daeca2d6b1fe397a8e242931a81e90fd6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.calendar.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i32f1c0e0a1af3233116ddc476f2459674db4e63222c4a5762ca7049d50e06a83.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty%2Did"] = id
    }
    return i32f1c0e0a1af3233116ddc476f2459674db4e63222c4a5762ca7049d50e06a83.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
