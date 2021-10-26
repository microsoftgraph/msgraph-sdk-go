package calendar

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2b51dfa3f8db575bfd1cf417b5dfb8c4cdf908c75a9564dea49c87ff2ba2bc39 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/getschedule"
    i4f49218b52ed641a8e7705f1fb4925e481940a5b9ec0118b430e4970a475cb08 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/events"
    i6553f3264d76a6863408a3af5fa86774eca3765b06f0fe036194fdf2db926905 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/allowedcalendarsharingroleswithuser"
    i844db35faf6567fd503449f547a64e3eaf7595c64b8add5f98c2aae3d90720dd "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/multivalueextendedproperties"
    i95a6afe86bb5e487bdceb9de1027d748d8804afbbe3a5a922d5104a170ab443e "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarview"
    ib5ad14f01725ed0675ea8c11b4c0ebeb325f9d573303a28f98d07748efeefe6b "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/singlevalueextendedproperties"
    ice38b6c9a46bf41898aa9a53d410ef37559ddffc1ce48af8b23f8bb621a3c500 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarpermissions"
    i0bf17eb670631e415b3b43afeb7544b7d1bc464841e409b7640a111e77646b89 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/events/item"
    i96bbb9a0a554809b814e788939046bf8755e4295f5545b98afe2270cbaf69a6b "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/multivalueextendedproperties/item"
    id90b80c691e5bb6700f3c8133c8918bccebec820a5565c9cd5cb4bc196d9cc99 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarpermissions/item"
    idfb0e47169e27daf47cf9b6e884bf998866b87ca4170a08a1586cf046dd4ad7c "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/calendarview/item"
    ief6caf4d252fe8fbf215adf140ff42d10ff0adf01eed1082a0075d857ca24315 "github.com/microsoftgraph/msgraph-sdk-go/me/events/item/calendar/singlevalueextendedproperties/item"
)

// Builds and executes requests for operations under \me\events\{event-id}\calendar
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
// Builds and executes requests for operations under \me\events\{event-id}\calendar\microsoft.graph.allowedCalendarSharingRoles(User='{User}')
// Parameters:
//  - User : Usage: User={User}
func (m *CalendarRequestBuilder) AllowedCalendarSharingRolesWithUser(user *string)(*i6553f3264d76a6863408a3af5fa86774eca3765b06f0fe036194fdf2db926905.AllowedCalendarSharingRolesWithUserRequestBuilder) {
    return i6553f3264d76a6863408a3af5fa86774eca3765b06f0fe036194fdf2db926905.NewAllowedCalendarSharingRolesWithUserRequestBuilderInternal(m.pathParameters, m.requestAdapter, user);
}
func (m *CalendarRequestBuilder) CalendarPermissions()(*ice38b6c9a46bf41898aa9a53d410ef37559ddffc1ce48af8b23f8bb621a3c500.CalendarPermissionsRequestBuilder) {
    return ice38b6c9a46bf41898aa9a53d410ef37559ddffc1ce48af8b23f8bb621a3c500.NewCalendarPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.calendar.calendarPermissions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarPermissionsById(id string)(*id90b80c691e5bb6700f3c8133c8918bccebec820a5565c9cd5cb4bc196d9cc99.CalendarPermissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["calendarPermission_id"] = id
    }
    return id90b80c691e5bb6700f3c8133c8918bccebec820a5565c9cd5cb4bc196d9cc99.NewCalendarPermissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CalendarRequestBuilder) CalendarView()(*i95a6afe86bb5e487bdceb9de1027d748d8804afbbe3a5a922d5104a170ab443e.CalendarViewRequestBuilder) {
    return i95a6afe86bb5e487bdceb9de1027d748d8804afbbe3a5a922d5104a170ab443e.NewCalendarViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.calendar.calendarView.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) CalendarViewById(id string)(*idfb0e47169e27daf47cf9b6e884bf998866b87ca4170a08a1586cf046dd4ad7c.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return idfb0e47169e27daf47cf9b6e884bf998866b87ca4170a08a1586cf046dd4ad7c.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new CalendarRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewCalendarRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CalendarRequestBuilder) {
    m := &CalendarRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/events/{event_id}/calendar{?select}";
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
func (m *CalendarRequestBuilder) Events()(*i4f49218b52ed641a8e7705f1fb4925e481940a5b9ec0118b430e4970a475cb08.EventsRequestBuilder) {
    return i4f49218b52ed641a8e7705f1fb4925e481940a5b9ec0118b430e4970a475cb08.NewEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.calendar.events.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) EventsById(id string)(*i0bf17eb670631e415b3b43afeb7544b7d1bc464841e409b7640a111e77646b89.EventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["event_id1"] = id
    }
    return i0bf17eb670631e415b3b43afeb7544b7d1bc464841e409b7640a111e77646b89.NewEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) GetSchedule()(*i2b51dfa3f8db575bfd1cf417b5dfb8c4cdf908c75a9564dea49c87ff2ba2bc39.GetScheduleRequestBuilder) {
    return i2b51dfa3f8db575bfd1cf417b5dfb8c4cdf908c75a9564dea49c87ff2ba2bc39.NewGetScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CalendarRequestBuilder) MultiValueExtendedProperties()(*i844db35faf6567fd503449f547a64e3eaf7595c64b8add5f98c2aae3d90720dd.MultiValueExtendedPropertiesRequestBuilder) {
    return i844db35faf6567fd503449f547a64e3eaf7595c64b8add5f98c2aae3d90720dd.NewMultiValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.calendar.multiValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) MultiValueExtendedPropertiesById(id string)(*i96bbb9a0a554809b814e788939046bf8755e4295f5545b98afe2270cbaf69a6b.MultiValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["multiValueLegacyExtendedProperty_id"] = id
    }
    return i96bbb9a0a554809b814e788939046bf8755e4295f5545b98afe2270cbaf69a6b.NewMultiValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
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
func (m *CalendarRequestBuilder) SingleValueExtendedProperties()(*ib5ad14f01725ed0675ea8c11b4c0ebeb325f9d573303a28f98d07748efeefe6b.SingleValueExtendedPropertiesRequestBuilder) {
    return ib5ad14f01725ed0675ea8c11b4c0ebeb325f9d573303a28f98d07748efeefe6b.NewSingleValueExtendedPropertiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.events.item.calendar.singleValueExtendedProperties.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *CalendarRequestBuilder) SingleValueExtendedPropertiesById(id string)(*ief6caf4d252fe8fbf215adf140ff42d10ff0adf01eed1082a0075d857ca24315.SingleValueLegacyExtendedPropertyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["singleValueLegacyExtendedProperty_id"] = id
    }
    return ief6caf4d252fe8fbf215adf140ff42d10ff0adf01eed1082a0075d857ca24315.NewSingleValueLegacyExtendedPropertyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
