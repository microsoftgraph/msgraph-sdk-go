package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1180f93b60933e34ef33f96afc07a4be77403819faa65b894e5a0de04b3d9b48 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/multivalueextendedproperties"
    i3b3f84caa9fefd43d17756b51c5dbcf6296d76f57030fe4a59e608b0fbf480f6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/singlevalueextendedproperties"
    i4ee725269f1619c8d17d5002191fa563cc224a20987a3d28e8f7c0a96937f1e3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/events"
    i83307b0bb4dd3693dd9c3c2631b500576b5c1ad4c1c2c585c66fabe5e123655d "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/calendarpermissions"
    ia2ad2550d956b2e29823e9535842c3e00fe9e13e1c590951bae51011909f9271 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/calendarview"
    ia82327010b7ae094735a40ff755cb51be477fa9a5001db67817db9b81a7bc86c "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/allowedcalendarsharingroleswithuser"
    if9b4dbf3da3810313f51551c9d652cb7ebdc325eb7e638afecac9865bd2aa93d "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/getschedule"
    i095112975e3f30d2ec944a5cf3affbfd9ed66695f6a8e608e40afe80c3b69439 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/multivalueextendedproperties/item"
    i605d24cfa872d4b9a5c04521c9664e6c19f5407d888763d87eed35f0b9eca0c6 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/events/item"
    i6fd9c916708d5edf8f8ecadef507faef614e89fb24c563ba8592ca658364a4be "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/calendarview/item"
    i7f98b0c4658cddc42b780b24f99de687366d794f9cd6272cc1d2f8c4204302f7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/calendarpermissions/item"
    if2bed076e6d8a8919430bd29899d593f734ba6dbe96fe4760e4bbb926619a281 "github.com/microsoftgraph/msgraph-sdk-go/users/item/events/item/calendar/singlevalueextendedproperties/item"
)

// Builds and executes requests for operations under \users\{user-id}\events\{event-id}\calendar
type CalendarRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// The calendar that contains the event. Navigation property. Read-only.
type CalendarRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Select properties to be returned
    Select_escaped []string;
}
// Builds and executes requests for operations under \users\{user-id}\events\{event-id}\calendar\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
// Parameters:
//  - User : Usage: User={User}
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*ia82327010b7ae094735a40ff755cb51be477fa9a5001db67817db9b81a7bc86c.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return ia82327010b7ae094735a40ff755cb51be477fa9a5001db67817db9b81a7bc86c.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*i83307b0bb4dd3693dd9c3c2631b500576b5c1ad4c1c2c585c66fabe5e123655d.CalendarPermissionsRequestBuilder) {
    return i83307b0bb4dd3693dd9c3c2631b500576b5c1ad4c1c2c585c66fabe5e123655d.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.calendar.calendarPermissions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*i7f98b0c4658cddc42b780b24f99de687366d794f9cd6272cc1d2f8c4204302f7.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return i7f98b0c4658cddc42b780b24f99de687366d794f9cd6272cc1d2f8c4204302f7.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*ia2ad2550d956b2e29823e9535842c3e00fe9e13e1c590951bae51011909f9271.CalendarViewRequestBuilder) {
    return ia2ad2550d956b2e29823e9535842c3e00fe9e13e1c590951bae51011909f9271.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.calendar.calendarView.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*i6fd9c916708d5edf8f8ecadef507faef614e89fb24c563ba8592ca658364a4be.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i6fd9c916708d5edf8f8ecadef507faef614e89fb24c563ba8592ca658364a4be.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new CalendarRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/users/{user_id}/events/{event_id}/calendar{?select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new CalendarRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCalendarRequestBuilderInternal(urlParams, requestAdapter)
}
// The calendar that contains the event. Navigation property. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *CalendarRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The calendar that contains the event. Navigation property. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *CalendarRequestBuilder) CreateGetRequestInformation(q func (value *CalendarRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(CalendarRequestBuilderGetQueryParameters)
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
// The calendar that contains the event. Navigation property. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *CalendarRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The calendar that contains the event. Navigation property. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *CalendarRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CalendarRequestBuilder) Events()(*i4ee725269f1619c8d17d5002191fa563cc224a20987a3d28e8f7c0a96937f1e3.EventsRequestBuilder) {
    return i4ee725269f1619c8d17d5002191fa563cc224a20987a3d28e8f7c0a96937f1e3.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.calendar.events.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) EventsById(id string)(*i605d24cfa872d4b9a5c04521c9664e6c19f5407d888763d87eed35f0b9eca0c6.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i605d24cfa872d4b9a5c04521c9664e6c19f5407d888763d87eed35f0b9eca0c6.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendar that contains the event. Navigation property. Read-only.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *CalendarRequestBuilder) Get(q func (value *CalendarRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewCalendar() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar), nil
}
func (m *CalendarRequestBuilder) GetSchedule()(*if9b4dbf3da3810313f51551c9d652cb7ebdc325eb7e638afecac9865bd2aa93d.GetScheduleRequestBuilder) {
    return if9b4dbf3da3810313f51551c9d652cb7ebdc325eb7e638afecac9865bd2aa93d.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i1180f93b60933e34ef33f96afc07a4be77403819faa65b894e5a0de04b3d9b48.MultiValueExtendedPropertiesRequestBuilder) {
    return i1180f93b60933e34ef33f96afc07a4be77403819faa65b894e5a0de04b3d9b48.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.calendar.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i095112975e3f30d2ec944a5cf3affbfd9ed66695f6a8e608e40afe80c3b69439.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i095112975e3f30d2ec944a5cf3affbfd9ed66695f6a8e608e40afe80c3b69439.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The calendar that contains the event. Navigation property. Read-only.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *CalendarRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Calendar, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*i3b3f84caa9fefd43d17756b51c5dbcf6296d76f57030fe4a59e608b0fbf480f6.SingleValueExtendedPropertiesRequestBuilder) {
    return i3b3f84caa9fefd43d17756b51c5dbcf6296d76f57030fe4a59e608b0fbf480f6.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.events.item.calendar.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*if2bed076e6d8a8919430bd29899d593f734ba6dbe96fe4760e4bbb926619a281.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return if2bed076e6d8a8919430bd29899d593f734ba6dbe96fe4760e4bbb926619a281.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
