package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2ae013fc81ba91d5d10b80be62caed4e99956daabb3b84179a6a184b7884aaf8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/calendarpermissions"
    i41f68cbd031e280cd9314fdf343ad01c6a463068816993b60b1b3d68d2f1b24e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/allowedcalendarsharingroleswithuser"
    i7b51844bfff56f40e3c810f1df986993e8a799ee9ff748a156ed0888561aafd7 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/multivalueextendedproperties"
    i86d724659a1e73542b6faea46e7462e1b2091578e5af51927ce32ab4d5b72b04 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/events"
    iafe04be5f81ab9a203d3b164994a8ec8b303f90370d5275183abdbfb786a2567 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/singlevalueextendedproperties"
    ibc4f63369782e235d92efacb7d8a92adc982ce6f4a8ed922d04a38c22c6c0114 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/calendarview"
    ie1b540015d2b7e7d0eff958b47f3eeb992ed46aebfa9e1012148252c6dc74666 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/getschedule"
    i328cea8068bc37e2162686f5b4ad08308cd2499e95e867f9b4c15b900d5e9006 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/multivalueextendedproperties/item"
    i8160ed636b0320b13f78035c28a4339c68773bb308d2c3681038908f53b01e2d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/events/item"
    i9100fbc12a05f6ed9305e5c77c0cfd560c6bc9604a44b2dbaeb08d315472300a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/calendarpermissions/item"
    ib02bf8b32e07e0960de050dbacc51102f64adc61ff32ae669a33693db8f1143d "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/calendarview/item"
    ib26a1e13e7005f479d4a7a3f39052b8b5c3b81388b28d37a10a7684cac9c9699 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/events/item/calendar/singlevalueextendedproperties/item"
)

// CalendarRequestBuilder builds and executes requests for operations under \groups\{group-id}\events\{event-id}\calendar
type CalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CalendarRequestBuilderDeleteOptions options for Delete
type CalendarRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CalendarRequestBuilderGetOptions options for Get
type CalendarRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CalendarRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CalendarRequestBuilderGetQueryParameters the calendar that contains the event. Navigation property. Read-only.
type CalendarRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select_escaped []string;
}
// CalendarRequestBuilderPatchOptions options for Patch
type CalendarRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// AllowedCalendarSharingRolesWithUser builds and executes requests for operations under \groups\{group-id}\events\{event-id}\calendar\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i41f68cbd031e280cd9314fdf343ad01c6a463068816993b60b1b3d68d2f1b24e.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i41f68cbd031e280cd9314fdf343ad01c6a463068816993b60b1b3d68d2f1b24e.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i2ae013fc81ba91d5d10b80be62caed4e99956daabb3b84179a6a184b7884aaf8.CalendarPermissionsRequestBuilder) {
    return i2ae013fc81ba91d5d10b80be62caed4e99956daabb3b84179a6a184b7884aaf8.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarPermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.calendar.calendarPermissions.item collection
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i9100fbc12a05f6ed9305e5c77c0cfd560c6bc9604a44b2dbaeb08d315472300a.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i9100fbc12a05f6ed9305e5c77c0cfd560c6bc9604a44b2dbaeb08d315472300a.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ibc4f63369782e235d92efacb7d8a92adc982ce6f4a8ed922d04a38c22c6c0114.CalendarViewRequestBuilder) {
    return ibc4f63369782e235d92efacb7d8a92adc982ce6f4a8ed922d04a38c22c6c0114.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CalendarViewById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.calendar.calendarView.item collection
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*ib02bf8b32e07e0960de050dbacc51102f64adc61ff32ae669a33693db8f1143d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return ib02bf8b32e07e0960de050dbacc51102f64adc61ff32ae669a33693db8f1143d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewCalendarRequestBuilderInternal instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/events/{event_id}/calendar{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCalendarRequestBuilder instantiates a new CalendarRequestBuilder and sets the default values.
func NewCalendarRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) CreateDeleteRequestInformation(options *CalendarRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) CreateGetRequestInformation(options *CalendarRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(options *CalendarRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) Delete(options *CalendarRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarRequestBuilder) Events()(*i86d724659a1e73542b6faea46e7462e1b2091578e5af51927ce32ab4d5b72b04.EventsRequestBuilder) {
    return i86d724659a1e73542b6faea46e7462e1b2091578e5af51927ce32ab4d5b72b04.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// EventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.calendar.events.item collection
func (m *CalendarRequestBuilder) EventsById(id string)(*i8160ed636b0320b13f78035c28a4339c68773bb308d2c3681038908f53b01e2d.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i8160ed636b0320b13f78035c28a4339c68773bb308d2c3681038908f53b01e2d.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) Get(options *CalendarRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewCalendar() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar), nil
}
func (m *CalendarRequestBuilder) GetSchedule()(*ie1b540015d2b7e7d0eff958b47f3eeb992ed46aebfa9e1012148252c6dc74666.GetScheduleRequestBuilder) {
    return ie1b540015d2b7e7d0eff958b47f3eeb992ed46aebfa9e1012148252c6dc74666.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i7b51844bfff56f40e3c810f1df986993e8a799ee9ff748a156ed0888561aafd7.MultiValueExtendedPropertiesRequestBuilder) {
    return i7b51844bfff56f40e3c810f1df986993e8a799ee9ff748a156ed0888561aafd7.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MultiValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.calendar.multiValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i328cea8068bc37e2162686f5b4ad08308cd2499e95e867f9b4c15b900d5e9006.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i328cea8068bc37e2162686f5b4ad08308cd2499e95e867f9b4c15b900d5e9006.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the calendar that contains the event. Navigation property. Read-only.
func (m *CalendarRequestBuilder) Patch(options *CalendarRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*iafe04be5f81ab9a203d3b164994a8ec8b303f90370d5275183abdbfb786a2567.SingleValueExtendedPropertiesRequestBuilder) {
    return iafe04be5f81ab9a203d3b164994a8ec8b303f90370d5275183abdbfb786a2567.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SingleValueExtendedPropertiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.events.item.calendar.singleValueExtendedProperties.item collection
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ib26a1e13e7005f479d4a7a3f39052b8b5c3b81388b28d37a10a7684cac9c9699.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ib26a1e13e7005f479d4a7a3f39052b8b5c3b81388b28d37a10a7684cac9c9699.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
