package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i050c51f104551907881dca47a5b68eb85f83281bd2389b166735b279d46d5f4f "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/getschedule"
    i19f3c4225d6a53db2dbce2440df3ebb4533353f2c041c409dff7fd46d98df2be "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/multivalueextendedproperties"
    i2a669a13e7e0aed9933bce9ec015402b9da48e02775df04fe699ebf7e3813d13 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview"
    i765812ce9f375064376d8d08bdf56e7add0ab6cae4509707e547382247b0e4c5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events"
    i960ec40c7cec3d169b8676496546d8f7fc83a56efa118a6f891140469e03dc1b "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarpermissions"
    ib3e2d60ac4866f39298101caac31b65749b22f8324553621dd8ab89cda479ac5 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/allowedcalendarsharingroleswithuser"
    ie0575c19ab0c699d8e90db2fc7c7e71a798afd6e62fd458a4b24d77e3abb05f4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/singlevalueextendedproperties"
    i3ae5e02d4fe14108f67419bd2b06c86a94c5d708680ddf307e13e5626b6e4f90 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/events/item"
    i63cfd09f19b0541d7a22da8c42a45e2111a15329049eafa66f5b578b523446f0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarpermissions/item"
    i6bf629fd3755ecd71e269a2dd0099d187090dbef283089d7daad2429ce8064d7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/multivalueextendedproperties/item"
    i8018757b2ab9a734b5b294b2487ab2d8bc9f74ae3afdbc9532dcb7bc7a9c1825 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/singlevalueextendedproperties/item"
    ia02c9708adeba39a0c7bb01cfa028dde8889b5c1460ddc0736d84c28c8080558 "github.com/microsoftgraph/msgraph-sdk-go/users/item/calendars/item/calendarview/item"
)

// CalendarItemRequestBuilder provides operations to manage the calendars property of the microsoft.graph.user entity.
type CalendarItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CalendarItemRequestBuilderDeleteOptions options for Delete
type CalendarItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// CalendarItemRequestBuilderGetOptions options for Get
type CalendarItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *CalendarItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// CalendarItemRequestBuilderGetQueryParameters the user's calendars. Read-only. Nullable.
type CalendarItemRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string;
}
// CalendarItemRequestBuilderPatchOptions options for Patch
type CalendarItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AllowedCalendarSharingRolesWithUser provides operations to call the allowedCalendarSharingRoles method.
func (m *CalendarItemRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ib3e2d60ac4866f39298101caac31b65749b22f8324553621dd8ab89cda479ac5.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ib3e2d60ac4866f39298101caac31b65749b22f8324553621dd8ab89cda479ac5.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
// CalendarPermissions the calendarPermissions property
func (m *CalendarItemRequestBuilder) CalendarPermissions()(*i960ec40c7cec3d169b8676496546d8f7fc83a56efa118a6f891140469e03dc1b.CalendarPermissionsRequestBuilder) {
    return i960ec40c7cec3d169b8676496546d8f7fc83a56efa118a6f891140469e03dc1b.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarPermissions.item collection
func (m *CalendarItemRequestBuilder) CalendarPermissionsById(id string)(*i63cfd09f19b0541d7a22da8c42a45e2111a15329049eafa66f5b578b523446f0.CalendarPermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i63cfd09f19b0541d7a22da8c42a45e2111a15329049eafa66f5b578b523446f0.NewCalendarPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CalendarView the calendarView property
func (m *CalendarItemRequestBuilder) CalendarView()(*i2a669a13e7e0aed9933bce9ec015402b9da48e02775df04fe699ebf7e3813d13.CalendarViewRequestBuilder) {
    return i2a669a13e7e0aed9933bce9ec015402b9da48e02775df04fe699ebf7e3813d13.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.calendarView.item collection
func (m *CalendarItemRequestBuilder) CalendarViewById(id string)(*ia02c9708adeba39a0c7bb01cfa028dde8889b5c1460ddc0736d84c28c8080558.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return ia02c9708adeba39a0c7bb01cfa028dde8889b5c1460ddc0736d84c28c8080558.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarItemRequestBuilderInternal instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarItemRequestBuilder) {
    m := &CalendarItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/calendars/{calendar_id}{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarItemRequestBuilder instantiates a new CalendarItemRequestBuilder and sets the default values.
func NewCalendarItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CalendarItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calendars for users
func (m *CalendarItemRequestBuilder) CreateDeleteRequestInformation(options *CalendarItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the user's calendars. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) CreateGetRequestInformation(options *CalendarItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calendars in users
func (m *CalendarItemRequestBuilder) CreatePatchRequestInformation(options *CalendarItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property calendars for users
func (m *CalendarItemRequestBuilder) Delete(options *CalendarItemRequestBuilderDeleteOptions)(error) {
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
func (m *CalendarItemRequestBuilder) Events()(*i765812ce9f375064376d8d08bdf56e7add0ab6cae4509707e547382247b0e4c5.EventsRequestBuilder) {
    return i765812ce9f375064376d8d08bdf56e7add0ab6cae4509707e547382247b0e4c5.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.events.item collection
func (m *CalendarItemRequestBuilder) EventsById(id string)(*i3ae5e02d4fe14108f67419bd2b06c86a94c5d708680ddf307e13e5626b6e4f90.EventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id"] = id
    }
    return i3ae5e02d4fe14108f67419bd2b06c86a94c5d708680ddf307e13e5626b6e4f90.NewEventItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the user's calendars. Read-only. Nullable.
func (m *CalendarItemRequestBuilder) Get(options *CalendarItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Calendarable, error) {
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
func (m *CalendarItemRequestBuilder) GetSchedule()(*i050c51f104551907881dca47a5b68eb85f83281bd2389b166735b279d46d5f4f.GetScheduleRequestBuilder) {
    return i050c51f104551907881dca47a5b68eb85f83281bd2389b166735b279d46d5f4f.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedProperties the multiValueExtendedProperties property
func (m *CalendarItemRequestBuilder) MultiValueExtendedProperties()(*i19f3c4225d6a53db2dbce2440df3ebb4533353f2c041c409dff7fd46d98df2be.MultiValueExtendedPropertiesRequestBuilder) {
    return i19f3c4225d6a53db2dbce2440df3ebb4533353f2c041c409dff7fd46d98df2be.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.multiValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i6bf629fd3755ecd71e269a2dd0099d187090dbef283089d7daad2429ce8064d7.MultiValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i6bf629fd3755ecd71e269a2dd0099d187090dbef283089d7daad2429ce8064d7.NewMultiValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calendars in users
func (m *CalendarItemRequestBuilder) Patch(options *CalendarItemRequestBuilderPatchOptions)(error) {
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
func (m *CalendarItemRequestBuilder) SingleValueExtendedProperties()(*ie0575c19ab0c699d8e90db2fc7c7e71a798afd6e62fd458a4b24d77e3abb05f4.SingleValueExtendedPropertiesRequestBuilder) {
    return ie0575c19ab0c699d8e90db2fc7c7e71a798afd6e62fd458a4b24d77e3abb05f4.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.calendars.item.singleValueExtendedProperties.item collection
func (m *CalendarItemRequestBuilder) SingleValueExtendedPropertiesById(id string)(*i8018757b2ab9a734b5b294b2487ab2d8bc9f74ae3afdbc9532dcb7bc7a9c1825.SingleValueLegacyExtendedPropertyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return i8018757b2ab9a734b5b294b2487ab2d8bc9f74ae3afdbc9532dcb7bc7a9c1825.NewSingleValueLegacyExtendedPropertyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
